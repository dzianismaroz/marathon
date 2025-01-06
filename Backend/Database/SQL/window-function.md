
---

### **2. Window Functions**

**Window functions** allow you to perform calculations across a set of table rows that are related to the current row, within a defined window. This is useful for running totals, ranking, and more.

#### **Syntax**:
```sql
SELECT column1, column2, 
       WINDOW_FUNCTION() OVER (PARTITION BY column3 ORDER BY column4)
FROM table_name;
```

Common window functions include:
- `ROW_NUMBER()`: Assigns a unique row number to each row within a partition.
- `RANK()`/`DENSE_RANK()`: Assigns ranks to rows in a partition.
- `LEAD()`/`LAG()`: Accesses data from subsequent or preceding rows.

#### **Example**:

Suppose we want to rank employees based on their salary within each department.

```sql
SELECT name, department, salary,
       RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS salary_rank
FROM employees;
```

This query will rank employees within each department by their salary, with the highest-paid employee receiving rank 1.

#### **Go Example**:

Here’s how you can execute the window function query in Go:

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

	// Window function query to rank employees
	rows, err := db.Query(`
		SELECT name, department, salary,
               RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS salary_rank
		FROM employees
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		var name, department string
		var salary float64
		var rank int
		if err := rows.Scan(&name, &department, &salary, &rank); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Employee: %s, Department: %s, Salary: %.2f, Rank: %d\n", name, department, salary, rank)
	}

	// Check for any error after iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
```

---

### **3. Optimizing Pagination Queries: Cursor-Based Pagination**

Pagination is essential when dealing with large datasets. Traditional **offset-based pagination** (e.g., `LIMIT 10 OFFSET 20`) can be inefficient for large datasets as it requires scanning through all the previous rows. **Cursor-based pagination** is more efficient as it uses the last retrieved row’s identifier (cursor) to fetch the next set of results.

#### **Cursor-based Pagination**:
Cursor pagination uses the **unique identifier** (usually a primary key) to keep track of the last item fetched. Instead of using `OFFSET`, the query is executed with a condition like `WHERE id > last_id`.

#### **Example**:

Let’s assume we want to paginate through the `employees` table based on `id`.

```sql
SELECT id, name, department, salary
FROM employees
WHERE id > ?  -- cursor is the last id from the previous page
ORDER BY id
LIMIT 10;
```

In the query above, the `?` will be replaced with the ID of the last employee from the previous page. The query will return the next 10 employees with higher IDs.

#### **Go Example**:

Here’s how you can implement cursor-based pagination in Go:

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func getEmployeesPage(db *sql.DB, lastID int, limit int) {
	query := `
		SELECT id, name, department, salary
		FROM employees
		WHERE id > $1
		ORDER BY id
		LIMIT $2;
	`

	rows, err := db.Query(query, lastID, limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, department string
		var salary float64
		if err := rows.Scan(&id, &name, &department, &salary); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %.2f\n", id, name, department, salary)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "user=username password=password dbname=company sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// First page (starting with ID = 0)
	getEmployeesPage(db, 0, 10)

	// Second page (starting with the last ID from the first page)
	getEmployeesPage(db, 10, 10)
}
```

In the `getEmployeesPage` function:
- We pass the `lastID` to fetch the next page of results.
- This ensures that the pagination is more efficient, especially for large datasets.

---

### **4. Other Essential Topics**

#### **Joins**:
SQL Joins allow you to combine rows from two or more tables based on a related column.

- **INNER JOIN**: Returns rows when there is a match in both tables.
- **LEFT JOIN (OUTER JOIN)**: Returns all rows from the left table, and matching rows from the right table.
- **RIGHT JOIN (OUTER JOIN)**: Returns all rows from the right table, and matching rows from the left table.
- **FULL JOIN**: Returns all rows when there is a match in one of the tables.

#### **Example**: 
```sql
SELECT employees.name, departments.name
FROM employees
INNER JOIN departments ON employees.department_id = departments.id;
```

#### **Indexes**:
Indexes are used to speed up query performance by allowing the database to find rows more efficiently. Creating indexes on columns frequently used in `WHERE`, `JOIN`, or `ORDER BY` clauses can greatly improve performance.

#### **Transactions**:
A transaction is a sequence of operations that are executed as a single unit. In SQL, you can use `BEGIN`, `COMMIT`, and `ROLLBACK` to handle transactions.

```sql
BEGIN;
-- Perform some updates
COMMIT; -- or ROLLBACK in case of error
```

#### **Aggregation Functions**:
SQL offers aggregation functions such as `COUNT()`, `SUM()`, `AVG()`, `MIN()`, `MAX()`, which can be used to perform operations on a set of rows and return a single result.

```sql
SELECT department, AVG(salary) FROM employees GROUP BY department;
```
