package main 
import ("fmt" ; "net/http" ; "encoding/json" ; "time" ; "context" ; "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" ; "strings" ; "strconv")

var client *mongo.Client
var stdcollection *mongo.Collection
var courscollection *mongo.Collection
var enrollcollection *mongo.Collection

type std struct {
    ID   int    `bson:"id"   json:"id"`
    Name string `bson:"name" json:"name"`
    Age  int    `bson:"age"  json:"age"`
}

type course struct {
	ID int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type enroll struct {
    StdID    int `json:"stdid" bson:"stdid"`
    CourseID int `json:"courseid" bson:"courseid"`
}


type StudentDetail struct {
    ID      int      `json:"id"`
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Courses []string `json:"courses"`
}

type enrollmentRequest struct {
    StdName    string `json:"stdname"`
    CourseName string `json:"coursename"`
}


func studentsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getstd(w, r)
    case "POST":
        createstd(w, r)
    case "OPTIONS":
        w.WriteHeader(http.StatusOK) // ✅ let CORS preflight pass
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func createstd(w http.ResponseWriter , r *http.Request) {
	if r.Method != "POST" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed)
		return ;
	}
	var newstd std ;
	err := json.NewDecoder(r.Body).Decode(&newstd);
	if err != nil {
		http.Error(w , "Invalid request body" , http.StatusBadRequest)
		return ;
	}
	count , err := stdcollection.CountDocuments(context.Background() , bson.M{"id": newstd.ID});
	if err != nil {
		http.Error(w , "Error checking for existing student" , http.StatusInternalServerError)
		return ;
	}

	if count != 0 {
		http.Error(w , "Student with this ID already exists" , http.StatusBadRequest)
		return ;
	}
	
	reslut , err := stdcollection.InsertOne(context.Background() , newstd);
	if err != nil {
		http.Error(w , "Error creating student" , http.StatusInternalServerError)
		return ;
	}
	fmt.Printf("Inserted student with ID: %v\n" , reslut.InsertedID);

}


func getstd(w http.ResponseWriter , r *http.Request) {
	if r.Method != "GET" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed)
		return ;
	}
	cursor , err := stdcollection.Find(context.Background() , bson.M{});
	if err != nil {
		http.Error(w , "Error fetching students" , http.StatusInternalServerError)
		return ;
	}
	var students []std ;
	for cursor.Next(context.Background()) {
		var student std ;
		err := cursor.Decode(&student);
		if err != nil {
			http.Error(w , "Error decoding student data" , http.StatusInternalServerError)
			return ;
		}
		students = append(students , student);
	}
	cursor.Close(context.Background());
	w.Header().Set("Content-Type" , "application/json");
	json.NewEncoder(w).Encode(students);
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next(w, r)
    }
}

func deleteStudent(w http.ResponseWriter , r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed)
		return ;
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/students/");
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w , "Invalid student ID" , http.StatusBadRequest)
		return ;
	}
	result , err := stdcollection.DeleteOne(context.Background() , bson.M{"id": id});
	if err != nil {
		http.Error(w , "Error deleting student" , http.StatusInternalServerError)
		return ;
	}
	if result.DeletedCount == 0 {
		http.Error(w , "Student not found" , http.StatusNotFound)
		return ;
	}

}


func getStudentDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w , "method not allowed" , http.StatusMethodNotAllowed)
		return ;
	}
	
	idStr := strings.TrimPrefix(r.URL.Path, "/StudentDetail/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }
	
	
	var studentDetails StudentDetail ;
	var student std ;
	err = stdcollection.FindOne(context.Background() , bson.M{"id": id}).Decode(&student);
	if err != nil {
		http.Error(w , "student not found" , http.StatusNotFound)
		return ;
	}
	cursor , err := enrollcollection.Find(context.Background() , bson.M{"stdid": id});
	if err != nil {
		http.Error(w , "error fetching enrollments" , http.StatusInternalServerError)
		return ;
	}
	defer cursor.Close(context.Background());
	var coursenames []string ;
	for cursor.Next(context.Background()) {
		var enrollment enroll ;
		err := cursor.Decode(&enrollment);
		if err != nil {
			http.Error(w , "error decoding enrollment data" , http.StatusInternalServerError)
			return ;
		}
		var course course ;
		err = courscollection.FindOne(context.Background() , bson.M{"id": enrollment.CourseID}).Decode(&course);
		if err != nil {
			http.Error(w , "error fetching course data" , http.StatusInternalServerError)
			return ;
		}
		coursenames = append(coursenames , course.Name);
	}
	
	
	studentDetails.ID = student.ID ;
	studentDetails.Name = student.Name;
	studentDetails.Age = student.Age ;
	studentDetails.Courses = coursenames ;
	w.Header().Set("Content-Type" , "application/json");
	json.NewEncoder(w).Encode(studentDetails);
}

