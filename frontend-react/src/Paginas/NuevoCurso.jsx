import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Importa useNavigate
import styles from './NuevoCurso.module.css';
import axios from '../axiosConfig';

const NuevoCurso = () => {
  const navigate = useNavigate(); // Inicializa navigate

  const [curso, setCurso] = useState({
    name: '',
    description: '',
    category: '',
    requisitos: '',
    duracion: '',
    imagen_url: '',
  });

  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setCurso({ ...curso, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      const token = localStorage.getItem('token');
      const response = await axios.post(
        'http://localhost:8080/courses/create_course',
        curso,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      if (response.status === 201) {
        setSuccess('Curso creado exitosamente');
        setCurso({
          name: '',
          description: '',
          category: '',
          requisitos: '',
          duracion: '',
          imagen_url: '',
        });
        navigate('/admin'); // Redirige a la página de administración
      }
    } catch (err) {
      console.error('Error al crear el curso:', err);
      setError('Hubo un error al crear el curso. Intente nuevamente.');
    }
  };

  return (
    <div className={styles.fondo}>
      <h1>Crear Nuevo Curso</h1>
      <form className={styles.formulario} onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Nombre del Curso</label>
          <input
            type="text"
            name="name"
            value={curso.name}
            onChange={handleInputChange}
            required
          />
        </div>

        <div className="form-group">
          <label>Descripción</label>
          <textarea
            name="description"
            value={curso.description}
            onChange={handleInputChange}
            required
          ></textarea>
        </div>

        <div className="form-group">
          <label>Categoría</label>
          <select
            name="category"
            value={curso.category}
            onChange={handleInputChange}
            required
          >
            <option value="">Seleccione una categoría</option>
            <option value="Programación">Programación</option>
            <option value="Análisis de datos">Análisis de datos</option>
            <option value="Base de datos">Base de datos</option>
            <option value="Desarrollo web">Desarrollo web</option>
          </select>
        </div>

        <div className="form-group">
          <label>Requisitos</label>
          <textarea
            name="requisitos"
            value={curso.requisitos}
            onChange={handleInputChange}
            required
          ></textarea>
        </div>

        <div className="form-group">
          <label>Duración</label>
          <input
            type="text"
            name="duracion"
            value={curso.duracion}
            onChange={handleInputChange}
            required
          />
        </div>

        <div className="form-group">
          <label>URL de la Imagen</label>
          <input
            type="url"
            name="imagen_url"
            value={curso.imagen_url}
            onChange={handleInputChange}
          />
        </div>

        {error && <p className="text-danger">{error}</p>}
        {success && <p className="text-success">{success}</p>}

        <button type="submit" className={`${styles.btnPrimary} btn-primary`}>
          Crear Curso
        </button>
      </form>
    </div>
  );
};

export default NuevoCurso;
