# Api-Cart

Mini Service Shopping Cart with Gorilla/Mux HTTP Respons Framework, GORM for Object relation model, PostgreSQL for database.

## 🔗 Description

This Backend Service Shopping Cart is used for simple add product to cart.
There are 3 main routers :

1. Add Cart (Can add product if product is no exist also update quantity product if name product exist)
   > Attributes ProductCode, ProductName, Quantity
2. Delete Product (Delete by product code)
3. Get Product (Get all product, Get by product name, Get by quantity)

## Notes :

1. For answers no. 01 & 02, i put it in a folder task1 & task2, change directory to folder and run

```bash
go run .
```

2. For answer no. 03 - 05 i put in a folder task_3-5, open the .md file

## 🛠️ Installation Steps

1. Clone the repository

```bash
https://github.com/zazhedho/telkom-test.git
```

2. Install dependencies

```bash
go mod tidy
```

> Wait a minute, if still error run :

```bash
go mod vendor
```

3. Add Env File

```sh
    APP_PORT = Your DB Port
    DB_HOST = Your DB Host
    DB_NAME = Your DB Name
    DB_USER = Your DB User
    DB_PASS = Your DB Password
```

4. Database Migration and Rollback

```bash
go run main.go migrate --up //for database migration table
# or
go run main.go migrate --down //for rollback the database
```

5. Run the app

```bash
go run . serve
```

### 🚀 You are all set

## 🔗 REST API Endpoints

### GET /api/v1/products

> Get Data Cart

_Request Body_

```
not needed
```

_Request Query Params_

```
not needed
```

### POST /api/v1/products

> Post Data Cart

_Request Body_

```
{
    "kode_product": "ATK-2",
    "name": "Pensil",
    "kuantitas": 12
}
```

_Request Query Params_

```
no need
```

### PUT /api/v1/products/{id}

> Update Data Cart

_Request Body_

```
{
    "kuantitas": 10
}
```

_Request Query Params_

```
no need
```

### DELETE /api/v1/products/{kode}

> Delete Data Cart by Product Code

_Request Body_

```
no need
```

_Request Query Params_

```
no need
```

### GET /api/v1/products/name

> Search Data Cart by product name

_Request Body_

```
no need
```

_Request Query Params_

```
name = (input product name)
```

### GET api/v1/products/qty

> Search Data Cart by quantity

_Request Body_

```
no need
```

_Request Query Params_

```
kuantitas = (input quantity)
```

## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [gorilla/mux](https://github.com/gorilla/mux): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS

## 🚀 About Me

- Linkedin : [Zaidus Zhuhur](https://www.linkedin.com/in/zaidus-zhuhur/)
