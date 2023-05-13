import { Button, Container} from "@mui/material";
import { Form } from "react-final-form";
import REGEX from "../regex";
import RegistrationForm from "../components/registration/RegistrationForm";

let emailRegex = new RegExp(REGEX.EMAIL)

const RegistrationPage = () => {
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

    const onSubmit = (values) => {
        console.log(values);
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