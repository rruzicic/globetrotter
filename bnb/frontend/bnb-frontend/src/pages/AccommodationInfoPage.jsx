import { useContext, useEffect, useState } from "react";
import { useParams } from "react-router";
import { Grid, Typography, Button, Dialog, DialogTitle, DialogContent, TextField, Stack, Box, Container } from "@mui/material";
import RequestCard from "../components/accommodationManagement/RequestCard";
import AuthContext from "../config/authContext";
import ImageGallery from 'react-image-gallery';
import theme from "../theme";
import { axiosInstance } from "../config/interceptor";
import CONSTANTS from "../config/constants";
import BenefitsSelectionGrid from "../components/common/BenefitSelectionGrid";
import { toast } from "react-toastify";

const AccommodationInfoPage = () => {
    const { id } = useParams()
    const [objectInfo, setObjectInfo] = useState(null)
    const [requests, setRequests] = useState([])
    const [open, setOpen] = useState(false)
    const [benefits, setBenefits] = useState([])
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
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                setObjectInfo(response.data)
                if (response.data.availableCommodations) {
                    setBenefits(response.data.availableCommodations)
                } else {
                    setBenefits([])
                }
            })
        axiosInstance.get(`${CONSTANTS.GATEWAY}/reservation/accommodation/${id}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                setRequests(response.data)
            })
    }, [])

    const updateBenefits = () => {
        axiosInstance.put(`${CONSTANTS.GATEWAY}/accommodation/`, { ...objectInfo, availableCommodations: benefits })
            .catch((e) => {
                console.error(e)
            })
            .then((response) => {
                toast("Benefits updated!")
            })
    }

    const acceptReservation = (id) => {
        axiosInstance.post(`${CONSTANTS.GATEWAY}/reservation/approve/${id}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                setRequests((prevRequests) =>
                    prevRequests.map((request) =>
                        request.id === id ? { ...request, isApproved: true } : request
                    )
                );
            })
    }
    const declineReservation = (id) => {
        axiosInstance.post(`${CONSTANTS.GATEWAY}/reservation/reject/${id}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                setRequests((prev) => prev.filter((a) => a.id !== id))
            })
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
            .catch((error) => {
                console.error(error)
                return
            })
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
            .catch((error) => {
                console.error(error)
                return
            })
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
            <Grid container mb={4} justifyContent={"center"} spacing={1}>
                {
                    objectInfo && (
                        <>
                            <Typography variant="h3" color="initial" mt={8}>
                                {objectInfo.name}
                            </Typography>
                            <Grid item xs={12}>
                                <ImageGallery items={images} />
                            </Grid>
                            <Grid item xs={3} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem', textAlign: 'center', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
                                <Typography variant="h6" color="secondary">
                                    Accommodation info
                                </Typography>
                                <Stack sx={{ marginTop: '2rem', color: 'white' }} spacing={2}>
                                    <Box>
                                        Name: {objectInfo.name}
                                    </Box>
                                    <Box>
                                        Maximal number of guests: {objectInfo.guests}
                                    </Box>
                                    <Box>
                                        Price per unit: {objectInfo.unitPrice.amount.toString()}
                                    </Box>
                                </Stack>
                            </Grid>
                            <Grid item xs={3} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem', textAlign: 'center', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
                                <Typography variant="h6" color="secondary">
                                    Location info
                                </Typography>
                                <Stack sx={{ marginTop: '2rem', color: 'white' }} spacing={2}>
                                    <Box>
                                        Street: {objectInfo.location.street} {objectInfo.location.streetNum}
                                    </Box>
                                    <Box>
                                        ZIP Code: {objectInfo.location.zip}
                                    </Box>
                                    <Box>
                                        Country: {objectInfo.location.country}
                                    </Box>
                                </Stack>
                            </Grid>
                            <Grid item xs={3} sx={{ display: 'grid', placeItems: 'center', paddingLeft: '0', paddingRight: '0' }}>
                                <BenefitsSelectionGrid selected={benefits} setSelected={setBenefits} disabled={!authCtx.isHost()} />
                                {authCtx.isHost() && (
                                    <Button variant="contained" color="primary" sx={{ marginTop: '1rem' }} onClick={updateBenefits}>
                                        Update
                                    </Button>
                                )}
                            </Grid>
                        </>
                    )
                }
                {
                    requests && authCtx.isHost() && (
                        <Container>
                            <Typography variant="h4" mt={4}>
                                Reservation requests
                            </Typography>
                            <Grid container spacing={2} mt={4}>
                                {
                                    requests.map((request, index) => {
                                        return (
                                            <Grid item xs={12} key={index}>
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
                        </Container>
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