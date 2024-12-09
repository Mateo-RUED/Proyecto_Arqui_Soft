import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./Paginas/Home";
import NavBar from './components/NavBar';
import Cursos from './Paginas/Cursos';
import Login from './Paginas/Login';
import IniciarSesion from './Paginas/IniciarSesion';
import Alumno from './Paginas/Alumno'; // Página de alumno
import Admin from './Paginas/Admin'; // Página de administrador
import Comentario from "./Paginas/Comentario"; // Página de comentarios
import ArchivoAdmin from "./Paginas/ArchivoAdmin"; // Archivos para administrador
import ArchivoAlumno from "./Paginas/ArchivoAlumno"; // Archivos para alumno
import NuevoCurso from './Paginas/NuevoCurso'; // Página para nuevo curso
import { useState } from 'react';

function App() {
  const [searchTerm, setSearchTerm] = useState('');

  return (
    <div>
      <Router>
        <NavBar onSearch={setSearchTerm} />
        <Routes>
          {/* Página principal */}
          <Route exact path="/" element={<Home searchTerm={searchTerm} />} />

          {/* Páginas generales */}
          <Route exact path="/Cursos" element={<Cursos />} />
          <Route exact path="/Login" element={<Login />} />
          <Route exact path="/IniciarSesion" element={<IniciarSesion />} />

          {/* Páginas del alumno */}
          <Route exact path="/Alumno" element={<Alumno />} />
          <Route exact path="/comentario/:cursoID" element={<Comentario />} />
          <Route exact path="/alumno/archivos/:cursoID" element={<ArchivoAlumno />} />

          {/* Páginas del administrador */}
          <Route exact path="/admin" element={<Admin />} />
          <Route exact path="/nuevo-curso" element={<NuevoCurso />} />
          <Route exact path="/admin/archivos/:cursoID" element={<ArchivoAdmin />} />
        </Routes>
      </Router>
      <footer />
    </div>
  );
}

export default App;
