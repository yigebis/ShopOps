# ShopOps API Documentation

## Overview
The `ShopOps` API allows shop owners to register, log in, verify emails, manage employees, and activate accounts. All requests and responses use JSON format, and the API requires authentication via JWT tokens for specific routes.

### Base URL
`http://localhost:8080/`

---

## Endpoints

### **User Registration**

#### `POST /register`
Registers a new user (shop owner).

- **Request Body:**
  ```json
  {
    "first_name": "string",
    "last_name": "string",
    "sex" : "string",
    "phone_number" : "string",
    "email": "string",
    "password": "string"
  }

- **Validation:**
  -first_name : required, minimum_length = 1, maximum_length = 50
  -last_name : required, minimum_length = 1, maximum_length = 50
  -sex : required, can be of types "M" and "F"
  -phone_number : required, length=13, must start with '+', then all should be numbers
  -email : required, must be an email
  -password : required, length >= 8, lower_case >= 2, upper_case >= 1, special >= 1


- **Responses:**
    -200 OK - Registration successful, verification email sent.
    -400 Bad Request - Invalid request payload.
    -409 Conflict - User already exists or pending verification.


### **User Login**

#### `POST /login`
Logins a user (shop owner/employee).

- **Request Body:**
  ```json
  {
    "email": "string",
    "password": "string"
  }

- **Validation:**
  "email" : required
  "password" : required


- **Responses:**
    -200 OK - Returns JWT access and refresher tokens.
    ```json
    {
        "token" : "string",
        "refresher" : "string"
    }
    -400 Bad Request - Invalid request payload
                     - account not activated
                     - account not verified 
                     - invalid email or password

### **Email verification**

#### `POST /verify`
verifies a registered user (shop owner).

- **Query Parameters:**
  -email
  -token

- **Responses:**
    -200 OK - Email verified successfully.
    -400 Bad Request - Invalid verification token.
    -409 Conflict - Email already verified.


### **All Employees Fetch**

#### `GET /employees`
Returns a list of employees with their data.

- **Headers:**
    - Authorization : Bearer <token>

- **Responses:**
    -200 OK - List of employees.
    -401 Unauthorized - Invalid or missing token.


#### **Get Specific Employee**

#### `GET /employee/:email`
Retrieves details of a specific employee. Requires JWT authentication and ownership.

- **Headers:**
    -Authorization: Bearer <token>

- **Path Parameters:**
    -email: Email of the employee.

- **Responses:**
    -200 OK - Employee details.
    -401 Unauthorized - Invalid or missing token.
    -404 Not Found - Employee not found.

#### **Add Employee**

#### `POST /employee/add`
Adds a new employee to the shop. Requires JWT authentication and ownership.

- **Headers:**
    -Authorization: Bearer <token>

- **Request Body:**
    ```json
    {
        "first_name": "string",
        "last_name": "string",
        "sex" : "string",
        "phone_number" : "string",
        "email": "string",
        "password": "string"
    }

- **Validation:**
    "first_name" : required, minimum_length = 1, maximum_length = 50
    "last_name" : required, minimum_length = 1, maximum_length = 50
    "sex" : required, can be of types "M" and "F"
    "phone_number" : required, length=13, must start with '+', then all should be numbers
    "email" : required, must be an email
    "password" : required, length >= 8, lower_case >= 2, upper_case >= 1, special >= 1

- **Responses:**
    201 Created - Employee added successfully.
    400 Bad Request - Invalid request payload.
    409 Conflict - Employee already exists.

### **Edit Employee**

#### `PUT /employee/edit`
Edits an existing employee's information. Requires JWT authentication and ownership.

- **Headers:**
    -Authorization: Bearer <token>

- **Request Body:**
    ```json
    {
        "email": "string",
        "first_name": "string",
        "last_name": "string"
    }

- **Responses:**
    -200 OK - Employee updated successfully.
    -400 Bad Request - Invalid request payload.
    -404 Not Found - Employee not found.

### **Delete Employee**
#### `POST /employee/delete/:email`
Deletes an employee by their email. Requires JWT authentication and ownership.

- **Headers:**
    -Authorization: Bearer <token>

- **Path Parameters:**
    -email: Email of the employee.

- **Responses:**
    -200 OK - Employee deleted successfully.
    -404 Not Found - Employee not found.
    -401 Unauthorized - Invalid or missing token.

### **Activate Employee Account**
#### `POST /activate`
Activates an employee's account after receiving a temporary password.

- **Request Body:**

    ```json
    {
        "email": "string",
        "old_password": "string",
        "new_password": "string"
    }
- **Responses:**
    -200 OK - Account activated successfully.
    -400 Bad Request - Invalid request payload or temporary password.
    -401 Unauthorized - Invalid credentials.

