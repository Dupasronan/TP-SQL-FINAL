package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	EmployeeId   int    `json:"employee_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	BirthdayDate string `json:"birthday_date"`
	HireDate     string `json:"hire_date"`
	Mail         string `json:"mail"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	City         string `json:"city"`
	PostalCode   string `json:"postal_code"`
	JobId        int    `json:"job_id"`
	DepartmentId int    `json:"department_id"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./Entreprise.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/employees", employeesHandler)
	http.HandleFunc("/add-employee", addEmployeeHandler)
	http.HandleFunc("/delete-employee", deleteEmployeeHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var tmpl = template.Must(template.ParseFiles("employees.html"))

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	var rows *sql.Rows
	var err error

	if searchQuery != "" {
		query := `SELECT EmployeeId, FirstName, LastName, BirthdayDate, HireDate, Mail, Phone, Address, City, PostalCode 
                  FROM employees 
                  WHERE FirstName LIKE ? OR LastName LIKE ?`
		rows, err = db.Query(query, "%"+searchQuery+"%", "%"+searchQuery+"%")
	} else {
		rows, err = db.Query("SELECT EmployeeId, FirstName, LastName, BirthdayDate, HireDate, Mail, Phone, Address, City, PostalCode FROM employees")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.EmployeeId, &emp.FirstName, &emp.LastName, &emp.BirthdayDate, &emp.HireDate, &emp.Mail, &emp.Phone, &emp.Address, &emp.City, &emp.PostalCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}

	err = tmpl.Execute(w, employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")
		birthdayDate := r.FormValue("birthday_date")
		hireDate := r.FormValue("hire_date")
		mail := r.FormValue("mail")
		phone := r.FormValue("phone")
		address := r.FormValue("address")
		city := r.FormValue("city")
		postalCode := r.FormValue("postal_code")
		jobId := r.FormValue("job_id")
		departmentId := r.FormValue("department_id")

		_, err := db.Exec("INSERT INTO employees (FirstName, LastName, BirthdayDate, HireDate, Mail, Phone, Address, City, PostalCode, JobId, DepartmentId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			firstName, lastName, birthdayDate, hireDate, mail, phone, address, city, postalCode, jobId, departmentId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employees", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "gestion.html")
	}
}

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		employeeId := r.FormValue("employee_id")

		_, err := db.Exec("DELETE FROM employees WHERE EmployeeId = ?", employeeId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employees", http.StatusSeeOther)
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
