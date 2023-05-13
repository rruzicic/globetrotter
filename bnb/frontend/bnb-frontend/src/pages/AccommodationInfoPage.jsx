import { useContext, useEffect, useState } from "react";
import { useParams } from "react-router";
import { Grid, Typography, Button } from "@mui/material";
import RequestCard from "../components/accommodationManagement/RequestCard";
import AuthContext from "../config/authContext";
import ImageGallery from 'react-image-gallery';
import theme from "../theme";

const AccommodationInfoPage = () => {
    const { id } = useParams()
    const [objectInfo, setObjectInfo] = useState()
    const [requests, setRequests] = useState([])
    let authCtx = useContext(AuthContext)

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
        //TODO: fetch object info by id
        setObjectInfo(
            {
                name: 'Village home',
                minGuestNumber: '2',
                maxGuestNumber: '8',
                streetName: 'Balzakova',
                streetNumber: '64',
                zipCode: '21000',
                country: 'Srbija'
            }
        )
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

    return (
        <Grid container justifyContent={"center"} spacing={4}>
            {
                objectInfo && (
                    <>
                        <Typography variant="h6" color="initial" mt={8}>
                            {objectInfo.name} (id: {id})
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
                                <Grid item xs={6}>
                                    Minimal number of guests: {objectInfo.minGuestNumber}
                                </Grid>
                                <Grid item xs={12}>
                                    Maximal number of guests: {objectInfo.maxGuestNumber}
                                </Grid>
                            </Grid>
                        </Grid>
                        <Grid item xs={5} sx={{ backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '0 0.5rem' }}>
                            <Typography variant="h6" color="initial">
                                Location info
                            </Typography>
                            <Grid container>
                                <Grid item xs={8}>
                                    Street name: {objectInfo.streetName}
                                </Grid>
                                <Grid item xs={4}>
                                    Street number: {objectInfo.streetNumber}
                                </Grid>
                                <Grid item xs={12}>
                                    ZIP Code: {objectInfo.zipCode}
                                </Grid>
                                <Grid item xs={12}>
                                    Country: {objectInfo.country}
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
                        <Button variant="contained" color="primary" disabled>
                            Change info (coming soon..)
                        </Button>
                    </Grid>
                )
            }
        </Grid>
    );
}

export default AccommodationInfoPage;