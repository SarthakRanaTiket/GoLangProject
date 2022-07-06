package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SarthakRanaTiket/projectname/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	Id int `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//create connection with postgresql :/

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil{
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the db!")
	return db
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	users,err := getAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all users.%v", err)
	}
	json.NewEncoder(w).Encode(users)
}

//handlers ----------------->>>>>>>>>>>>>>>

func getAllUsers()([]models.User,error){
	db := createConnection()

	defer db.Close()

	var users []models.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query.%v", err)
	}

	defer rows.Close()

	for rows.Next(){
		var user models.User
		err = rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Email, &user.Age)
		if err != nil {
			log.Fatalf("unable to scan the row.%v", err)
		}
		
		users = append(users, user)
	}
	return users, err
}