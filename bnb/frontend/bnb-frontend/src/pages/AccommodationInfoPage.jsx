import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { Grid, Typography, Button } from "@mui/material";

const AccommodationInfoPage = () => {
    const { id } = useParams()
    const [objectInfo, setObjectInfo] = useState()
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
    }, [])

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
            <Grid item xs={12}>
                <Button variant="contained" color="primary" disabled>
                    Change info (coming soon..)
                </Button>
            </Grid>
        </Grid>
    );
}

export default AccommodationInfoPage;