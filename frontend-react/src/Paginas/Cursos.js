import React, { useEffect, useState } from 'react';
import axios from 'axios';
import styles from './Cursos.module.css';

const CursoCard = ({ curso }) => (
  <div className={styles.card}>
    <div className="card mb-4 bg-secondary-subtle">
      <div className="row g-0">
        <div className="col-md-2">
          <img src={curso.imagen_url} className="card-img-top" alt={`Imagen del curso ${curso.name}`} />
        </div>
        <div className="col-md-8">
          <div className="card-body">
            <h5 className="card-title">{curso.name}</h5>
            <p className="card-text">{curso.description}</p>
            <p className="card-text">
            <small className="text-body-secondary">
              • <b>Requisitos:</b> {curso.requisitos} <br />
              • <b>Duración:</b> {curso.duracion} <br />
              • <b>Categoría:</b> {curso.category}
            </small>
          </p>
          </div>
        </div>
      </div>
    </div>
  </div>
);

const Cursos = () => {
  const [cursos, setCursos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCursos = async () => {
      try {
        const response = await axios.get('http://localhost:8080/courses/all');
        setCursos(response.data.courses);
      } catch (err) {
        setError('Error al cargar los cursos.');
      } finally {
        setLoading(false);
      }
    };
    fetchCursos();
  }, []);

  if (loading) return <div className={styles.loader}>Cargando...</div>;
  if (error) return <div className={styles.error}>{error}</div>;

  return (
    <div className={styles.fondo}>
      <div className="container">
        <div className="card text-bg-dark">
          <img src={require('../img/fondoCursos.jpg')} className="card-img" alt="Fondo de Cursos" />
          <div className="card-img-overlay">
            <h1 className="card-title">Bienvenido a la Sección Cursos</h1>
            <h3 className="card-text">En este apartado te mostraremos todos los cursos con los que contamos.</h3>
          </div>
        </div>
      </div>

      <div className={styles.tarjetas}>
        {cursos.map(curso => (
          <CursoCard key={curso.id} curso={curso} />
        ))}
      </div>

      <div className={styles.piePagina}>
        <fieldset>
          <legend>Sobre el Curso</legend>
          <p>
            Udemy es una <b>plataforma de educación online en vivo</b> creada con la misión de democratizar la educación
            de calidad en toda Latinoamérica. Queremos que nuestros alumnos estén preparados de la mejor manera para su
            salida al mundo laboral.
          </p>
        </fieldset>
      </div>
    </div>
  );
};

export default Cursos;
