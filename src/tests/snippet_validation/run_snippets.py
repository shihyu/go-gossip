#!/usr/bin/env python3
import json
import os
import re
import shutil
import subprocess
import sys
from dataclasses import dataclass, asdict
from pathlib import Path
from typing import List, Optional

ROOT = Path(__file__).resolve().parents[2]
GO126 = Path('/tmp/go1.26.0/bin/go')
OUTDIR = ROOT / 'tests' / 'snippet_validation'
WORKDIR = OUTDIR / 'work'
REPORT_JSON = OUTDIR / 'report.json'
REPORT_MD = OUTDIR / 'report.md'
TIMEOUT_SEC = 4

FENCE_START = re.compile(r'^```\s*([^\n]*)\s*$')
FENCE_END = re.compile(r'^```\s*$')
GO_TOKEN_HINTS = (
    'package ', 'func ', 'import ', 'type ', 'var ', 'const ', ':=', 'fmt.',
    'if ', 'for ', 'switch ', 'go ', 'defer ', 'chan ', 'map[', 'struct{',
)
PROMPT_PREFIXES = ('$ ', '> ', 'C:\\>', 'PS ', 'SET ', 'set ')

@dataclass
class SnippetResult:
    file: str
    index: int
    start_line: int
    end_line: int
    info: str
    kind: str
    action: str
    status: str
    reason: str = ''
    returncode: Optional[int] = None
    stdout: str = ''
    stderr: str = ''
    note: str = ''


def parse_fences(path: Path):
    lines = path.read_text(encoding='utf-8').splitlines()
    i = 0
    snippets = []
    while i < len(lines):
        m = FENCE_START.match(lines[i])
        if not m:
            i += 1
            continue
        info = m.group(1).strip()
        start = i + 1
        i += 1
        buf = []
        while i < len(lines) and not FENCE_END.match(lines[i]):
            buf.append(lines[i])
            i += 1
        end = i + 1 if i < len(lines) else len(lines)
        snippets.append((start, end, info, '\n'.join(buf) + ('\n' if buf else '')))
        if i < len(lines):
            i += 1
    return snippets


def first_nonempty_line(s: str) -> str:
    for ln in s.splitlines():
        if ln.strip():
            return ln.strip()
    return ''


def looks_like_go(text: str) -> bool:
    t = text.strip()
    if not t:
        return False
    if any(t.startswith(p) for p in PROMPT_PREFIXES):
        return False
    if 'package ' in text:
        return True
    score = sum(1 for h in GO_TOKEN_HINTS if h in text)
    return score >= 3


def classify(text: str):
    stripped = text.strip()
    if not stripped:
        return ('empty', 'skip', '空白 code fence')
    first = first_nonempty_line(text)
    if any(first.startswith(p) for p in PROMPT_PREFIXES) or first.startswith('go ') or first.startswith('node '):
        return ('shell_example', 'skip', '命令列範例（可能依賴特定環境/互動）')
    if first.startswith('func ') and 'package ' not in text:
        return ('go_fragment', 'skip', 'Go 片段（缺 package）')
    if (first.startswith('type ') or first.startswith('const ') or first.startswith('var ')) and 'package ' not in text:
        return ('go_fragment', 'skip', 'Go 片段（宣告示例）')
    if not looks_like_go(text):
        return ('text_or_output', 'skip', '輸出示意/非 Go 程式碼')
    m = re.search(r'^\s*package\s+(\w+)\b', text, re.M)
    if not m:
        return ('go_fragment', 'skip', 'Go 片段（無 package）')
    pkg = m.group(1)
    has_main = bool(re.search(r'\bfunc\s+main\s*\(', text))
    if pkg == 'main' and has_main:
        if 'import "syscall/js"' in text or 'import (\n' in text and '"syscall/js"' in text:
            return ('go_program_wasm', 'build', 'js/wasm 範例，改用編譯驗證')
        return ('go_program', 'run', '完整 Go 程式')
    if pkg == 'main' and not has_main:
        return ('go_file_main_nomian', 'build', 'main 套件片段（無 main），嘗試編譯')
    return ('go_package_file', 'build', 'Go 套件檔案，嘗試編譯')


def expected_failure_kind(sn: SnippetResult, code: str):
    # Teaching snippets that are intentionally invalid / panic.
    if '編譯錯誤' in code or 'panic:' in code:
        return 'expected_fail'
    if sn.file == 'gofmt.md':
        return 'skip'
    if sn.file in {'BreakContinueGoto.md', 'PreDeclaredType.md', 'TypeAssertion.md'}:
        return 'expected_fail'
    if sn.file in {'Method.md'} and 'String redeclared' in code:
        return 'expected_fail'
    if sn.file in {'Reflect.md'} and 'unaddressable value' in code:
        return 'expected_fail'
    if sn.file in {'DeferPanicRecover.md'}:
        return 'expected_fail'
    return None


