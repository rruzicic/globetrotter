import { Stack } from "@mui/system";
import Button from '@mui/material/Button'
import LoginIcon from '@mui/icons-material/Login';
import HomeIcon from '@mui/icons-material/Home';
import LogoutIcon from '@mui/icons-material/Logout';
import AirplanemodeActiveIcon from '@mui/icons-material/AirplanemodeActive';
import AddIcon from '@mui/icons-material/Add';
import KeyIcon from '@mui/icons-material/Key';
import PersonAddIcon from '@mui/icons-material/PersonAdd';
import { Link } from "react-router-dom";
import { useContext } from "react";
import AuthContext from "config/authContext";

const NavBar = () => {
    const authCtx = useContext(AuthContext)
    const styles = {
        container: {
            justifyContent: 'space-between',
            padding: '1rem 4rem',
            boxShadow: 'rgba(0, 0, 0, 0.25) 0px 20px 25px -5px, rgba(0, 0, 0, 0.04) 0px 10px 10px -5px'
        }
    }

    return (
        <Stack direction='row' sx={styles.container} bgcolor='primary.main'>
            <Stack spacing={4} direction={'row'}>
                <Link to='/flights'>
                    <Button variant="contained" color="secondary" startIcon={<AirplanemodeActiveIcon />}>
                        All flights
                    </Button>
                </Link>
                {
                    authCtx.isAdmin() && (
                        <Link to='/flights/create'>
                            <Button variant="contained" color="secondary" startIcon={<AddIcon />}>
                                Create flight
                            </Button>
                        </Link>
                    )
                }
                {
                    authCtx.isUser() && (
                        <Link to='/api'>
                            <Button variant="contained" color="secondary" startIcon={<KeyIcon />}>
                                API key
                            </Button>
                        </Link>
                    )
                }
            </Stack>
            {
                authCtx.isLoggedIn && (
                    <Button variant="contained" color="secondary" startIcon={<LogoutIcon />} onClick={() => { authCtx.logout() }}>
                        Logout
                    </Button>
                )
            }
            {
                !authCtx.isLoggedIn && (
                    <Stack spacing={4} direction={'row'}>
                        <Link to='/login'>
                            <Button variant="contained" color="secondary" startIcon={<LoginIcon />}>
                                Login
                            </Button>
                        </Link>
                        <Link to='/register'>
                            <Button variant="contained" color="secondary" startIcon={<PersonAddIcon />}>
                                Register
                            </Button>
                        </Link>
                    </Stack>
                )
            }
        </Stack>
    );
}

export default NavBar;