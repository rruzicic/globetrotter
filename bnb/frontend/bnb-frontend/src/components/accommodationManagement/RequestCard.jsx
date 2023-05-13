import { Stack, Typography, Button } from "@mui/material";
import theme from "../../theme";

const RequestCard = ({ requestId, userId, startDate, endDate, guestNumber, accept, decline }) => {
    return (
        <Stack spacing={4} direction={"row"} bgcolor={theme.palette.primary.main} sx={{ padding: '0.5rem' }}>
            <Typography variant="body1">
                User: {userId}
            </Typography>
            <Typography variant="body1">
                From: {startDate.toISOString()}
            </Typography>
            <Typography variant="body1">
                To: {endDate.toISOString()}
            </Typography>
            <Typography variant="body1">
                Guest number: {guestNumber}
            </Typography>
            <Button variant="contained" color="secondary" onClick={() => accept(requestId)}>
                Accept
            </Button>
            <Button variant="contained" color="secondary" onClick={() => decline(requestId)}>
                Decline
            </Button>
        </Stack>
    );
}

export default RequestCard;