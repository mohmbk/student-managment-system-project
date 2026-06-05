import './stdt.css'
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import Courses from '../courses/course';

interface StdDetail {
  id: number;
  name: string;
  age: number;
  courses: string[];
}


function StudentDetail() {
  const [stdetail, setstdetail] = useState<StdDetail | null>(null);
  const {id} = useParams();
  useEffect(() => {
  fetch(`http://localhost:8080/StudentDetail/${id}`)
    .then((res) => res.json())
    .then((data) => setstdetail(data))
    .catch((err) => console.error(err));
  }, [id]);
  return (
    <>
      <div className='container'>
        <div className='subcontainer'>
          <div className='stdinfo'>
            <h1>student info :</h1><br />
            <h3>id : {stdetail?.id}</h3><br />
            <h3>age : {stdetail?.age}</h3><br />
            <h3>name : {stdetail?.name}</h3> <br /><br /><br />
            <h1>course enrolled in :</h1><br />
            {stdetail?.courses?.map((course, index) => (
              <h3 key={index}>{course}</h3>
            ))}
          </div>

          <div className='imgdiv'>
            <img src="/cs img.jpg" alt="" className='img'/>
          </div>
        </div>
      </div>
    </>
  )
}

export default StudentDetail