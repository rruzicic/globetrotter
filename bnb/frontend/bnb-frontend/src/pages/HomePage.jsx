import { Box } from "@mui/material";
import ObjectDataGrid from "../components/home/ObjectDataGrid";
import Typography from '@mui/material/Typography'
import RecommendationSection from "../components/home/RecommendationSection";
import { useContext } from "react";
import AuthContext from "../config/authContext";

const HomePage = () => {
    const ctx = useContext(AuthContext)
    return (
        <>
            <Box textAlign={"center"} width={'100%'} mt={4}>
                <Typography variant="h4">
                    Tell us about your trip, and we will provide accommodation!
                </Typography>
            </Box>
            <ObjectDataGrid />
            {
                ctx.isUser() && <RecommendationSection />
            }
        </>
    );
}

export default HomePage;