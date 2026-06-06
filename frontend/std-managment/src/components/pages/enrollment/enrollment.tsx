import './enrollment.css';
import { useEffect, useState } from "react";
function Enrollment() {
 
  interface enroll{
    stdname : string;
    coursename : string;
  }
  const [enrollments , setenrollments] = useState<enroll[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/enrollments")
      .then((res) => res.json())
      .then((data) => {
        console.log("Data:", data);
        setenrollments(data);
      })
      .catch((err) => console.error(err));
  }, []);
  return (
    <>
      <section>
        { enrollments.map((enroll) => (
            <div>
              <h3> student name : {enroll.stdname}</h3>
              <h3> course name : {enroll.coursename}</h3>
            </div>
            
        ))}
      </section>
    </>
  )
}

export default Enrollment