import './std.css'
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

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

  async function deletestd( id : number){
      try {
        const res = await fetch(`http://localhost:8080/students/${id}` , {
          method : "DELETE",
        });

        if(!res.ok){
          throw new Error("failed to delete the student");
        }
    } catch (error) {
      console.error(error);
    }
  }

  const navigate = useNavigate();
  async function veiwdetails(id : number) {
    navigate(`/students/${id}`);
  }

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
                <button className='stdbtn'  onClick={() => deletestd(student.id)}>delete</button>
                <button className='stdbtn' onClick={() => veiwdetails(student.id)}>show details</button>
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