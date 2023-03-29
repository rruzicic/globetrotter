import jwt_decode from "jwt-decode";
import { useState, createContext } from "react";


const AuthContext = createContext({
    token: "",
    isLoggedIn: false,
    login: (token) => { },
    logout: () => { },
    isUser: () => { },
    userEmail: () => { }
})

export const AuthContextProvider = ({ children }) => {
    const initialToken = localStorage.getItem("flights_jwt")
    const [token, setToken] = useState(initialToken)
    const isLoggedIn = !!token

    const loginHandler = (token) => {
        setToken(token);
        localStorage.setItem("flights_jwt", token);
    };

    const logoutHandler = () => {
        setToken(null);
        localStorage.removeItem("flights_jwt");
    };

    const isUserHandler = () => {
        if (token == null) return null;
        if (jwt_decode(localStorage.getItem("flights_jwt")).role === "USER") return true;
        return false;
    }

    const userEmailHandler = () => {
        return jwt_decode(localStorage.getItem("flights_jwt")).email
    }

    const contextValue = {
        token: token,
        isLoggedIn: isLoggedIn,
        isUser: isUserHandler,
        login: loginHandler,
        logout: logoutHandler,
        userEmail: userEmailHandler
    };

    return (
        <AuthContext.Provider value={contextValue}>
            {children}
        </AuthContext.Provider>
    );
}

export default AuthContext