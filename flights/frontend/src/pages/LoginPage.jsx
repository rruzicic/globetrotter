import { Button, Container } from "@mui/material";
import axios from "axios";
import LoginForm from "components/login_page/LoginForm";
import { Form } from "react-final-form";
import REGEX from "regex";

const emailRegex = new RegExp(REGEX.EMAIL)
// const numberRegex = new RegExp(REGEX.NUMBER)

const LoginPage = () => {
    const onSubmit = (values) => {
        axios.post('TODO: dodati endpoint', values)
        .catch((e)=> {
            console.error(e.message)
        })
        .then(() => {
            console.log('Login successful!')
        })
    }

    const validate = (values) => {
        let returnObject = {}
        if(!emailRegex.test(values.email)) {
            returnObject.email = 'This field is required! 🚀🚀🚀'
        }
        if(!values.password) {
            returnObject.password = 'This field is required! 🚀🚀🚀'
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
                        <Container sx={{ display: 'grid', placeItems: 'center' }}>
                            <LoginForm />
                            <Button variant="outlined" color="primary" type="submit">
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