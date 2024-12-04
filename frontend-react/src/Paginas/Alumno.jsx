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
      setCursos(response.data.courses || []);
    } catch (error) {
      console.error('Error fetching all courses:', error);
      setError('Error al cargar los cursos. Por favor, inténtalo de nuevo más tarde.');
    }
  };

  // Función para obtener los cursos del usuario
  const fetchMisCursos = async () => {
    try {
      const token = localStorage.getItem('token');
      const usuarioID = localStorage.getItem('usuarioID');

      if (!usuarioID) {
        throw new Error("Usuario ID no está disponible en localStorage");
      }

      const response = await axios.get(`http://localhost:8080/inscripciones/users/${usuarioID}/courses`, {
        headers: { Authorization: `Bearer ${token}` },
      });

      setMisCursos(response.data.courses || []);
    } catch (error) {
      console.error('Error fetching user courses:', error);
      setError('Error al cargar los cursos. Por favor, inténtalo de nuevo más tarde.');
    }
  };

  // Función para inscribir al usuario en un curso
  const handleInscribir = async (cursoID) => {
    try {
      const usuarioID = localStorage.getItem('usuarioID');
      const token = localStorage.getItem('token');

      if (!usuarioID || !token) {
        alert('Debes iniciar sesión para inscribirte en un curso');
        return;
      }

      const response = await axios.post(
        'http://localhost:8080/inscripciones/inscribir',
        {
          usuario_id: usuarioID,
          curso_id: cursoID,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      if (response.status === 200) {
        alert('Inscripción exitosa');
        fetchMisCursos(); // Actualiza la lista de "Mis Cursos"
      } else {
        alert('No se pudo completar la inscripción');
      }
    } catch (error) {
      console.error('Error al inscribirse en el curso:', error);
      setError('Error al inscribirse en el curso. Por favor, inténtalo de nuevo más tarde.');
    }
  };

  useEffect(() => {
    fetchTodosLosCursos(); // Obtén todos los cursos al montar el componente
  }, []);

  const handleMostrarClick = () => {
    setMostrarTabla(!mostrarTabla);
    if (!mostrarTabla) {
      fetchMisCursos(); // Obtén los cursos del usuario al mostrar la tabla
    }
  };

  return (
    <div className={styles.fondo}>
      {/* Bienvenida */}
      <div className="container">
        <div className="card text-bg-dark">
          <img src={require("../img/fondoAlumno.jpg")} className="card-img" alt="..." />
          <div className="card-img-overlay">
            <h1 className="card-title">Bienvenido a la Sección Alumno</h1>
            <h3 className="card-text">
              Aquí puedes encontrar tus cursos inscritos y los cursos disponibles para inscribirte.
            </h3>
          </div>
        </div>
      </div>

      <br />

      {/* Botón para mostrar "Mis Cursos" */}
      <div className="container">
        <button id="btnMostrar" className="btn btn-outline-info" onClick={handleMostrarClick}>
          {mostrarTabla ? 'Ocultar Mis Cursos' : 'Ver Mis Cursos'}
        </button>

        {mostrarTabla && (
          <div id={styles.tabla}>
            <table className="table table-dark table-striped">
              <thead>
                <tr>
                  <th>#</th>
                  <th>Nombre</th>
                  <th>Duración</th>
                  <th>Requisitos</th>
                </tr>
              </thead>
              <tbody>
                {misCursos.map((curso, index) => (
                  <tr key={index}>
                    <th>{index + 1}</th>
                    <td>{curso.name}</td>
                    <td>{curso.duracion}</td>
                    <td>{curso.requisitos}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>

      <br />

      {/* Lista de cursos disponibles */}
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
                      <p className="card-text">{curso.description}</p>
                      <p className="card-text">
                        <small className="text-body-secondary">
                          <b>Requisitos:</b> {curso.requisitos}
                          <br />
                          <b>Duración:</b> {curso.duracion}
                        </small>
                      </p>
                      <button
                        className="btn btn-outline-info"
                        onClick={() => handleInscribir(curso.id)}
                      >
                        Inscribirme
                      </button>
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
    </div>
  );
};

export default Alumno;
