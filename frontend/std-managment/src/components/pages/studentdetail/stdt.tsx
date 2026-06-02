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
      <h3>{stdetail?.id}</h3>
      <h3>{stdetail?.age}</h3>
      <h3>{stdetail?.name}</h3>
      {stdetail?.courses?.map((course, index) => (
        <h3 key={index}>{course}</h3>
      ))}
    </>
  )
}

export default StudentDetail