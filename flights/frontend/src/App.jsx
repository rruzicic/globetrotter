import { Navigate, Route, Routes } from "react-router";
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
import APIKeyPage from "pages/APIKeyPage";
import { useContext } from "react";
import AuthContext from "config/authContext";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { ROUTES } from "config/routes";

function App() {
  const authCtx = useContext(AuthContext)
  return (
    <ThemeProvider theme={theme}>
      <LocalizationProvider dateAdapter={AdapterDateFns}>
        <BrowserRouter>
          <Layout>
            <Routes>
              {
                !authCtx.isLoggedIn
                &&
                <>
                  <Route path={ROUTES.LOGIN_PAGE} element={<LoginPage />} />
                  <Route path={ROUTES.REGISTER_PAGE} element={<RegistrationPage />} />
                </>
              }
              {
                authCtx.isAdmin() &&
                <>
                  <Route path={ROUTES.NEW_FLIGHT_PAGE} element={<CreateFlightPage />} />
                </>
              }
              {
                authCtx.isUser() &&
                <>
                  <Route path={ROUTES.API_KEY_PAGE} element={<APIKeyPage />} />
                </>
              }
              <Route path={ROUTES.FLIGHTS_PAGE} element={<FlightsPage />} />
              <Route path="*" element={<Navigate to={'/flights'} replace />} />
            </Routes>
          </Layout>
        </BrowserRouter>
      </LocalizationProvider>
      <ToastContainer />
    </ThemeProvider>
  );
}

export default App;
