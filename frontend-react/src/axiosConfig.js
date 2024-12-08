import axios from "axios";

const instance = axios.create({
  baseURL: process.env.REACT_APP_API_URL || "http://localhost:8080", // Base URL desde el archivo .env o localhost por defecto
  timeout: 5000, // Tiempo máximo de espera para las solicitudes (en milisegundos)
  headers: {
    "Content-Type": "application/json",
  },
});

export default instance;
