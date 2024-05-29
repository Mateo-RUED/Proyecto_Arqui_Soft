import React from 'react'
import styles from './Home.module.css'


const Home = () => {
  return (
    <div class="bg-secondary-subtle">

      {/*<-- -------------------------------------------- Descripcion --------------------------------------- -->*/}

       <div className={styles.descripcion}>

        <h1>Unite a la comunidad de aprendizaje online en vivo más grande de Latinoamérica</h1>
        <p>Clases online en vivo dictadas por expertos de la industria, enfoque 100% práctico, mentorías personalizadas y acceso a una comunidad de +100.000 estudiantes.</p>

      </div> <br></br>
      
      {/*<-- -------------------------------------------- Profesores --------------------------------------- -->*/}

      <div className={styles.profesores}>

        <h1>Educación online que funciona </h1>
      
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
      
          <img src={require("../img/profesor.jpg")} /> 
      

        </div> <br></br><br></br><br></br>
      </div>

      {/*<-- -------------------------------------------- Estudiantes --------------------------------------- -->*/}

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
          <img src={require("../img/estudiantes.jpg")} />
        </div>

      </div>
      
      <div className={styles.cursos}>

        <section id={styles.cards}>

          <div class="container">

            <div className={styles.title}>
              <h1>Contamos con una amplia selección de cursos:</h1>
              <p></p>
            </div>

            <div class="row">

              <div class="col-md-3 g-5">

                <div class="card text-center">

                  <img src={require("../img/OIP.jpg")} class="card-img-top"/>

                  <div className={styles.cardbody}>
                    <h5 className={styles.cardtitle}>HTML y CSS</h5>
                    <p >Duración: 10 Semanas <br />
                        Clases: 2 Clases semanales <br />
                        Profesor: <br />
                        Requisitos:</p> 
                        <br />
                  </div>

                  <button class="btn btn-outline-info">Inscribirte</button>

                </div>

              </div>

              <div class="col-md-3 g-5">

                <div class="card text-center">
                  <img src={require("../img/mysql.png")} class="card-img-top"/>

                  <div className={styles.cardbody}>
                    <h5 className={styles.cardtitle}>MySQL</h5>
                    <p class="card-text">Duración: 10 Semanas <br />
                        Clases: 2 Clases semanales <br />
                        Profesor: <br />
                        Requisitos:</p> 
                        <br />
                  </div>

                  <button class="btn btn-outline-info">Inscribirte</button>

                </div>

              </div>

              <div class="col-md-3 g-5">

                <div class="card text-center">

                  <img src={require("../img/JavaScript.png")} class="card-img-top"/>

                  <div className={styles.cardbody}>
                    <h5 className={styles.cardtitle}>JavaScript</h5>
                    <p class="card-text">Duración: 8 Semanas<br />
                        Clases: 2 Clases semanales <br />
                        Profesor: <br />
                        Requisitos:</p> 
                        <br />
                  </div>

                  <button class="btn btn-outline-info">Inscribirte</button>

                </div>
                
              </div>
              
              <div class="col-md-3 g-5">
                
                <div class="card text-center">
                
                  <img src={require("../img/powerby.jpg")} class="card-img-top"/>

                  <div className={styles.cardbody}>
                    <h5 className={styles.cardtitle}>Power Bi</h5>
                    <p class="card-text">Duración: 8 Semanas, <br />
                        Clases: 2 Clases semanales <br />
                        Profesor: <br />
                        Requisitos:</p> 
                        <br />
                  </div>

                  <button class="btn btn-outline-info">Inscribirte</button>

                </div>

              </div>

              <div class="col-md-3 g-5">

                <div class="card text-center ">

                  <img src={require("../img/Java.jpg")} class="card-img-top"/>

                  <div className={styles.cardbody}>
                    <h5 className={styles.cardtitle}>Java</h5>
                    <p >Duración: 14 Semanas <br />
                        Clases: 3 Clases semanales <br />
                        Profesor: <br />
                        Requisitos: </p>
                    <br />
                  </div> 

                  <button class="btn btn-outline-info">Inscribirte</button>

                </div>
              </div>

              <div class="col-md-3 g-5">

                <div class="card text-center">

                  <img src={require("../img/python.jpg")} class="card-img-top"/>

                  <div className={styles.cardbody}>
                    <h5 className={styles.cardtitle}>Python</h5>
                    <p class="card-text">Duración:  13 Semanas<br />
                        Clases: 3 Clases semanales <br />
                        Profesor: <br />
                        Requisitos:</p> 
                        <br />
                  </div>

                  <button class="btn btn-outline-info">Inscribirte</button>

                </div>

              </div>

            </div>

          </div>

        </section>

      </div> 

      <div className={styles.piePagina}>

            <fieldset>
                <legend>Sobre el Curso</legend>
            
                <p>Udemy es una <b>plataforma de educación online en vivo</b> creada con la misión de democratizar la educacion de calidad en toda Latinoamérica. Queremos
                    que nuestros alumnos estén preparados de la mejor manera para su salidad al mundo laboral.</p>
            
            </fieldset>

      </div>

    </div>

  )
}

export default Home