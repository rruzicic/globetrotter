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
      const token = localStorage.getItem('bnb_jwt');
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
