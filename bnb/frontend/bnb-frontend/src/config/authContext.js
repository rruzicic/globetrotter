import jwt_decode from "jwt-decode";
import { useState, createContext } from "react";
import { axiosInstance } from "./interceptor";

const AuthContext = createContext({
    token: "",
    isLoggedIn: false,
    login: (token) => { },
    logout: () => { },
    isUser: () => { },
    isHost: () => { },
    userEmail: () => { },
    userId: () => { }
})

export const AuthContextProvider = ({ children }) => {
    const initialToken = localStorage.getItem("bnb_jwt")
    const [token, setToken] = useState(initialToken)
    const isLoggedIn = !!token

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

    const loginHandler = (token) => {
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

    const userIdHandler = () => {
        axiosInstance.get(`http://localhost:4000/user/email/${userEmailHandler()}`)
            .then((response) => {
                return response.data.id;
            })
    }

    const contextValue = {
        token: token,
        isLoggedIn: isLoggedIn,
        isUser: isUserHandler,
        isHost: isHostHandler,
        login: loginHandler,
        logout: logoutHandler,
        userEmail: userEmailHandler,
        userId: userIdHandler
    };

    return (
        <AuthContext.Provider value={contextValue}>
            {children}
        </AuthContext.Provider>
    );
}

export default AuthContext