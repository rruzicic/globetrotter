import { useContext, useEffect, useState } from "react";
import { axiosInstance } from "../../config/interceptor";
import CONSTANTS from "../../config/constants";
import { Button, Card, CardActionArea, CardActions, CardContent, CardMedia, TextField, Typography } from "@mui/material";
import AuthContext from "../../config/authContext";


const AccommodationRatingCard = ({ id }) => {
    const [accommodationData, setAccommodationData] = useState()
    const [rating, setRating] = useState('');

    const ctx = useContext(AuthContext)

    const handleChange = (event) => {
        setRating(event.target.value);
    };

    const submitRating = () => {
        const dto = { rating: parseInt(rating), userId: ctx.userId(), accommodationId: id }
        console.log(dto);
        axiosInstance.post(`${CONSTANTS.GATEWAY}/feedback/AccommodationFeedback/`, dto)
            .catch((err) => {
                console.error(err);
            })
            .then((response) => {
                console.log(response);
            })
    }

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/accommodation/${id}`)
            .catch((err) => {
                console.error(err);
            })
            .then((response) => {
                setAccommodationData(response.data)
            })
    }, [])

    return (
        <>
            {
                accommodationData && (
                    <Card sx={{ width: 345 }}>
                        <CardActionArea>
                            <CardMedia
                                component="img"
                                height="200"
                                image={accommodationData.photos[0]}
                                alt="green iguana"
                            />
                            <CardContent>
                                <Typography gutterBottom variant="h5" component="div">
                                    {accommodationData.name}
                                </Typography>
                            </CardContent>
                        </CardActionArea>
                        <CardActions>
                            <TextField onChange={handleChange} value={rating} type="number" inputProps={{
                                min: 1,
                                max: 5,
                            }} />
                            <Button size="small" color="primary" onClick={submitRating}>
                                Rate
                            </Button>
                        </CardActions>
                    </Card>
                )
            }
        </>
    );
}

export default AccommodationRatingCard;