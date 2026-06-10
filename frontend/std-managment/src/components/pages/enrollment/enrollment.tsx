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
      const res = await fetch(`http://localhost:8080/enrollments/${stdid}/${coursename}`, {
         method: "DELETE"
    });

        if(!res.ok){
          throw new Error("failed to delete the enrollment");
        }
      
    } catch (error) {
      console.log(error);
    }
  }


  const [stdid , setstdid] = useState("");
  const [stdname , setstdname] = useState("");
  const [csname , setcsname] = useState("");

  const createenroll = async ( e : React.MouseEvent) => {
    e.preventDefault() ;

    try {

      const response = await fetch("http://localhost:8080/enrollments" , {
        method : "POST",
        headers : {
          "Content-Type" : "application/json" ,
        }, 

        body : JSON.stringify({
          stdid : Number(stdid),
          stdname : stdname ,
          coursename : csname ,
        }),
      }) 


      if (!response.ok){
        alert(await response.text());
        return ;
      }

      alert("enrollment created");

      setcsname("");
      setstdid("");
      setstdname("");

      
      
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

        <br /><br /><br />

        <form className='enrolldiv'>
          <input type="text" placeholder='enter the student id' className='enrollinp' onChange={(e) => setstdid(e.target.value)}/>
          <input type="text" placeholder='enter the student name' className='enrollinp' onChange={(e) => setstdname(e.target.value)}/>
          <input type="text" placeholder='enter the course you want' className='enrollinp' onChange={(e) => setcsname(e.target.value)}/>
          <input type="button" value="create enroll !!" className='enrollbtn' onClick={createenroll}/>
        </form>
      
    </>
  )
}

export default Enrollment