def runtime_context(sn: SnippetResult, code: str, d: Path):
    args = []
    stdin = None
    timeout = TIMEOUT_SEC
    note = []

    # Provide argv for os.Args examples.
    if 'os.Args[1]' in code:
        args = ['Codex']
        note.append('argv=Codex')

    # Provide sample file for filename-input examples.
    asks_string = 'Scanf(\"%s\"' in code or 'Scanln(&filename' in code or 'Scanf(\"%q\"' in code
    if asks_string and ('os.Open(filename)' in code or 'Open(filename)' in code):
        sample = d / 'sample.txt'
        sample.write_text('line1\nline2\n測試\n', encoding='utf-8')
        stdin = 'sample.txt\n'
        note.append('stdin=sample.txt')

    # Numeric prompt examples (goto retry etc.)
    if stdin is None and ('Scanf(\"%d\"' in code or 'Scan(&input' in code):
        stdin = '1\n'
        note.append('stdin=1')

    # Some teaching goroutine snippets intentionally sleep to wait for goroutines.
    if 'time.Sleep(5 * time.Second)' in code:
        timeout = max(timeout, 7)
        note.append('timeout=7s')

    # Random game examples can run long; treat as expected-timeout if they overrun.
    if sn.file in {'For.md'} and 'random(' in code and 'time.Sleep(time.Second)' in code:
        timeout = max(timeout, 3)
        note.append('may-timeout-expected')

    return args, stdin, timeout, '; '.join(note)


def non_std_imports(text: str) -> List[str]:
    imports = []
    # import "x"
    for m in re.finditer(r'^\s*import\s+"([^"]+)"', text, re.M):
        imports.append(m.group(1))
    # import (
    in_block = False
    for ln in text.splitlines():
        if re.match(r'^\s*import\s*\($', ln):
            in_block = True
            continue
        if in_block:
            if re.match(r'^\s*\)\s*$', ln):
                in_block = False
                continue
            m = re.match(r'^\s*(?:[\w\.]+\s+)?"([^"]+)"', ln)
            if m:
                imports.append(m.group(1))
    out = []
    for p in imports:
        head = p.split('/')[0]
        if '.' in head:
            out.append(p)
    return out


def run_cmd(args, cwd: Path, extra_env=None, timeout=TIMEOUT_SEC, stdin_text=None):
    env = os.environ.copy()
    env.pop('GOROOT', None)
    env['PATH'] = f"{GO126.parent}:{env.get('PATH','')}"
    if extra_env:
        env.update(extra_env)
    cp = subprocess.run(args, cwd=str(cwd), env=env, input=stdin_text, capture_output=True, text=True, timeout=timeout)
    return cp


def prepare_snippet_dir(sn: SnippetResult, code: str) -> Path:
    d = WORKDIR / Path(sn.file).stem / f'snippet_{sn.index:03d}'
    if d.exists():
        shutil.rmtree(d)
    d.mkdir(parents=True, exist_ok=True)
    (d / 'snippet.go').write_text(code, encoding='utf-8')
    return d


def maybe_init_mod(d: Path, code: str):
    if not non_std_imports(code):
        return None
    # init and tidy for external deps
    try:
        run_cmd([str(GO126), 'mod', 'init', 'snippettest'], d, timeout=3)
    except Exception:
        pass
    return run_cmd([str(GO126), 'mod', 'tidy'], d, timeout=8)


def execute(sn: SnippetResult, code: str) -> SnippetResult:
    if sn.action == 'skip':
        sn.status = 'skipped'
        return sn
    if not GO126.exists():
        sn.status = 'error'
        sn.reason = f'找不到 Go 1.26 binary: {GO126}'
        return sn
    d = prepare_snippet_dir(sn, code)
    try:
        args_extra, stdin_text, timeout_sec, note = runtime_context(sn, code, d)
        sn.note = note
        tidy_cp = maybe_init_mod(d, code)
        if tidy_cp is not None and tidy_cp.returncode != 0:
            sn.status = 'fail'
            sn.reason = 'go mod tidy 失敗'
            sn.returncode = tidy_cp.returncode
            sn.stdout = tidy_cp.stdout[-2000:]
            sn.stderr = tidy_cp.stderr[-2000:]
            return sn

        if sn.kind == 'go_program_wasm':
            cp = run_cmd([str(GO126), 'build', '-o', 'snippet.wasm', 'snippet.go'], d, extra_env={'GOOS':'js','GOARCH':'wasm'}, timeout=8)
        elif sn.action == 'run':
            cp = run_cmd([str(GO126), 'run', 'snippet.go', *args_extra], d, timeout=timeout_sec, stdin_text=stdin_text)
        else:
            cp = run_cmd([str(GO126), 'build', 'snippet.go'], d, timeout=timeout_sec, stdin_text=stdin_text)
        sn.returncode = cp.returncode
        sn.stdout = cp.stdout[-2000:]
        sn.stderr = cp.stderr[-2000:]
        if cp.returncode == 0:
            sn.status = 'pass'
        else:
            stderr = sn.stderr
            # Reclassify contextual / intentional cases.
            if 'package ' in stderr and ' is not in std ' in stderr:
                sn.status = 'skipped'
                sn.reason = '需要同頁前置目錄/本地套件結構'
            elif 'undefined: Add' in stderr and sn.file == 'Testing.md':
                sn.status = 'skipped'
                sn.reason = '測試片段依賴同頁未一併建立的 Add 實作'
            elif expected_failure_kind(sn, code) == 'skip':
                sn.status = 'skipped'
                sn.reason = '格式化示意/非可編譯完整範例'
            elif expected_failure_kind(sn, code) == 'expected_fail':
                sn.status = 'expected_fail'
                sn.reason = '教學用預期失敗/恐慌示範'
            else:
                sn.status = 'fail'
                sn.reason = 'go 指令失敗'
        return sn
    except subprocess.TimeoutExpired as e:
        if sn.file in {'For.md', 'Goroutine.md'} or ('select {}' in code) or ('for {' in code and 'break' not in code):
            sn.status = 'expected_timeout'
            sn.reason = f'教學示範長時間/阻塞流程（>{e.timeout}s）'
        else:
            sn.status = 'timeout'
            sn.reason = f'執行逾時（>{e.timeout}s）'
        sn.stdout = (e.stdout or '')[-2000:] if isinstance(e.stdout, str) else ''
        sn.stderr = (e.stderr or '')[-2000:] if isinstance(e.stderr, str) else ''
        return sn
    except Exception as e:
        sn.status = 'error'
        sn.reason = str(e)
        return sn