func getcourse(w http.ResponseWriter , r *http.Request) {
	if r.Method != "GET" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed)
		return ;
	}
	cursor , err := courscollection.Find(context.Background() , bson.M{});
	if err != nil {
		http.Error(w , "Error fetching courses" , http.StatusInternalServerError)
		return ;
	}
	var courses []course ;
	for cursor.Next(context.Background()) {
		var course course ;
		err := cursor.Decode(&course);
		if err != nil {
			http.Error(w , "Error decoding course data" , http.StatusInternalServerError)
			return ;
		}
		courses = append(courses , course);
	}
	cursor.Close(context.Background());
	w.Header().Set("Content-Type" , "application/json");
	json.NewEncoder(w).Encode(courses);
}

func getenroll(w http.ResponseWriter , r *http.Request) {
	if r.Method != "GET" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed)
		return ;
	}
	cursor , err := enrollcollection.Find(context.Background() , bson.M{});
	if err != nil {
		http.Error(w , "Error fetching enrollments" , http.StatusInternalServerError)
		return ;
	}
	var enrollments []enrollmentRequest ;
	for cursor.Next(context.Background()) {
		var enrollment enroll ;
		err := cursor.Decode(&enrollment);
		if err != nil {
			http.Error(w , "Error decoding enrollment data" , http.StatusInternalServerError)
			return ;
		}
		var student std ;
		err = stdcollection.FindOne(context.Background() , bson.M{"id": enrollment.StdID}).Decode(&student);
		if err != nil {
			http.Error(w , "Error fetching student data" , http.StatusInternalServerError)
			return ;
		}
		var course course ;
		err = courscollection.FindOne(context.Background() , bson.M{"id": enrollment.CourseID}).Decode(&course);
		if err != nil {
			http.Error(w , "Error fetching course data" , http.StatusInternalServerError)
			return ;
		}
		var enrollmentReq enrollmentRequest ;
		enrollmentReq.StdName = student.Name ;
		enrollmentReq.CourseName = course.Name ;
		enrollments = append(enrollments , enrollmentReq);
	}
	cursor.Close(context.Background());
	w.Header().Set("Content-Type" , "application/json");
	json.NewEncoder(w).Encode(enrollments);
}




func main() {
	ctx , cancel := context.WithTimeout(context.Background() , 10*time.Second);
	defer cancel();
	var err error;
	client , err = mongo.Connect(ctx , options.Client().ApplyURI("mongodb://localhost:27017"));
	if err != nil {
		fmt.Println("Error connecting to MongoDB: " , err);
		return ;
	}

	stdcollection = client.Database("stdmsDB").Collection("stdcollection");
	courscollection = client.Database("stdmsDB").Collection("courscollection");
	enrollcollection = client.Database("stdmsDB").Collection("enrollcollection");
	
	// Vérifier si l'étudiant existe déjà
	count, err := stdcollection.CountDocuments(context.Background(), bson.M{"id": 1})
	if count == 0 {
    	result , err := stdcollection.InsertOne(context.Background(), std{ID: 1, Name: "Mohamed", Age: 20});
		if err != nil {
			fmt.Println("Error inserting student: ", err);
			return;
		}
    	fmt.Println("Student inserted" , result.InsertedID);
	} else {
    	fmt.Println("Student already exists, skipping...")
	}	

	count , err = courscollection.CountDocuments(context.Background(), bson.M{"id": 1});
	if count == 0 {
			result , err := courscollection.InsertOne(context.Background() , course{ID: 1 , Name: "Math"});
				if err != nil {
					fmt.Println("error inserting course:", err)
					return;
				}
				fmt.Println("Course inserted" , result.InsertedID);
	}else{
		fmt.Println("Course already exists, skipping...")
	}
		
	count , err = enrollcollection.CountDocuments(context.Background(), bson.M{"stdid": 1});
	if count == 0 {
		result , err := enrollcollection.InsertOne(context.Background() , enroll{StdID: 1, CourseID: 1});
				if err != nil {
					fmt.Println("error inserting course:", err)
					return ;
				}	
		fmt.Println("Enrollment inserted" , result.InsertedID);
	}else{
		fmt.Println("Enrollment already exists, skipping...")
	}
	
	http.HandleFunc("/enrollments", enableCORS(getenroll));
	http.HandleFunc("/courses", enableCORS(getcourse));
	http.HandleFunc("/students", enableCORS(studentsHandler));
	http.HandleFunc("/students/", enableCORS(deleteStudent));
	http.HandleFunc("/StudentDetail/", enableCORS(getStudentDetails));

	fmt.Println("Server running on http://localhost:8080");
	 http.ListenAndServe(":8080", nil);

}