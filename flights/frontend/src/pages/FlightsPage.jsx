import { Button, Checkbox, Dialog, DialogActions, DialogContent, DialogTitle, Paper, Stack, Table, TableBody, TableCell, TableHead, TablePagination, TableRow, TableSortLabel, TextField, Typography } from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers";
import { useContext, useEffect, useState } from "react";
import theme from "theme";
import { useDebounce } from "use-debounce";
import AuthContext from "config/authContext";
import { toast } from "react-toastify";
import { axiosInstance, stringAxiosInstance } from "config/interceptor";
import { formatLocaleDate } from "util";

const FlightsPage = () => {
    const fixDate = (date) => {
        date.setUTCMinutes(date.getUTCMinutes() - date.getTimezoneOffset())
        return date.toISOString().split('T')[0]
    }

    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);
    const [orderBy, setOrderBy] = useState('');
    const [order, setOrder] = useState('asc');

    const [flights, setFlights] = useState([])

    const [forFriend, setForFriend] = useState(false)
    const [friendAPIKey, setFriendAPIKey] = useState(false)
    const [myAPIKey, setMyAPIKey] = useState()

    const [departureSP, setDepartureSP] = useState("")
    const [departureDateSP, setDepartureDateSP] = useState(null)
    const [destinationSP, setDestinationSP] = useState("")
    const [arrivalDateSP, setArrivalDateSP] = useState(null)
    const [passengerNumSP, setPassengerNumSP] = useState("")


    const [debounceDepartureSP] = useDebounce(departureSP, 500)
    const [debounceDepartureDateSP] = useDebounce(
        (departureDateSP == null) ? null
            :
            fixDate(departureDateSP)
        , 500)
    const [debounceDestinationSP] = useDebounce(destinationSP, 500)
    const [debounceArrivalDateSP] = useDebounce(
        (arrivalDateSP == null) ? null
            :
            fixDate(arrivalDateSP)
        , 500)
    const [debouncePassengerNumSP] = useDebounce(passengerNumSP, 500)

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
    const handleForFriend = () => {
        setForFriend((prev) => !prev)
    }
    const handleFriendKey = (e) => {
        setFriendAPIKey(e.target.value)
    }

    useEffect(() => {
        axiosInstance.get(`/user/current`)
            .catch((e) => {
                console.error(e)
            })
            .then((res) => {
                setMyAPIKey(res.data.data.apiKey.key)
            })
    }, [])

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

    const getAll = () => {
        axiosInstance.get('/flights')
            .catch((error) => {
                toast('Unable to fetch flights, try again in a few seconds 😢')
                return
            })
            .then((response) => {
                setFlights(response.data.data)
            })
    }

    useEffect(() => {
        getAll()
    }, [])

    const styles = {
        row: {
            cursor: 'pointer',
            '&:hover': {
                backgroundColor: theme.palette.primary.light,
            }
        }
    }

    const deleteFlight = (id, event) => {
        event.stopPropagation()
        stringAxiosInstance.delete(`/flights/delete`, {
            params: {
                id: id
            }
        })
            .catch((e) => {
                toast('Could not delete flight! 😢')
            })
            .then((response) => {
                if (response !== undefined) {
                    toast('Flight successfully deleted!')
                    resetSearch()
                    getAll()
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
        if (forFriend) {
            axiosInstance.defaults.headers['x-api-key'] = myAPIKey;
            axiosInstance.post(`/flights/buy-ticket-for-friend`, {
                flightId: selectedFlight,
                apiKey: friendAPIKey,
                numOfTicketsOptional: [parseInt(ticketNumber)]
            }).catch((e) => {
                console.error(e);
                toast('Not enough tickets left 😢')
                return
            }).then((res) => {
                toast('Successfully bought tickets! 😊')
                resetSearch()
                getAll()
                handleClose()
                return
            })
            return
        }

        if (!forFriend) {
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
                        resetSearch()
                        getAll()
                    }
                })
            handleClose()
        }
    })

    const handleOpen = (flightId) => {
        if (authCtx.isUser()) {
            setSelectedFlight(flightId)
            setOpenModal(true)
        } else {
            return
        }
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
                    <Stack>
                        <TextField onChange={changeTicketNumber} label="Number of Tickets" />
                        <Stack direction={"row"} alignItems={'center'} mt={2}>
                            <Typography variant="subtitle1">
                                I'm buying for a friend:
                            </Typography>
                            <Checkbox value={forFriend} onChange={handleForFriend} label='For friend' size="small" />
                        </Stack>
                        {
                            forFriend && (
                                <TextField onChange={handleFriendKey} label="Friends API KEY" />
                            )
                        }
                    </Stack>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button onClick={() => { buyTicket() }}>Buy</Button>
                </DialogActions>
            </Dialog>
            <Typography variant="h3" align="center" sx={{ margin: '3rem 0rem 1rem 0rem' }}>Feel free to search for the flight you need: </Typography>
            <Paper elevation={4} sx={{ width: '80%', margin: '1rem auto', display: 'flex', justifyContent: 'space-around', padding: '0.5rem' }} >
                <TextField onChange={changeDeparture} label='Departure' value={departureSP} sx={{ margin: '0 1rem' }} />
                <DatePicker
                    disablePast
                    value={departureDateSP}
                    renderInput={(props) => <TextField {...props} />}
                    onChange={changeDepartureDate} label='Departure Date' sx={{ margin: '0 1rem' }} />
                <TextField onChange={changeDestination} label='Destination' value={destinationSP} sx={{ margin: '0 1rem' }} />
                <DatePicker
                    disablePast
                    value={arrivalDateSP}
                    renderInput={(props) => <TextField {...props} />}
                    onChange={changeArrivalDate} label='Arrival Date' sx={{ margin: '0 1rem' }} />
                <TextField onChange={changePassengerNumber} label='Remaining Tickets' value={passengerNumSP} sx={{ margin: '0 1rem' }} />
                <Button variant="contained" color="primary" onClick={resetSearch} sx={{ margin: '0 1rem', whiteSpace: 'nowrap', width: '13rem' }}>
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
                                    Remaining tickets
                                </TableSortLabel>
                            </TableCell>
                            <TableCell>
                                <TableSortLabel
                                    active={orderBy === 'duration'}
                                    direction={orderBy === 'duration' ? order : 'asc'}
                                    onClick={() => handleSort('duration')}
                                >
                                    Duration(min)
                                </TableSortLabel>
                            </TableCell>
                            {
                                authCtx.isAdmin() &&
                                <TableCell>
                                    Admin Controls
                                </TableCell>
                            }
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {sortedFlights.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((flight) => (
                            <TableRow key={flight.id} sx={styles.row} onClick={() => { handleOpen(flight.id) }}>
                                <TableCell>{formatLocaleDate(flight.departureDateTime)}</TableCell>
                                <TableCell>{flight.departure}</TableCell>
                                <TableCell>{formatLocaleDate(flight.arrivalDateTime)}</TableCell>
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