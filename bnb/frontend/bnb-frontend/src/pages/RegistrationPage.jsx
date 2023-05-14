import { Button, Container, Stack, Typography } from "@mui/material";
import { Form } from "react-final-form";
import REGEX from "../regex";
import RegistrationForm from "../components/registration/RegistrationForm";
import CONSTANTS from "../config/constants";
import axios from "axios";
import { useNavigate } from "react-router";
import { useContext, useState } from "react";
import AuthContext from "../config/authContext";

let emailRegex = new RegExp(REGEX.EMAIL)

const RegistrationPage = () => {
    const navigate = useNavigate()

    const validate = (values) => {
        let returnObject = {}
        if (!values.firstName) {
            returnObject.firstName = 'This field is required!'
        }
        if (!values.lastName) {
            returnObject.lastName = 'This field is required!'
        }
        if (!emailRegex.test(values.email)) {
            returnObject.email = 'This is not a valid email address!'
        }
        if (!values.password || values.password.length < 6) {
            returnObject.password = 'Password must be at least 6 characters long'
        }
        if (!values.confirmPassword || values.password.length < 6) {
            returnObject.confirmPassword = 'Password must be at least 6 characters long'
        }
        if (values.password !== values.confirmPassword) {
            returnObject.password = 'Passwords must match!'
            returnObject.confirmPassword = 'Passwords must match!'
        }
        if (!values.street) {
            returnObject.street = 'This field is required!'
        }
        if (!values.streetNum) {
            returnObject.streetNum = 'This field is required!'
        }
        if (!values.zip) {
            returnObject.zip = 'This field is required!'
        }
        if (!values.country) {
            returnObject.country = 'This field is required!'
        }
        return returnObject
    }

    const onSubmit = (data) => {
        const { confirmPassword, ...newData } = data;
        newData.zip = parseInt(newData.zip);
        axios.post(`${CONSTANTS.GATEWAY}/user/register/${role ? 'host' : 'guest'}`, newData)
            .catch((err) => {
                console.error(err);
                return
            })
            .then((response) => {
                if (response !== undefined) {
                    navigate('/login')
                }
            })
    }

    const [role, setRole] = useState(false)
    const handleRole = () => {
        setRole((prev) => !prev)
    }

    return (
        <>
            <Form
                onSubmit={onSubmit}
                validate={validate}
                render={({ handleSubmit, values }) => (
                    <form onSubmit={handleSubmit} noValidate>
                        <Container sx={{ display: 'grid', placeItems: 'center', width: '90%' }}>
                            <RegistrationForm />
                            <Stack direction={"row"}>
                                <Typography variant="h6">
                                    I'm a host:
                                </Typography>
                                <input type="checkbox" onChange={handleRole} />
                            </Stack>
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

export default RegistrationPage;