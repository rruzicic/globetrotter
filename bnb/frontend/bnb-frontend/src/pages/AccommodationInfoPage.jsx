import { useContext, useEffect, useState } from "react";
import { useParams } from "react-router";
import { Grid, Typography, Button, Dialog, DialogTitle, DialogContent, TextField, Stack } from "@mui/material";
import RequestCard from "../components/accommodationManagement/RequestCard";
import AuthContext from "../config/authContext";
import ImageGallery from 'react-image-gallery';
import theme from "../theme";
import { axiosInstance } from "../config/interceptor";
import CONSTANTS from "../config/constants";

const AccommodationInfoPage = () => {
    const { id } = useParams()
    const [objectInfo, setObjectInfo] = useState(null)
    const [requests, setRequests] = useState([])
    const [open, setOpen] = useState(false)
    let authCtx = useContext(AuthContext)

    const [price, setPrice] = useState('');
    const [isSinglePerson, setIsSinglePerson] = useState(false);
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');
    const [aStartDate, setAStartDate] = useState('');
    const [aEndDate, setAEndDate] = useState('');

    const handlePriceChange = (event) => {
        setPrice(event.target.value);
    };

    const handleSinglePersonChange = (event) => {
        setIsSinglePerson(event.target.checked);
    };

    const handleStartDateChange = (event) => {
        setStartDate(event.target.value);
    };

    const handleEndDateChange = (event) => {
        setEndDate(event.target.value);
    };
    const handleAStartDateChange = (event) => {
        setAStartDate(event.target.value);
    };

    const handleAEndDateChange = (event) => {
        setAEndDate(event.target.value);
    };

    const images = [
        {
            original: '/home.jpg',
            thumbnail: '/home1.jpg',
        },
        {
            original: '/home1.jpg',
            thumbnail: '/home1.jpg',
        },
        {
            original: '/home2.jpg',
            thumbnail: '/home2.jpg',
        },
    ];

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/accommodation/${id}`)
            .then((response) => {
                setObjectInfo(response.data)
            })
        axiosInstance.get(`${CONSTANTS.GATEWAY}/reservation/accommodation/${id}`)
            .then((response) => {
                console.log(response.data);
                setRequests(response.data)
            })
    }, [])

    const acceptReservation = (id) => {
        axiosInstance.post(`${CONSTANTS.GATEWAY}/reservation/approve/${id}`)
            .then((response) => {
                window.location.reload()
            })
        console.log('Accepted ' + id);
    }
    const declineReservation = (id) => {
        axiosInstance.post(`${CONSTANTS.GATEWAY}/reservation/reject/${id}`)
            .then((response) => {
                window.location.reload()
            })
        console.log('Declined ' + id);
    }
    const handleOpen = () => {
        setOpen((prev) => !prev)
    }
    const submitPriceChange = () => {
        let dto = {
            accommodationId: id,
            newPrice: parseFloat(price),
            newInterval: {
                start: new Date(startDate).toISOString(),
                end: new Date(endDate).toISOString()
            },
            priceForPerson: isSinglePerson
        }
        axiosInstance.put(`${CONSTANTS.GATEWAY}/accommodation/price`, dto)
            .then((response) => {
                handleOpen()
            })
    }
    const submitAvailabilityChange = () => {
        let dto = {
            accommodationId: id,
            newInterval: {
                start: new Date(aStartDate).toISOString(),
                end: new Date(aEndDate).toISOString(),
            }
        }
        axiosInstance.put(`${CONSTANTS.GATEWAY}/accommodation/availability`, dto)
            .then((response) => {
                handleOpen()
            })
    }
    return (
        <>

            <Dialog open={open} onClose={handleOpen}>
                <DialogTitle>Intervals and Price management</DialogTitle>
                <DialogContent>
                    <Grid container spacing={8}>
                        <Grid item xs={6} >
                            <Typography variant="subtitle1" mb={2}>
                                Update price
                            </Typography>
                            <Stack spacing={2}>
                                <TextField name="price" label="Price" onChange={handlePriceChange} />
                                <Stack direction={"row"}>
                                    <Typography variant="body1">
                                        Single person
                                    </Typography>
                                    <input type="checkbox" name="singlePerson" onChange={handleSinglePersonChange} />
                                </Stack>
                                <input type="date" name='start' onChange={handleStartDateChange} />
                                <input type="date" name='end' onChange={handleEndDateChange} />
                                <Button variant="contained" color="primary" onClick={submitPriceChange}>
                                    Submit
                                </Button>
                            </Stack>
                        </Grid>
                        <Grid item xs={6} >
                            <Typography variant="subtitle1" mb={2}>
                                Update availability
                            </Typography>
                            <Stack spacing={2}>
                                <input type="date" onChange={handleAStartDateChange} />
                                <input type="date" onChange={handleAEndDateChange} />
                                <Button variant="contained" color="primary" onClick={submitAvailabilityChange}>
                                    Submit
                                </Button>
                            </Stack>
                        </Grid>
                    </Grid>
                </DialogContent>

            </Dialog>
            <Grid container mb={4} justifyContent={"center"} spacing={4}>
                {
                    objectInfo && (
                        <>
                            <Typography variant="h3" color="initial" mt={8}>
                                {objectInfo.name}
                            </Typography>
                            <Grid item xs={12}>
                                <ImageGallery items={images} />
                            </Grid>
                            <Grid item xs={5} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem' }}>
                                <Typography variant="h6" color="initial">
                                    Accommodation info
                                </Typography>
                                <Grid container>
                                    <Grid item xs={6}>
                                        Name: {objectInfo.name}
                                    </Grid>
                                    <Grid item xs={12}>
                                        Maximal number of guests: {objectInfo.guests}
                                    </Grid>
                                    <Grid item xs={12}>
                                        Price per unit: {objectInfo.unitPrice.amount.toString()}
                                    </Grid>
                                </Grid>
                            </Grid>
                            <Grid item xs={5} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem' }}>
                                <Typography variant="h6" color="initial">
                                    Location info
                                </Typography>
                                <Grid container>
                                    <Grid item xs={12}>
                                        Street: {objectInfo.location.street} {objectInfo.location.streetNum}
                                    </Grid>
                                    <Grid item xs={12}>
                                        ZIP Code: {objectInfo.location.zip}
                                    </Grid>
                                    <Grid item xs={12}>
                                        Country: {objectInfo.location.country}
                                    </Grid>
                                </Grid>
                            </Grid>
                        </>
                    )
                }
                {
                    requests && authCtx.isHost() && (
                        <>
                            <Typography variant="h4" mt={4}>
                                Reservation requests
                            </Typography>
                            <Grid container spacing={2} mt={4}>
                                {
                                    requests.map((request) => {
                                        return (
                                            <Grid item xs={12}>
                                                <RequestCard
                                                    requestId={request.id}
                                                    accept={acceptReservation}
                                                    decline={declineReservation}
                                                    userId={request.userId}
                                                    startDate={request.dateInterval.start}
                                                    endDate={request.dateInterval.end}
                                                    guestNumber={request.numOfGuests}
                                                    isApproved={request.isApproved} />

                                            </Grid>
                                        )
                                    })
                                }
                            </Grid>
                        </>
                    )
                }
                {
                    authCtx.isHost() && (
                        <Grid item xs={12} mt={4}>
                            <Button variant="contained" color="primary" onClick={handleOpen}>
                                Change info
                            </Button>
                        </Grid>
                    )
                }
            </Grid>
        </>

    );
}

export default AccommodationInfoPage;