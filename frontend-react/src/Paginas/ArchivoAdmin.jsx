import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "../axiosConfig";
import styles from "./ArchivoAdmin.module.css";

const ArchivoAdmin = () => {
  const { cursoID } = useParams();
  const [archivos, setArchivos] = useState([]);
  const [archivo, setArchivo] = useState(null);
  const navigate = useNavigate();

  // Obtener archivos del curso
  const fetchArchivos = async () => {
    try {
      const response = await axios.get(`/archivos/${cursoID}/listar`);
      setArchivos(response.data.files || []);
    } catch (error) {
      console.error("Error al obtener archivos:", error);
    }
  };

  // Manejar cambios en el archivo seleccionado
  const handleArchivoChange = (e) => {
    setArchivo(e.target.files[0]);
  };

  // Subir un archivo
  const handleSubirArchivo = async () => {
    if (!archivo) {
      alert("Por favor, selecciona un archivo.");
      return;
    }

    const formData = new FormData();
    formData.append("file", archivo);

    try {
      const token = localStorage.getItem("token");
      await axios.post(`/archivos/subir/${cursoID}`, formData, {
        headers: {
          "Content-Type": "multipart/form-data",
          Authorization: `Bearer ${token}`,
        },
      });
      alert("Archivo subido correctamente.");
      fetchArchivos(); // Actualizar lista de archivos
      setArchivo(null); // Limpiar archivo seleccionado
    } catch (error) {
      console.error("Error al subir el archivo:", error);
      alert("Hubo un problema al subir el archivo.");
    }
  };

  useEffect(() => {
    fetchArchivos();
  }, []);

  return (
    <div className={styles.fondo}>
      <h1 className={styles.titulo}>Archivos del Curso</h1>
      <div className={styles.archivos}>
        <input type="file" onChange={handleArchivoChange} />
        <button onClick={handleSubirArchivo}>Subir Archivo</button>
        <button onClick={() => navigate("/admin")}>Volver</button>
        <ul>
          {archivos.map((archivo, index) => (
            <li key={index}>
              <a
                href={`${process.env.REACT_APP_API_URL}/archivos/${cursoID}/descargar/${archivo}`}
                download
              >
                {archivo}
              </a>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default ArchivoAdmin;
