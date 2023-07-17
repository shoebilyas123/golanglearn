package main

// Model for courses - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}


type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
}


var courses = []Course{
	{
CourseId: "key1",
CourseName: "ReactJS",
Price: 250,
	},
}


// middleware - helper file
func isEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main(){}


// API: 
//  - Courses - CRUD with verifiying with schema
// 	- 