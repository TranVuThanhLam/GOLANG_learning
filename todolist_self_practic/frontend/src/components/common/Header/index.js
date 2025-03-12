function Header() {
  return (
    <header className="fixed-top">
      <nav class="py-2 bg-dark border-bottom">
        <div class="container d-flex flex-wrap ">
          <ul class="nav me-auto">
            <li class="nav-item">
              <a
                href="#"
                class="nav-link link-body-emphasis px-2 active text-white"
                aria-current="page"
              >
                Home
              </a>
            </li>
            <li class="nav-item">
              <a href="#" class="nav-link link-body-emphasis px-2 text-white">
                Features
              </a>
            </li>
            <li class="nav-item">
              <a href="#" class="nav-link link-body-emphasis px-2 text-white">
                Pricing
              </a>
            </li>
            <li class="nav-item">
              <a href="#" class="nav-link link-body-emphasis px-2 text-white">
                FAQs
              </a>
            </li>
            <li class="nav-item">
              <a href="#" class="nav-link link-body-emphasis px-2 text-white">
                About
              </a>
            </li>
          </ul>
          <ul class="nav">
            <li class="nav-item">
              <a href="#" class="nav-link link-body-emphasis px-2 text-white">
                Login
              </a>
            </li>
            <li class="nav-item">
              <a href="#" class="nav-link link-body-emphasis px-2 text-white">
                Sign up
              </a>
            </li>
          </ul>
        </div>
      </nav>

      <header class="py-3 mb-4 border-bottom bg-body-tertiary">
        <div class="container d-flex flex-wrap justify-content-center">
          <a
            href="/"
            class="d-flex align-items-center mb-3 mb-lg-0 me-lg-auto link-body-emphasis text-decoration-none"
          >
            <svg class="bi me-2" width="40" height="32">
              <use href="#bootstrap" />
            </svg>
            <span class="fs-4">Todolist</span>
          </a>
          <form class="col-12 col-lg-auto mb-3 mb-lg-0" role="search">
            <input
              type="search"
              class="form-control"
              placeholder="Search..."
              aria-label="Search"
            />
          </form>
        </div>
      </header>
    </header>
  );
}
export default Header;
