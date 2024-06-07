import React from 'react'
import styles from './Alumno.module.css'

const Alumno = () => {
    return (
     //La forma de poder darle stilos a las pag en react es a traves de className={styles.}, en este caso le estasmos dando el estilo al fondo de la misma
     <div className={styles.fondo}> <br />

     {/*La forma de encuadrar un contenido es a través del bloque de construcción container*/}
     <div class="container">

       <div class="card text-bg-dark">
         
         {/*En react las imagenes se agregan a través del codigo src={require("...")}*/}
         <img src={require("../img/fondoAlumno.jpg")} class="card-img" alt="..."/>

         <div class="card-img-overlay">

           <h1 class="card-title">Bienvenido a la Sección Alumno</h1>
           <h3 class="card-text">A continuación podras encontrar una sección MisCursos, donde apareceran todos los cursos en los cuales te encuentras 
                inscripto y tambien los cursos disponibles para inscribirse.</h3>

         </div>

       </div>

     </div>

     <br/><br/><br/>

     <div class="container">
        <div class="row">
            <div class="col-lg-4 text-center">
                <img src={require("../img/profesor1.webp")} className={styles.imgCirc} alt="..." />
                <br /><br /><p class='fw-normal text-white'>Contamos con profesores altamente calificados.</p>
            </div>
            <div class="col-lg-4 text-center">
                <img src={require("../img/certificado.jpeg")} className={styles.imgCirc} alt="..." />
                <br /><br /><p class='fw-normal text-white'>Todos nuestros cursos cuentan con una certificación a nivel internacional.</p>
            </div>
            <div class="col-lg-4 text-center">
                <img src={require("../img/compañeros.jpg")} className={styles.imgCirc} alt="..." />
                <br /><br /><p class='fw-normal text-white'>Está demostrado que aprender en grupo es más eficiente y motivador.</p>
            </div>
        </div>
    </div>

     <br/><br/><br/><br/><br/><br/>

    <div class='container'>
         {/*La idea de este boton es que sea capaz de mostrar/ocultar la tabla que contiene a los cursos inscriptos de cada alumno*/}

        <div>
            <button id='btnMostrar' class="btn btn-outline-info" >Mis Cursos:</button>
        </div>

        {/*Crea una barra de división entre el boton y la tabla*/}
        <hr class="featurette-divider" /> <br/>

        {/*Tabla que va a contener los datos de los cursos los cuales el alumno se inscrbió*/}
        <div id={styles.tabla}>
            <table class="table table-dark table-striped">
                <thead>
                    <tr>
                    <th scope="col">#</th>
                    <th scope="col">Nombre</th>
                    <th scope="col">Duración</th>
                    <th scope="col">Profesor</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                    <th scope="row">1</th>
                    <td>Mark</td>
                    <td>Otto</td>
                    <td>@mdo</td>
                    </tr>
                    <tr>
                    <th scope="row">2</th>
                    <td>Jacob</td>
                    <td>Thornton</td>
                    <td>@fat</td>
                    </tr>
                    <tr>
                    <th scope="row">3</th>
                    <td colspan="2">Larry the Bird</td>
                    <td>@twitter</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>

    <br/><br/><br/>

     {/*En este caso el className={styles.tarjetas} se encargar de darle el tamaño necesario a nuestras tarjetas*/}
     <div className={styles.tarjetas}>

       {/*En className={styles.card} lo que hacemos es darle ese efecto de zoom a la hora de pasar el cursor por encima de la tarjeta*/}
       <div className={styles.card}>

         <div class="card mb-4 bg-secondary-subtle">

           <div class="row g-0">
             
             {/*Con col-md-2 modificamos el tamaño de la tarjeta*/}
             <div class="col-md-2">

               <img src={require("../img/nuevojava.png")} class="card-img-top"/>

             </div>

             <div class="col-md-8">

               <div class="card-body">

                 <h5 class="card-title">Java</h5>

                 <p class="card-text">Nuestro programa está diseñado para equiparte con habilidades avanzadas en desarrollo backend utilizando JAVA,
                   incluyendo frameworks como Spring Boot, manejo de bases de datos, y mejores prácticas en seguridad y diseño de software.</p>
                 
                 <p class="card-text"><small class="text-body-secondary">
                   • <b>Requisitos:</b> <br/>
                   &nbsp;&nbsp;- Entender programación básica, incluyendo variables, tipos de datos y estructuras de control. <br/>
                   &nbsp;&nbsp;- Poseer una noción básica sobre desarrollo backend. <br/>
                   &nbsp;&nbsp;- Familiaridad con IDs y sistemas de control de versiones como Git. <br/>
                   • <b>Duración:</b> <br/>
                   &nbsp;&nbsp;- 30 Horas de contenido en video.
                 </small></p>

                 <button class="btn btn-outline-info">Inscribirte</button>

               </div>

             </div>

           </div>

         </div>

       </div>
       
       <div className={styles.card}>

         <div class="card mb-4 bg-secondary-subtle">

           <div class="row g-0">

             <div class="col-md-2">

               <img src={require("../img/nuevohtml.png")} class="card-img-top"/>

             </div>

             <div class="col-md-8">

               <div class="card-body">

                 <h5 class="card-title">HTML y CSS</h5>

                 <p class="card-text">En este curso aprenderás a crear tu sitio web partiendo del prototipo en papel. Te sumergirás en las mejores
                   prácticas del desarrollo web, trabajando con HTML y CSS.</p>
                 
                 {/*EL text-body-secondary se utiliza para darle el color de fondo a las tarjetas*/}
                 <p class="card-text"><small class="text-body-secondary">
                   • <b>Requisitos:</b> <br/>
                   &nbsp;&nbsp;- Para este curso no es necesario tener conocimientos previos. <br/>
                   &nbsp;&nbsp;- Para mejorar tu experiencia de cursada, te aconsejamos tener una computadora con 2 GB de memoria RAM, <br/>
                   &nbsp;&nbsp;&nbsp; procesador de dos núcleos y GPU de 2 GB de RAM. <br/>
                   • <b>Duración:</b> <br/>
                   &nbsp;&nbsp;- 7 Horas de contenido en video.
                 </small></p>

                 <button class="btn btn-outline-info">Inscribirte</button>

               </div>

             </div>

           </div>

         </div>

       </div>

       <div className={styles.card}>

         <div class="card mb-4 bg-secondary-subtle">

           <div class="row g-0">

             <div class="col-md-2">

               <img src={require("../img/JavaScript.png")} class="card-img-top"/>
               
             </div>

             <div class="col-md-8">

               <div class="card-body">

                 <h5 class="card-title">JavaScript</h5>

                 <p class="card-text">En este curso aprenderás los fundamentos del lenguaje de programación más usado en la actualidad, con el cual
                   es posible crear aplicaciones de todo tipo. Explorarás inicialmente herramientas propias del mismo, indagando casos prácticos de
                   aplicación.</p>

                 <p class="card-text"><small class="text-body-secondary">
                   • <b>Requisitos:</b> <br/>
                   &nbsp;&nbsp;- Tener conocimientos intermedios de HTML, CSS. <br/>
                   &nbsp;&nbsp;- Poseer una noción básica sobre programación lógica. <br/>
                   • <b>Duración:</b> <br/>
                   &nbsp;&nbsp;- 6 Horas de contenido en video.
                 </small></p>

                 <button class="btn btn-outline-info">Inscribirte</button>

               </div>

             </div>

           </div>

         </div>

       </div>
       
       <div className={styles.card}>

         <div class="card mb-4 bg-secondary-subtle">

           <div class="row g-0">

             <div class="col-md-2">

               <img src={require("../img/sql.png")} class="card-img-top"/>

             </div>

             <div class="col-md-8">

               <div class="card-body">

                 <h5 class="card-title">SQL</h5>

                 <p class="card-text">En este curso aprenderás las nociones centrales de las bases de datos relacionales, las cuales son implementadas
                   por todas las organizaciones para poder tomar decisiones con base en la información que generan en su modelo de negocio.</p>
                 
                 <p class="card-text"><small class="text-body-secondary">
                   • <b>Requisitos:</b> <br/>
                   &nbsp;&nbsp;- Tener conocimientos sobre otros lenguajes de programación. <br/>
                   &nbsp;&nbsp;- PC con 4 GB de memoria RAM, procesador de cuatro núcleos y GPU de 2 GB de RAM. <br/>
                   • <b>Duración:</b> <br/>
                   &nbsp;&nbsp;- 4 Horas de contenido en video.
                 </small></p>

                 <button class="btn btn-outline-info">Inscribirte</button>

               </div>
             </div>

           </div>

         </div>

       </div>

       <div className={styles.card}>

         <div class="card mb-4 bg-secondary-subtle">

           <div class="row g-0">

             <div class="col-md-2">

               <img src={require("../img/nuevopowerby.png")} class="card-img-top"/>

             </div>

             <div class="col-md-8">

               <div class="card-body">

                 <h5 class="card-title">Powerby</h5>

                 <p class="card-text">Power BI es líder del análisis y la visualización de datos gracias a su completo conjunto de soluciones
                     que pueden ayudar a las organizaciones a tomar mejores decisiones. El programa de formación está destinado a todas aquellas
                     personas que quieran incursionar en el mundo del análisis de datos.</p>
                 
                 <p class="card-text"><small class="text-body-secondary">
                   • <b>Requisitos:</b> <br/>
                   &nbsp;&nbsp;- Conocimiento básico de navegación en la web y uso de un sistema operativo. <br/>
                   &nbsp;&nbsp;- Comprender conceptos matemáticos fundamentales, como operaciones aritméticas y estadística básica. <br/>
                   &nbsp;&nbsp;- Habilidad para descomponer problemas y analizar datos para tomar decisiones informadas. <br/>
                   • <b>Duración:</b> <br/>
                   &nbsp;&nbsp;- 8.5 Horas de contenido en video.
                 </small></p>
                 
                 <button class="btn btn-outline-info">Inscribirte</button>

               </div>

             </div>

           </div>

         </div>

       </div>


       <div className={styles.card}>

         <div class="card mb-4 bg-secondary-subtle">

           <div class="row g-0">

             <div class="col-md-2">

               <img src={require("../img/nuevopython.png")} class="card-img-top"/>

             </div>

             <div class="col-md-8">

               <div class="card-body">

                 <h5 class="card-title">Python</h5>

                 <p class="card-text">Te sumergirás en el mundo de la programación con Python, aprendiendo desde los conceptos básicos hasta la
                   programación orientada a objetos. Desarrollarás habilidades en manejo de datos, definición de funciones y resolución de problemas
                   complejos mediante el diseño de algoritmos.</p>
                 
                 <p class="card-text"><small class="text-body-secondary">
                   • <b>Requisitos:</b> <br/>
                   &nbsp;&nbsp;- Comprender conceptos matemáticos fundamentales, como operaciones aritméticas. <br/>
                   &nbsp;&nbsp;- Habilidad para descomponer y analizar problemas. <br/>
                   &nbsp;&nbsp;- Capacidad para analizar y resolver problemas de programación de nivel intermedio. <br/>
                   • <b>Duración:</b> <br/>
                   &nbsp;&nbsp;- 13 Horas de contenido en video.
                 </small></p>

                 <button class="btn btn-outline-info">Inscribirte</button>

               </div>

             </div>

           </div>

         </div>

       </div>

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

export default Alumno