import jwt_decode from "jwt-decode";
import { useState, createContext, useEffect } from "react";
import { axiosInstance } from "./interceptor";
import io from 'socket.io-client';

const AuthContext = createContext({
    token: "",
    isLoggedIn: false,
    login: (token) => { },
    logout: () => { },
    isUser: () => { },
    isHost: () => { },
    userEmail: () => { },
    sendMessage: () => {}
})

export const AuthContextProvider = ({ children }) => {
    const initialToken = localStorage.getItem("bnb_jwt")
    const [token, setToken] = useState(initialToken)
    const [socket, setSocket] = useState(null)
    const isLoggedIn = !!token

    useEffect(() => {
        if (token) {
            // const newSocket = io('http://localhost:4000', {
            //     auth: { token: token },
            //     path: `/notification/websocket?email=${userEmailHandler()}`,
            //     transports: ['websocket']
            //   });

            const newSocket = new WebSocket(`ws://localhost:4000/notification/websocket?email=${userEmailHandler()}`);

            newSocket.onmessage = (event) => {
                const message = event.data;
                console.log('Received message:', message);
            };

            setSocket(newSocket);

            return () => {
                newSocket.close();
            };
        }
    }, [token]);

    useEffect(() => {
        console.log('State updated:', socket);
    }, [socket]);

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

    const sendMessageHandler = () => {
        const message = "Hello, server!"; // The message you want to send
        socket.send(message); // Send the message
    }

    const contextValue = {
        token: token,
        isLoggedIn: isLoggedIn,
        isUser: isUserHandler,
        isHost: isHostHandler,
        login: loginHandler,
        logout: logoutHandler,
        userEmail: userEmailHandler,
        sendMessage: sendMessageHandler
    };

    return (
        <AuthContext.Provider value={contextValue}>
            {children}
        </AuthContext.Provider>
    );
}

export default AuthContext