import { useContext, useEffect, useState } from "react";
import { useParams } from "react-router";
import { Grid, Typography, Button } from "@mui/material";
import RequestCard from "../components/accommodationManagement/RequestCard";
import AuthContext from "../config/authContext";

const AccommodationInfoPage = () => {
    const { id } = useParams()
    const [objectInfo, setObjectInfo] = useState()
    const [requests, setRequests] = useState([])
    let authCtx = useContext(AuthContext)

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
        <Grid container>
            {
                objectInfo && (
                    <>
                        <Grid item xs={6}>
                            <Typography variant="h6" color="initial">
                                Object info with id: {id}
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
                        <Grid item xs={6}>
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
            <Grid item xs={12}>
                React image gallery npm
            </Grid>
            <Grid container spacing={2} mt={4}>
                {
                    requests && authCtx.isHost() && requests.map((request) => {
                        return (
                            <Grid item xs={12}>
                                <RequestCard requestId={request.requestId} accept={acceptReservation} decline={declineReservation} userId={request.userId} startDate={request.startDate} endDate={request.endDate} guestNumber={request.numberOfGuests} />
                            </Grid>
                        )
                    })
                }
            </Grid>
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