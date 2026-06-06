import './course.css'
import { useEffect, useState } from "react";

function Courses() {
 
  interface course {
    id : number ;
    name : string ;
  }

  const [courses , setcourses] = useState<course[]>([]);

 useEffect(() => {
    fetch("http://localhost:8080/courses")
      .then((res) => res.json())
      .then((data) => {
        console.log("Data:", data);
        setcourses(data);
      })
      .catch((err) => console.error(err));
  }, []);
  
  return (
    <>
      <section>
        {courses.map((course) => (
          <div key={course.id} className="course-item">
            <h3>{course.name}</h3>
            <h3>ID: {course.id}</h3>
          </div>
        ))}
      </section>
    </>
  )
}

export default Courses