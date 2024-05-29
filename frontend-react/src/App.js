import './App.css';
import { BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Home from "./Paginas/Home"
import NavBar from './components/NavBar';
import Cursos from './Paginas/Cursos';
import Login from './Paginas/Login';
import IniciarSesion from './Paginas/IniciarSesion';



function App() {
  return (
    <div>
      <Router>
      <NavBar />
      <Routes>
        <Route exact path = "/" element={<Home/>}/>
        <Route exact path = "/Cursos" element={<Cursos/>}/>
        <Route exact path = "/Login" element={<Login/>}/>
        <Route exact path = "/IniciarSesion" element={<IniciarSesion/>}/>

      </Routes>
      </Router>
      <footer />
    </div>
  
        
  );

}

export default App;