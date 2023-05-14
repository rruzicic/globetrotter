import { Container } from "@mui/material";
import MyReservationCard from "../components/myReservations/MyReservationCard";
import { axiosInstance } from "../config/interceptor"
import AuthContext from "../config/authContext"
import { useContext, useEffect, useState } from "react";
import CONSTANTS from "../config/constants";

const MyReservationsPage = () => {
    const authCtx = useContext(AuthContext)
    const [data, setData] = useState(null)

    useEffect(() => {
        axiosInstance.get(`http://localhost:4000/user/email/${authCtx.userEmail()}`)
            .then((response) => {
                axiosInstance.get(`${CONSTANTS.GATEWAY}/reservation/user/${response.data.id}`)
                    .then((ret) => {
                        setData(ret.data)
                        console.log(ret.data);
                    })
            })
    }, [])

    const handleCancel = (id) => {
        axiosInstance.delete(`${CONSTANTS.GATEWAY}/reservation/${id}`)
            .then((response) => {
                console.log(response);
            })
    }

    return (
        <Container sx={{ marginTop: '2rem' }}>
            {
                data && data.map((reservation) => {
                    return (
                        <MyReservationCard key={reservation.id}
                            reservationId={reservation.id}
                            handleCancel={handleCancel}
                            objectId={reservation.accommodationId}
                            start={reservation.dateInterval.start}
                            end={reservation.dateInterval.end}
                            guestNum={reservation.numOfGuests}
                            totalPrice={reservation.totalPrice}
                            status={reservation.isApproved}
                            image='/home.jpg' />
                    )
                })
            }
        </Container>
    );
}

export default MyReservationsPage;