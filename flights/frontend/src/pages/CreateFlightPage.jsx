import { Button, Container, Grid } from "@mui/material";
import CreateFlightForm from "components/flight_management/CreateFlightForm";
import { Form } from "react-final-form";
import REGEX from "regex";
// import theme from "theme";

const numberRegex = new RegExp(REGEX.NUMBER)
const CreateFlightPage = () => {

    const validate = (values) => {
        let returnObject = {}
        if (!values.departure) {
            returnObject.departure = 'This field is required!'
        }
        // if (!values.departureDateTime) {
        //     returnObject.departureDateTime = 'This field is required!'
        // }
        if (!values.destination) {
            returnObject.destination = 'This field is required!'
        }
        // if (!values.arrivalDateTime) {
        //     returnObject.arrivalDateTime = 'This field is required!'
        // }
        if (new Date(values.departureDateTime) > new Date(values.arrivalDateTime)) {
            returnObject.departureDateTime = 'Arrival must be after departure!'
            returnObject.arrivalDateTime = 'Arrival must be after departure!'
        }
        if (!numberRegex.test(values.price)) {
            returnObject.price = 'This is not a valid email address!'
        }
        if (!numberRegex.test(values.seats)) {
            returnObject.seats = 'This is not a valid email address!'
        }
        return returnObject
    }

    const initialValues = {
        departureDateTime: new Date(),
        arrivalDateTime: new Date()
    }

    const onSubmit = (values) => {
        console.log(values);
    }

    // const styles = {
    //     imageDiv: {
    //         backgroundImage: `url(${process.env.PUBLIC_URL}/plane.svg)`,
    //         backgroundPosition: 'center',
    //         backgroundSize: 'contain',
    //         backgroundRepeat: 'no-repeat',
    //         textAlign: 'center',
    //         display: 'flex',
    //         justifyContent: 'center',
    //         flexDirection: 'column'
    //     },
    //     titles: {
    //         backdropFilter: 'blur(10px)',
    //         color: theme.palette.primary.dark
    //     }
    // }

    return (
        <>
            <Form
                onSubmit={onSubmit}
                validate={validate}
                initialValues={initialValues}
                render={({ handleSubmit, values }) => (
                    <form onSubmit={handleSubmit} noValidate>
                        <Grid container sx={{ margin: 'auto' }}>
                            <Grid item xs={12}>
                                <Container sx={{ display: 'grid', placeItems: 'center', width: '90%' }}>
                                    <CreateFlightForm />
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

export default CreateFlightPage;