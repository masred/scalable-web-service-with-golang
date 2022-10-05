package session_6

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"testing"
)

func TestWebServerHTMLTemplate(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Masred",
			"Message": "Have a nice day",
		}

		var t, err = template.ParseFiles("./template/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})
	fmt.Println("Listening http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{
		ID:       1,
		Name:     "Masred",
		Age:      22,
		Division: "IT",
	},
	{
		ID:       2,
		Name:     "Clara",
		Age:      17,
		Division: "Finance",
	},
	{
		ID:       3,
		Name:     "Ujang",
		Age:      19,
		Division: "Security",
	},
}

var PORT = ":8000"

func TestWebServer(t *testing.T) {
	http.HandleFunc("/", greet)

	fmt.Println("Listening on http://localhost:8000")
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

func TestWebServerAPI(t *testing.T) {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Listening on http://localhost:8000")
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
