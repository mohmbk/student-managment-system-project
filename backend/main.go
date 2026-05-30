package main 
import ("fmt" ; "net/http" ; "encoding/json" ; "time" ; "context" ; "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" ; "strconv"
    "strings" )

var client *mongo.Client
var stdcollection *mongo.Collection
var courscollection *mongo.Collection
var enrollcollection *mongo.Collection

type std struct {
	id int `json:"id" bson:"id"`
	name string `json:"name" bson:"name"`
	age int `json:"age" bson:"age"`
}

type course struct {
	id int `json:"id" bson:"id"`
	name string `json:"name" bson:"name"`
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

}


func getstd(w http.ResponseWriter , r *http.Request) {
	if r.Method != "GET" {
		http.Error(w , "Method not allowed" , http.StatusMethodNotAllowed)
		return ;
	}
	cursor , err := stdcollection.Find(context.Background() , bson.M{});
	if err != nil {
		http.Error(w , "Error fetching students" , http.StatusInternalServerError)
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

	w.Header().Set("Content-Type" , "application/json");
	json.NewEncoder(w).Encode(students);
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

	stdcollection = client.Database("stdmsDB").Collection("prayercollection");
	courscollection = client.Database("stdmsDB").Collection("courscollection");
	enrollcollection = client.Database("stdmsDB").Collection("enrollcollection");
	
	result , err := stdcollection.InsertOne(context.Background() , std{id: 1 , name: "Mohamed" , age: 20});
	if err != nil {
		fmt.Println("Error inserting student: " , err);
		return ;
	}

fmt.Printf("Inserted student with ID: %v\n" , result.InsertedID);
	http.HandleFunc("/students" , getstd);
	fmt.Println("Server running on http://localhost:8080");
	 http.ListenAndServe(":8080", nil);

}