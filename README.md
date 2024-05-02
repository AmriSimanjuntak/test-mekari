This Project Contains All Configuration, Enums, etc to be sharing each of microservices

Prerequisites
1. Install latest golang version. Visit https://golang.org/doc/install for detail instructions.
2. Add bin folder from GOPATH to your environment. 
   For Mac User, add "export PATH=$HOME/go/bin:$PATH" to $HOME/.zshrc
   For Windows User, please visit https://www.architectryan.com/2018/08/31/how-to-change-environment-variables-on-windows-10
   For Linux User, add "export PATH=$HOME/go/bin:$PATH" to $HOME/.bash_profile
3. Install docker. 


DOCKER DB
1. Install Docker: Make sure Docker is installed on your system. You can download and install Docker Desktop from the official website
2. Run PostgreSQL Container: Open a terminal and run the following command to start a PostgreSQL container:
   docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
3. Connect to PostgreSQL: You can connect to the PostgreSQL container using any PostgreSQL client tool or by running a new container with psql
   docker run -it --rm --link postgres:postgres postgres psql -h postgres -U postgres
4. Create Table with Columns: Once connected to PostgreSQL, you can create the employees table with the specified columns:
   CREATE TABLE employees (
      id SERIAL PRIMARY KEY,
      first_name TEXT,
      last_name TEXT,
      email TEXT,
      hire_date DATE
   );
5. To start the PostgreSQL container, navigate to the directory containing the docker-compose.yml
   Build the Golang Application Docker Image:
   docker-compose build
   Start the Containers: Once the Docker image for your Golang application is built, you can start the containers using the following command:
   docker-compose up -d
   Access Your Golang Application:
   Your Golang application should now be running as a Docker container, along with the PostgreSQL container.
   You can access your application at http://localhost:8080



Spec Doc
Endpoints
1. Get Employee List
      Description: Retrieve a list of all employees from the employees table.
      Method: Post
      URL: http://localhost:8080/employee/v1/list
      Response:
      Status Code: 200 OK
      Content Type: application/json
      Body: JSON array containing employee objects.
      Error Handling:
      If no employees are found, return an empty array with status code 200 OK.
      If an error occurs during retrieval, return an appropriate error response with the corresponding status code.
2. Get Employee By Id
      Description: Retrieve an employee by their id.
      Method: GET
      URL:http://localhost:8080/employee/v1/view{id}
      Parameters:
      {id} (path parameter): The unique identifier of the employee.
      Response:
      Status Code: 200 OK
      Content Type: application/json
      Body: JSON object representing the employee details.
      Error Handling:
      If the specified employee id does not exist, return a 404 Not Found error.
      If an error occurs during retrieval, return an appropriate error response with the corresponding status code.
3. Create Employee Data
      Description: Create a new employee record in the database.
      Method: POST
      URL: http://localhost:8080/employee/v1/add
      Request Body: JSON object representing the new employee details.
      Response:
      Status Code: 201 Created
      Content Type: application/json
      Body: JSON object representing the newly created employee details, including the auto-generated id.
      Error Handling:
      If the request body is invalid or missing required fields, return a 400 Bad Request error.
      If an error occurs during creation, return an appropriate error response with the corresponding status code.
4. Update Employee Data
      Description: Update an employee record by their id.
      Method: PUT
      URL: http://localhost:8080/employee/v1/edit
      Request Body: JSON object representing the updated employee details.
      Response:
      Status Code: 200 OK
      Content Type: application/json
      Body: JSON object representing the updated employee details.
      Error Handling:
      If the specified employee id does not exist, return a 404 Not Found error.
      If the request body is invalid or missing required fields, return a 400 Bad Request error.
      If an error occurs during update, return an appropriate error response with the corresponding status code.
5. Delete Employee Data
      Description: Delete an employee record by their id.
      Method: DELETE
      URL:http://localhost:8080/employee/v1/delete{id}
      Parameters:
      {id} (path parameter): The unique identifier of the employee to be deleted.
      Response:
      Status Code: 204 No Content
      Error Handling:
      If the specified employee id does not exist, return a 404 Not Found error.
      If an error occurs during deletion, return an appropriate error response with the corresponding status code.

Error Handling
   All endpoints should implement appropriate error handling mechanisms to provide meaningful error messages and status codes in case of errors.
   Use standard HTTP status codes to indicate the success or failure of the API requests.
   Include detailed error messages in the response body to help clients understand and troubleshoot issues.