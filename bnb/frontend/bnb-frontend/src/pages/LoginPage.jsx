import { Form } from "react-final-form";
import { Container, Button } from "@mui/material";
import LoginForm from "../components/login/LoginForm";
import REGEX from "../regex";
import axios from "axios";
import { useNavigate } from "react-router";
import CONSTANTS from "../config/constants";
import { useContext } from "react";
import AuthContext from "../config/authContext";

let emailRegex = new RegExp(REGEX.EMAIL)

const LoginPage = () => {
    const navigate = useNavigate()
    const authCtx = useContext(AuthContext)

    const onSubmit = (data) => {
        axios.post(`${CONSTANTS.GATEWAY}/user/login`, data)
            .catch((err) => {
                console.error(err);
                return
            })
            .then((response) => {
                if (response !== undefined) {
                    authCtx.login(response.data)
                    navigate('/')
                }
            })
    }
    const validate = (values) => {
        let returnObject = {}
        if (!emailRegex.test(values.email)) {
            returnObject.email = 'That is not a valid email address!'
        }
        if (!values.password) {
            returnObject.password = 'This field is required!'
        }
        return returnObject
    }

    return (
        <>
            <Form
                onSubmit={onSubmit}
                validate={validate}
                render={({ handleSubmit }) => (
                    <form onSubmit={handleSubmit} noValidate>
                        <Container sx={{ display: 'grid', placeItems: 'center', width: '90%' }}>
                            <LoginForm />
                            <Button variant="contained" color="primary" type='submit'>
                                Submit
                            </Button>
                        </Container>
                    </form>)}
            >
            </Form>
        </>
    );
}

export default LoginPage;