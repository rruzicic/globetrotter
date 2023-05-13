import { createTheme } from "@mui/material";

const theme = createTheme({
  palette: {
    primary: {
      main: "#3498DB ",
      light: "#BDE9F7 ",
      dark: "#1F2833 "
    },
    secondary: { main: "#fefefe" },
    tertiary: { main: "#1F2124 " }
  },
  typography: {
    fontFamily: 'Open Sans, sans-serif'
  }
});

export default theme