import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./Paginas/Home";
import NavBar from './components/NavBar';
import Cursos from './Paginas/Cursos';
import Login from './Paginas/Login';
import IniciarSesion from './Paginas/IniciarSesion';
import Alumno from './Paginas/Alumno';
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
          <Route exact path="/Login" element={<Login />} />
          <Route exact path="/IniciarSesion" element={<IniciarSesion />} />
        </Routes>
      </Router>
      <footer />
    </div>
  );
}

export default App;
