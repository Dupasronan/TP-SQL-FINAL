package main

import (
	"database/sql"
	"fmt"
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

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/employees", employeesHandler)
	http.HandleFunc("/add-employee", addEmployeeHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.EmployeeId, &emp.FirstName, &emp.LastName, &emp.BirthdayDate, &emp.HireDate, &emp.Mail, &emp.Phone, &emp.Address, &emp.City, &emp.PostalCode, &emp.JobId, &emp.DepartmentId)
		if err != nil {
			log.Println(err)
			continue
		}
		employees = append(employees, emp)
	}

	html := `<table class="table table-striped">
        <thead>
            <tr>
                <th>ID</th>
                <th>Prénom</th>
                <th>Nom</th>
                <th>Date de naissance</th>
                <th>Date d'embauche</th>
                <th>Email</th>
                <th>Téléphone</th>
                <th>Adresse</th>
                <th>Ville</th>
                <th>Code Postal</th>
            </tr>
        </thead>
        <tbody>`

	for _, emp := range employees {
		html += fmt.Sprintf(`
            <tr>
                <td>%d</td>
                <td>%s</td>
                <td>%s</td>
                <td>%s</td>
                <td>%s</td>
                <td>%s</td>
                <td>%s</td>
                <td>%s</td>
                <td>%s</td>
                <td>%s</td>
            </tr>`, emp.EmployeeId, emp.FirstName, emp.LastName, emp.BirthdayDate, emp.HireDate, emp.Mail, emp.Phone, emp.Address, emp.City, emp.PostalCode)
	}

	html += `</tbody></table>`
	w.Write([]byte(html))
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
		http.ServeFile(w, r, "add_employee.html")
	}
}
