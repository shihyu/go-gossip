// Populate the sidebar
//
// This is a script, and not included directly in the page, to control the total size of the book.
// The TOC contains an entry for each page, so if each page includes a copy of the TOC,
// the total size of the page becomes O(n**2).
class MDBookSidebarScrollbox extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.innerHTML = '<ol class="chapter"><li class="chapter-item expanded "><a href="index.html"><strong aria-hidden="true">1.</strong> 首頁</a></li><li class="chapter-item expanded "><a href="HelloWorld.html"><strong aria-hidden="true">2.</strong> 來個 Hello, World</a></li><li class="chapter-item expanded "><a href="Package.html"><strong aria-hidden="true">3.</strong> Go 套件管理</a></li><li class="chapter-item expanded "><a href="gofmt.html"><strong aria-hidden="true">4.</strong> gofmt 格式化原始碼</a></li><li class="chapter-item expanded "><a href="godoc.html"><strong aria-hidden="true">5.</strong> go doc 文件即註解</a></li><li class="chapter-item expanded "><a href="Testing.html"><strong aria-hidden="true">6.</strong> Go 測試套件</a></li><li class="chapter-item expanded "><a href="PreDeclaredType.html"><strong aria-hidden="true">7.</strong> 認識預定義型態</a></li><li class="chapter-item expanded "><a href="VariableConstantDeclaration.html"><strong aria-hidden="true">8.</strong> 變數宣告、常數宣告</a></li><li class="chapter-item expanded "><a href="String.html"><strong aria-hidden="true">9.</strong> 位元組構成的字串</a></li><li class="chapter-item expanded "><a href="Array.html"><strong aria-hidden="true">10.</strong> 身為複合值的陣列</a></li><li class="chapter-item expanded "><a href="Slice.html"><strong aria-hidden="true">11.</strong> 底層為陣列的 slice</a></li><li class="chapter-item expanded "><a href="Map.html"><strong aria-hidden="true">12.</strong> 成對鍵值的 map</a></li><li class="chapter-item expanded "><a href="Operator.html"><strong aria-hidden="true">13.</strong> 運算子</a></li><li class="chapter-item expanded "><a href="IfElseSwitch.html"><strong aria-hidden="true">14.</strong> if ... else、switch 條件式</a></li><li class="chapter-item expanded "><a href="For.html"><strong aria-hidden="true">15.</strong> for 迴圈</a></li><li class="chapter-item expanded "><a href="BreakContinueGoto.html"><strong aria-hidden="true">16.</strong> break、continue、goto</a></li><li class="chapter-item expanded "><a href="Function.html"><strong aria-hidden="true">17.</strong> 函式入門</a></li><li class="chapter-item expanded "><a href="FirstClassFunction.html"><strong aria-hidden="true">18.</strong> 一級函式</a></li><li class="chapter-item expanded "><a href="Closure.html"><strong aria-hidden="true">19.</strong> 匿名函式與閉包</a></li><li class="chapter-item expanded "><a href="DeferPanicRecover.html"><strong aria-hidden="true">20.</strong> defer、panic、recover</a></li><li class="chapter-item expanded "><a href="Struct.html"><strong aria-hidden="true">21.</strong> 結構入門</a></li><li class="chapter-item expanded "><a href="Method.html"><strong aria-hidden="true">22.</strong> 結構與方法</a></li><li class="chapter-item expanded "><a href="StructComposition.html"><strong aria-hidden="true">23.</strong> 結構組合</a></li><li class="chapter-item expanded "><a href="Interface.html"><strong aria-hidden="true">24.</strong> 介面入門</a></li><li class="chapter-item expanded "><a href="TypeAssertion.html"><strong aria-hidden="true">25.</strong> 型態斷言</a></li><li class="chapter-item expanded "><a href="InterfaceComposition.html"><strong aria-hidden="true">26.</strong> 介面組合</a></li><li class="chapter-item expanded "><a href="StdOutInErr.html"><strong aria-hidden="true">27.</strong> 從標準輸入、輸出認識 io</a></li><li class="chapter-item expanded "><a href="ReaderWriter.html"><strong aria-hidden="true">28.</strong> io.Reader、io.Writer</a></li><li class="chapter-item expanded "><a href="bufio.html"><strong aria-hidden="true">29.</strong> bufio 套件</a></li><li class="chapter-item expanded "><a href="File.html"><strong aria-hidden="true">30.</strong> 檔案操作</a></li><li class="chapter-item expanded "><a href="ErrNil.html"><strong aria-hidden="true">31.</strong> err 是否 nil？</a></li><li class="chapter-item expanded "><a href="ErrorComparison.html"><strong aria-hidden="true">32.</strong> 錯誤的比對</a></li><li class="chapter-item expanded "><a href="errors.html"><strong aria-hidden="true">33.</strong> errors 套件</a></li><li class="chapter-item expanded "><a href="Sort.html"><strong aria-hidden="true">34.</strong> sort 套件</a></li><li class="chapter-item expanded "><a href="List.html"><strong aria-hidden="true">35.</strong> list 套件</a></li><li class="chapter-item expanded "><a href="Heap.html"><strong aria-hidden="true">36.</strong> heap 套件</a></li><li class="chapter-item expanded "><a href="Ring.html"><strong aria-hidden="true">37.</strong> ring 套件</a></li><li class="chapter-item expanded "><a href="StrconvStrings.html"><strong aria-hidden="true">38.</strong> strconv、strings 套件</a></li><li class="chapter-item expanded "><a href="Bytes.html"><strong aria-hidden="true">39.</strong> bytes 套件</a></li><li class="chapter-item expanded "><a href="Unicode.html"><strong aria-hidden="true">40.</strong> unicode 套件</a></li><li class="chapter-item expanded "><a href="XText.html"><strong aria-hidden="true">41.</strong> 編碼轉換</a></li><li class="chapter-item expanded "><a href="Reflect.html"><strong aria-hidden="true">42.</strong> 反射入門</a></li><li class="chapter-item expanded "><a href="FieldTag.html"><strong aria-hidden="true">43.</strong> 結構欄位標籤</a></li><li class="chapter-item expanded "><a href="Goroutine.html"><strong aria-hidden="true">44.</strong> Goroutine</a></li><li class="chapter-item expanded "><a href="Channel.html"><strong aria-hidden="true">45.</strong> Channel</a></li><li class="chapter-item expanded "><a href="Vendor.html"><strong aria-hidden="true">46.</strong> vendor</a></li><li class="chapter-item expanded "><a href="Module.html"><strong aria-hidden="true">47.</strong> 模組入門</a></li><li class="chapter-item expanded "><a href="WebAssembly.html"><strong aria-hidden="true">48.</strong> 哈囉！WebAssembly！</a></li><li class="chapter-item expanded "><a href="JavaScript.html"><strong aria-hidden="true">49.</strong> Go 呼叫 JavaScript</a></li><li class="chapter-item expanded "><a href="Callback.html"><strong aria-hidden="true">50.</strong> JavaScript 回呼 Go</a></li></ol>';
        // Set the current, active page, and reveal it if it's hidden
        let current_page = document.location.href.toString().split("#")[0].split("?")[0];
        if (current_page.endsWith("/")) {
            current_page += "index.html";
        }
        var links = Array.prototype.slice.call(this.querySelectorAll("a"));
        var l = links.length;
        for (var i = 0; i < l; ++i) {
            var link = links[i];
            var href = link.getAttribute("href");
            if (href && !href.startsWith("#") && !/^(?:[a-z+]+:)?\/\//.test(href)) {
                link.href = path_to_root + href;
            }
            // The "index" page is supposed to alias the first chapter in the book.
            if (link.href === current_page || (i === 0 && path_to_root === "" && current_page.endsWith("/index.html"))) {
                link.classList.add("active");
                var parent = link.parentElement;
                if (parent && parent.classList.contains("chapter-item")) {
                    parent.classList.add("expanded");
                }
                while (parent) {
                    if (parent.tagName === "LI" && parent.previousElementSibling) {
                        if (parent.previousElementSibling.classList.contains("chapter-item")) {
                            parent.previousElementSibling.classList.add("expanded");
                        }
                    }
                    parent = parent.parentElement;
                }
            }
        }
        // Track and set sidebar scroll position
        this.addEventListener('click', function(e) {
            if (e.target.tagName === 'A') {
                sessionStorage.setItem('sidebar-scroll', this.scrollTop);
            }
        }, { passive: true });
        var sidebarScrollTop = sessionStorage.getItem('sidebar-scroll');
        sessionStorage.removeItem('sidebar-scroll');
        if (sidebarScrollTop) {
            // preserve sidebar scroll position when navigating via links within sidebar
            this.scrollTop = sidebarScrollTop;
        } else {
            // scroll sidebar to current active section when navigating via "next/previous chapter" buttons
            var activeSection = document.querySelector('#sidebar .active');
            if (activeSection) {
                activeSection.scrollIntoView({ block: 'center' });
            }
        }
        // Toggle buttons
        var sidebarAnchorToggles = document.querySelectorAll('#sidebar a.toggle');
        function toggleSection(ev) {
            ev.currentTarget.parentElement.classList.toggle('expanded');
        }
        Array.from(sidebarAnchorToggles).forEach(function (el) {
            el.addEventListener('click', toggleSection);
        });
    }
}
window.customElements.define("mdbook-sidebar-scrollbox", MDBookSidebarScrollbox);
