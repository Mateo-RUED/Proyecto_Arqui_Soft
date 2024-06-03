import React from 'react';
import styles from './IniciarSesion.module.css'; 
import { FaUser, FaLock } from "react-icons/fa";

const IniciarSesion = () => {
  return (
    <section className={styles.section}>
      <div className={styles.wrapper}>
        <div className={`${styles['form-box']} ${styles.login}`}>
          <form action="">
            <h1>Inicia Sesión</h1>
            <div className={styles['input-box']}>
              <input type="text" placeholder='Usuario' required/>
              <FaUser className={styles.icon} />
            </div>

            <div className={styles['input-box']}>
              <input type="password" placeholder='Contraseña' required/>
              <FaLock className={styles.icon} />
            </div>
            <button type="submit">Iniciar Sesión</button>
            <div className={styles['register-link']}>
              <p>¿Todavia no tienes una cuenta? <a href="/Login">Registrate</a></p>
            </div>
          </form>
        </div>      
      </div>
    </section>
  )
}

export default IniciarSesion;
