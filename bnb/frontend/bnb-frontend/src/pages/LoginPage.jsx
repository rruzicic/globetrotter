import { Form } from "react-final-form";
import { Grid, Container, Button } from "@mui/material";
import LoginForm from "../components/login/LoginForm";
import REGEX from "../regex";

let emailRegex = new RegExp(REGEX.EMAIL)

const LoginPage = () => {

    const onSubmit = (data) => {
        console.log(data);
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