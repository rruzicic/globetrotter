import { Container } from "@mui/material";
import MyReservationCard from "../components/myReservations/MyReservationCard";

const MyReservationsPage = () => {

    const reservations = [
        {
            reservationId: '1',
            objectId: '1',
            start: new Date(),
            end: new Date(),
            guestNum: '8',
            totalPrice: '120',
            objectName: "Village home",
            image: '/home.jpg'
        },
        {
            reservationId: '2',
            objectId: '2',
            start: new Date(),
            end: new Date(),
            guestNum: '8',
            totalPrice: '120',
            objectName: "Village home",
            image: '/home1.jpg'
        },
        {
            reservationId: '3',
            objectId: '3',
            start: new Date(),
            end: new Date(),
            guestNum: '8',
            totalPrice: '120',
            objectName: "Village home",
            image: '/home2.jpg'
        },
        {
            reservationId: '4',
            objectId: '4',
            start: new Date(),
            end: new Date(),
            guestNum: '8',
            totalPrice: '120',
            objectName: "Village home",
            image: '/home.jpg'
        },
    ]

    const handleCancel = (id) => {
        console.log('Cancel res with id ' + id);
    }

    return (
        <Container sx={{marginTop: '2rem'}}>
            {
                reservations.map((reservation) => {
                    return (<MyReservationCard key={reservation.reservationId} reservationId={reservation.reservationId} handleCancel={handleCancel} objectId={reservation.objectId} start={reservation.start} end={reservation.end} guestNum={reservation.guestNum} totalPrice={reservation.totalPrice} objectName={reservation.objectId} image={reservation.image} />)
                })
            }
        </Container>
    );
}

export default MyReservationsPage;