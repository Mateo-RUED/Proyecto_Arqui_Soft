import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "../axiosConfig";
import styles from "./ArchivoAlumno.module.css";

const ArchivoAlumno = () => {
  const { cursoID } = useParams(); // Obtener el ID del curso desde la URL
  const [archivos, setArchivos] = useState([]); // Lista de archivos del curso
  const [mensaje, setMensaje] = useState(""); // Mensajes de error
  const navigate = useNavigate();

  // Obtener los archivos existentes para el curso
  const fetchArchivos = async () => {
    try {
      const response = await axios.get(`/archivos/${cursoID}/listar`);
      setArchivos(response.data.files || []); // Actualiza la lista de archivos
      setMensaje(""); // Limpiar mensajes previos
    } catch (error) {
      console.error("Error al obtener archivos:", error);
      setMensaje("Error al cargar los archivos.");
    }
  };

  // Cargar los archivos al montar el componente
  useEffect(() => {
    fetchArchivos(); // Cargar los archivos cuando el componente se monta
  }, [cursoID]); // Dependencia en cursoID, para recargar cuando cambie el curso

  return (
    <div className={styles.fondo}>
      <h1 className={styles.titulo}>Archivos del Curso</h1>
      <div className={styles.archivos}>
        {mensaje && <p className={styles.mensaje}>{mensaje}</p>}
        {archivos.length === 0 ? (
          <p className={styles.mensaje}>No hay archivos subidos para este curso.</p>
        ) : (
          <ul>
            {archivos.map((archivo, index) => (
              <li key={index}>
                <a
                  href={`${process.env.REACT_APP_API_URL || "http://localhost:8080"}/archivos/${cursoID}/descargar/${archivo.nombre}`}
                  download
                  className={styles.enlace}
                >
                  {archivo.nombre}
                </a>
              </li>
            ))}
          </ul>
        )}
        <button onClick={() => navigate(-1)} className={styles.boton}>
          Volver
        </button>
      </div>
    </div>
  );
};

export default ArchivoAlumno;
