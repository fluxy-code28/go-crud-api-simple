# Go CRUD REST API with MySql

## Description

Project ini saya gunakan untuk melatih pemahaman go dengan menggunakan pendekatan project.

## Prerequisites

1. Go
2. MySql
3. Postman
   
## Package 

- github.com/go-sql-driver/mysql
- github.com/gorilla/mux

## Database

menggunakan database mysql dengan nama go_crud_api 
```sql
USE go_crud_api;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```



## Source

- [Medium/rameshfadatare-Golang CRUD REST API with MySQL](https://rameshfadatare.medium.com/golang-crud-rest-api-with-mysql-36713f43f470)