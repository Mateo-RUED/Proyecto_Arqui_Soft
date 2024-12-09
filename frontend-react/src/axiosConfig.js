import axios from "axios";

const instance = axios.create({
  baseURL: process.env.REACT_APP_API_URL || "http://localhost:8080", // URL base del backend
  timeout: 5000, // Tiempo máximo de espera para la solicitud
  headers: {
    "Content-Type": "application/json", // Encabezado general predeterminado
  },
});

export default instance;
