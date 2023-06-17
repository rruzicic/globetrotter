import { useContext, useEffect, useState } from "react";
import AuthContext from "../config/authContext";
import { axiosInstance } from "../config/interceptor";
import Container from '@mui/material/Container'
import { Stack, useTheme } from "@mui/material";

const NotificationsPage = () => {
    const ctx = useContext(AuthContext)
    const [notifications, setNotifications] = useState([])
    const theme = useTheme()

    useEffect(() => {
        ctx.clearNotificationCount()
    }, [])
    
    useEffect(()=>{
        axiosInstance.get(`http://localhost:4000/notification/user/${ctx.userId()}`)
            .catch((err) => {
                console.error(err)
            })
            .then((response) => {
                setNotifications(response.data.sort((a, b) => b.createdOn - a.createdOn))
            })
    }, [])

    const timestampToDate = (timestamp) => {
        const milliseconds = timestamp * 1000
        const date = new Date(milliseconds)
        return date.toLocaleString()
    }

    const generateNotificationText = (notification) => {
        switch (notification.type) {
            case 'RESERVATION':
                return (`${timestampToDate(notification.createdOn)}:   You have a new reservation request for ${notification.accommodationName}`)
            case 'CANCELLATION':
                return (`${timestampToDate(notification.createdOn)}:   Reservation at ${notification.accommodationName} has been canceled`)
            case 'RATING':
                return ("TODO")
            case 'A_RATING':
                return ("TODO")
            case 'HOST_STATUS':
                return ("TODO")
            case 'RESPONSE':
                return ("TODO")
            default:
                console.error('Unknown notification type');
        }
    }


    return (
        <Stack spacing={4} mt={4} mb={4}>
            {
                notifications && notifications.map((notification) => {
                    return (
                        <Container sx={{backgroundColor: theme.palette.primary.main, padding: '1.5rem', width: '60%', margin: '0'}}>
                            {generateNotificationText(notification)}
                        </Container>
                    )
                })
            }
        </Stack>
    );
}

export default NotificationsPage;