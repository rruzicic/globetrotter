import { Paper, Table, TableBody, TableCell, TableHead, TablePagination, TableRow, TableSortLabel, TextField, Typography } from "@mui/material";
import { useState } from "react";
import theme from "theme";
import formatDate from "util";

const FlightsPage = () => {

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

    return (
        <>
            <Typography variant="h2" align="center" sx={{margin: '1rem 0'}}>List of all flights</Typography>
            <Paper elevation={4} sx={{ width: '60%', margin: '1rem auto', display: 'flex', justifyContent: 'space-around', padding: '0.5rem 0' }} >
                    <TextField size='small' label='Date'/>
                    <TextField size='small' label='Departure'/>
                    <TextField size='small' label='Destination'/>
                    <TextField size='small' label='Passenger Number'/>
            </Paper>
            <Paper elevation={4} sx={{ width: '90%', margin: '1rem auto' }}>
            <Table>
                <TableHead>
                    <TableRow sx={{backgroundColor: theme.palette.primary.main, color: theme.palette.secondary.main}}>
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
                    </TableRow>
                </TableHead>
                <TableBody>
                    {sortedFlights.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((flight) => (
                        <TableRow key={flight.id} sx={styles.row}>
                            <TableCell>{formatDate(flight.departureDateTime.toISOString())}</TableCell>
                            <TableCell>{flight.departure}</TableCell>
                            <TableCell>{formatDate(flight.arrivalDateTime.toISOString())}</TableCell>
                            <TableCell>{flight.destination}</TableCell>
                            <TableCell>{flight.price}</TableCell>
                            <TableCell>{flight.seats}</TableCell>
                            <TableCell>{flight.duration}</TableCell>
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