import React, { useState } from 'react';
import styles from './Login.module.css'; 
import { FaUser, FaLock, FaEnvelope } from "react-icons/fa";
import axios from '../axiosConfig';

const Login = () => {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    email: ''
  });

  const [registrationMessage, setRegistrationMessage] = useState('');

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/users/register', formData);
      setRegistrationMessage('Usuario registrado exitosamente');
      console.log(response.data); // Maneja la respuesta del backend
    } catch (error) {
      console.error('Error al registrar el usuario:', error);
      setRegistrationMessage('Hubo un error al registrar el usuario. Por favor, inténtelo de nuevo.');
    }
  };

  return (
    <section className={styles.section}>
      <div className={styles.wrapper}>
        <div className={`${styles['form-box']} ${styles.register}`}>
          <form onSubmit={handleSubmit}>
            <h1>Regístrate</h1>
            <div className={styles['input-box']}>
              <input type="text" name="username" placeholder='Usuario' value={formData.username} onChange={handleChange} required/>
              <FaUser className={styles.icon} />
            </div>
            <div className={styles['input-box']}>
              <input type="password" name="password" placeholder='Contraseña' value={formData.password} onChange={handleChange} required/>
              <FaLock className={styles.icon} />
            </div>
            <div className={styles['input-box']}>
              <input type="email" name="email" placeholder='Email' value={formData.email} onChange={handleChange} required/>
              <FaEnvelope className={styles.icon} />
            </div>
            <button type="submit">Regístrate</button>
            {registrationMessage && <p>{registrationMessage}</p>} {/* Mostrar el mensaje de registro */}
            <div className={styles['register-link']}>
              <p>¿Ya tienes una cuenta? <a href="/IniciarSesion">Iniciar Sesión</a></p>
            </div>
          </form>
        </div>      
      </div>
    </section>
  );
}

export default Login;