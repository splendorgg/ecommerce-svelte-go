<h1 align="center">E-commerce Site With Svelte and Go </h1>

# Overview
This project is a user authentication system developed with a Go backend and Svelte frontend, using modern tools and libraries for a seamless and efficient user experience. The application allows users to sign up, log in, and access protected routes through JWT-based authentication. The frontend is built with Svelte, a fast and lightweight framework that ensures smooth interactions and minimal overhead. The backend is developed using Go, known for its performance and scalability, providing a secure API for user authentication and management.

# How It Works
This diagram represents the flow of a request from the frontend to the backend in a typical user authentication process. It outlines the sequence of steps taken by the backend to process the request, validate the user data, interact with the database, and return a response containing a JWT (JSON Web Token) for authentication.

```text
+-------------------------------+
|   Frontend (Request)          |
|  - User sends a request       |
+-------------------------------+
               |
               v
+-------------------------------+
| Backend Router                |
|  - The request is routed to   |
|    the appropriate controller |
+-------------------------------+
               |
               v
+-------------------------------+
| Controller (Signup/Login)     |
|  - The request is processed,  |
|    data is retrieved          |
+-------------------------------+
               |
               v
+-------------------------------+
| Data Validation and Processing|
|  - Email and password are     |
|    validated, password is     |
|    hashed                     |
+-------------------------------+
               |
               v
+-------------------------------+
| Database Operations           |
|  - User is created in the     |
|    database                   |
|  - User is queried for login  |
+-------------------------------+
               |
               v
+-------------------------------+
| Token Generation (JWT)        |
|  - JWT token is created using |
|    user data                  |
+-------------------------------+
               |
               v
+-------------------------------+
| Response to Frontend          |
|  - Token is sent back in the  |
|    response                   |
+-------------------------------+
```


## Getting Started
```bash
# Clone this repository
git clone https://github.com/splendorgg/ecommerce-svelte-go.git
# Run the Go server
go run main.go
# Go into svelte directory and install dependencies
npm install
# Run the project
npm run dev
```

# Technologies Used
- **Svelte**(v4.2.12)
- **Go**(v1.22.3)
- **Flowbite**
- **Svelte-Spa-Router**
- **Axios**
- **TailwindCSS**
- **GORM**
- **JWT**

# Conclusion
This project demonstrates the power of combining Svelte for the frontend and Go for the backend, enabling developers to build fast, secure, and scalable applications. By integrating JWT for authentication and leveraging modern technologies like TailwindCSS, Flowbite, and GORM, the application delivers an intuitive user experience and efficient backend management.
