import { Divider, Grid, Typography } from "@mui/material";
import { DateTimePicker, TextField } from "mui-rff";
import { AdapterDateFns } from '@mui/x-date-pickers-pro/AdapterDateFns';
import theme from "theme";

const CreateFlightForm = () => {
    //model for reference
    // DepartureDateTime time.Time     `json:"departureDateTime" bson:"departure_date_time"`
    // ArrivalDateTime   time.Time     `json:"arrivalDateTime" bson:"arrival_date_time"`
    // Departure         string        `json:"departure" bson:"departure" `
    // Destination       string        `json:"destination" bson:"destination"`
    // Price             float32       `json:"price" bson:"price"`
    // Seats             int           `json:"seats" bson:"seats"`
    // Duration          int           `json:"duration" bson:"duration"`

    return (
        <>
            <Grid container spacing={1}>
                <Typography variant="h5" sx={{ margin: '2rem 0', color: theme.palette.primary.dark }}>
                        New Flight Information:
                </Typography>
                <Divider variant='fullWidth' sx={{ borderColor: theme.palette.primary.dark, width: '100%' }} />
                <Grid item xs={12} sm={8}>
                    <TextField
                        autoComplete='off'
                        fullWidth
                        required
                        margin="normal"
                        name="departure"
                        label="Departure Location"
                    />
                </Grid>
                <Grid item xs={12} sm={4} sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                        <DateTimePicker
                            inputVariant="outlined"
                            format="dd.MM.yyyy hh:mm a"
                            ampm={true}
                            fullWidth
                            required
                            margin="normal"
                            name="departureDateTime"
                            label="Departure Date and Time"
                            adapter={AdapterDateFns}
                        />
                </Grid>
                <Grid item xs={12} sm={8}>
                    <TextField
                        autoComplete='off'
                        fullWidth
                        required
                        margin="normal"
                        name="destination"
                        label="Destination Location"
                    />
                </Grid>
                <Grid item xs={12} sm={4} sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                        <DateTimePicker
                            inputVariant="outlined"
                            format="dd.MM.yyyy hh:mm a"
                            ampm={true}
                            fullWidth
                            required
                            margin="normal"
                            name="arrivalDateTime"
                            label="Arrival Date and Time"
                            adapter={AdapterDateFns}
                        />
                </Grid>
                <Grid item xs={12} sm={4}>
                    <TextField
                        autoComplete='off'
                        fullWidth
                        required
                        margin="normal"
                        name="price"
                        label="Ticket Price"
                    />
                </Grid>
                <Grid item xs={12} sm={4}>
                    <TextField
                        autoComplete='off'
                        fullWidth
                        required
                        margin="normal"
                        name="seats"
                        label="Number of seats"
                    />
                </Grid>
                <Grid item xs={12} sm={4}>
                    <TextField
                        autoComplete='off'
                        fullWidth
                        required
                        margin="normal"
                        name="duration"
                        label="Flight Duration"
                    />
                </Grid>
            </Grid>
        </>
    );
}

export default CreateFlightForm;