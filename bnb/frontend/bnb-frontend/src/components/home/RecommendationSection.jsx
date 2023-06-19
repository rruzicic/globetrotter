import { Box, Stack } from "@mui/material";
import { useContext, useEffect, useState } from "react";
import { axiosInstance } from "../../config/interceptor";
import CONSTANTS from "../../config/constants";
import AuthContext from "../../config/authContext";
import AccommodationCard from "../accommodationManagement/AccomodationCard";
import { useNavigate } from "react-router";

const RecommendationSection = () => {
    const ctx = useContext(AuthContext)
    const [accommodations, setAccommodations] = useState()
    const navigate = useNavigate()

    useEffect(() => {
        axiosInstance.post(`${CONSTANTS.GATEWAY}/recommendation/accommodations`, { name: "what", mongoId: ctx.userId() })
            .catch((err) => {
                console.error(err);
                return;
            })
            .then((response) => {
                setAccommodations(response.data)
            })
    }, [])

    const moreInfo = (id) => {
        navigate(`/accommodationInfo/${id}`)
    }

    return (
        <Stack spacing={2} mt={4} mb={4} sx={{ width: '90vw', overflowX: "scroll", padding: '1rem', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
            {
                accommodations && accommodations.map((obj, index) => {
                    return (
                        <Box onClick={() => moreInfo(obj.mongoId)} key={index} sx={{cursor: 'pointer'}}>
                            <AccommodationCard name={obj.name} location={obj.location} image={'/home2.jpg'} />
                        </Box>
                    )
                })
            }
        </Stack>
    );
}

export default RecommendationSection;