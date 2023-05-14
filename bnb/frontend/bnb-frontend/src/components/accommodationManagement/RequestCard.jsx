import { Stack, Typography, Button, Grid } from "@mui/material";
import theme from "../../theme";
import { useEffect, useState } from "react";
import axios from "axios";
import { axiosInstance } from "../../config/interceptor";
import CONSTANTS from "../../config/constants";

const RequestCard = ({ requestId, userId, startDate, endDate, guestNumber, accept, decline, isApproved }) => {
    const [email, setEmail] = useState(null)

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/user/id/${userId}`)
            .then((response) => {
                setEmail(response.data.email)
            })
    })


    return (
        <Grid container direction={"row"} bgcolor={theme.palette.primary.main} sx={{ padding: '0.5rem', borderRadius: '10px' }}>
            <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0' }}>
                <Typography variant="body1">
                    User: {email && <>{email}</>}
                </Typography>
            </Grid>
            <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0' }}>
                <Typography variant="body1">
                    From: {new Date(startDate).toLocaleDateString()}
                </Typography>
            </Grid>
            <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0' }}>
                <Typography variant="body1">
                    To: {new Date(endDate).toLocaleDateString()}
                </Typography>
            </Grid>
            <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0' }}>
                <Typography variant="body1">
                    Guest number: {guestNumber}
                </Typography>
            </Grid>
            {
                !isApproved ? (
                    <>
                        <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0' }}>

                            <Button variant="contained" color="secondary" onClick={() => accept(requestId)}>
                                Accept
                            </Button>
                        </Grid>
                        <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0' }}>

                            <Button variant="contained" color="secondary" onClick={() => decline(requestId)}>
                                Decline
                            </Button>
                        </Grid>
                    </>
                ) : <>Approved 🔥</>
            }
        </Grid>
    );
}

export default RequestCard;