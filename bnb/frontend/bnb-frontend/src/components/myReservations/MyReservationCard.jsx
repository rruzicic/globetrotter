import { Stack, Typography, Button } from "@mui/material";
import theme from "../../theme";

const MyReservationCard = ({ objectId, start, end, guestNum, totalPrice, objectName, image, handleCancel, reservationId }) => {

    return (
        <Stack direction={"row"} spacing={4} mt={4} bgcolor={theme.palette.primary.main} sx={{padding: '0.5rem 1rem', borderRadius: '10px'}}>
            <img src={image} alt="object" height={'60px'} width={'auto'}/>
            <Typography variant="h6" >
                Object: {objectName}
            </Typography>
            <Typography variant="h6" >
                Dates: {start.toISOString()} {end.toISOString()}
            </Typography>
            <Typography variant="h6" >
                Guests: {guestNum}
            </Typography>
            <Typography variant="h6" >
                Price: {totalPrice}
            </Typography>
            <Button variant="contained" onClick={() => handleCancel(reservationId)}>
                Cancel
            </Button>
        </Stack>
    );
}

export default MyReservationCard;