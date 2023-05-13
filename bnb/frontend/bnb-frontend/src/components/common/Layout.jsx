import NavBar from "./NavBar";

const Layout = ({children}) => {
    return ( 
        <>
        <NavBar />
        <main style={{padding: '0 3rem'}}>
            {children}
        </main>
        </>
     );
}
 
export default Layout;