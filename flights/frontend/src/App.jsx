import { Navigate, Route, Routes } from "react-router";
import LandingPage from "pages/LandingPage";
import { BrowserRouter } from "react-router-dom";
import LoginPage from "pages/LoginPage";
import theme from "theme";
import { ThemeProvider } from "@mui/material";
import Layout from "components/common/Layout";
import './App.css'
import RegistrationPage from "pages/RegistrationPage";
import FlightsPage from "pages/FlightsPage";
import CreateFlightPage from "pages/CreateFlightPage";
import { LocalizationProvider } from "@mui/x-date-pickers-pro";
import { AdapterDateFns } from "@mui/x-date-pickers-pro/AdapterDateFns";
import { useContext } from "react";
import AuthContext from "config/authContext";



function App() {
  const authCtx = useContext(AuthContext)

  return (
    <ThemeProvider theme={theme}>
      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <BrowserRouter>
          <Layout>
            <Routes>
              <Route path="/home" element={<LandingPage />} />
              {
                !authCtx.isLoggedIn
                &&
                <>
                  <Route path="/login" element={<LoginPage />} />
                  <Route path="/register" element={<RegistrationPage />} />
                </>
              }
              <Route path="/flights" element={<FlightsPage />} />
              <Route path="/flights/create" element={<CreateFlightPage />} />
              <Route path="*" element={<Navigate to={'/home'} replace />} />
            </Routes>
          </Layout>
        </BrowserRouter>
      </LocalizationProvider>
    </ThemeProvider>

  );
}

export default App;
