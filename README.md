# Bin Management System

## Overview

The Bin Management System is a backend application designed to efficiently manage and track bins, inventory, and related operations. Built with Go (Golang), the system leverages a modular architecture with clearly defined controllers, models, and routes to handle different aspects of the bin management process. This project also includes database initialization, validation mechanisms, and route management to ensure smooth data flow.

## Features

- **Bin Management**: Handle bin creation, updates, and removal.
- **Database Integration**: Supports easy integration with databases.
- **Validation**: Built-in data validation to ensure consistent and reliable input.
- **Modular Structure**: Organized into separate folders for controllers, models, routes, and more.
- **Scalable**: Designed with scalability in mind, easily extendable for future enhancements.

## Technologies Used

- **Go (Golang)**: Backend logic and routing
- **PostgreSQL/MySQL**: Database management (Depending on your configuration)
- **Git**: Version control
- **Air**: Live reloading for Go applications

## Installation

### 1. Clone the repository:

   ```bash
   git clone https://github.com/Ya-s-h/Bin-Management.git
   cd Bin-Management
   ```

### 2. Initialize GO Environment:

   ```bash
    go mod tidy
```

### 3. Start Server
> There are two ways to start the server.
- Directly Running `main.go` using `go run .` or `go run main.go`
- Using `air`
  
### 4. Initialise DB:
> Initialise db by using API endpoint `/api/init/db`



## APIs
### Create User `POST` `/api/create/user`
Payload:
```json
{
	username: str,
	role_id: int,
	email_address: str,
	password: str,

}
```

### Delete User `POST` `/api/delete/user?user_id={}`

### Update User `POST` `/api/update/user?user_id={}`
Payload:
```json
{
	username: str,
	role_id: int,
	email_address: str,
	password: str,

}
```

### Add Waste `POST` `/api/add/waste`
Payload:
```json
{
	bin_id: int,
	weight_in_kgs: int

}
```
### Create Bin `POST` `/api/create/bin`
Payload:
```json
{
	
	area_id: int,
	user_id: int,

}
```

### Delete Bin `POST` `/api/delete/bin`
Payload:
```json
{
	bin_id: int,
}
```


### Assign bin to area `POST` `/api/assign_bin/area`
Payload:
```json
{
	bin_id: int,
	area_id: int

}
```

### Assign bin to user `POST` `/api/assign_bin/user`
Payload:
```json
{
	bin_id: int,
	user_id: int

}
```


### Create Area `POST` `/api/create/area`
Payload:
```json
{
	name: str,
	location: str,
	user_id: int,

}
```

### Delete Area `POST` `/api/delete/area`
Payload:
```json
{
	area_id: int,
}
```


### Assign area to user `POST` `/api/assign_area/user`
Payload:
```json
{
	user_id: int,
	area_id: int

}
```
