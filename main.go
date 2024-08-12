package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type employee struct{
	ID		string	`json: "id"`
	Name 	string	`json: "name"`
	Title	string	`json: "title"`
	Status	string	`json: "status"`
}

//Pseudo Database
var employees = []employee{
	{ID: "1", Name: "Naru Muraleedharan", Title: "CEO", Status: "active"},
	{ID: "2", Name: "Lindsay Muraleedharan", Title: "Partner", Status: "active"},
	{ID: "3", Name: "Kyle Johnson", Title: "Lead Developer", Status: "active"},
	{ID: "4", Name: "Melissa Sutphin", Title: "Executive Assistant", Status: "active"},
}

//Routes
func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

func getEmployeeById(c *gin.Context) {
	id := c.Param("id")
	employee, err := employeeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, employee)
}

func employeeById(id string) (*employee, error) {
	for i, e := range employees {
		if e.ID == id {
			return &employees[i], nil
		}
	}

	return nil, errors.New("Employee not found")
}

func createEmployee(c *gin.Context) {
	var newEmployee employee

	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}

	employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func removeEmployeeById(c *gin.Context) {
	id := c.Param("id")

	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			c.IndentedJSON(http.StatusOK, e)
			return 
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}

func changeStatusById(c *gin.Context) {
	id := c.Param("id")
	status, ok := c.GetQuery("status")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing status query parameter."})
		return
	}

	employee, err := employeeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	employee.Status = status
	c.IndentedJSON(http.StatusOK, employee)
}


//Endpoints
func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", getEmployeeById)
	router.POST("/employees", createEmployee)
	router.DELETE("/remove/:id", removeEmployeeById)
	router.PATCH("/status/:id", changeStatusById)
	router.Run(":8080")
}