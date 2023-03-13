import { Divider, Grid, Typography } from '@mui/material';
import { TextField } from 'mui-rff';
import theme from 'theme';


const RegistrationForm = () => {
    // FirstName string`json:"firstName" bson:"first_name" binding:"required"`
	// LastName  string`json:"lastName" bson:"last_name" binding:"required"`
	// EMail     string`json:"email" bson:"email" binding:"required,email"`
	// Password  string`json:"password" bson:"password"`
    // Country   string`json:"country" bson:"country"`
	// Street    string`json:"street" bson:"street"`
	// StreetNum string`json:"streetNum" bson:"street_num"`
	// ZIPCode   int`json:"zip" bson:"zip"`

    return (
        <>
        <Grid container spacing={1}>
            <Typography variant="subtitle1" sx={{ marginTop: '3rem', color: theme.palette.primary.dark }}>
                Personal information:
            </Typography>
            <Divider variant='fullWidth' sx={{ borderColor: theme.palette.primary.dark, width: '100%' }} />
            <Grid item xs={12} sm={6}>
                <TextField
                    autoComplete='off'
                    fullWidth
                    required
                    margin="normal"
                    name="firstName"
                    label="First Name"
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                    fullWidth
                    required
                    margin="normal"
                    name="lastName"
                    label="Last Name"
                />
            </Grid>
            <Grid item xs={12}>
                <TextField
                    fullWidth
                    required
                    margin="normal"
                    name="email"
                    type='email'
                    label="Email Address"
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                    fullWidth
                    required
                    margin="normal"
                    name="password"
                    label="Password"
                    type='password'
                />
            </Grid>
            <Grid item xs={12} sm={6}>
                <TextField
                    fullWidth
                    required
                    margin="normal"
                    name="confirmPassword"
                    label="Confirm password"
                    type='password'
                />
            </Grid>
            <Typography variant="subtitle1" color="tertiary" sx={{ marginTop: '3rem', color: theme.palette.primary.dark }}>
                Address information:
            </Typography>
            <Divider variant='fullWidth' sx={{ borderColor: theme.palette.primary.dark, width: '100%' }} />
            <Grid item xs={12} sm={10}>
                <TextField
                    autoComplete='off'
                    fullWidth
                    required
                    margin="normal"
                    name="street"
                    label="Street Name"
                />
            </Grid>
            <Grid item xs={12} sm={2}>
                <TextField
                    fullWidth
                    required
                    margin="normal"
                    name="streetNum"
                    label="Number"
                />
            </Grid>
            <Grid item xs={12} sm={4}>
                <TextField
                    fullWidth
                    required
                    margin="normal"
                    name="zip"
                    label="ZIP Code"
                />
            </Grid>
            <Grid item xs={12} sm={8}>
                <TextField
                    fullWidth
                    required
                    margin="normal"
                    name="country"
                    label="Country"
                />
            </Grid>
        </Grid>
    </>
    );
}

export default RegistrationForm;