import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { axiosInstance } from "../config/interceptor";
import CONSTANTS from "../config/constants";
import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Paper, Stack, Table, TableBody, TableCell, TableHead, TableRow, TableSortLabel, TextField, useTheme, Typography } from "@mui/material";
import { toast } from "react-toastify";

const ReservationInfoPage = () => {
    const theme = useTheme()
    const { id } = useParams();
    const [reservation, setReservation] = useState()
    const [accommodation, setAccommodation] = useState()
    const [open, setOpen] = useState(false);
    const [startLocation, setStartLocation] = useState('');
    const [endLocation, setEndLocation] = useState('');
    const [flights, setFlights] = useState()

    const handleStartLocationChange = (event) => {
        setStartLocation(event.target.value);
    };

    const handleEndLocationChange = (event) => {
        setEndLocation(event.target.value);
    };

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/reservation/${id}`)
            .catch((err) => {
                console.error(err)
            })
            .then((res) => {
                setReservation(res.data)
                axiosInstance.get(`${CONSTANTS.GATEWAY}/accommodation/${res.data.accommodationId}`)
                    .catch((err) => {
                        console.error(err)
                    })
                    .then((res) => {
                        setAccommodation(res.data)
                    })
            })
    }, [])

    const handleRecommend = () => {
        const dto = {
            reservationStartDate: reservation.dateInterval.start,
            reservationEndDate: reservation.dateInterval.end,
            departureLocationToReservation: startLocation,
            arrivalLocationAtReservation: accommodation.location.city,
            departureLocationFromReservation: accommodation.location.city,
            arrivalLocationAtHome: endLocation,
            people: reservation.numOfGuests
        }
        axiosInstance.post(`${CONSTANTS.GATEWAY}/recommendation/flights`, dto)
            .catch((err) => {
                console.error(err)
            })
            .then((res) => {
                setFlights(res.data)
            })
        handleClose()
    }

    const styles = {
        row: {
            cursor: 'pointer',
            '&:hover': {
                backgroundColor: theme.palette.primary.light,
            }
        }
    }

    const [apiKey, setApiKey] = useState()
    const [openApi, setOpenApi] = useState(false)
    const [selectedFlightId, setSelectedFlightId] = useState()
    const handleApiKeyChange = (e) => {
        setApiKey(e.target.value)
    }
    const handleOpenApi = (id) => {
        setSelectedFlightId(id)
        setOpenApi((prev) => !prev)
    }
    const handleBook = (id) => {
        const dto = {
            apiKey: apiKey,
            flightId: selectedFlightId,
            numOfTicketsOptional: [parseInt(reservation.numOfGuests)]
        }
        axiosInstance.post(`${CONSTANTS.GATEWAY}/recommendation/flights/buy-ticket`, dto)
            .catch((err) => {
                console.error(err)
                return
            })
            .then((res) => {
                toast("Bought ticket!")
            })

        handleOpenApi()
    }

    return (
        <>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Tell us more about your journey</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        To recommend flights that align with your reservation we need a bit more information.
                    </DialogContentText>
                    <TextField
                        id="startLocation"
                        label="Where are you starting your journey?"
                        fullWidth
                        variant="standard"
                        onChange={handleStartLocationChange}
                    />
                    <TextField
                        id="endLocation"
                        label="Where are you ending your journey?"
                        fullWidth
                        variant="standard"
                        onChange={handleEndLocationChange}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button onClick={handleRecommend}>Recommend</Button>
                </DialogActions>
            </Dialog>
            <Dialog open={openApi} onClose={handleOpenApi}>
                <DialogTitle>IDENTIFY YOURSELF</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        To book this flight we need a reference to your globetrotter account
                    </DialogContentText>
                    <TextField
                        id="startLocation"
                        label="Enter you globetrotter API key (found on your profile)"
                        fullWidth
                        variant="standard"
                        onChange={handleApiKeyChange}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleOpenApi}>Cancel</Button>
                    <Button onClick={handleBook}>Book</Button>
                </DialogActions>
            </Dialog>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Tell us more about your journey</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        To recommend flights that align with your reservation we need a bit more information.
                    </DialogContentText>
                    <TextField
                        id="startLocation"
                        label="Where are you starting your journey?"
                        fullWidth
                        variant="standard"
                        onChange={handleStartLocationChange}
                    />
                    <TextField
                        id="endLocation"
                        label="Where are you ending your journey?"
                        fullWidth
                        variant="standard"
                        onChange={handleEndLocationChange}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button onClick={handleRecommend}>Recommend</Button>
                </DialogActions>
            </Dialog>
            <Stack mt={4}>
                <Typography variant="h2" alignSelf={"center"} sx={{ margin: '0 0 2rem 0' }}>
                    Reservation info
                </Typography>
                {
                    accommodation && reservation && (
                        <>
                            <Stack direction={"row"} spacing={2} sx={{ color: theme.palette.secondary.main, backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem', textAlign: 'center', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
                                <Stack sx={{ width: '50%' }} justifyContent={"center"} alignItems={"center"}>
                                    <img src={accommodation.photos[0]} alt="preview" style={{ width: '400px', height: '250px' }} />
                                    <Typography variant="h4" >
                                        {accommodation.name}
                                    </Typography>
                                    <Typography variant="h6" >
                                        Country: {accommodation.location.country}
                                    </Typography>
                                    <Typography variant="h6" >
                                        City: {accommodation.location.city}
                                    </Typography>
                                    <Typography variant="h6" >
                                        Street: {accommodation.location.street} {accommodation.location.streetNum}
                                    </Typography>
                                </Stack>
                                <Stack spacing={4} justifyContent={"center"} alignItems={"center"} sx={{ width: '50%', backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem', textAlign: 'center', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
                                    <Typography variant="h6" >
                                        Starting date: {new Date(reservation.dateInterval.start).toLocaleDateString()}
                                    </Typography>
                                    <Typography variant="h6" >
                                        Ending date: {new Date(reservation.dateInterval.end).toLocaleDateString()}
                                    </Typography>
                                    <Typography variant="h6" >
                                        This reservation is {reservation.isApproved ? "approved" : "not approved"}
                                    </Typography>
                                    <Typography variant="h6" >
                                        Number of passengers: {reservation.numOfGuests}
                                    </Typography>
                                </Stack>
                            </Stack>
                            <Button variant="contained" color="primary" onClick={handleClickOpen} sx={{ marginTop: '2rem' }}>
                                Recommend flights that match this reservation
                            </Button>
                        </>
                    )
                }
                {
                    flights && (
                        <Paper elevation={4} sx={{ width: '90%', margin: '1rem auto' }}>
                            <Table>
                                <TableHead>
                                    <TableRow sx={{ backgroundColor: theme.palette.primary.main, color: theme.palette.secondary.main }}>
                                        <TableCell>
                                            <TableSortLabel>
                                                Departure Time
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                Departure Location
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                Arrival Time
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                Destination
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                Price(eur)
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                Remaining tickets
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                Duration(min)
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                Passenger Number
                                            </TableSortLabel>
                                        </TableCell>
                                        <TableCell>
                                            <TableSortLabel>
                                                User action
                                            </TableSortLabel>
                                        </TableCell>
                                    </TableRow>
                                </TableHead>
                                <TableBody>
                                    {flights.map((flight) => (
                                        <TableRow key={flight.id} sx={styles.row}>
                                            <TableCell>{new Date(flight.departureDateTime).toLocaleString()}</TableCell>
                                            <TableCell>{flight.departure}</TableCell>
                                            <TableCell>{new Date(flight.arrivalDateTime).toLocaleString()}</TableCell>
                                            <TableCell>{flight.destination}</TableCell>
                                            <TableCell>{flight.price}</TableCell>
                                            <TableCell>{flight.seats}</TableCell>
                                            <TableCell>{flight.duration}</TableCell>
                                            <TableCell>{reservation.numOfGuests}</TableCell>
                                            <TableCell>
                                                <Button variant="contained" color="primary" onClick={() => handleOpenApi(flight.id)}>
                                                    Book!
                                                </Button>
                                            </TableCell>
                                        </TableRow>
                                    ))}
                                </TableBody>
                            </Table>
                        </Paper>
                    )
                }
            </Stack >
        </>
    );
}

export default ReservationInfoPage;