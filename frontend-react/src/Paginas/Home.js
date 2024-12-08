import React, { useEffect, useState } from 'react';
import axios from 'axios';
import styles from './Home.module.css';

const Home = ({ searchTerm }) => {
  const [filteredCursos, setFilteredCursos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Obtener los cursos desde la base de datos
  useEffect(() => {
    const fetchCursos = async () => {
      try {
        const response = await axios.get('http://localhost:8080/courses/all');
        setFilteredCursos(response.data.courses);
      } catch (err) {
        setError('Error al cargar los cursos.');
      } finally {
        setLoading(false);
      }
    };
    fetchCursos();
  }, []);

  useEffect(() => {
    if (searchTerm) {
      const term = searchTerm.toLowerCase();
      const filtered = filteredCursos.filter(curso =>
        curso.name.toLowerCase().includes(term) ||
        curso.category.toLowerCase().includes(term)
      );
      setFilteredCursos(filtered);
    }
  }, [searchTerm, filteredCursos]);

  if (loading) return <div className={styles.loader}>Cargando...</div>;
  if (error) return <div className={styles.error}>{error}</div>;

  return (
    <div className={styles.fondo}>
      {/* Descripción */}
      <div className={styles.descripcion}>
        <h1>Unite a la comunidad de aprendizaje online en vivo más grande de Latinoamérica</h1>
        <p>Clases online en vivo dictadas por expertos de la industria, enfoque 100% práctico, mentorías personalizadas y acceso a una comunidad de +100.000 estudiantes.</p>
      </div>
      <br />

      {/* Profesores */}
      <div className={styles.profesores}>
        <h1>Educación online que funciona</h1>
        <div className={styles.profesores_descr}>
          <h2>Clases online en vivo con Profesores Expertos</h2>
          <p>En nuestros cursos online en vivo vas a tener interacción real e instantánea con tu profesor, tutor y compañeros, para resolver todas tus dudas y avanzar en tus proyectos prácticos.</p>
          <ul>
            <li>Interacción en vivo.</li>
            <li>Corrección de proyectos prácticos.</li>
            <li>Networking con tus compañeros.</li>
          </ul>
        </div>
        <div>
          <img src={require("../img/profesor.jpg")} alt="Profesor" />
        </div>
        <br /><br /><br />
      </div>

      {/* Estudiantes */}
      <div className={styles.estudiantes}>
        <div className={styles.estudiantes_descr}>
          <h2>Aprendé junto a tus compañeros</h2>
          <p>Está demostrado que aprender en grupo es más eficiente y motivador. El networking con tus compañeros de clase ayuda a que puedas tener nuevas ideas y hacer mejores proyectos.</p>
          <ul>
            <li>Clases en grupo</li>
            <li>Canales de chat</li>
            <li>Foro privado</li>
          </ul>
        </div>
        <div>
          <img src={require("../img/estudiantes.jpg")} alt="Estudiantes" />
        </div>
      </div>

      <div className={styles.cursos}>
        <section id={styles.cards}>
          <div className="container">
            <div className={styles.title}>
              <h1>Contamos con una amplia selección de cursos:</h1>
              <p></p>
            </div>
            <div className="row">
              {filteredCursos.map((curso) => (
                <div key={curso.id} className="col-md-3 g-5">
                  <div className="card text-center">
                    <img src={curso.imagen_url} className="card-img-top" alt={curso.name} />
                    <div className={styles.cardbody}>
                      <h5 className={styles.cardtitle}>{curso.name}</h5>
                      <p>
                        Duración: {curso.duracion} <br />
                        Categoría: {curso.category} <br />
                      </p>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>
      </div>

      <div className={styles.piePagina}>
        <fieldset>
          <legend>Sobre el Curso</legend>
          <p>
            Udemy es una <b>plataforma de educación online en vivo</b> creada con la misión de democratizar la educación de calidad en toda Latinoamérica. Queremos
            que nuestros alumnos estén preparados de la mejor manera para su salida al mundo laboral.
          </p>
        </fieldset>
      </div>
    </div>
  );
};

export default Home;
