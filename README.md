# gobook

### Structure of Project

#### 1. Folder

###### a. cmd (project)

###### b. pkg (package)

- **config**: connect with database
- **controllers**: process data and response
- **models**: create structs and models used for database
- **routes**: create routing
- **utils**: marshall, unmarshall for json request, response

#### 2. Routing

| controller function | route          | method |
| :------------------ | :------------- | :----- |
| GET BOOK            | /book/         | GET    |
| CREATE BOOK         | /book/         | POST   |
| GET BOOK BY ID      | /hook/{bookID} | GET    |
| UPDATE BOOK         | /book/{bookID} | PUT    |
| DELETE BOOK         | /book/{bookID} | DELETE |

#### 3. Tech stack

- **Database:** MySQL
- **ORM**: GORM
- **JSON**: marshall, unmarshall
- **Routing**: Gorilla Mux

#### Step coding

routes ->
