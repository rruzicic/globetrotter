import { Paper, TableContainer, Table, TableHead, TableRow, TableCell, TableBody, TablePagination, Stack, TextField, Button, Grid, Typography, Divider, Slider, Checkbox } from "@mui/material";
import { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { axiosInstance } from '../../config/interceptor'
import CONSTANTS from '../../config/constants'
import AuthContext from "../../config/authContext";
import BenefitsSelectionGrid from "../common/BenefitSelectionGrid";

const ObjectDataGrid = () => {
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);

    const [locationSP, setLocationSP] = useState('')
    const [guestNumberSP, setGuestNumberSP] = useState('')
    const [startDateSP, setStartDateSP] = useState()
    const [endDateSP, setEndDateSP] = useState()
    const [priceRange, setPriceRange] = useState([10, 100])
    const [objects, setObjects] = useState(null)
    const [shownObjects, setShownObjects] = useState(null)
    const [benefits, setBenefits] = useState([])


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
                setShownObjects(response.data)
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
    const handlePriceRangeChange = (event, newValue) => {
        setPriceRange(newValue)
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
                    .catch((e) => {
                        console.error(e)
                        return
                    })
                    .then((response) => {
                        axiosInstance.post(`${CONSTANTS.GATEWAY}/reservation/accommodation/${id}/reservation/${response.data.id}`)
                            .catch((e) => {
                                console.error(e)
                                return
                            })
                            .then((response) => {
                            })
                    })
            })
    }

    const isSubArray = (objectBenefits) => {
        return benefits.every((benefit) => objectBenefits.includes(benefit))
    }

    const applyFilter = () => {
        setShownObjects(() => objects.filter((object) => object.unitPrice.amount >= priceRange[0] && object.unitPrice.amount <= priceRange[1] && isSubArray(object.availableCommodations ? object.availableCommodations : [])))
    }

    return (
        <>
            <Stack direction={"row"} sx={{ width: '100%', justifyContent: 'center' }} spacing={4} mt={4} mb={2}>
                <TextField label="Location" variant="outlined" onChange={handleLocationChange} />
                <TextField type="number" label="Guest number" variant="outlined" onChange={handleGuestNumberChange} />
                <input type="date" onChange={handleStartDateChange} label='Start Date' />
                <input type="date" onChange={handleEndDateChange} label='End Date' />
                <Button variant="contained" color="primary" onClick={search} disabled={!(locationSP && startDateSP && endDateSP && guestNumberSP)}>
                    Search
                </Button>
            </Stack>
            <Grid container spacing={1}>
                <Grid item xs={3}>
                    <Paper sx={{ padding: '1rem', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
                        <Typography variant="h5" >
                            Filters:
                        </Typography>
                        <Divider />
                        <Typography variant="subtitle1" >
                            Price range per day
                        </Typography>
                        <Stack direction={"row"} spacing={2}>
                            <Typography variant="subtitle1">
                                {priceRange[0]}
                            </Typography>
                            <Slider
                                value={priceRange}
                                valueLabelDisplay="auto"
                                min={0}
                                max={500}
                                onChange={handlePriceRangeChange}
                                aria-labelledby="range-slider"
                            />
                            <Typography variant="subtitle1">
                                {priceRange[1]}
                            </Typography>
                        </Stack>
                        <Stack direction={"row"}>
                            <Typography variant="subtitle1" alignSelf={'center'}>
                                Super host:
                            </Typography>
                            <Checkbox size="small" />
                        </Stack>
                        <Stack sx={{ display: 'grid', placeItems: 'center' }}>
                            <BenefitsSelectionGrid selected={benefits} setSelected={setBenefits} />
                        </Stack>
                        <Button fullWidth variant="contained" color="primary" sx={{ marginTop: '2rem' }} disabled={!objects} onClick={applyFilter}>
                            Apply filters
                        </Button>
                    </Paper>
                </Grid>
                {
                    shownObjects && (
                        <Grid item xs={9}>
                            <Paper sx={{ width: '100%', overflow: 'hidden', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
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
                                            {shownObjects
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
                        </Grid>
                    )
                }
            </Grid>
        </>
    );
}

export default ObjectDataGrid;