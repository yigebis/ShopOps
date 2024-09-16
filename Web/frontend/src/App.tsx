import './App.css'
import Register from './pages/register'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './pages/login';


function App() {

  return (
    <>
    <Router>
      <Routes>
        <Route path = "/login" element = {<Login />}/>
        <Route path = "/register" element = {<Register />}/>
      </Routes>
    </Router>
    


    </>
  )
}

export default App
