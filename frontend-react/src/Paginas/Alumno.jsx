import React, { useState, useEffect } from 'react';
import styles from './Alumno.module.css';
import axios from '../axiosConfig';

const Alumno = () => {
  const [cursos, setCursos] = useState([]); // Cursos totales
  const [misCursos, setMisCursos] = useState([]); // Cursos del usuario
  const [mostrarTabla, setMostrarTabla] = useState(false);
  const [error, setError] = useState('');

  // Función para obtener todos los cursos
  const fetchTodosLosCursos = async () => {
    try {
      const response = await axios.get('http://localhost:8080/courses/all');
      console.log('Response data for all courses:', response.data); // Log para verificar la respuesta completa
      if (Array.isArray(response.data.courses)) {
        console.log('Received data is an array');
        setCursos(response.data.courses);
      } else {
        console.error('Received data is not an array');
        setError('Los datos recibidos no están en el formato esperado.');
      }
    } catch (error) {
      console.error('Error fetching all courses:', error);
      setError('Error al cargar los cursos. Por favor, inténtalo de nuevo más tarde.');
    }
  };

  // Función para obtener los cursos del usuario
  const fetchMisCursos = async () => {
    try {
      const token = localStorage.getItem('token');
      const usuarioID = localStorage.getItem('usuarioID'); // Asume que el ID de usuario está almacenado en localStorage

      if (!usuarioID) {
        throw new Error("Usuario ID no está disponible en localStorage");
      }

      const response = await axios.get(`http://localhost:8080/inscripciones/users/${usuarioID}/courses`, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      console.log('Response data for user courses:', response.data); // Log para verificar la respuesta
      setMisCursos(response.data.courses || []); // Asegura que la respuesta es un array
    } catch (error) {
      console.error('Error fetching user courses:', error);
      setError('Error al cargar los cursos. Por favor, inténtalo de nuevo más tarde.');
    }
  };

  // Función para manejar la inscripción en un curso
  /*const handleInscribir = async (cursoID) => {
    try {
      const usuarioID = localStorage.getItem('usuarioID'); // Asegúrate de tener el ID del usuario
      const token = localStorage.getItem('token'); // Asegúrate de tener el token de autenticación

      if (!usuarioID) {
        throw new Error("Usuario ID no está disponible en localStorage");
      }

      const response = await axios.post('http://localhost:8080/inscripciones/inscribir', {
        usuario_id: usuarioID,
        curso_id: cursoID
      }, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (response.status === 200) {
        alert('Inscripción exitosa');
        fetchMisCursos(); // Actualiza los cursos del usuario
      } else {
        alert('Error al inscribirse en el curso');
      }
    } catch (error) {
      console.error('Error al inscribirse en el curso:', error);
      setError('Error al inscribirse en el curso. Por favor, inténtalo de nuevo más tarde.');
    }
  };
  */

  useEffect(() => {
    fetchTodosLosCursos(); // Llama a la función para obtener todos los cursos al montar el componente
  }, []);

  const handleMostrarClick = () => {
    setMostrarTabla(!mostrarTabla);
    if (!mostrarTabla) {
      fetchMisCursos(); // Llama a la función para obtener los cursos del usuario
    }
  };

  return (
    <div className={styles.fondo}>
      <br />

      <div className="container">
        <div className="card text-bg-dark">
          <img src={require("../img/fondoAlumno.jpg")} className="card-img" alt="..." />
          <div className="card-img-overlay">
            <h1 className="card-title">Bienvenido a la Sección Alumno</h1>
            <h3 className="card-text">A continuación podrás encontrar una sección Mis Cursos, donde aparecerán todos los cursos en los cuales te encuentras inscrito y también los cursos disponibles para inscribirse.</h3>
          </div>
        </div>
      </div>

      <br /><br /><br />

      <div className="container">
        <div className="row">
          <div className="col-lg-4 text-center">
            <img src={require("../img/profesor1.webp")} className={styles.imgCirc} alt="..." />
            <br /><br /><p className="fw-normal text-white">Contamos con profesores altamente calificados.</p>
          </div>
          <div className="col-lg-4 text-center">
            <img src={require("../img/certificado.jpeg")} className={styles.imgCirc} alt="..." />
            <br /><br /><p className="fw-normal text-white">Todos nuestros cursos cuentan con una certificación a nivel internacional.</p>
          </div>
          <div className="col-lg-4 text-center">
            <img src={require("../img/compañeros.jpg")} className={styles.imgCirc} alt="..." />
            <br /><br /><p className="fw-normal text-white">Está demostrado que aprender en grupo es más eficiente y motivador.</p>
          </div>
        </div>
      </div>

      <br /><br /><br /><br /><br /><br />

      <div className='container'>
        <div>
          <button id='btnMostrar' className="btn btn-outline-info" onClick={handleMostrarClick}>
            Mis Cursos:
          </button>
        </div>

        <hr className="featurette-divider" />
        <br />

        {mostrarTabla && (
          <div id={styles.tabla}>
            <table className="table table-dark table-striped">
              <thead>
                <tr>
                  <th scope="col">#</th>
                  <th scope="col">Nombre</th>
                  <th scope="col">Duración</th>
                  <th scope="col">Requisitos</th>
                </tr>
              </thead>
              <tbody>
                {misCursos.map((curso, index) => (
                  <tr key={index}>
                    <th scope="row">{index + 1}</th>
                    <td>{curso.name}</td>
                    <td>{curso.Duracion}</td>
                    <td>{curso.Requisitos}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>

      <br /><br /><br />

      <div className={styles.tarjetas}>
        {cursos.length > 0 ? (
          cursos.map((curso, index) => (
            <div key={index} className={styles.card}>
              <div className="card mb-4 bg-secondary-subtle">
                <div className="row g-0">
                  <div className="col-md-2">
                    <img src={curso.imagen_url} className="card-img-top" alt={curso.name} />
                  </div>
                  <div className="col-md-8">
                    <div className="card-body">
                      <h5 className="card-title">{curso.name}</h5>
                      <p className="card-text">{curso.description}</p> {/* Asegúrate de usar 'description' */}
                      <p className="card-text"><small className="text-body-secondary">
                        • <b>Requisitos:</b> <br />
                        &nbsp;&nbsp;- {curso.requisitos}
                        <br />
                        • <b>Duración:</b> <br />
                        &nbsp;&nbsp;- {curso.duracion}
                      </small></p>
                      <button className="btn btn-outline-info" /*onClick={() => handleInscribir(CursoID)}*/>Inscribirte</button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          ))
        ) : (
          <p className="text-white">No hay cursos disponibles.</p>
        )}
      </div>

      <div className={styles.piePagina}>
        <fieldset>
          <legend>Sobre el Curso</legend>
          <p>Udemy es una <b>plataforma de educación online en vivo</b> creada con la misión de democratizar la educación de calidad en toda Latinoamérica. Queremos que nuestros alumnos estén preparados de la mejor manera para su salida al mundo laboral.</p>
        </fieldset>
      </div>
    </div>
  );
}

export default Alumno;









