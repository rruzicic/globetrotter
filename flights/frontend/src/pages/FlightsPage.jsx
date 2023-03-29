import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Paper, Table, TableBody, TableCell, TableHead, TablePagination, TableRow, TableSortLabel, TextField, Typography } from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers";
import { useCallback, useContext, useEffect, useState } from "react";
import axios from "axios";
import theme from "theme";
import { useDebounce } from "use-debounce";
import formatDate from "util";
import AuthContext from "config/authContext";
import { toast } from "react-toastify";
import { axiosInstance } from "config/interceptor";
import { useNavigate } from "react-router";

const FlightsPage = () => {

    //model for reference
    // Id                bson.ObjectId `json:"id" bson:"_id,omitempty"`
    // DepartureDateTime time.Time     `json:"departureDateTime" bson:"departure_date_time"`
    // ArrivalDateTime   time.Time     `json:"arrivalDateTime" bson:"arrival_date_time"`
    // Departure         string        `json:"departure" bson:"departure" `
    // Destination       string        `json:"destination" bson:"destination"`
    // Price             float32       `json:"price" bson:"price"`
    // Seats             int           `json:"seats" bson:"seats"`
    // Duration          int           `json:"duration" bson:"duration"`

    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);
    const [orderBy, setOrderBy] = useState('');
    const [order, setOrder] = useState('asc');
    const [flights, setFlights] = useState([])
    //tracks values
    const [departureSP, setDepartureSP] = useState("")
    const [departureDateSP, setDepartureDateSP] = useState(null)
    const [destinationSP, setDestinationSP] = useState("")
    const [arrivalDateSP, setArrivalDateSP] = useState(null)
    const [passengerNumSP, setPassengerNumSP] = useState("")
    //prevents sending api requests for every char typed, sends it 500ms after last char is typed
    const [debounceDepartureSP] = useDebounce(departureSP, 500)
    const [debounceDepartureDateSP] = useDebounce(
        (departureDateSP == null) ? null
            :
            departureDateSP.toISOString().split('T')[0]
        , 500)
    const [debounceDestinationSP] = useDebounce(destinationSP, 500)
    const [debounceArrivalDateSP] = useDebounce(
        (arrivalDateSP == null) ? null
            :
            arrivalDateSP.toISOString().split('T')[0]
        , 500)
    const [debouncePassengerNumSP] = useDebounce(passengerNumSP, 500)

    //had to do it like this because of non matching versions of npm packages
    const changeDeparture = (e) => {
        setDepartureSP(e.target.value);
    }
    const changeDepartureDate = (e) => {
        console.log(e);
        setDepartureDateSP(e)
    }
    const changeDestination = (e) => {
        setDestinationSP(e.target.value);
    }
    const changeArrivalDate = (e) => {
        console.log(e);
        setArrivalDateSP(e);
    }
    const changePassengerNumber = (e) => {
        setPassengerNumSP(e.target.value);
    }

    const sortedFlights = flights.sort((a, b) => {
        if (orderBy === 'departureDateTime' || orderBy === 'arrivalDateTime') {
            return order === 'asc'
                ? new Date(a[orderBy]).getTime() - new Date(b[orderBy]).getTime()
                : new Date(b[orderBy]).getTime() - new Date(a[orderBy]).getTime();
        } else if (orderBy === 'departure' || orderBy === 'destination') {
            return order === 'asc'
                ? a[orderBy].localeCompare(b[orderBy])
                : b[orderBy].localeCompare(a[orderBy]);
        } else {
            return order === 'asc' ? a[orderBy] - b[orderBy] : b[orderBy] - a[orderBy];
        }
    });

    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    const handleSort = (field) => {
        const isAsc = (orderBy === field && order === 'asc');
        setOrderBy(field);
        setOrder(isAsc ? 'desc' : 'asc');
    };

    useEffect(() => {
        axiosInstance.get('/flights')
            .catch((error) => {
                toast('Unable to fetch flights, try again in a few seconds 😢')
                return
            })
            .then((response) => {
                setFlights(response.data.data)
            })
    }, [])

    const styles = {
        row: {
            cursor: 'pointer',
            '&:hover': {
                backgroundColor: theme.palette.primary.light,
            }
        }
    }
    const navigate = useNavigate()

    const deleteFlight = (id, event) => {
        event.stopPropagation()
        console.log('Should delete flight with id: ' + id);
        axiosInstance.delete('/flights/delete', id)
        .catch((e)=> {
            toast('Could not delete flight! 😢')
        })
        .then((response) => {
            if(response !== undefined) {
                toast('Flight successfully deleted!')
                navigate('/flights')
            }
        })
    }

    const resetSearch = () => {
        setDepartureSP("")
        setDepartureDateSP(null)
        setArrivalDateSP(null)
        setDestinationSP("")
        setPassengerNumSP("")
    }

    const [openModal, setOpenModal] = useState(false)
    const [selectedFlight, setSelectedFlight] = useState(false)
    const [ticketNumber, setTicketNumber] = useState(false)
    const authCtx = useContext(AuthContext)

    const buyTicket = (() => {
        axiosInstance.post('/flights/buy-ticket', {
            flightId: selectedFlight,
            userEmail: authCtx.userEmail(),
            numOfTicketsOptional: [parseInt(ticketNumber)]
        })
            .catch((err) => {
                toast('Not enough tickets left 😢')
                return
            })
            .then((response) => {
                if (response !== undefined) {
                    toast('Successfully bought tickets! 😊')
                }
            })
        handleClose()
    })

    const handleOpen = (flightId) => {
        setSelectedFlight(flightId)
        setOpenModal(true)
    }
    const handleClose = () => {
        setSelectedFlight(null)
        setTicketNumber(0)
        setOpenModal(false)
    }

    const changeTicketNumber = (e) => {
        setTicketNumber(e.target.value)
    }

    useEffect(() => {
        axiosInstance.get('/flights/search', {
            params: {
                departure: debounceDepartureSP,
                departureDateTime: debounceDepartureDateSP,
                destination: debounceDestinationSP,
                arrivalDateTime: debounceArrivalDateSP,
                passengerNumber: debouncePassengerNumSP
            }
        }).catch((err) => {
            toast('No flights match that criteria 😢')
            return
        }).then((response) => {
            if (response.data.data == null) {
                setFlights([])
            } else {
                setFlights(response.data.data)
            }
        })
    }, [debounceDepartureSP, debounceDepartureDateSP, debounceDestinationSP, debounceArrivalDateSP, debouncePassengerNumSP])

    return (
        <>
            <Dialog
                open={openModal}
                keepMounted
                onClose={handleClose}
            >
                <DialogTitle>{"How many tickets would you like to buy?"}</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        <TextField onChange={changeTicketNumber} label="Number of Tickets">

                        </TextField>
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button onClick={() => { buyTicket() }}>Buy</Button>
                </DialogActions>
            </Dialog>
            <Typography variant="h2" align="center" sx={{ margin: '1rem 0' }}>List of all flights </Typography>
            <Paper elevation={4} sx={{ width: '60%', margin: '1rem auto', display: 'flex', justifyContent: 'space-around', padding: '0.5rem' }} >
                <TextField onChange={changeDeparture} label='Departure' value={departureSP} />
                <DatePicker
                    disablePast
                    value={departureDateSP}
                    renderInput={(props) => <TextField {...props} />}
                    onChange={changeDepartureDate} label='Departure Date' />
                <TextField onChange={changeDestination} label='Destination' value={destinationSP} />
                <DatePicker
                    disablePast
                    value={arrivalDateSP}
                    renderInput={(props) => <TextField {...props} />}
                    onChange={changeArrivalDate} label='Arrival Date' />
                <TextField onChange={changePassengerNumber} label='Passenger Number' value={passengerNumSP} />
                <Button variant="outlined" color="primary" onClick={resetSearch}>
                    Reset Search
                </Button>
            </Paper>
            <Paper elevation={4} sx={{ width: '90%', margin: '1rem auto' }}>
                <Table>
                    <TableHead>
                        <TableRow sx={{ backgroundColor: theme.palette.primary.main, color: theme.palette.secondary.main }}>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'departureDateTime'}
                                    direction={orderBy === 'departureDateTime' ? order : 'asc'}
                                    onClick={() => handleSort('departureDateTime')}
                                >
                                    Departure Time
                                </TableSortLabel>
                            </TableCell>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'departure'}
                                    direction={orderBy === 'departure' ? order : 'asc'}
                                    onClick={() => handleSort('departure')}
                                >
                                    Departure Location
                                </TableSortLabel>
                            </TableCell>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'arrivalDateTime'}
                                    direction={orderBy === 'arrivalDateTime' ? order : 'asc'}
                                    onClick={() => handleSort('arrivalDateTime')}
                                >
                                    Arrival Time
                                </TableSortLabel>
                            </TableCell>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'destination'}
                                    direction={orderBy === 'destination' ? order : 'asc'}
                                    onClick={() => handleSort('destination')}
                                >
                                    Destination
                                </TableSortLabel>
                            </TableCell>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'price'}
                                    direction={orderBy === 'price' ? order : 'asc'}
                                    onClick={() => handleSort('price')}
                                >
                                    Price(eur)
                                </TableSortLabel>
                            </TableCell>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'seats'}
                                    direction={orderBy === 'seats' ? order : 'asc'}
                                    onClick={() => handleSort('seats')}
                                >
                                    Seats
                                </TableSortLabel>
                            </TableCell>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'duration'}
                                    direction={orderBy === 'duration' ? order : 'asc'}
                                    onClick={() => handleSort('duration')}
                                >
                                    Duration(hr)
                                </TableSortLabel>
                            </TableCell>
                            {
                                authCtx.isAdmin() &&
                                <TableCell>
                                    Admin
                                </TableCell>
                            }
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {sortedFlights.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((flight) => (
                            <TableRow key={flight.id} sx={styles.row} onClick={() => { handleOpen(flight.id) }}>
                                <TableCell>{formatDate(flight.departureDateTime)}</TableCell>
                                <TableCell>{flight.departure}</TableCell>
                                <TableCell>{formatDate(flight.arrivalDateTime)}</TableCell>
                                <TableCell>{flight.destination}</TableCell>
                                <TableCell>{flight.price}</TableCell>
                                <TableCell>{flight.seats}</TableCell>
                                <TableCell>{flight.duration}</TableCell>
                                {
                                    authCtx.isAdmin() &&
                                    <TableCell>
                                        <Button variant='contained' color='primary' onClick={(event) => deleteFlight(flight.id, event)}>
                                            Delete
                                        </Button>
                                    </TableCell>
                                }
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
                <TablePagination
                    rowsPerPageOptions={[5, 10, 25]}
                    component="div"
                    count={flights.length}
                    rowsPerPage={rowsPerPage}
                    page={page}
                    onPageChange={handleChangePage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                />
            </Paper>
        </>
    );
}

export default FlightsPage;