
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import './App.css'
import Enrollment from './components/pages/enrollment/enrollment';
import Students from './components/pages/student/std';
import StudentDetail from './components/pages/studentdetail/stdt';
import Courses from './components/pages/courses/course';
import Nav from "./components/navbar/nav";

function App() {
 

  return (
    <>
      <Router>
        <Nav />
        <Routes>
          <Route path="/students" element={<Students />} />
          <Route path="/students/:id" element={<StudentDetail />} />
          <Route path="/courses" element={<Courses />} />
          <Route path="/enroll" element={<Enrollment/>} />
        </Routes>
    </Router>
    </>
  )
}

export default App
