import Header from "../components/common/Header";
import Footer from "../components/common/Footer";
function Layout({ children }) {
  return (
    <div className="d-flex flex-column h-100">
      <Header />
      <div className="my-5" />
      <div className="container mt-5 mb-5">{children}</div>
      <Footer />
    </div>
  );
}
export default Layout;
