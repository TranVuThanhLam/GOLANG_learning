import Header from "../components/common/Header";
import Footer from "../components/common/Footer";
function Layout({ children }) {
  return (
    <div className="jumbotron">
      <Header />
      {children}
      <Footer />
    </div>
  );
}
export default Layout;
