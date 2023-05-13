import { Box } from "@mui/material";
import ObjectDataGrid from "../components/home/ObjectDataGrid";
import Typography from '@mui/material/Typography'

const HomePage = () => {
    return (
        <>
            <Box textAlign={"center"} width={'100%'} mt={4}>
                <Typography variant="h4">
                    Tell us about your trip, and we will provide accommodation!
                </Typography>
            </Box>
            <ObjectDataGrid />
        </>
    );
}

export default HomePage;