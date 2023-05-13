import { useState } from "react";
import AccommodationCard from "../components/accommodationManagement/AccomodationCard";
import { Grid } from "@mui/material";

const AccommodationManagementPage = () => {
    const [objects, setObjects] = useState([
        {
            name: 'Village home',
            location: 'Zlatibor',
            image: '/home.jpg'
        },
        {
            name: 'Mountain home',
            location: 'Tara',
            image: '/home.jpg'
        }
    ])

    const styles = {
        card: {
            display: 'grid',
            placeItems: 'center',
            '&:hover': {
                transform: 'scale(1.1)'
            },
            cursor: 'pointer',
            transition: 'all 1s'
        }
    }

    return (
        <Grid container spacing={2} mt={6}>
            {
                objects && objects.map((object) => {
                    return (
                        <Grid item xs={6} sx={styles.card}>
                            <AccommodationCard name={object.name} location={object.location} image={object.image} />
                        </Grid>
                    )
                }
                )
            }
        </Grid>
    );
}

export default AccommodationManagementPage;