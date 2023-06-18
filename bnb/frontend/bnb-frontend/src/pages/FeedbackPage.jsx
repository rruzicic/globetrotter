import { useContext, useEffect, useState } from "react";
import { axiosInstance } from "../config/interceptor";
import CONSTANTS from "../config/constants";
import AuthContext from "../config/authContext";
import { Container, Stack, Typography } from "@mui/material";
import HostRatingCard from "../components/feedback/HostRatingCard";
import AccommodationRatingCard from "../components/feedback/AccommodationRatingCard";

const FeedbackPage = () => {
    const ctx = useContext(AuthContext)
    const [hosts, setHosts] = useState()
    const [accommodations, setAccommodations] = useState()


    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/feedback/pastAccommodations/${ctx.userId()}`)
            .catch((err) => {
                console.error(err)
            })
            .then((response) => {
                setAccommodations(response.data)
            })
        axiosInstance.get(`${CONSTANTS.GATEWAY}/feedback/pastHosts/${ctx.userId()}`)
            .catch((err) => {
                console.error(err)
            })
            .then((response) => {
                setHosts(response.data)
            })
    }, [])

    return (
        <Container width="100%">
            <Typography variant="h4" sx={{margin: '2rem 0'}}>
                Let us know how you feel about past hosts and accommodations!
            </Typography>
            <Stack direction={"row"} mt={4} >

                <Stack sx={{ width: '50%' }} alignItems={"center"}>
                    {
                        hosts && hosts.map((host, index) => <HostRatingCard key={index} id={host} />)
                    }
                </Stack>
                <Stack alignItems={"center"}>
                    {
                        accommodations && accommodations.map((accommodation, index) => <AccommodationRatingCard key={index} id={accommodation} />)
                    }
                </Stack>
            </Stack>
        </Container>
    );
}

export default FeedbackPage;