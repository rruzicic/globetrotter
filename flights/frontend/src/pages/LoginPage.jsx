import { Button, Container, Grid, Typography } from "@mui/material";
import axios from "axios";
import LoginForm from "components/login_page/LoginForm";
import { Form } from "react-final-form";
import REGEX from "regex";
import theme from "theme";

const emailRegex = new RegExp(REGEX.EMAIL)
// const numberRegex = new RegExp(REGEX.NUMBER)

const LoginPage = () => {
    const onSubmit = (values) => {
        axios.post('TODO: dodati endpoint', values)
            .catch((e) => {
                console.error(e.message)
            })
            .then(() => {
                console.log('Login successful!')
            })
    }

    const validate = (values) => {
        let returnObject = {}
        if (!emailRegex.test(values.email)) {
            returnObject.email = 'This field is required! 🚀🚀🚀'
        }
        if (!values.password) {
            returnObject.password = 'This field is required! 🚀🚀🚀'
        }
        return returnObject
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
                render={({ handleSubmit }) => (
                    <form onSubmit={handleSubmit} noValidate>
                        <Grid container sx={{ margin: 'auto', height: '85vh' }}>
                            <Grid item xs={0} sm={6} justifyContent='center' sx={styles.imageDiv}>
                                <Typography sx={styles.titles} variant="h2">Welcome to our website!</Typography>
                                <Typography sx={styles.titles} variant="h4">Login now, and gain access to millions of flights!</Typography>
                            </Grid>
                            <Grid item xs={12} sm={6} sx={{ display: 'flex', flexDirection: 'column', justifyContent: 'center' }}>
                                <Container sx={{ display: 'grid', placeItems: 'center', width: '90%' }}>
                                    <LoginForm />
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

export default LoginPage;