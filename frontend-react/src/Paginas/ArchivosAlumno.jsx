import React, { useState, useEffect } from "react";
import axios from "../axiosConfig";
import { useNavigate, useParams } from "react-router-dom";

const ArchivosAlumno = () => {
  const [files, setFiles] = useState([]); // Lista de archivos
  const { courseID } = useParams(); // ID del curso desde la URL
  const navigate = useNavigate(); // Para navegación

  // Obtener los archivos del curso al montar el componente
  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await axios.get(`/archivos/${courseID}/listar`);
        console.log("Archivos recibidos:", response.data.files); // Verificar los datos
        setFiles(response.data.files || []);
      } catch (error) {
        console.error("Error al obtener los archivos:", error);
      }
    };
  
    fetchFiles();
  }, [courseID]);
  

  return (
    <div>
      <h1>Archivos del Curso</h1>
      {files.length === 0 ? (
        <p>No hay archivos subidos para este curso.</p>
      ) : (
        <ul>
          {files.map((file, index) => (
            <li key={index}>
              <a
                href={`${process.env.REACT_APP_API_URL || "http://localhost:8080"}${file.url}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                {file.nombre}
              </a>
            </li>
          ))}
        </ul>
      )}
      <button onClick={() => navigate(-1)}>Volver</button>
    </div>
  );
  
};

export default ArchivosAlumno;
