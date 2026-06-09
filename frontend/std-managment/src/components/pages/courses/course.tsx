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
  

  async function deleatecourse(id : Number) {
    try {
      const res = await fetch(`http://localhost:8080/courses/${id}` , {
          method : "DELETE",
        });

        if(!res.ok){
          throw new Error("failed to delete the course");
        }
      
    } catch (error) {
      console.log(error);
    }
  }
  return (
    <>
      <section>
        {courses.map((course) => (
          <div key={course.id} className="courseitem">
            
              <h3>{course.name}</h3>
              <h3>ID: {course.id}</h3>
            <button className='coursedlt' onClick={() => deleatecourse(course.id)}>delete</button>
          </div>
        ))}
      </section>
    </>
  )
}

export default Courses