import { Box, Button, useTheme } from "@mui/material";
import { Link } from "react-router-dom";
import { useContext } from "react";
import AuthContext from "../../config/authContext";
const NavBar = () => {
    const authCtx = useContext(AuthContext)
    const theme = useTheme()
    return (
        <Box style={{ backgroundColor: theme.palette.primary.main, width: '100%', padding: '1rem 0', display: 'flex', justifyContent: 'space-around', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
            <Link to={'/'}>
                <Button variant="contained" color="secondary">
                    Home
                </Button>
            </Link>
            {
                authCtx.isLoggedIn && (
                    <>
                    <Link to={'/account'}>
                        <Button variant="contained" color="secondary">
                            Account
                        </Button>
                    </Link>
                    <Link to={'/notifications'}>
                        <Button variant="contained" sx={{backgroundColor: authCtx.countNewNotifications() === 0 ? theme.palette.secondary.main : "red", color: authCtx.countNewNotifications() === 0 ? "black" : "white"}}>
                            Notifications ({authCtx.countNewNotifications()})
                        </Button>
                    </Link>
                    </>
                )
            }
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
                authCtx.isUser() && (
                    <Link to={'/myReservations'}>
                        <Button variant="contained" color="secondary">
                            Reservations
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