import { Grid, Typography, Button } from "@mui/material";
import theme from "../../theme";

const AccountInfo = ({setUpdate, userInfo}) => {

    const handleChangeState = () => {
        setUpdate(true)
    }

    return (
        <Grid container spacing={2} justifyContent={"center"} mt={4}>
            {
                userInfo && (
                    <>
                        <Grid item xs={5} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem' }}>
                            <Typography variant="h6" color="initial">
                                User info
                            </Typography>
                            <Grid container>
                                <Grid item xs={6}>
                                    First name: {userInfo.firstName}
                                </Grid>
                                <Grid item xs={6}>
                                    Last name: {userInfo.lastName}
                                </Grid>
                                <Grid item xs={12}>
                                    Email: {userInfo.email}
                                </Grid>
                            </Grid>
                        </Grid>
                        <Grid item xs={5} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem' }}>
                            <Typography variant="h6" color="initial">
                                Location info
                            </Typography>
                            <Grid container>
                                <Grid item xs={8}>
                                    Street name: {userInfo.street}
                                </Grid>
                                <Grid item xs={4}>
                                    Street number: {userInfo.streetNum}
                                </Grid>
                                <Grid item xs={12}>
                                    ZIP Code: {userInfo.zip}
                                </Grid>
                                <Grid item xs={12}>
                                    Country: {userInfo.country}
                                </Grid>
                            </Grid>
                        </Grid>
                    </>
                )
            }

            <Grid item xs={12} sx={{display: 'flex', justifyContent: 'center', width: '100%'}}>
                <Button variant="contained" color="primary" onClick={handleChangeState}>
                    Change info
                </Button>
            </Grid>
        </Grid>
    );
}

export default AccountInfo;