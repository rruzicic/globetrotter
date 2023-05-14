import { Card, CardMedia, CardContent, Typography } from "@mui/material";

const AccommodationCard = ({ name, location, image }) => {
    return (
        <Card sx={{ width: 345 }}>
            <CardMedia
                sx={{ height: 140 }}
                image='/home.jpg'
                title={name}
            />
            <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                    {name}
                </Typography>
                <Typography gutterBottom variant="h6" component="div">
                    {location}
                </Typography>
            </CardContent>
        </Card>
    );
}

export default AccommodationCard;