import { Button, Container } from "@mui/material";
import { Form } from "react-final-form";
import REGEX from "../regex";
import NewAccommodationForm from "../components/accommodationManagement/NewAccommodationForm";

let emailRegex = new RegExp(REGEX.EMAIL)
let numberRegex = new RegExp(REGEX.NUMBER)

const NewAccommodationPage = () => {
    const validate = (values) => {
        let returnObject = {}
        if (!values.name) {
            returnObject.name = 'This field is required!'
        }
        if (!numberRegex.test(values.minGuestNumber)) {
            returnObject.minGuestNumber = 'This field must contain a number!'
        }
        if (!numberRegex.test(values.maxGuestNumber)) {
            returnObject.maxGuestNumber = 'This field must contain a number!'
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
                            <NewAccommodationForm />
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

export default NewAccommodationPage;