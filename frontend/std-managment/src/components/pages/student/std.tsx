import './std.css'
import { useEffect, useState } from "react";

interface Student {
    id: number;
    name: string;
    age: number;
}

function Students() {

  const [students, setStudents] = useState<Student[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/students")
      .then((res) => res.json())
      .then((data) => {
        console.log("Data:", data);
        setStudents(data);
      })
      .catch((err) => console.error(err));
  }, []);

return (
    <>
      <div className='students'>
        {students.map((student) => (
          <div key={student.id} className='std'>     
              <div className='info'>
                <h3>id :{student.id}</h3>    
                <h3>Name :{student.name}</h3>
                <h3>Age :{student.age}</h3>
              </div>

              <div className='stdbtndiv'>
                <button className='stdbtn'>delete</button>
                <button className='stdbtn'>show details</button>
              </div>
          </div>
        ))}
      </div>

      <section className=''>

      </section>
      
    </>
  )
}

export default Students