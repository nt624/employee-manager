package main

import (
	"example/employee-manager/db"
	"example/employee-manager/router"
)

func main() {
	dbConn := db.Init()
	router.Router(dbConn)

}

/*
// EmployeeByDept queries for employee that have the specified department name.
func EmployeeByDept(name string) ([]Employee, error) {
	// An employee slice to hold data from returned rows.
	var employee []Employee

	rows, err := db.Query("SELECT * FROM employee WHERE dept = ?", name)
	if err != nil {
		return nil, fmt.Errorf("EmployeeByDept %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Dept, &emp.Age); err != nil {
			return nil, fmt.Errorf("EmployeeByDept %q: %v", name, err)
		}
		employee = append(employee, emp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("EmployeeByDept %q: %v", name, err)
	}
	return employee, nil
}

// employeeByID queries for the employee with the specified ID.
func employeeByID(id int64) (Employee, error) {
	// An employee to hold data from the returned row.
	var emp Employee

	row := db.QueryRow("SELECT * FROM employee WHERE id = ?", id)
	if err := row.Scan(&emp.ID, &emp.Name, &emp.Dept, &emp.Age); err != nil {
		if err == sql.ErrNoRows {
			return emp, fmt.Errorf("employeeByID %d: no such employee", id)
		}
		return emp, fmt.Errorf("employeeByID %d: %v", id, err)
	}
	return emp, nil
}

// addEmployee adds the specified employee to the database,
// returning the employee ID of the new entry
func addEmployee(emp Employee) (int64, error) {
	result, err := db.Exec("INSERT INTO employee (name, dept, age) VALUES (?, ?, ?)", emp.Name, emp.Dept, emp.Age)
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	return id, nil
}
*/
