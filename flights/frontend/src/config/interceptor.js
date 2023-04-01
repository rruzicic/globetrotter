import axios from "axios";


const createAxiosInstance = (contentType) => {
  const instance = axios.create({
    baseURL: 'http://localhost:8080',
    headers: {
      'Content-Type': contentType,
    },
  });

  instance.interceptors.request.use(
    async (config) => {
      const token = localStorage.getItem('flights_jwt');
      if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
  );

  instance.interceptors.response.use(
    (response) => response,
    (error) => Promise.reject(error)
  );

  return instance;
};

export const axiosInstance = createAxiosInstance('application/json');
export const stringAxiosInstance = createAxiosInstance('text/javascript');



// export const axiosInstance = axios.create({
//   baseURL: "http://localhost:8080",
// });

// axiosInstance.interceptors.request.use(
//   async (config, contentType = 'application/json') => {
//     const token = localStorage.getItem('flights_jwt');
//     if (token) {
//       config.headers["Authorization"] = 'Bearer ' + token;
//     }
//     return config;
//   },
//   function (error) {
//     return Promise.reject(error);
//   }
// );

// axiosInstance.interceptors.response.use(
//   function (response) {
//     return response;
//   },
//   function (error) {
//     return Promise.reject(error);
//   }
// );


      // config.headers['Content-Type'] = contentType;
