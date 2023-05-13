import { useEffect, useState } from "react";
import AccommodationCard from "../components/accommodationManagement/AccomodationCard";
import { Box, Grid, Button, Typography } from "@mui/material";
import { Stack } from "@mui/system";
import { Link } from "react-router-dom";

const AccommodationManagementPage = () => {
    const [objects, setObjects] = useState([])

    useEffect(() => {
        //TODO: get hosts objects and setObjects
        setObjects([
            {
                name: 'Village home',
                location: 'Zlatibor',
                image: '/home.jpg'
            },
            {
                name: 'Mountain home',
                location: 'Tara',
                image: '/home.jpg'
            },
            {
                name: 'Costal home',
                location: 'Nice',
                image: '/home.jpg'
            },
            {
                name: 'City home',
                location: 'NYC',
                image: '/home.jpg'
            }
        ])
    }, [])

    const styles = {
        grid: {
            display: 'grid',
            placeItems: 'center',
        },
        card: {
            '&:hover': {
                transform: 'scale(1.1)'
            },
            cursor: 'pointer',
            transition: 'all 1s'
        }
    }

    return (
        <Grid container spacing={2} mt={6}>
            <Grid item xs={12}>
                <Stack direction={"row"} spacing={4} mb={4}>
                    <Typography variant="subtitle1" >
                        Have another accommodation?
                    </Typography>
                    <Link to={"/newAccommodation"}>
                        <Button variant="contained" color="primary">
                            Create new accommodation
                        </Button>
                    </Link>
                </Stack>

            </Grid>
            {
                objects && objects.map((object) => {
                    return (
                        <Grid item xs={4} sx={styles.grid}>
                            <Box sx={styles.card}>
                                <AccommodationCard name={object.name} location={object.location} image={object.image} />
                            </Box>
                        </Grid>
                    )
                }
                )
            }
        </Grid>
    );
}

export default AccommodationManagementPage;