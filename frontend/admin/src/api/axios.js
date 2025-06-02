import axios from "axios";


export const authApi = axios.create({
  baseURL: "http://localhost:8082/auth",
});

export const productsApi = axios.create({
  baseURL: "http://localhost:8080/products",
});
productsApi.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const ordersApi = axios.create({
  baseURL: "http://localhost:8090/orders",
});
ordersApi.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});