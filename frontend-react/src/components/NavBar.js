import React from 'react'
import {Link} from 'react-router-dom'



const NavBar = () => {
  return (
    <div><nav class="navbar navbar-expand-lg bg-success-subtle" >
    <div class="container-fluid">
        <Link class="navbar-brand" to='/'>Udemy</Link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarColor01">
            <ul class="navbar-nav me-auto">
            
            <li class="nav-item">
                <Link class="nav-link" to='/Cursos'>Cursos</Link>
            </li>
            
            <li class="nav-item ">
                <a class=" btn btn-outline-secondary bg-success-subtle btn-sm" href="/IniciarSesion" role="button">Iniciar Sesión</a>
            </li>
            <li class="nav-item ">
                <a class="btn btn-outline-secondary bg-success-subtle btn-sm" href="/Login" role="button">Registrate</a>
            </li>
            </ul>
            <form class="d-flex">
            <input class="form-control me-sm-2" type="search" placeholder="Search" />
            <button class="btn btn-secondary my-2 my-sm-0" type="submit">Buscar</button>
            </form>
        </div>
    </div>
    </nav></div>
  )
}

export default NavBar