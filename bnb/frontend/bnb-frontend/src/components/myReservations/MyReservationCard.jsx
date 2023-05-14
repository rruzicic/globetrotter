import { Stack, Typography, Button, Grid } from "@mui/material";
import theme from "../../theme";
import { axiosInstance } from "../../config/interceptor";
import CONSTANTS from "../../config/constants";
import { useEffect, useState } from "react";

const MyReservationCard = ({ objectId, start, end, guestNum, totalPrice, image, handleCancel, reservationId, status }) => {

    const [name, setName] = useState(null)

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/accommodation/${objectId}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                setName(response.data.name)
            })
    }, [])

    return (
        <Grid container direction={"row"} spacing={4} mt={4} bgcolor={theme.palette.primary.main} sx={{ padding: '0.5rem 1rem', borderRadius: '10px' }}>
            <Grid item xs={1} sx={{ paddingTop: '0', paddingLeft: '0', padding: '0.5rem' }}>
                <img src={image} alt="object" height={'60px'} width={'auto'} />
            </Grid>
            <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0', padding: '0.5rem' }}>
                <Typography variant="h6" >
                    Object: {name && <>{name}</>}
                </Typography>
            </Grid>
            <Grid item xs={3} sx={{ paddingTop: '0', paddingLeft: '0', padding: '0.5rem' }}>
                <Typography variant="h6" >
                    Dates: {new Date(start).toLocaleDateString()} - {new Date(end).toLocaleDateString()}
                </Typography>
            </Grid>
            <Grid item xs={1} sx={{ paddingTop: '0', paddingLeft: '0', padding: '0.5rem' }}>
                <Typography variant="h6" >
                    Guests: {guestNum}
                </Typography>
            </Grid>
            <Grid item xs={1} sx={{ paddingTop: '0', paddingLeft: '0', padding: '0.5rem' }}>
                <Typography variant="h6" >
                    Price: {totalPrice}
                </Typography>
            </Grid>
            <Grid item xs={2} sx={{ paddingTop: '0', paddingLeft: '0', padding: '0.5rem' }}>
                <Typography variant="h6" >
                    Approved: {status.toString()}
                </Typography>
            </Grid>
            <Grid item xs={1} sx={{ paddingTop: '0', paddingLeft: '0', padding: '0.5rem' }}>
                <Button variant="contained" onClick={() => handleCancel(reservationId)}>
                    Cancel
                </Button>
            </Grid>
        </Grid>
    );
}

export default MyReservationCard;