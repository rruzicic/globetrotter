import { useContext, useEffect, useState } from "react";
import { axiosInstance } from "../../config/interceptor";
import CONSTANTS from "../../config/constants";
import { Button, Card, CardActionArea, CardActions, CardContent, CardMedia, TextField, Typography } from "@mui/material";
import AuthContext from "../../config/authContext";
import { toast } from "react-toastify";

const HostRatingCard = ({ id }) => {
    const [hostData, setHostData] = useState()
    const [rating, setRating] = useState('');

    const ctx = useContext(AuthContext)

    const handleChange = (event) => {
        setRating(event.target.value);
    };

    const submitRating = () => {
        const dto = { rating: parseInt(rating), userId: ctx.userId(), hostId: id }
        axiosInstance.post(`${CONSTANTS.GATEWAY}/feedback/HostFeedback/`, dto)
            .catch((err) => {
                console.error(err);
            })
            .then((response) => {
                toast('Successfully rated 🔥')
            })
    }

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/user/id/${id}`)
            .catch((err) => {
                console.error(err);
            })
            .then((response) => {
                setHostData(response.data)
            })
    }, [])


    return (
        <>
            {
                hostData && (
                    <Card sx={{ width: 345 }}>
                        <CardActionArea>
                            <CardMedia
                                component="img"
                                height="200"
                                image="account.png"
                                alt="green iguana"
                            />
                            <CardContent>
                                <Typography gutterBottom variant="h5" component="div">
                                    {hostData.firstName} {hostData.lastName}
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

export default HostRatingCard;