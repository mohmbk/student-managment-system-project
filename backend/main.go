package main 
import ("fmt" ; "net/http" ; "encoding/json" ; "time" ; "context" ; "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" ;  )

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
	stdid int `json:"stdid" bson:"stdid"`
	courseid int `json:"courseid" bson:"courseid"`
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


	http.HandleFunc("/students", enableCORS(getstd));

	fmt.Println("Server running on http://localhost:8080");
	 http.ListenAndServe(":8080", nil);

}