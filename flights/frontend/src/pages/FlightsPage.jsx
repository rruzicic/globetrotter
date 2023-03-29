import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Paper, Table, TableBody, TableCell, TableHead, TablePagination, TableRow, TableSortLabel, TextField, Typography } from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers";
import { useCallback, useContext, useEffect, useState } from "react";
import theme from "theme";
import { useDebounce } from "use-debounce";
import formatDate from "util";

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
    //TODO: remove after adding role recognition
    const [admin, setAdmin] = useState(true)

    //temporary dummy data, should fetch flights from API inside useEffect hook with no
    //parameters in dependency array
    const flights = [
        {
            id: 0,
            departureDateTime: new Date('2024-04-15T18:00:00'),
            arrivalDateTime: new Date('2024-04-16T02:00:00'),
            departure: 'Belgrade',
            destination: 'Bangkok',
            price: 290.33,
            seats: 150,
            duration: 8
        },
        {
            id: 1,
            departureDateTime: new Date('2024-05-15T12:00:00'),
            arrivalDateTime: new Date('2024-05-15T19:00:00'),
            departure: 'Belgrade',
            destination: 'Shanghai',
            price: 130.33,
            seats: 200,
            duration: 7
        },
        {
            id: 2,
            departureDateTime: new Date('2024-04-23T08:00:00'),
            arrivalDateTime: new Date('2024-04-23T09:00:00'),
            departure: 'Belgrade',
            destination: 'Kharkiv',
            price: 20.33,
            seats: 100,
            duration: 1
        },
        {
            id: 3,
            departureDateTime: new Date('2024-04-22T16:00:00'),
            arrivalDateTime: new Date('2024-04-22T17:00:00'),
            departure: 'Belgrade',
            destination: 'Donbas',
            price: 20.33,
            seats: 100,
            duration: 1
        },
        {
            id: 4,
            departureDateTime: new Date('2024-04-22T17:00:00'),
            arrivalDateTime: new Date('2024-04-22T19:00:00'),
            departure: 'Belgrade',
            destination: 'Munich',
            price: 120.33,
            seats: 200,
            duration: 2
        },
        {
            id: 5,
            departureDateTime: new Date('2024-04-22T16:00:00'),
            arrivalDateTime: new Date('2024-04-22T19:00:00'),
            departure: 'Belgrade',
            destination: 'Paris',
            price: 300.00,
            seats: 200,
            duration: 3
        },
        {
            id: 6,
            departureDateTime: new Date('2024-04-22T08:00:00'),
            arrivalDateTime: new Date('2024-04-22T12:00:00'),
            departure: 'Belgrade',
            destination: 'Mogadishu',
            price: 250.33,
            seats: 50,
            duration: 4
        },
    ]

    const handleSort = (field) => {
        const isAsc = (orderBy === field && order === 'asc');
        setOrderBy(field);
        setOrder(isAsc ? 'desc' : 'asc');
    };

    //probably spaghetti code but the best i could do
    const sortedFlights = flights.sort((a, b) => {
        if (orderBy === 'departureDateTime' || orderBy === 'arrivalDateTime') {
            return order === 'asc'
                ? a[orderBy].getTime() - b[orderBy].getTime()
                : b[orderBy].getTime() - a[orderBy].getTime();
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

    const styles = {
        row: {
            cursor: 'pointer',
            '&:hover': {
                backgroundColor: theme.palette.primary.light,
            }
        }
    }

    const deleteFlight = (id) => {
        console.log('Should delete flight with id: ' + id);
        //should send API request to remove flight
    }

    //TODO: remove after adding role recognition
    const changeRole = () => {
        setAdmin((prev) => !prev)
    }

    //tracks values
    const [departureSP, setDepartureSP] = useState()
    const [departureDateSP, setDepartureDateSP] = useState()
    const [destinationSP, setDestinationSP] = useState()
    const [arrivalDateSP, setArrivalDateSP] = useState()
    const [passengerNumSP, setPassengerNumSP] = useState()
    //prevents sending api requests for every char typed, sends it 500ms after last char is typed
    const [debounceDepartureSP] = useDebounce(departureSP, 500)
    const [debounceDepartureDateSP] = useDebounce(departureDateSP, 500)
    const [debounceDestinationSP] = useDebounce(destinationSP, 500)
    const [debounceArrivalDateSP] = useDebounce(arrivalDateSP, 500)
    const [debouncePassengerNumSP] = useDebounce(passengerNumSP, 500)

    //had to do it like this because of non matching versions of npm packages
    const changeDeparture = (e) => {
        setDepartureSP(e.target.value);
    }
    const changeDepartureDate = (e) => {
        setDepartureDateSP(e)
    }
    const changeDestination = (e) => {
        setDestinationSP(e.target.value);
    }
    const changeArrivalDate = (e) => {
        setArrivalDateSP(e);
    }
    const changePassengerNumber = (e) => {
        setPassengerNumSP(e.target.value);
    }

    const [openModal, setOpenModal] = useState(false)
    const [selectedFlight, setSelectedFlight] = useState(false)
    const [ticketNumber, setTicketNumber] = useState(false)
    // const authCntx = useContext(Auth)

    const buyTicket = (() => {
        console.log('Bought ' +  ticketNumber + ' tickets for flight with id ' + selectedFlight);
        //TOOD: get user id from context
        // axiosInstance.post('http://localhost:8080/flights/buy-ticket', {
        //     flightId: selectedFlight,
        //     userId: 1,
        //     numOfTicketsOptional: ticketNumber
        // })
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

    const search = useCallback(() => {
        console.log(
            {
                departure: debounceDepartureSP,
                departureDateTime: debounceDepartureDateSP,
                destination: debounceDestinationSP,
                arrivalDateTime: debounceArrivalDateSP,
                passengerNumber: debouncePassengerNumSP
            }
            //should not be logged, should send request to search api with these params
        );
    }, [debounceDepartureSP, debounceDepartureDateSP, debounceDestinationSP, debounceArrivalDateSP, debouncePassengerNumSP])

    useEffect(() => {
        search()
    }, [search])

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
                    <Button onClick={() => {buyTicket()}}>Buy</Button>
                </DialogActions>
            </Dialog>
            <Typography variant="h2" align="center" sx={{ margin: '1rem 0' }}>List of all flights <Button onClick={changeRole}>Is admin: {String(admin)}</Button></Typography>
            <Paper elevation={4} sx={{ width: '60%', margin: '1rem auto', display: 'flex', justifyContent: 'space-around', padding: '0.5rem 0' }} >
                <TextField onChange={changeDeparture} label='Departure' />
                <DatePicker
                    value={departureDateSP}
                    renderInput={(props) => <TextField {...props} />}
                    onChange={changeDepartureDate} label='Departure Date' />
                <TextField onChange={changeDestination} label='Destination' />
                <DatePicker
                    value={arrivalDateSP}
                    renderInput={(props) => <TextField {...props} />}
                    onChange={changeArrivalDate} label='Arrival Date' />
                <TextField onChange={changePassengerNumber} label='Passenger Number' />
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
                                admin &&
                                <TableCell />
                            }
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {sortedFlights.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((flight) => (
                            <TableRow key={flight.id} sx={styles.row} onClick={() => { handleOpen(flight.id) }}>
                                <TableCell>{formatDate(flight.departureDateTime.toISOString())}</TableCell>
                                <TableCell>{flight.departure}</TableCell>
                                <TableCell>{formatDate(flight.arrivalDateTime.toISOString())}</TableCell>
                                <TableCell>{flight.destination}</TableCell>
                                <TableCell>{flight.price}</TableCell>
                                <TableCell>{flight.seats}</TableCell>
                                <TableCell>{flight.duration}</TableCell>
                                {
                                    admin &&
                                    <TableCell>
                                        <Button variant='contained' color='primary' onClick={() => deleteFlight(flight.id)}>
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