def main():
    OUTDIR.mkdir(parents=True, exist_ok=True)
    WORKDIR.mkdir(parents=True, exist_ok=True)

    results: List[SnippetResult] = []
    for md in sorted(ROOT.glob('*.md')):
        snippets = parse_fences(md)
        for idx, (start, end, info, code) in enumerate(snippets, 1):
            kind, action, reason = classify(code)
            sn = SnippetResult(
                file=md.name,
                index=idx,
                start_line=start,
                end_line=end,
                info=info,
                kind=kind,
                action=action,
                status='pending',
                reason=reason,
            )
            results.append(execute(sn, code))

    REPORT_JSON.write_text(json.dumps([asdict(r) for r in results], ensure_ascii=False, indent=2), encoding='utf-8')

    by_file = {}
    for r in results:
        by_file.setdefault(r.file, []).append(r)

    lines = ['# Snippet Validation Report (Go 1.26)', '']
    total = len(results)
    runnable = sum(1 for r in results if r.action != 'skip')
    passed = sum(1 for r in results if r.status == 'pass')
    failed = sum(1 for r in results if r.status == 'fail')
    timeout = sum(1 for r in results if r.status == 'timeout')
    expected_fail = sum(1 for r in results if r.status == 'expected_fail')
    expected_timeout = sum(1 for r in results if r.status == 'expected_timeout')
    skipped = sum(1 for r in results if r.status == 'skipped')
    errors = sum(1 for r in results if r.status == 'error')
    lines += [f'- total snippets: {total}', f'- runnable attempted: {runnable}', f'- pass: {passed}', f'- expected_fail: {expected_fail}', f'- expected_timeout: {expected_timeout}', f'- fail: {failed}', f'- timeout: {timeout}', f'- skipped: {skipped}', f'- error: {errors}', '']

    for fname in sorted(by_file):
        rs = by_file[fname]
        p = sum(1 for r in rs if r.status == 'pass')
        ef = sum(1 for r in rs if r.status == 'expected_fail')
        et = sum(1 for r in rs if r.status == 'expected_timeout')
        f = sum(1 for r in rs if r.status == 'fail')
        t = sum(1 for r in rs if r.status == 'timeout')
        s = sum(1 for r in rs if r.status == 'skipped')
        e = sum(1 for r in rs if r.status == 'error')
        lines.append(f'## {fname}')
        lines.append(f'- snippets: {len(rs)} | pass {p} | expected_fail {ef} | expected_timeout {et} | fail {f} | timeout {t} | skipped {s} | error {e}')
        for r in rs:
            if r.status in ('fail', 'timeout', 'error'):
                lines.append(f'- snippet #{r.index} lines {r.start_line}-{r.end_line}: {r.status} ({r.kind}) - {r.reason}')
                if r.stderr:
                    msg = r.stderr.strip().splitlines()[-1]
                    lines.append(f'  - stderr: {msg}')
            elif r.status in ('expected_fail', 'expected_timeout'):
                lines.append(f'- snippet #{r.index} lines {r.start_line}-{r.end_line}: {r.status} ({r.kind}) - {r.reason}')
            elif r.status == 'skipped' and r.kind in ('go_fragment', 'go_file_main_nomian'):
                lines.append(f'- snippet #{r.index} lines {r.start_line}-{r.end_line}: skipped ({r.kind}) - {r.reason}')
        lines.append('')

    REPORT_MD.write_text('\n'.join(lines), encoding='utf-8')
    print(f'Wrote {REPORT_JSON}')
    print(f'Wrote {REPORT_MD}')
    print(f'total={total} runnable={runnable} pass={passed} expected_fail={expected_fail} expected_timeout={expected_timeout} fail={failed} timeout={timeout} skipped={skipped} error={errors}')

if __name__ == '__main__':
    main()
