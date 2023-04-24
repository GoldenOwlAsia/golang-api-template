# Golden Owl Golang Gin API Template

<p align="center">
<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="180" alt="accessibility text">
</p>

### Why Using Gin For Golang Backend?

Gin allows you to build web applications and microservices in Go. It contains a set of commonly used functionalities (e.g., routing, middleware support, rendering, etc.) that reduce boilerplate code and make it simpler to build web applications.


### Prerequisites

1. Install go. You can download the Golang in this [page](https://go.dev/doc/install). You should install version 1.20
2. Install Postgres database. You can download the Postgres in this [page](https://www.postgresql.org/download/). You should install version 14.1
3. Make an `.env` file from `.env.example`
4. Go to `pgadmin` create a database. Note the name of it and add to `.env`
5. Install **Air - live reload fo Go apps**. You can visit this [page](https://github.com/cosmtrek/air).

### 💿 Installation

#### Via `go`

1. You run this command to install packages
   ```sh
   go mod download && go mod tidy
   ```
2. Create `.env` file from `.env.example` file.
3. run this command to start (hot reload):
   ```sh
   make watch
   ```
   run without hot reload
   ```sh
   make run
   ```
4. Visit: http://localhost:8080/swagger/index.html to access the API interface.
#### Via `docker`

1. Run by docker
   ```sh
   docker-compose up
   ```

### Build go executables
build api executable file
```sh
make build
```

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Golang]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://go.dev/doc/


#### 📌 Database Diagram
[![](https://mermaid.ink/img/pako:eNqtkk1PwzAMhv9K5HM3pZ90uU0bBw5IiLEL6iVqvRGpTabEFYyu_53QgTq2cUDCJ_uxI7-x3UFpKgQBaJdKbq1sCs28rR1ad-Ky7hh8Wqs0sbslO7Ux68gqvR1eadngb3n2IJ17Nba6zNw2UtWX-NHUeElXJKl1Iyfley4sSsJqTmd8vauu8iXWeML7Qh-duSVV1uh-Rn-eBXtSdE37wmhCTf8uftza4TCZHLpRuWCl7ymVdhBAg9ZPuvLLHz5UAL1ggwUI71a4kW1NBRS696WyJbPa6xIE2RYDaActXxcDYiNr5-lO6mdjmu8iH4Lo4A1EnOXTWZImYRjFPL7hWQB7EJMkn0ZpksY5z_Mw4rM-gPfhfTjlaZ7FnKchD7M8ytIAsFJk7P3xWoej7T8A9wvK7w?type=png)](https://mermaid.live/edit#pako:eNqtkk1PwzAMhv9K5HM3pZ90uU0bBw5IiLEL6iVqvRGpTabEFYyu_53QgTq2cUDCJ_uxI7-x3UFpKgQBaJdKbq1sCs28rR1ad-Ky7hh8Wqs0sbslO7Ux68gqvR1eadngb3n2IJ17Nba6zNw2UtWX-NHUeElXJKl1Iyfley4sSsJqTmd8vauu8iXWeML7Qh-duSVV1uh-Rn-eBXtSdE37wmhCTf8uftza4TCZHLpRuWCl7ymVdhBAg9ZPuvLLHz5UAL1ggwUI71a4kW1NBRS696WyJbPa6xIE2RYDaActXxcDYiNr5-lO6mdjmu8iH4Lo4A1EnOXTWZImYRjFPL7hWQB7EJMkn0ZpksY5z_Mw4rM-gPfhfTjlaZ7FnKchD7M8ytIAsFJk7P3xWoej7T8A9wvK7w)

#### 🔗 Converts Go annotations to Swagger Documentation 2.0
Run `swag init` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).
```sh
swag init
```


#### 💉 Dependency injection with Wire
1. change configuration in file ```wire.go ```
2. run the following command to automatically generate the code
```sh
make di
```

#### 🧪 testing
```sh
make test
# or
go test ./... -cover
# with cover
go test ./... -cover
# with verbose
go test -v ./... -cover
# specific folder
go test -v ./utils -cover
# specific test file
go test ./utils/array_test.go ./utils/array.go
# one unit test
# - api/utils is a package name
# - TestChunkSlice is a testcase
go test api/utils -run TestChunkSlice
```

#### 🧪 Improve code with lint checks
```sh
make lint
```

### Demo
#### Login
```sh
curl -L 'localhost:8080/api/v1/user/login' \
-H 'Content-Type: application/json' \
-d '{
  "username": "admin",
  "password": "1234"
}'
```
response:
```json
{
   "status": "success",
   "message": "welcome back",
   "data": {
      "user": {
         "id": 1,
         "created_at": "2023-04-19T14:32:21.978531Z",
         "updated_at": "2023-04-19T14:32:21.978531Z",
         "username": "admin",
         "email": "admin@example.com",
         "role": "Admin",
         "status": ""
      },
      "access_token": "xxx.xxx.xxx",
      "refresh_token": "yyy.yyy.yyy"
   }
}
```
#### Articles
```shell
curl -L 'localhost:8080/api/v1/articles' -H 'Authorization: Bearer xxx.xxx.xxx-xxx'
```
response:
```json
{
   "_metadata": {
      "limit": 10,
      "total": 54,
      "total_pages": 6,
      "per_page": 10,
      "page": 1,
      "sort": "created_at DESC"
   },
   "records": [
      {
         "id": 59,
         "created_at": "2023-04-23T10:38:29.80017Z",
         "updated_at": "2023-04-23T10:38:29.80017Z",
         "user": {
            "id": 1,
            "created_at": "2023-04-19T14:32:21.978531Z",
            "updated_at": "2023-04-19T14:32:21.978531Z",
            "username": "admin",
            "email": "admin@example.com",
            "role": "Admin",
            "status": ""
         },
         "title": "test",
         "content": "test content"
      }
   ]
}
```
