package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-nouveau-postgres-api/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {

	//load .env file

	err := godotenv.Load(".env")

	if err != nil {

		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	//db, err = Open("mysql", "gorm:gorm@/gorm?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}

	// Check the connection
	err = db.Ping()

	if err != nil {

		panic(err)

	}

	fmt.Println("Successfully connected ")

	// return the connection

	return db
}

func CreateVenture(w http.ResponseWriter, r *http.Request) {

	var venture models.Venture

	err := json.NewDecoder(r.Body).Decode(&venture)

	if err != nil {

		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// Call insert venture function and pass the venture

	insertID := insertVenture(venture)

	// format a response object

	res := response{
		ID:      insertID,
		Message: "Venture created successfully",
	}

	// Send the response
	json.NewEncoder(w).Encode(res)

}

func GetVenture(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {

		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// Call the getVenture function with venture  id to retrieve a single Venture

	venture, err := getVenture(int64(id))

	if err != nil {

		log.Fatalf("Unable to get venture, %v", err)
	}

	json.NewEncoder(w).Encode(venture)

}

func GetAllVentures(w http.ResponseWriter, r *http.Request) {

	// Get all the ventures in the DB

	ventures, err := getAllVentures()

	if err != nil {

		log.Fatalf("Unable to get all ventures, %v", err)
	}

	// Send all the vendorts as response

	json.NewEncoder(w).Encode(ventures)

}

func UpdateVenture(w http.ResponseWriter, r *http.Request) {

	// Get the venture id from the request params , key is id

	params := mux.Vars(r)

	// Convert the id type from string to int

	id, err := strconv.Atoi(params["id"])

	if err != nil {

		log.Fatalf("Unable to convert the string into int, %v", err)
	}

	// Create empty venture of type models.Venture

	var venture models.Venture

	// decode the json request to venture

	err = json.NewDecoder(r.Body).Decode(&venture)

	if err != nil {

		log.Fatalf("Unable to decode the request body, %v", err)
	}

	// Call updateVenture() function to update the venture

	updateRows := updateVenture(int64(id), venture)

	// format the message string

	msg := fmt.Sprintf("Stock updated successfulluy. total rows/record afffected by", updateRows)

	// format the response message

	res := response{

		ID:      int64(id),
		Message: msg,
	}

	// send the response

	json.NewEncoder(w).Encode(res)

}

func DeleteVenture(w http.ResponseWriter, r *http.Request) {

	// Get the venture id from the request params

	params := mux.Vars(r)

	// convert the id in string to int

	id, err := strconv.Atoi(params["id"])

	if err != nil {

		log.Fatalf("Unable to convert the string into int, %v", err)
	}

	// call the deleteVenture, convert the int to int64

	deleteRows := deleteVenture(int64(id))

	// format the message string

	msg := fmt.Sprintf("Venture updated successfully, Total rows/record affected %v", deleteRows)

	// format the response  message

	res := response{

		ID: int64(id),

		Message: msg,
	}

	// Send the response

	json.NewEncoder(w).Encode(res)

}

// Insert one venture in DB
func insertVenture(venture models.Venture) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := "INSERT INTO ventures (name, domain, revenue_estimation) VALUES ($1,$2,$3) RETURNING  ventureid"

	var id int64

	err := db.QueryRow(sqlStatement, venture.Name, venture.Domain, venture.Revenue).Scan(&id)

	if err != nil {

		log.Fatalf("Unable to execute the query, %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id

}

// Get one venture from the DB by its ventureID

func getVenture(id int64) (models.Venture, error) {

	// Create the Postgres DB connection

	db := createConnection()

	// Close the db connecvtion

	defer db.Close()

	// Create a venture of models.Venture type

	var venture models.Venture

	// Create the SELECT SQL query

	sqlStatement := `SELECT * FROM ventures WHERE ventureid=$1`

	row := db.QueryRow(sqlStatement, id)

	// Create the SQL statement

	err := row.Scan(&venture.VentureID, &venture.Name, &venture.Domain, &venture.Revenue)

	switch err {

	case sql.ErrNoRows:

		fmt.Println("No rows are returned")

		return venture, nil

	case nil:

		return venture, nil

	default:

		log.Fatalf("Unable to scan the row. %v", err)

	}

	return venture, err

}

func getAllVentures() ([]models.Venture, error) {

	db := createConnection()

	defer db.Close()

	var ventures []models.Venture

	// CREATE the SQL Query

	sqlStatement := `SELECT * FROM ventures`

	// Execute SQL query

	rows, err := db.Query(sqlStatement)

	if err != nil {

		log.Fatalf("Unable to execute the Query. %v", err)

	}

	defer rows.Close()

	// Iterate over the rows

	for rows.Next() {

		var venture models.Venture

		// Unmarshal the rows object to venture
		err = rows.Scan(&venture.VentureID, &venture.Name, &venture.Domain, &venture.Revenue)

		if err != nil {

			log.Fatalf("Unable to scan the row, %v", err)
		}

		// append the venture in Ventures slioce

		ventures = append(ventures, venture)

	}

	return ventures, err

}

func updateVenture(id int64, venture models.Venture) int64 {

	// Create the Postgres DB connection

	db := createConnection()

	defer db.Close()

	// Create the update query

	sqlStatement := `UPDATE ventures SET name=$2, domain=$3, revenue_estimation=$4 WHERE ventureid=$1`

	res, err := db.Exec(sqlStatement, id, venture.Name, venture.Domain, venture.Revenue)

	if err != nil {

		log.Fatalf("Unable to execute the query. %v", err)

	}

	// Check how many rows are affected

	rowsaffected, err := res.RowsAffected()

	if err != nil {

		log.Fatalf("Error while checking the affected rows. %v", err)

	}

	fmt.Printf("Total rows are affected, %v", rowsaffected)

	return rowsaffected
}


// delete venture from DB

func deleteVenture(id int64) int64 {

	// Create the postgress db connection

	db := createConnection()

	// close the connection

	defer db.Close()

	// Create the Delete sql query

	sqlStatement := `Delete FROM ventures WHERE ventureid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {

		log.Fatalf("Unable to execute the query. %v", err)

	}

	// Check rows are affected

	rowsAffected, err := res.RowsAffected()

	if err != nil {

		log.Fatalf("Error while checking affected rows, %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected

}
