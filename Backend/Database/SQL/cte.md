
### **1. Common Table Expressions (CTE)**

A **Common Table Expression (CTE)** is a temporary result set in a SQL query, which can be referenced within `SELECT`, `INSERT`, `UPDATE`, or `DELETE` statements. CTEs make complex queries easier to read and write.

#### **Syntax**:
```sql
WITH cte_name AS (
    -- CTE query here
)
SELECT * FROM cte_name;
```

#### **Example**:

Let’s consider a simple example where we have a table `employees`:

| id  | name     | department  | salary |
| --- | -------- | ----------- | ------ |
| 1   | John     | IT          | 60000  |
| 2   | Jane     | HR          | 50000  |
| 3   | Jim      | IT          | 65000  |
| 4   | Jack     | Finance     | 70000  |

Now, let’s say we want to find the employees with salaries greater than the average salary in the company, using a CTE.

```sql
WITH average_salary AS (
    SELECT AVG(salary) AS avg_salary FROM employees
)
SELECT name, salary 
FROM employees, average_salary 
WHERE salary > average_salary.avg_salary;
```

#### **Go Example**:

Here’s how you can execute the above query using Go and `database/sql`:

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "user=username password=password dbname=company sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query using CTE
	rows, err := db.Query(`
		WITH average_salary AS (
			SELECT AVG(salary) AS avg_salary FROM employees
		)
		SELECT name, salary 
		FROM employees, average_salary 
		WHERE salary > average_salary.avg_salary;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		var name string
		var salary float64
		if err := rows.Scan(&name, &salary); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Employee: %s, Salary: %.2f\n", name, salary)
	}

	// Check for any error after iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
```
