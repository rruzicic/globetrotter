import { useContext, useEffect, useState } from "react";
import AccommodationCard from "../components/accommodationManagement/AccomodationCard";
import { Box, Grid, Button, Typography } from "@mui/material";
import { Stack } from "@mui/system";
import { Link } from "react-router-dom";
import { axiosInstance } from "../config/interceptor"
import CONSTANTS from "../config/constants"
import AuthContext from "../config/authContext";

const AccommodationManagementPage = () => {
    const [objects, setObjects] = useState([])
    const authCtx = useContext(AuthContext)

    useEffect(() => {
        axiosInstance.get(`http://localhost:4000/user/email/${authCtx.userEmail()}`)
            .then((response) => {
                axiosInstance.get(`${CONSTANTS.GATEWAY}/accommodation/host/${response.data.id}`)
                    .catch((error) => {
                        console.error(error)
                        return
                    })
                    .then((response) => {
                        setObjects(response.data)
                    })
            })
    }, [])

    const styles = {
        grid: {
            display: 'grid',
            placeItems: 'center',
        },
        card: {
            '&:hover': {
                transform: 'scale(1.05)'
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
                objects && objects.map((object, index) => {
                    return (
                        <Grid item xs={4} sx={styles.grid} key={object.id}>
                            <Link to={`/accommodationInfo/${object.id}`}>
                                <Box sx={styles.card}>
                                    <AccommodationCard name={object.name} location={object.location.city} image={object.photos ? object.photos.at(0) : "/home.jpg"}/>
                                </Box>
                            </Link>
                        </Grid>
                    )
                }
                )
            }
        </Grid>
    );
}

export default AccommodationManagementPage;