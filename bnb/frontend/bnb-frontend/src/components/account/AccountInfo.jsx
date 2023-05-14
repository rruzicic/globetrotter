import { Grid, Typography, Button, Stack } from "@mui/material";
import theme from "../../theme";
import { useNavigate } from "react-router"
import { useContext } from "react";
import AuthContext from "../../config/authContext";
import { axiosInstance } from "../../config/interceptor";
import CONSTANTS from "../../config/constants";

const AccountInfo = ({ setUpdate, userInfo }) => {
    const navigate = useNavigate()
    const authCtx = useContext(AuthContext)
    const handleChangeState = () => {
        setUpdate(true)
    }
    //TODO: 1.4
    const handleDeleteAccount = () => {
        axiosInstance.get(`http://localhost:4000/user/email/${authCtx.userEmail()}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                axiosInstance.delete(`${CONSTANTS.GATEWAY}/user/delete/${response.data.id}`)
                    .catch((e) => {
                        console.error(e)
                        return
                    })
                    .then((response) => {
                        authCtx.logout()
                    })
            })

        navigate('/')
    }

    return (
        <Grid container spacing={2} justifyContent={"center"} mt={4}>
            <Grid item xs={12} sx={{ display: 'grid', placeItems: 'center' }}>
                <img src="/account.png" alt="account" height={'300px'} width={'auto'} />
            </Grid>
            {
                userInfo && (
                    <>
                        <Grid item xs={5} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem' }}>
                            <Typography variant="h6" color="initial">
                                User info
                            </Typography>
                            <Grid container>
                                <Grid item xs={12}>
                                    First name: {userInfo.firstName}
                                </Grid>
                                <Grid item xs={12}>
                                    Last name: {userInfo.lastName}
                                </Grid>
                                <Grid item xs={12}>
                                    Email: {userInfo.email}
                                </Grid>
                                <Grid item xs={12}>
                                    Role: {userInfo.role}
                                </Grid>
                            </Grid>
                        </Grid>
                        <Grid item xs={5} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem' }}>
                            <Typography variant="h6" color="initial">
                                Location info
                            </Typography>
                            <Grid container>
                                <Grid item xs={12}>
                                    Street: {userInfo.street} {userInfo.streetNum}
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

            <Grid item xs={6} sx={{ display: 'flex', justifyContent: 'center', width: '100%' }}>
                <Stack direction={"row"} spacing={4}>
                    <Button variant="contained" color="primary" onClick={handleChangeState}>
                        Change info
                    </Button>
                    <Button variant="contained" color="primary" onClick={handleDeleteAccount}>
                        Delete Account
                    </Button>
                </Stack>
            </Grid>
        </Grid>
    );
}

export default AccountInfo;