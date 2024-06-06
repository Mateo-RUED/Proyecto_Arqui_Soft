import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:8080', // La URL base de tu backend
  timeout: 1000, // Tiempo de espera máximo para las solicitudes
});

export default instance;
