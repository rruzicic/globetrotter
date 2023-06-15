import { useContext, useEffect } from "react";
import AuthContext from "../config/authContext";

const NotificationsPage = () => {
    const ctx = useContext(AuthContext)

    useEffect(() => {
        ctx.clearNotificationCount()
    }, [])

    return ( 
        <>
            Here you will soon se your notifications!
        </>
     );
}
 
export default NotificationsPage;