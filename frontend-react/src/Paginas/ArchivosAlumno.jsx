import React, { useState, useEffect } from "react";
import axios from "../axiosConfig";
import { useNavigate, useParams } from "react-router-dom";

const ArchivosAlumno = () => {
  const [files, setFiles] = useState([]);
  const { courseID } = useParams();
  const navigate = useNavigate();

  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await axios.get(`/archivos/${courseID}/listar`);
        setFiles(response.data.files || []);
      } catch (error) {
        console.error("Error fetching files:", error);
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
              {/* Agregar atributo `download` para descargar directamente */}
              <a href={`/archivos/${courseID}/descargar/${file}`} download>
                {file}
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
