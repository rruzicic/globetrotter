import { Box, Container } from "@mui/material";
import MyReservationCard from "../components/myReservations/MyReservationCard";
import { axiosInstance } from "../config/interceptor"
import AuthContext from "../config/authContext"
import { useContext, useEffect, useState } from "react";
import CONSTANTS from "../config/constants";
import { useNavigate } from "react-router";

const MyReservationsPage = () => {
    const authCtx = useContext(AuthContext)
    const [data, setData] = useState(null)

    useEffect(() => {
        axiosInstance.get(`http://localhost:4000/user/email/${authCtx.userEmail()}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                axiosInstance.get(`${CONSTANTS.GATEWAY}/reservation/user/${response.data.id}`)
                    .catch((error) => {
                        console.error(error)
                        return
                    })
                    .then((ret) => {
                        setData(ret.data)
                    })
            })
    }, [])
    const handleCancel = (id) => {
        axiosInstance.delete(`${CONSTANTS.GATEWAY}/reservation/${id}`)
            .catch((error) => {
                console.error(error)
                return
            })
            .then((response) => {
                setData((prev) => prev.filter((a) => a.id !== id))
            })
    }

    const navigate = useNavigate()
    const handleNavigate = (id) => {
        navigate(`/reservationInfo/${id}`)
    }

    const styles = {
        box: {
            '&:hover': {
                transform: 'scale(1.02)'
            },
            cursor: 'pointer',
            transition: 'all 1s',
            width: '100%',
            boxShadow: 'rgba(0, 0, 0, 0.35) 0px 5px 15px',
            borderRadius: '10px'
        }
    }

    return (
        <Box sx={{ marginTop: '2rem', width: '90%', margin: '1rem auto' }}>
            {
                data && data.map((reservation) => {
                    return (
                        <Box key={reservation.id} onClick={() => handleNavigate(reservation.id)} sx={styles.box}>
                            <MyReservationCard
                                reservationId={reservation.id}
                                handleCancel={handleCancel}
                                objectId={reservation.accommodationId}
                                start={reservation.dateInterval.start}
                                end={reservation.dateInterval.end}
                                guestNum={reservation.numOfGuests}
                                totalPrice={reservation.totalPrice}
                                status={reservation.isApproved}
                                image='/home.jpg' />
                        </Box>
                    )
                })
            }
        </Box>
    );
}

export default MyReservationsPage;