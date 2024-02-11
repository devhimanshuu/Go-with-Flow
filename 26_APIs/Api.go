package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//model for courses - file
type Course struct {
	CourseId string `json:"courseid`
	CourseName string `json:"coursename`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

//fake DB 
var courses []Course

//middleware, helper - file
func(c *Course) IsEmpty() bool{
	// return c.CourseId==""&&c.CourseName==""
	return c.CourseName==""
}
type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//controllers 

//serve home route

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte ("<h1>Welcome to Apis in Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type" , "application/json")

	//grab id from request
	params:=mux.Vars(r)

	//loop through courses , find matching id and return responses
	for _,course:=range courses{
		if course.CourseId==params["id"]{
		json.NewEncoder(w).Encode(course)
		return 
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return 
}
func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type" , "application/json")

	//what if: Body is empty
	if r.Body==nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about -{}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id , String
	//append course into courses

}










func main() {
	fmt.Println("Build APIs in Golang")


}
