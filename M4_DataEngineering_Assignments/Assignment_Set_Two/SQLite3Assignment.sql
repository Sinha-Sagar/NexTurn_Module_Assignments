------------------------SQLite3 Exercises: Employee Management System-----------------------------------

//	CREATING DATABASE AND COLLECTIONS
		
		sqlite3 EmployeeManagementDB.db
		
//	CREATING AND INSERTING SAMPLE DATA IN TABLE
		
		CREATE TABLE Employees (
			EmployeeID INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL,
			HireDate DATE NOT NULL,
			Salary REAL NOT NULL,
			DepartmentID INTEGER,
			FOREIGN KEY (DepartmentID) REFERENCES Departments (DepartmentID)
		);
		
		CREATE TABLE Departments (
			DepartmentID INTEGER PRIMARY KEY AUTOINCREMENT,
			DepartmentName TEXT NOT NULL
		);
		
		INSERT INTO Departments (DepartmentName) VALUES ('HR'), ('Finance'), ('IT');

		INSERT INTO Employees (Name, HireDate, Salary, DepartmentID)
		VALUES 
			('Alice', '2021-01-15', 70000, 1),
			('Bob', '2020-03-10', 60000, 2),
			('Charlie', '2022-05-20', 80000, 1),
			('Diana', '2019-07-25', 75000, 3)
			
//	QUERIES

1.	Write a query to list the names of employees hired after January 1, 2021. 
		
		SELECT Name, HireDate 
		FROM Employees 
		WHERE HireDate > '2021-01-01';

2.	Write a query to calculate the average salary of employees in each department.

		SELECT DepartmentID, AVG(Salary) AS AverageSalary
		FROM Employees
		GROUP BY DepartmentID;

		
3.	Write a query to find the department name where the total salary is the highest.

		SELECT d.DepartmentName
		FROM Departments d
		JOIN (
			SELECT DepartmentID, SUM(Salary) AS TotalSalary
			FROM Employees
			GROUP BY DepartmentID
		) e ON d.DepartmentID = e.DepartmentID
		ORDER BY e.TotalSalary DESC
		LIMIT 1;

 
4. 	Write a query to list all departments that currently have no employees assigned. 

		SELECT DepartmentName
		FROM Departments
		WHERE DepartmentID NOT IN (SELECT DISTINCT DepartmentID FROM Employees);


5. Write a query to fetch all employee details along with their department names.

		SELECT e.*, d.DepartmentName
		FROM Employees e
		JOIN Departments d ON e.DepartmentID = d.DepartmentID;




