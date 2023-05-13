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
          </Routes>
        </Layout>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
