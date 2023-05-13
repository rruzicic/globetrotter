import { Box, Button } from "@mui/material";
import theme from "../../theme";
import { Link } from "react-router-dom";
import { useContext } from "react";
import AuthContext from "../../config/authContext";
const NavBar = () => {
    const authCtx = useContext(AuthContext)
    return (
        <Box style={{ backgroundColor: theme.palette.primary.main, width: '100%', padding: '1rem 0', display: 'flex', justifyContent: 'space-around' }}>
            <Link to={'/'}>
                <Button variant="contained" color="secondary">
                    Home
                </Button>
            </Link>
            <Link to={'/account'}>
                <Button variant="contained" color="secondary">
                    Account
                </Button>
            </Link>
            {
                authCtx.isHost() && (
                    <Link to={'/myAccommodation'}>
                        <Button variant="contained" color="secondary">
                            Accommodations
                        </Button>
                    </Link>
                )
            }
            {
                !authCtx.isLoggedIn ? (
                    <>
                        <Link to={'/login'}>
                            <Button variant="contained" color="secondary">
                                Login
                            </Button>
                        </Link>
                        <Link to={'/register'}>
                            <Button variant="contained" color="secondary">
                                Register
                            </Button>
                        </Link>
                    </>
                ) : (
                    <Link to={'/'}>
                        <Button variant="contained" color="secondary" onClick={authCtx.logout}>
                            Logout
                        </Button>
                    </Link>
                )
            }
        </Box>
    );
}

export default NavBar;