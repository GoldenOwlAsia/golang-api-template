## Golden Owl Golang Gin API

## Getting Started

### Prerequisites

1. Install go. You can download the Golang in this [page](https://go.dev/doc/install). You should install version 1.20
2. Install Postgres database. You can download the Postgres in this [page](https://www.postgresql.org/download/). You should install version 14.1
3. Make an `.env` file from `.env.example`
4. Go to `pgadmin` create a database. Note the name of it and add to `.env`
5. Migrations & Seeding data

   create all database tables
   ```sh
   go run cmd/run_migration.go
   ```
   migrate and seed your database with dummy data
   ```sh
   go run cmd/run_migration.go --seed
   ```
6. Run command to create user for login

   `curl --location 'pathtohost' \
   --header 'Content-Type: application/json' \
   --data-raw '{
   "confirm_password": "123",
   "email": "admin@gmail.com",
   "password": "123",
   "username": "admin"
   }'`
7. Install **Air - live reload fo Go apps**. You can visit this [page](https://github.com/cosmtrek/air).

### 💿 Installation

#### Via `go`

1. You run this command to install packages
   ```sh
   go mod download && go mod tidy
   ```
2. Create `.env` file from `.env.example` file.
3. ▶ run this command to start (hot reload):
   ```sh
   make watch
   ```
   ▶ run without hot reload
   ```sh
   make run
   ```
4. Visit: http://localhost:8080/swagger/index.html to access the API interface.
#### Via `docker`

1. Run by docker
   ```sh
   docker-compose up
   ```

### ▶️  Build go executables
build api executable file
```sh
make build
```

run it
```sh
api
```

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Golang]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://go.dev/doc/


#### 📌 Database Diagram
[![](https://mermaid.ink/img/pako:eNp1klFrgzAQx79KuGcRNdVq3rq1D3sYjEpfhi-hubaCJhKTbU787kutW0fFg8Dd7x_-yeXSw1EJBAaotyU_a14Xkrg45Lv9PSP9Lb-GLaUhL4L8j7vaGl3KMzm0qCWvcUknb7xtP5UWc2VX87Ka472qcE5zw41t53zTNFp9oFjQnzVyg-Kpm1-8EQvKFit8UEzpOpy8NuaBT04zPvn88gE8qFG7poWbwvjOBZgL1lgAc6nAE7eVKaCQ163cGpV38gjsxKsWPbDjMdPs_mjD5btSrjba3kpgPXwBo5T6YZwGiVthkAZp7EEHbBX5MXUgWaU0joI1HTz4Hg0CP4uzLA6zLEkDGqV07QGK0ij9evs44_8ZfgA7ZKoW?type=png)](https://mermaid.live/edit#pako:eNp1klFrgzAQx79KuGcRNdVq3rq1D3sYjEpfhi-hubaCJhKTbU787kutW0fFg8Dd7x_-yeXSw1EJBAaotyU_a14Xkrg45Lv9PSP9Lb-GLaUhL4L8j7vaGl3KMzm0qCWvcUknb7xtP5UWc2VX87Ka472qcE5zw41t53zTNFp9oFjQnzVyg-Kpm1-8EQvKFit8UEzpOpy8NuaBT04zPvn88gE8qFG7poWbwvjOBZgL1lgAc6nAE7eVKaCQ163cGpV38gjsxKsWPbDjMdPs_mjD5btSrjba3kpgPXwBo5T6YZwGiVthkAZp7EEHbBX5MXUgWaU0joI1HTz4Hg0CP4uzLA6zLEkDGqV07QGK0ij9evs44_8ZfgA7ZKoW)
#### 🔗 Converts Go annotations to Swagger Documentation 2.0
Run `swag init` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).
```sh
swag init
```


#### 💉 Dependency injection with Wire
1. change configuration in file ```wire.go ```
2. run the following command to automatically generate the code
```sh
wire
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
#or
golangci-lint run -v
```