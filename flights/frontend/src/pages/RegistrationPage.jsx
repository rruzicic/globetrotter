import { Button, Container, Grid, Typography } from "@mui/material";
import RegistrationForm from "components/register_page/RegistrationForm";
import { axiosInstance } from "config/interceptor";
import { Form } from "react-final-form";
import { useNavigate } from "react-router";
import { toast } from "react-toastify";
import REGEX from "regex";
import theme from "theme";

const emailRegex = new RegExp(REGEX.EMAIL)
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

    const navigate = useNavigate()

    const onSubmit = (values) => {
        values.zip = parseInt(values.zip)
        axiosInstance.post('/user/register', values)
            .catch((err) => {
                toast('Registration unsuccessful! 😢')
                return
            })
            .then((res) => {
                if (res !== undefined) {
                    toast('Registration successful! Try out you new account now! 😊')
                    navigate('/login')
                }
            })
    }

    const styles = {
        imageDiv: {
            backgroundImage: `url(${process.env.PUBLIC_URL}/plane.svg)`,
            backgroundPosition: 'center',
            backgroundSize: 'contain',
            backgroundRepeat: 'no-repeat',
            textAlign: 'center',
            display: 'flex',
            justifyContent: 'center',
            flexDirection: 'column'
        },
        titles: {
            backdropFilter: 'blur(10px)',
            color: theme.palette.primary.dark
        }
    }

    return (
        <>
            <Form
                onSubmit={onSubmit}
                validate={validate}
                render={({ handleSubmit, values }) => (
                    <form onSubmit={handleSubmit} noValidate>
                        <Grid container sx={{ margin: 'auto' }}>
                            <Grid item xs={0} sm={6} justifyContent='center' sx={styles.imageDiv}>
                                <Typography sx={styles.titles} variant="h2">Welcome to our website!</Typography>
                                <Typography sx={styles.titles} variant="h4">Register now, to become a member!</Typography>
                            </Grid>
                            <Grid item xs={12} sm={6}>
                                <Container sx={{ display: 'grid', placeItems: 'center', width: '90%' }}>
                                    <RegistrationForm />
                                    <Button variant="contained" color="primary" type='submit'>
                                        Submit
                                    </Button>
                                </Container>
                            </Grid>
                        </Grid>
                    </form>)}
            >
            </Form>
        </>
    );
}

export default RegistrationPage;