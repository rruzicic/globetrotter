import { useContext, useEffect, useState } from "react";
import { useParams } from "react-router";
import { Grid, Typography, Button, Dialog, DialogTitle, DialogContent, DialogContentText, TextField, DialogActions, Stack, Divider } from "@mui/material";
import RequestCard from "../components/accommodationManagement/RequestCard";
import AuthContext from "../config/authContext";
import ImageGallery from 'react-image-gallery';
import theme from "../theme";
import { axiosInstance } from "../config/interceptor";
import CONSTANTS from "../config/constants";
import { Link } from "react-router-dom"

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
                console.log(response);
                setObjectInfo(response.data)
            })
        setRequests(
            [{
                userId: '1',
                requestId: '1',
                startDate: new Date(),
                endDate: new Date(),
                numberOfGuests: 8
            },
            {
                userId: '2',
                requestId: '2',
                startDate: new Date(),
                endDate: new Date(),
                numberOfGuests: 4
            },
            {
                userId: '3',
                requestId: '3',
                startDate: new Date(),
                endDate: new Date(),
                numberOfGuests: 6
            }]
        )
    }, [])

    const acceptReservation = (id) => {
        //TODO: send to BE
        console.log('Accepted ' + id);
    }
    const declineReservation = (id) => {
        //TODO: send to BE
        console.log('Declined ' + id);
    }
    const handleOpen = () => {
        setOpen((prev) => !prev)
    }
    //TODO: hit endpoint
    const submitPriceChange = () => {
        console.log({
            accommodationId: id,
            newPrice: price,
            newInterval: {
                start: new Date(startDate).toISOString(),
                end: new Date(endDate).toISOString()
            },
            priceForPerson: isSinglePerson
        });
    }
    const submitAvailabilityChange = () => {
        console.log({
            accommodationId: id,
            timeInterval: {
                start: new Date(aStartDate).toISOString(),
                end: new Date(aEndDate).toISOString(),
            },
        });
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
                                                <RequestCard requestId={request.requestId} accept={acceptReservation} decline={declineReservation} userId={request.userId} startDate={request.startDate} endDate={request.endDate} guestNumber={request.numberOfGuests} />
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