import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./Paginas/Home";
import NavBar from './components/NavBar';
import Cursos from './Paginas/Cursos';
import Login from './Paginas/Login';
import IniciarSesion from './Paginas/IniciarSesion';
import Alumno from './Paginas/Alumno'; // Nueva página de alumno
import Admin from './Paginas/Admin'; // Nueva página de administrador
import Comentario from "./Paginas/Comentario"; // Nueva página de comentarios
import ArchivoAdmin from "./Paginas/ArchivoAdmin";
import ArchivoAlumno from "./Paginas/ArchivoAlumno";
import { useState } from 'react';

function App() {
  const [searchTerm, setSearchTerm] = useState('');

  return (
    <div>
      <Router>
        <NavBar onSearch={setSearchTerm} />
        <Routes>
          <Route exact path="/" element={<Home searchTerm={searchTerm} />} />
          <Route exact path="/Cursos" element={<Cursos />} />
          <Route exact path="/Alumno" element={<Alumno />} />
          <Route exact path="/comentario/:cursoID" element={<Comentario />} />
          <Route exact path="/Login" element={<Login />} />
          <Route exact path="/IniciarSesion" element={<IniciarSesion />} />
          <Route exact path="/admin" element={<Admin />} /> {/* Ruta para administrador */}
          <Route exact path="/admin/archivos/:cursoID" element={<ArchivoAdmin />} />
          <Route exact path="/alumno/archivos/:cursoID" element={<ArchivoAlumno />} />
        </Routes>
      </Router>
      <footer />
    </div>
  );
}

export default App;
