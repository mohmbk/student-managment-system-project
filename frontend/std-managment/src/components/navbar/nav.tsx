import { Link } from "react-router-dom";
import './nav.css';

function Nav() {
 

  return (
    <>
      <section className='navbar'>
        <h2>Student Management System</h2>
        <ul className="routes">
          <li className="item"><Link to="/students">Students</Link></li>
          <li className="item"><Link to="/courses">Courses</Link></li>
          <li className="item"><Link to="/enrollments">Enrollments</Link></li>
        </ul>
      </section>
    </>
  )
}

export default Nav