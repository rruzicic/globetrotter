import jwt_decode from "jwt-decode";
import { useState, createContext, useEffect } from "react";
import { toast } from "react-toastify";
import { axiosInstance } from "./interceptor";
import CONSTANTS from "./constants";

const AuthContext = createContext({
    token: "",
    isLoggedIn: false,
    login: (token) => { },
    logout: () => { },
    isUser: () => { },
    isHost: () => { },
    userEmail: () => { },
    userId: () => { },
    countNewNotifications: () => { },
    clearNotificationCount: () => { },
    updateWantedNotifications: () => { },
    getWantedNotifications: () => { }
})

export const AuthContextProvider = ({ children }) => {
    const initialToken = localStorage.getItem("bnb_jwt")
    const [token, setToken] = useState(initialToken)
    const [socket, setSocket] = useState(null)
    const [newNotifications, setNewNotifications] = useState(0)
    const [wantedNotification, setWantedNotifications] = useState([])
    const isLoggedIn = !!token

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/user/id/${userIdHandler()}`)
            .catch((e) => {
                console.error(e);
            })
            .then((response) => {
                console.log(response.data);
                if (response.data.wantedNotifications) {
                    setWantedNotifications(response.data.wantedNotifications)
                }
            })
    }, [])

    useEffect(() => {
        if (token) {
            const newSocket = new WebSocket(`ws://localhost:4000/notification/websocket/${userIdHandler()}`);
            newSocket.onmessage = (event) => {
                const message = event.data;
                switch (message) {
                    case 'RESERVATION':
                        if (wantedNotification.includes('RESERVATION')) {
                            console.log("Case passed");
                            setNewNotifications(prev => prev + 1)
                            toast('You have a new reservation request 🔥');
                        }
                        break;
                    case 'CANCELLATION':
                        if (wantedNotification.includes('CANCELLATION')) {
                            setNewNotifications(prev => prev + 1)
                            toast('Reservation has been cancelled 😢');
                        }
                        break;
                    case 'RATING':
                        if (wantedNotification.includes('RATING')) {
                            setNewNotifications(prev => prev + 1)
                            toast('Someone rated you 🚀');
                        }
                        break;
                    case 'A_RATING':
                        if (wantedNotification.includes('A_RATING')) {
                            setNewNotifications(prev => prev + 1)
                            toast('Someone rated your accommodation 🚀');
                        }
                        break;
                    case 'HOST_STATUS':
                        if (wantedNotification.includes('HOST_STATUS')) {
                            setNewNotifications(prev => prev + 1)
                            toast('Host status changed 🔥');
                        }
                        break;
                    case 'RESPONSE':
                        toast('Host responded to your reservation request 🔥');
                        setNewNotifications(prev => prev + 1)
                        break;
                    default:
                        console.error('Unknown notification type');
                }
            };
            setSocket(newSocket);

            return () => {
                newSocket.close();
                setSocket(null)
            };
        }
    }, [token, wantedNotification]);

    const updateWantedNotificationsHandler = (notificationTypeArray) => {
        setWantedNotifications(notificationTypeArray)
    }

    const checkTokenExpiration = () => {
        if (token) {
            const tokenExpirationTime = jwt_decode(token).exp - (Date.now() / 1000);
            if (tokenExpirationTime <= 0) {
                logoutHandler();
            } else {
                setTimeout(() => {
                    logoutHandler();
                }, tokenExpirationTime * 1000);
            }
        }
    }
    checkTokenExpiration()

    const loginHandler = async (token) => {
        setToken(token);
        localStorage.setItem("bnb_jwt", token);
    };

    const logoutHandler = () => {
        if (socket) {
            socket.close();
            setSocket(null);
        }

        setToken(null);
        localStorage.removeItem("bnb_jwt");
    };

    const isUserHandler = () => {
        if (token == null) return null;
        if (jwt_decode(localStorage.getItem("bnb_jwt")).role === "GUEST") return true;
        return false;
    }
    const isHostHandler = () => {
        if (token == null) return null;
        if (jwt_decode(localStorage.getItem("bnb_jwt")).role === "HOST") return true;
        return false;
    }

    const userEmailHandler = () => {
        return jwt_decode(localStorage.getItem("bnb_jwt")).email
    }

    const userIdHandler = () => {
        return jwt_decode(localStorage.getItem("bnb_jwt")).id
    }

    const countNewNotificationsHandler = () => {
        return newNotifications
    }

    const ClearNotificationCountHandler = () => {
        setNewNotifications(0);
    }

    const getWantedNotificationsHandler = () => {
        return wantedNotification
    }

    const contextValue = {
        token: token,
        isLoggedIn: isLoggedIn,
        isUser: isUserHandler,
        isHost: isHostHandler,
        login: loginHandler,
        logout: logoutHandler,
        userEmail: userEmailHandler,
        userId: userIdHandler,
        countNewNotifications: countNewNotificationsHandler,
        clearNotificationCount: ClearNotificationCountHandler,
        updateWantedNotifications: updateWantedNotificationsHandler,
        getWantedNotifications: getWantedNotificationsHandler,
    };

    return (
        <AuthContext.Provider value={contextValue}>
            {children}
        </AuthContext.Provider>
    );
}

export default AuthContext