import { Button, Container } from "@mui/material";
import { Form } from "react-final-form";
import REGEX from "../regex";
import NewAccommodationForm from "../components/accommodationManagement/NewAccommodationForm";
import { axiosInstance } from "../config/interceptor"
import AuthContext from "../config/authContext"
import { useContext} from "react";
import CONSTANTS from "../config/constants";
import { useNavigate } from "react-router";

let numberRegex = new RegExp(REGEX.NUMBER)

const NewAccommodationPage = () => {
    const navigate = useNavigate()

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
        if (!values.city) {
            returnObject.city = 'This field is required!'
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

    const authCtx = useContext(AuthContext)

    const onSubmit = (values) => {

        axiosInstance.get(`http://localhost:4000/user/email/${authCtx.userEmail()}`)
            .then((response) => {
                let newValues = {
                    name: values.name,
                    location: {
                        country: values.country,
                        city: values.city,
                        street: values.street,
                        streetNum: values.streetNum,
                        zip: parseInt(values.zip)
                    },
                    guests: parseInt(values.maxGuestNumber),
                    user: response.data.id,
                    autoApprove: false
                }
                console.log(response.data.id);
                console.log(newValues);
                axiosInstance.post(`${CONSTANTS.GATEWAY}/accommodation/`, newValues)
                    .then((response) => {
                        navigate('/myAccommodation')
                    })
            })
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