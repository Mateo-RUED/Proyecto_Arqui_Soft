import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "../axiosConfig";
import styles from "./ArchivoAdmin.module.css";

const ArchivoAdmin = () => {
  const { cursoID } = useParams(); // Obtener el ID del curso desde la URL
  const [archivos, setArchivos] = useState([]); // Lista de archivos del curso
  const [archivo, setArchivo] = useState(null); // Archivo seleccionado
  const [mensaje, setMensaje] = useState(""); // Mensajes de error o éxito
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

  // Manejar el archivo seleccionado
  const handleArchivoChange = (e) => {
    setArchivo(e.target.files[0]);
    setMensaje(""); // Limpiar cualquier mensaje previo
  };

  // Subir el archivo al servidor
  const handleSubirArchivo = async () => {
    if (!archivo) {
      setMensaje("Por favor, selecciona un archivo.");
      return;
    }

    const formData = new FormData();
    formData.append("file", archivo);

    try {
      const token = localStorage.getItem("token");
      const response = await axios.post(
        `/archivos/subir/${cursoID}`,
        formData,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      if (response.status === 200) {
        setMensaje("Archivo subido correctamente.");
        fetchArchivos(); // Actualizar la lista de archivos
        setArchivo(null); // Limpiar el archivo seleccionado
        document.querySelector("input[type='file']").value = ""; // Resetear el input
      }
    } catch (error) {
      console.error("Error al subir el archivo:", error.response || error.message);
      if (error.response) {
        console.error("Detalles del error:", error.response.data);
        setMensaje(`Error: ${error.response.data.message || "No se pudo subir el archivo"}`);
      } else {
        setMensaje("Hubo un problema al subir el archivo.");
      }
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
        <input
          type="file"
          onChange={handleArchivoChange}
          className={styles.input}
        />
        <button onClick={handleSubirArchivo} className={styles.boton}>
          Subir Archivo
        </button>
        <button onClick={() => navigate("/admin")} className={styles.boton}>
          Volver
        </button>
        {mensaje && <p className={styles.mensaje}>{mensaje}</p>}
        <ul>
          {archivos.map((archivo, index) => (
            <li key={index}>
              <a
                href={`${process.env.REACT_APP_API_URL || "http://localhost:8080"}/backend/backend/uploads/${cursoID}/${archivo.nombre}`} // URL correcta para descargar
                download // Esto asegura que el archivo se descargue
              >
                {archivo.nombre} {/* Nombre del archivo */}
              </a>
            </li>
          ))}
        </ul>

      </div>
    </div>
  );
};

export default ArchivoAdmin;
