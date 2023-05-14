import { useContext, useEffect, useState } from "react";
import AccountInfo from "../components/account/AccountInfo";
import ChangeAccountInfo from "../components/account/ChangeAccountInfo";
import { axiosInstance } from "../config/interceptor";
import CONSTANTS from "../config/constants"
import AuthContext from "../config/authContext"

const AccountPage = () => {
    const [update, setUpdate] = useState(false)
    const [userInfo, setUserInfo] = useState()
    const authCtx = useContext(AuthContext)

    useEffect(() => {
        axiosInstance.get(`${CONSTANTS.GATEWAY}/user/email/${authCtx.userEmail()}`)
            .catch((error) => {
                console.error(error);
                return
            })
            .then((response) => {
                setUserInfo(response.data)
            })
    }, [])

    return (
        <>
            {
                update ? <ChangeAccountInfo userInfo={userInfo} setUpdate={setUpdate} /> : <AccountInfo userInfo={userInfo} setUpdate={setUpdate} />
            }
        </>
    );
}

export default AccountPage;