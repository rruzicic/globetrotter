import { Divider, Grid, Typography, Button } from '@mui/material';
import { TextField } from 'mui-rff';
import theme from '../../theme'


const NewAccommodationForm = () => {
    return (
        <>
            <Grid container spacing={1}>
                <Typography variant="subtitle1" sx={{ marginTop: '2rem', color: theme.palette.primary.dark }}>
                    Object information:
                </Typography>
                <Divider variant='fullWidth' sx={{ borderColor: theme.palette.primary.dark, width: '100%' }} />
                <Grid item xs={12} sm={12}>
                    <TextField
                        autoComplete='off'
                        fullWidth
                        required
                        margin="normal"
                        name="name"
                        label="Object name"
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        fullWidth
                        required
                        margin="normal"
                        name="minGuestNumber"
                        label="Minimal number of guests"
                    />
                </Grid>
                <Grid item xs={6}>
                    <TextField
                        fullWidth
                        required
                        margin="normal"
                        name="maxGuestNumber"
                        label="Maximal number of guests"
                    />
                </Grid>
                <Grid item xs={6}>
                    <Button
                        variant="contained"
                        component="label"
                    >
                        Upload Main Photo
                        <input
                            type="file"
                            hidden
                        />
                    </Button>
                </Grid>
                <Grid item xs={6}>
                    <Button
                        variant="contained"
                        component="label"
                    >
                        Upload Other Photos
                        <input
                            type="file"
                            hidden
                        />
                    </Button>
                </Grid>


                {/* ************************************************************************************************************* */}


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

export default NewAccommodationForm;