import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "../axiosConfig";
import styles from "./Comentario.module.css";

const Comentario = () => {
  const { cursoID } = useParams();
  const [comentarios, setComentarios] = useState([]);
  const [nuevoComentario, setNuevoComentario] = useState("");
  const navigate = useNavigate();

  // Función para cargar los comentarios desde el backend
  const fetchComentarios = async () => {
    try {
      const response = await axios.get(`/courses/${cursoID}/comments`);
      setComentarios(response.data.comments || []);
    } catch (error) {
      console.error("Error al cargar comentarios:", error);
    }
  };

  // Función para agregar un nuevo comentario
  const handleAgregarComentario = async () => {
    if (!nuevoComentario.trim()) {
      alert("El comentario no puede estar vacío.");
      return;
    }

    try {
      const usuarioID = localStorage.getItem("usuarioID"); // Usuario autenticado
      const token = localStorage.getItem("token");

      if (!token || !usuarioID) {
        alert("No estás autenticado. Por favor, inicia sesión.");
        return;
      }

      // Asegurarse de que cursoID y usuarioID sean números
      const courseId = parseInt(cursoID, 10);
      const userId = parseInt(usuarioID, 10);

      if (isNaN(courseId) || isNaN(userId)) {
        alert("Error: cursoID o usuarioID no son válidos.");
        return;
      }

      // Enviar el nuevo comentario al backend
      const response = await axios.post(
        "/courses/add_comment",
        {
          course_id: courseId,
          user_id: userId,
          content: nuevoComentario.trim(),
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      // Verificar que la respuesta sea exitosa
      if (response.status === 200 || response.status === 201) {
        setNuevoComentario(""); // Limpiar el textarea
        fetchComentarios(); // Actualizar la lista de comentarios
      } else {
        alert("Hubo un problema al enviar tu comentario. Intenta nuevamente.");
      }
    } catch (error) {
      console.error("Error al agregar comentario:", error);

      // Mostrar errores más específicos
      if (error.response) {
        // Si el servidor devolvió un error, mostrar el mensaje
        alert(`Error del servidor: ${error.response.data.message}`);
      } else if (error.request) {
        // Si no se pudo conectar al servidor
        alert("No se pudo conectar al servidor. Verifica tu conexión.");
      } else {
        // Otros errores
        alert("Hubo un error al procesar tu solicitud.");
      }
    }
  };

  // Cargar los comentarios al montar el componente
  useEffect(() => {
    fetchComentarios();
  }, []);

  return (
    <div className={styles.fondo}>
      <h1 className={styles.titulo}>Comentarios del Curso</h1>
      <div className={styles.comentarios}>
        <ul>
          {comentarios.map((comentario) => (
            <li key={comentario.id}>
              {comentario.content} - Usuario: {comentario.user_id}
            </li>
          ))}
        </ul>
        <textarea
          className={styles.textareaComentario}
          value={nuevoComentario}
          onChange={(e) => setNuevoComentario(e.target.value)}
          placeholder="Escribe tu comentario..."
        />
        <button className={styles.botonEnviar} onClick={handleAgregarComentario}>
          Enviar Comentario
        </button>
        <button className={styles.botonVolver} onClick={() => navigate("/Alumno")}>
          Volver
        </button>
      </div>
    </div>
  );
};

export default Comentario;
