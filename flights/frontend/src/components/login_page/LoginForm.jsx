import { Grid } from "@mui/material";
import { TextField } from "mui-rff";

const LoginForm = () => {

    const styles = {
        container: {
            width: '80%',
            padding: '2rem',
            justifyContent: 'center'
        }
    }

    return (
        <Grid container sx={styles.container}>
            <Grid item xs={12}>
                <TextField
                    autoComplete='off'
                    fullWidth
                    required
                    margin="normal"
                    name="email"
                    label="E-mail"

                />
            </Grid>
            <Grid item xs={12}>
                <TextField
                    autoComplete='off'
                    fullWidth
                    required
                    margin="normal"
                    name="password"
                    label="Password"
                />
            </Grid>
        </Grid>
    );
}

export default LoginForm;