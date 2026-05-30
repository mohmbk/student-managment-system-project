import './stdt.css'
import { useState, useEffect } from 'react';
function StudentDetail() {
  
 interface Student {
    id: string;
    name: string;
    age: number;
  }

  const [students, setStudents] = useState<Student[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/students")
      .then((res) => res.json())
      .then((data) => setStudents(data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <>
      <div>
        {students.map((student) =>
          <div>
            <h3>{student.id}</h3>
            <h3>{student.name}</h3>
            <h3>{student.age}</h3>
          </div>
        )}
      </div>
    </>
  )
}

export default StudentDetail