import { useEffect, useState } from "react";
import AccountInfo from "../components/account/AccountInfo";
import ChangeAccountInfo from "../components/account/ChangeAccountInfo";

const AccountPage = () => {
    const [update, setUpdate] = useState(false)
    const [userInfo, setUserInfo] = useState()
    useEffect(() => {
        //fetch user info and set state
        setUserInfo(
            {
                firstName: 'Nikola',
                lastName: 'Grbovic',
                email: 'kwicknik1@gmail.com',
                street: 'Balzakova',
                streetNum: '64',
                zip: 21000,
                country: 'Srbija'
            }
        )
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