import { Typography, Grid } from "@mui/material";
import { TextField } from "mui-rff";

const LoginForm = () => {

    const styles = {
        container: {
            width: '40%',
            height: '50%',
            margin: '5% auto',
            padding: '2rem',
            justifyContent: 'center'
        }
    }

    return (
        <Grid container sx={styles.container}>
            <Grid item>
                <Typography variant="h3">Welcome to login page!</Typography>
            </Grid>
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
// InputLabelProps={{
//     style: {
//         color: 'white'
//     }
// }
// }
// inputProps={{
//     style: {
//         borderColor: 'white',
//         color: 'white',
//     }
// }}
// FormHelperTextProps={{
//     style: {
//         backgroundColor: 'white',
//         borderRadius: '10px',
//         textAlign: 'center',
//         marginTop: '10px',
//     }
// }}

export default LoginForm;