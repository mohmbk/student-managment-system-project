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


  const [id , setid] = useState("");
  const [age , setage] = useState("");
  const [name , setname] = useState("");

  const createStudent = async( e: React.MouseEvent) => {
    e.preventDefault();
    try {
      const response = await fetch("http://localhost:8080/students" , {
        method : "POST",
        headers : {
          "Content-Type" : "application/json",
        },

        body: JSON.stringify({
            id: Number(id),
            name: name,
            age: Number(age),
        }),
      })

      if(!response.ok){
        alert(await response.text());
        return ;
      }

      alert("student created");

      setid("");
      setname("");
      setage("");
    } catch (error) {
      console.error(error);
    }
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
        <br /><br />
        <h1 className='tittle'>create a student</h1>
      <section className='create'>
        <form className='inputdiv'>
          <div className='inp1'>
            <input type="text" placeholder=' id :' name='id' id='id' className='input' onChange={(e) => setid(e.target.value)}/>
            <input type="text" name='name' id='name' placeholder=' name :' className='input' onChange={(e) => setname(e.target.value)}/>
          </div>
          
          
          <div className='inp2'>
            
            <input type="text" name='age' id='age' placeholder=' age :' className='input' onChange={(e) => setage(e.target.value)}/>
          </div>

          <input type="button" value="create !!" className='createinp' onClick={createStudent}/>
        </form>
      </section>
      
    </>
  )
}

export default Students