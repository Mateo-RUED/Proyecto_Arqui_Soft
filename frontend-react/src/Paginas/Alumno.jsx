import React, { useState, useEffect } from "react";
import styles from "./Alumno.module.css";
import axios from "../axiosConfig";
import { useNavigate } from "react-router-dom";

const Alumno = () => {
  const [cursos, setCursos] = useState([]); // Cursos totales
  const [misCursos, setMisCursos] = useState([]); // Cursos del usuario
  const [mostrarTabla, setMostrarTabla] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  // Función para obtener todos los cursos
  const fetchTodosLosCursos = async () => {
    try {
      const response = await axios.get("http://localhost:8080/courses/all");
      setCursos(response.data.courses || []);
    } catch (error) {
      console.error("Error fetching all courses:", error);
      setError(
        "Error al cargar los cursos. Por favor, inténtalo de nuevo más tarde."
      );
    }
  };

  // Función para obtener los cursos del usuario
  const fetchMisCursos = async () => {
    try {
      const token = localStorage.getItem("token");
      const usuarioID = localStorage.getItem("usuarioID");

      if (!usuarioID) {
        throw new Error("Usuario ID no está disponible en localStorage");
      }

      const response = await axios.get(
        `http://localhost:8080/inscripciones/users/${usuarioID}/courses`,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      setMisCursos(response.data.courses || []);
    } catch (error) {
      console.error("Error fetching user courses:", error);
      setError(
        "Error al cargar los cursos. Por favor, inténtalo de nuevo más tarde."
      );
    }
  };

  // Función para inscribir al usuario en un curso
  const handleInscribir = async (cursoID) => {
    try {
      if (misCursos.some((curso) => curso.id === cursoID)) {
        alert("Ya estás inscrito en este curso.");
        return;
      }

      const usuarioID = localStorage.getItem("usuarioID");
      const token = localStorage.getItem("token");

      if (!usuarioID || !token) {
        alert("Debes iniciar sesión para inscribirte en un curso");
        return;
      }

      const response = await axios.post(
        "http://localhost:8080/inscripciones/inscribir",
        {
          usuario_id: Number(usuarioID),
          curso_id: Number(cursoID),
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      if (response.status === 200) {
        alert("Inscripción exitosa");
        fetchMisCursos();
      } else {
        alert("No se pudo completar la inscripción");
      }
    } catch (error) {
      console.error("Error al inscribirse en el curso:", error);
      setError(
        "Error al inscribirse en el curso. Por favor, inténtalo de nuevo más tarde."
      );
    }
  };

  // Redirige a la página de agregar comentario para el curso seleccionado
  const handleAgregarComentario = (cursoID) => {
    navigate(`/comentario/${cursoID}`);
  };

  // Redirige a la página de ver archivos del curso
  const handleVerArchivos = (cursoID) => {
    navigate(`archivos/${cursoID}`);
  };

  useEffect(() => {
    fetchTodosLosCursos();
  }, []);

  const handleMostrarClick = () => {
    setMostrarTabla(!mostrarTabla);
    if (!mostrarTabla) {
      fetchMisCursos();
    }
  };

  return (
    <div className={styles.fondo}>
      <div className="container">
        <div className="card text-bg-dark">
          <img
            src={require("../img/fondoAlumno.jpg")}
            className="card-img"
            alt="..."
          />
          <div className="card-img-overlay">
            <h1 className="card-title">Bienvenido a la Sección Alumno</h1>
            <h3 className="card-text">
              Aquí puedes encontrar tus cursos inscritos y los cursos
              disponibles para inscribirte.
            </h3>
          </div>
        </div>
      </div>

      <br />

      <div className="container">
        <button
          id="btnMostrar"
          className={`btn btn-outline-info ${styles.boton}`}
          onClick={handleMostrarClick}
        >
          {mostrarTabla ? "Ocultar Mis Cursos" : "Ver Mis Cursos"}
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
                  <tr key={curso.id}>
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

      <div className={styles.tarjetas}>
        {cursos.length > 0 ? (
          cursos.map((curso) => (
            <div key={curso.id} className={styles.card}>
              <div className="card mb-4 bg-secondary-subtle">
                <div className="row g-0">
                  <div className="col-md-2">
                    <img
                      src={curso.imagen_url}
                      className="card-img-top"
                      alt={curso.name}
                    />
                  </div>
                  <div className="col-md-8">
                    <div className="card-body">
                      <h5 className="card-title">{curso.name}</h5>
                      <p className="card-text">{curso.description}</p>
                      <button
                        className={styles.boton}
                        onClick={() => handleInscribir(curso.id)}
                        disabled={misCursos.some((misCurso) => misCurso.id === curso.id)}
                      >
                        {misCursos.some((misCurso) => misCurso.id === curso.id)
                          ? "Ya inscrito"
                          : "Inscribirme"}
                      </button>
                      <button
                        className={styles.boton}
                        onClick={() => handleAgregarComentario(curso.id)}
                      >
                        Agregar Comentario
                      </button>
                      <button
                        className={styles.boton}
                        onClick={() => handleVerArchivos(curso.id)}
                      >
                        Ver Archivo
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
