import { Paper, TableContainer, Table, TableHead, TableRow, TableCell, TableBody, TablePagination, Stack, TextField, Button } from "@mui/material";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { axiosInstance } from '../../config/interceptor'
import CONSTANTS from '../../config/constants'
import AuthContext from "../../config/authContext";

const ObjectDataGrid = () => {
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);

    const [locationSP, setLocationSP] = useState('')
    const [guestNumberSP, setGuestNumberSP] = useState('')
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
            id: '0',
            align: 'left',
            minWidth: '150px',
            label: 'Thumbnail',
        },
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
            label: 'User Action',
        }
    ]
    // localhost:4000/accommodation/search?cityName=&guestNum=0&startDate=2023-01-01&endDate=2023-12-01
    const search = () => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/accommodation/search?cityName=${locationSP ? locationSP : " "}&guestNum=${guestNumberSP ? parseInt(guestNumberSP) : " "}&startDate=${startDateSP}&endDate=${endDateSP}`)
            .catch((err) => {
                console.error(err);
                return
            })
            .then((response) => {
                setObjects(response.data)
            })
    }

    const authCtx = useContext(AuthContext)

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
        let dto = {
            accommodationId: id,
            userId: "",
            dateInterval: {
                start: new Date(startDateSP).toISOString(),
                end: new Date(endDateSP).toISOString()
            },
            numOfGuests: parseInt(guestNumberSP),
            totalPrice: 0
        }
        axiosInstance.get(`http://localhost:4000/user/email/${authCtx.userEmail()}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                dto.userId = response.data.id
                axiosInstance.post(`${CONSTANTS.GATEWAY}/reservation/`, dto)
                    .then((response) => {
                        axiosInstance.post(`${CONSTANTS.GATEWAY}/reservation/accommodation/${id}/reservation/${response.data.id}`)
                            .catch((e) => {
                                console.error(e)
                                return
                            })
                            .then((response) => {
                                console.log(response);
                            })
                    })
            })
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
                                                        <img src='/home.jpg' alt={object.name} height={'60px'} width={'auto'} />
                                                    </TableCell>
                                                    <TableCell>
                                                        {object.name}
                                                    </TableCell>
                                                    <TableCell>
                                                        {object.location.city}
                                                    </TableCell>
                                                    <TableCell>
                                                        {object.unitPrice.amount.toString()}
                                                    </TableCell>
                                                    <TableCell>
                                                        <Button variant="outlined" color="primary" onClick={(e) => handleBook(object.id, e)} disabled={authCtx.isHost()}>
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