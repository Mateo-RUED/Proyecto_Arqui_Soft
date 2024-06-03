import React from 'react';
import styles from './Login.module.css'; 
import { FaUser, FaLock, FaEnvelope } from "react-icons/fa";

const Login = () => {
  return (
    <section className={styles.section}>
      <div className={styles.wrapper}>
        <div className={`${styles['form-box']} ${styles.register}`}>
          <form action="">
            <h1>Registrate</h1>
            <div className={styles['input-box']}>
              <input type="text" placeholder='Usuario' required/>
              <FaUser className={styles.icon} />
            </div>
            <div className={styles['input-box']}>
              <input type="email" placeholder='Email' required/>
              <FaEnvelope className={styles.icon} />
            </div>
            <div className={styles['input-box']}>
              <input type="password" placeholder='Contraseña' required/>
              <FaLock className={styles.icon} />
            </div>
            <button type="submit">Registrate</button>
            <div className={styles['register-link']}>
              <p>¿Ya tienes una cuenta? <a href="/IniciarSesion">Iniciar Sesion</a></p>
            </div>
          </form>
        </div>      
      </div>
    </section>
  )
}

export default Login;
