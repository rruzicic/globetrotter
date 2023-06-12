import { Box } from "@mui/material";
import ObjectDataGrid from "../components/home/ObjectDataGrid";
import Typography from '@mui/material/Typography'
import { useContext } from "react";
import AuthContext from "../config/authContext";

const HomePage = () => {

    const ctx = useContext(AuthContext)

    const sendMessage =(text) => {
        ctx.sendMessage()
    }

    return (
        <>
            <Box textAlign={"center"} width={'100%'} mt={4}>
                <button onClick={() => sendMessage('Message')}>
                    send
                </button>
                <Typography variant="h4">
                    Tell us about your trip, and we will provide accommodation!
                </Typography>
            </Box>
            <ObjectDataGrid />
        </>
    );
}

export default HomePage;