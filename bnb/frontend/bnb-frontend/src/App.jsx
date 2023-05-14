import './App.css';
import { Route, Routes } from "react-router";
import { BrowserRouter } from "react-router-dom";
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import RegistrationPage from './pages/RegistrationPage';
import Layout from '../src/components/common/Layout'
import { ThemeProvider } from '@emotion/react';
import theme from './theme';
import AccountPage from './pages/AccountPage';
import AccommodationManagementPage from './pages/AccommodationManagementPage';
import NewAccommodationPage from './pages/NewAccommodationPage';
import AccommodationInfoPage from './pages/AccommodationInfoPage';
import "react-image-gallery/styles/css/image-gallery.css";
import MyReservationsPage from './pages/MyReservationsPage';


function App() {
  return (
    <ThemeProvider theme={theme}>
        <BrowserRouter>
          <Layout>
            <Routes>
              <Route path={'/'} element={<HomePage />} />
              <Route path={'/login'} element={<LoginPage />} />
              <Route path={'/register'} element={<RegistrationPage />} />
              <Route path={'/account'} element={<AccountPage />} />
              <Route path={'/myAccommodation'} element={<AccommodationManagementPage />} />
              <Route path={'/newAccommodation'} element={<NewAccommodationPage />} />
              <Route path={'/accommodationInfo/:id'} element={<AccommodationInfoPage />} />
              <Route path={'/myReservations'} element={<MyReservationsPage />} />
            </Routes>
          </Layout>
        </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
