import { Grid, Typography, Button } from "@mui/material";
import { useEffect, useState } from "react";

const AccountPage = () => {
    const [userInfo, setUserInfo] = useState()
    useEffect(() => {
        //fetch user info and set state
        setUserInfo(
            {
                firstName: 'Nikola',
                lastName: 'Grbovic',
                email: 'kwicknik1@gmail.com',
                streetName: 'Balzakova',
                streetNumber: '64',
                zipCode: '21000',
                country: 'Srbija'
            }
        )
    }, [])

    return (
        <Grid container>
            {
                userInfo && (
                    <>
                        <Grid item xs={6}>
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
                        <Grid item xs={6}>
                            <Typography variant="h6" color="initial">
                                Location info
                            </Typography>
                            <Grid container>
                                <Grid item xs={8}>
                                    Street name: {userInfo.streetName}
                                </Grid>
                                <Grid item xs={4}>
                                    Street number: {userInfo.streetNumber}
                                </Grid>
                                <Grid item xs={12}>
                                    ZIP Code: {userInfo.zipCode}
                                </Grid>
                                <Grid item xs={12}>
                                    Country: {userInfo.country}
                                </Grid>
                            </Grid>
                        </Grid>
                    </>
                )
            }

            <Grid item xs={12}>
                <Button variant="contained" color="primary">
                    Change info
                </Button>
            </Grid>
        </Grid>
    );
}

export default AccountPage;