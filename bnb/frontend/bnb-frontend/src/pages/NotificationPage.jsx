import { useContext, useEffect, useState } from "react";
import AuthContext from "../config/authContext";
import { axiosInstance } from "../config/interceptor";
import Container from '@mui/material/Container'
import { Checkbox, FormControlLabel, Stack, useTheme, Typography, Button } from "@mui/material";
import CONSTANTS from "../config/constants";

const NotificationsPage = () => {
    const ctx = useContext(AuthContext)
    const [notifications, setNotifications] = useState([])
    const theme = useTheme()
    const [wantedNotification, setWantedNotification] = useState([]);
    const [printedTypes, setPrintedTypes] = useState()

    const handleCheckboxChange = (event) => {
        if (wantedNotification.includes(event.target.name)) {
            setWantedNotification(wantedNotification.filter((item) => item !== event.target.name));
        } else {
            setWantedNotification([...wantedNotification, event.target.name]);
        }
    };

    useEffect(() => {
        ctx.clearNotificationCount()
        setPrintedTypes(ctx.getWantedNotifications())
    }, [])

    useEffect(() => {
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

    const arrayToString = (array) => {
        if (array.length === 0) {
            return '';
        } else if (array.length === 1) {
            return array[0];
        } else {
            const lastElement = array[array.length - 1];
            const otherElements = array.slice(0, array.length - 1);
            return `${otherElements.join(', ')} and ${lastElement}`;
        }
    }

    const submitChanges = () => {
        axiosInstance.patch(`${CONSTANTS.GATEWAY}/user/notificationPreferences`, { id: ctx.userId(), notificationList: wantedNotification })
            .catch((e) => {
                console.error(e)
            })
            .then((response) => {
                ctx.updateWantedNotifications(wantedNotification)
            })
    }

    const generateNotificationText = (notification) => {
        switch (notification.type) {
            case 'RESERVATION':
                return (`${timestampToDate(notification.createdOn)}:   You have a new reservation request for ${notification.accommodationName}`)
            case 'CANCELLATION':
                return (`${timestampToDate(notification.createdOn)}:   Reservation at ${notification.accommodationName} has been canceled`)
            case 'RATING':
                return (`${timestampToDate(notification.createdOn)}:   You have been rated with a ${notification.rating}`)
            case 'A_RATING':
                return (`${timestampToDate(notification.createdOn)}:   Your accommodation ${notification.accommodationName} has been rated with a ${notification.rating}`)
            case 'HOST_STATUS':
                return (`${timestampToDate(notification.createdOn)}: Your super host status has been updated!`)
            case 'RESPONSE':
                return (`${timestampToDate(notification.createdOn)}: Your reservation request for ${notification.accommodationName} has been ${notification.approved ? "approved!" : "rejected.."}`)
            default:
                console.error('Unknown notification type');
        }
    }

    return (
        <Stack direction={"row"}>
            <Stack spacing={4} mt={4} mb={4} sx={{ width: '65vw' }}>
                {
                    notifications && notifications.map((notification, index) => {
                        return (
                            <Container key={index} sx={{ backgroundColor: theme.palette.primary.main, color: theme.palette.secondary.main, fontWeight: 'bold', borderRadius: '1rem', padding: '1.5rem', width: '100%', margin: '0', boxShadow: 'rgba(0, 0, 0, 0.35) 0px 5px 15px' }}>
                                {generateNotificationText(notification)}
                            </Container>
                        )
                    })
                }
            </Stack>
            {ctx.isHost() &&
                (<Stack justifyContent={"space-between"} flex={"0 0 auto"} sx={{ color: theme.palette.secondary.main, height: '350px', backgroundColor: theme.palette.primary.main, borderRadius: '10px', paddingLeft: '0', paddingTop: '0', padding: '1rem', margin: '2rem 1rem', textAlign: 'center', boxShadow: "rgba(0, 0, 0, 0.35) 0px 5px 15px" }}>
                    <Typography variant="h6">
                        Check which notifications you would like to receive as a pop-up:
                    </Typography>
                    <FormControlLabel
                        control={
                            <Checkbox
                                color="secondary"
                                checked={wantedNotification.RESERVATION}
                                onChange={handleCheckboxChange}
                                name="RESERVATION"
                            />
                        }
                        label="New reservation request"
                    />
                    <FormControlLabel
                        control={
                            <Checkbox
                                color="secondary"
                                checked={wantedNotification.CANCELLATION}
                                onChange={handleCheckboxChange}
                                name="CANCELLATION"
                            />
                        }
                        label="Reservation cancellation"
                    />
                    <FormControlLabel
                        control={
                            <Checkbox
                                color="secondary"
                                checked={wantedNotification.RATING}
                                onChange={handleCheckboxChange}
                                name="RATING"
                            />
                        }
                        label="New host rating"
                    />
                    <FormControlLabel
                        control={
                            <Checkbox
                                color="secondary"
                                checked={wantedNotification.A_RATING}
                                onChange={handleCheckboxChange}
                                name="A_RATING"
                            />
                        }
                        label="New accommodation rating"
                    />
                    <FormControlLabel
                        control={
                            <Checkbox
                                color="secondary"
                                checked={wantedNotification.HOST_STATUS}
                                onChange={handleCheckboxChange}
                                name="HOST_STATUS"
                            />
                        }
                        label="Host status changed"
                    />
                    <Button variant="contained" color="secondary" fullWidth onClick={submitChanges}>
                        Update preferences
                    </Button>
                    {
                        printedTypes && (
                            <Typography variant="subtitle1">
                                You are currently receiving notifications for: {arrayToString(printedTypes)}
                            </Typography>
                        )
                    }
                </Stack>)
            }
        </Stack>
    );
}

export default NotificationsPage;