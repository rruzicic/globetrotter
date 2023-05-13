import { Paper, TableContainer, Table, TableHead, TableRow, TableCell, TableBody, TablePagination, Stack, TextField, Button } from "@mui/material";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const ObjectDataGrid = () => {
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);

    const [locationSP, setLocationSP] = useState()
    const [guestNumberSP, setGuestNumberSP] = useState()
    const [startDateSP, setStartDateSP] = useState()
    const [endDateSP, setEndDateSP] = useState()
    const [objects, setObjects] = useState(null)


    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(+event.target.value);
        setPage(0);
    };

    const columns = [
        {
            id: '1',
            align: 'left',
            minWidth: '150px',
            label: 'Name',
        },
        {
            id: '2',
            align: 'left',
            minWidth: '150px',
            label: 'Location',
        },
        {
            id: '3',
            align: 'left',
            minWidth: '150px',
            label: 'Price per night',
        },
        {
            id: '4',
            align: 'left',
            minWidth: '150px',
            label: 'Total price',
        },
        {
            id: '5',
            align: 'left',
            minWidth: '150px',
            label: 'User Action',
        }
    ]

    const search = () => {
        //TODO: call API for search
        setObjects([
            {
                id: 1,
                name: 'Village House',
                priceNight: 25,
                priceTotal: 125,
                location: 'Tara'
            },
            {
                id: 2,
                name: 'Village House',
                priceNight: 25,
                priceTotal: 125,
                location: 'Tara'
            },
            {
                id: 3,
                name: 'Village House',
                priceNight: 25,
                priceTotal: 125,
                location: 'Tara'
            },
            {
                id: 4,
                name: 'Village House',
                priceNight: 25,
                priceTotal: 125,
                location: 'Tara'
            },
        ])
        console.log({
            startDate: startDateSP,
            endDate: endDateSP,
            location: locationSP,
            guestNumber: guestNumberSP
        });
    }

    const handleLocationChange = (event) => {
        setLocationSP(event.target.value)
    }

    const handleGuestNumberChange = (event) => {
        setGuestNumberSP(event.target.value)
    }

    const handleStartDateChange = (event) => {
        setStartDateSP(event.target.value)
    }

    const handleEndDateChange = (event) => {
        setEndDateSP(event.target.value)
    }

    const navigate = useNavigate()
    const seeInfo = (id) => {
        navigate(`/accommodationInfo/${id}`)
    }

    const handleBook = (id, event) => {
        event.stopPropagation()
        console.log('Sent request for object with id: ' + id);
    }

    return (
        <>
            <Stack direction={"row"} sx={{ width: '100%', justifyContent: 'center' }} spacing={4} mt={4} mb={2}>
                <TextField label="Location" variant="outlined" onChange={handleLocationChange} />
                <TextField type="number" label="Guest number" variant="outlined" onChange={handleGuestNumberChange} />
                <input type="date" onChange={handleStartDateChange} label='Start Date' />
                <input type="date" onChange={handleEndDateChange} label='End Date' />
                <Button variant="contained" color="primary" onClick={search}>
                    Search
                </Button>
            </Stack>
            {
                objects && (
                    <Paper sx={{ width: '100%', overflow: 'hidden' }}>
                        <TableContainer sx={{ maxHeight: 440 }}>
                            <Table stickyHeader aria-label="sticky table">
                                <TableHead>
                                    <TableRow>
                                        {columns.map((column) => (
                                            <TableCell
                                                key={column.id}
                                                align={column.align}
                                                style={{ minWidth: column.minWidth }}
                                            >
                                                {column.label}
                                            </TableCell>
                                        ))}
                                    </TableRow>
                                </TableHead>
                                <TableBody>
                                    {objects
                                        .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                                        .map((object) => {
                                            return (
                                                <TableRow hover tabIndex={-1} key={object.id} onClick={() => seeInfo(object.id)}>
                                                    <TableCell>
                                                        {object.name}
                                                    </TableCell>
                                                    <TableCell>
                                                        {object.location}
                                                    </TableCell>
                                                    <TableCell>
                                                        {object.priceNight}
                                                    </TableCell>
                                                    <TableCell>
                                                        {object.priceTotal}
                                                    </TableCell>
                                                    <TableCell>
                                                        <Button variant="outlined" color="primary" onClick={(e) => handleBook(object.id, e)}>
                                                            Book!
                                                        </Button>
                                                    </TableCell>
                                                </TableRow>
                                            );
                                        })}
                                </TableBody>
                            </Table>
                        </TableContainer>
                        <TablePagination
                            rowsPerPageOptions={[10, 25, 100]}
                            component="div"
                            count={objects.length}
                            rowsPerPage={rowsPerPage}
                            page={page}
                            onPageChange={handleChangePage}
                            onRowsPerPageChange={handleChangeRowsPerPage}
                        />
                    </Paper>
                )
            }
        </>
    );
}

export default ObjectDataGrid;