import { Stack } from "@mui/system";
import Button from '@mui/material/Button'
import LoginIcon from '@mui/icons-material/Login';
import HomeIcon from '@mui/icons-material/Home';
import { Link } from "react-router-dom";

const NavBar = () => {
    const styles = {
        container: {
            width: '100%',
            justifyContent: 'space-evenly',
            padding: '1rem 0',
        }
    }

    return (
        <Stack direction='row' sx={styles.container} bgcolor='primary.main'>
            <Link to='/' sx={{textDecoration: 'none'}}>
                <Button  sx={{textDecoration: 'none'}} variant="contained" color="secondary" startIcon={<HomeIcon />}>
                    Home
                </Button>
            </Link>
            <Link to='/login'>
                <Button variant="contained" color="secondary" startIcon={<LoginIcon />}>
                    Login
                </Button>
            </Link>
            <Link to='/register'>
                <Button variant="contained" color="secondary" startIcon={<LoginIcon />}>
                    Register
                </Button>
            </Link>
            <Link to='/flights'>
                <Button variant="contained" color="secondary" startIcon={<LoginIcon />}>
                    All flights
                </Button>
            </Link>
            <Link to='/flights/create'>
                <Button variant="contained" color="secondary" startIcon={<LoginIcon />}>
                    Create flight
                </Button>
            </Link>
        </Stack>
    );
}

export default NavBar;