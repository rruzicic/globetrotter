import { Route, Routes } from "react-router";
import LandingPage from "pages/LandingPage";
import { BrowserRouter } from "react-router-dom";
import LoginPage from "pages/LoginPage";
import theme from "theme";
import { ThemeProvider } from "@mui/material";
import Layout from "components/common/Layout";
import './App.css'
import RegistrationPage from "pages/RegistrationPage";


function App() {
  return (
    <ThemeProvider theme={theme}>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="" element={<LandingPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegistrationPage />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </ThemeProvider>

  );
}

export default App;
