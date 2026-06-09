import './enrollment.css';
import { useEffect, useState } from "react";
function Enrollment() {
 
  interface enroll{
    stdid : number ;
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


  async function deleteenroll(stdid : number, coursename : string) {
    try {
      const res = await fetch(`http://localhost:8080/enroll/${stdid}/${coursename}`, {
         method: "DELETE"
    });

        if(!res.ok){
          throw new Error("failed to delete the enrollment");
        }
      
    } catch (error) {
      console.log(error);
    }
  }
  return (
    <>
      <section>
        { enrollments.map((enroll) => (
            <div className='enroll'>
              <h3>student id : {enroll.stdid}</h3>
              <h3> student name : {enroll.stdname}</h3>
              <h3> course name : {enroll.coursename}</h3>
              <button className='enrdlt' onClick={() => deleteenroll(enroll.stdid, enroll.coursename)}>delete</button>
            </div>
            
        ))}
      </section>
    </>
  )
}

export default Enrollment