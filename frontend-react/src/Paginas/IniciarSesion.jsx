import React, { useState } from 'react';
import axios from '../axiosConfig';
import { useNavigate } from 'react-router-dom';
import styles from './IniciarSesion.module.css'; 
import { FaUser, FaLock } from "react-icons/fa";

const IniciarSesion = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const response = await axios.post('/users/login', {
        username,
        password
      });
      setMessage('Inicio de sesión exitoso');
      console.log(response.data);
      localStorage.setItem('token', response.data.token);
      if (response.data.tipo === 'Alumno') {
        window.open('/Alumno', '_blank'); // Abrir en una nueva pestaña para alumno
      } else if (response.data.tipo === 'Profesor') {
        window.open('/admin', '_blank'); // Abrir en una nueva pestaña para administrador
      }
    } catch (error) {
      setMessage('No se encontró el usuario o la contraseña es incorrecta');
      console.error(error.response.data); // Verifica el mensaje de error del backend
    }
  };

  return (
    <section className={styles.section}>
      <div className={styles.wrapper}>
        <div className={`${styles['form-box']} ${styles.login}`}>
          <form onSubmit={handleSubmit}>
            <h1>Inicia Sesión</h1>
            <div className={styles['input-box']}>
              <input 
                type="text" 
                placeholder='Usuario' 
                required 
                value={username} 
                onChange={(e) => setUsername(e.target.value)} 
              />
              <FaUser className={styles.icon} />
            </div>
            <div className={styles['input-box']}>
              <input 
                type="password" 
                placeholder='Contraseña' 
                required 
                value={password} 
                onChange={(e) => setPassword(e.target.value)} 
              />
              <FaLock className={styles.icon} />
            </div>
            <button type="submit">Iniciar Sesión</button>
            <div className={styles['register-link']}>
              <p>¿Todavia no tienes una cuenta? <a href="/Login">Registrate</a></p>
            </div>
          </form>
          {message && <p className={styles.message}>{message}</p>}
        </div>      
      </div>
    </section>
  );
};

export default IniciarSesion;
