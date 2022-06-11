package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	// "time"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/nouveau05/nouveau-microservices-go/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
// func printFiles() {
// 	files, err := ioutil.ReadDir(".")
//     if err != nil {
//         log.Fatal(err)
//     }

//     for _, file := range files {
//         log.Println(file.Name(), file.IsDir())
//     }
// }
// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
// configureConnectionPool sets database connection pool properties.
// For more information, see https://golang.org/pkg/database/sql
// func configureConnectionPool(dbPool *sql.DB) {
// 	// [START cloud_sql_mysql_databasesql_limit]

// 	// Set maximum number of connections in idle connection pool.
// 	dbPool.SetMaxIdleConns(5)

// 	// Set maximum number of open connections to the database.
// 	dbPool.SetMaxOpenConns(7)

// 	// [END cloud_sql_mysql_databasesql_limit]

// 	// [START cloud_sql_mysql_databasesql_lifetime]

// 	// Set Maximum time (in seconds) that a connection can remain open.
// 	dbPool.SetConnMaxLifetime(1800 * time.Second)

// 	// [END cloud_sql_mysql_databasesql_lifetime]
// }
// func initTCPConnectionPool() (*sql.DB, error) {
// 	// [START cloud_sql_mysql_databasesql_create_tcp]
// 	var (
// 		dbUser    = mustGetenv("DB_USER") // e.g. 'my-db-user'
// 		dbPwd     = mustGetenv("DB_PASS") // e.g. 'my-db-password'
// 		dbTCPHost = mustGetenv("DB_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
// 		dbPort    = mustGetenv("DB_PORT") // e.g. '3306'
// 		dbName    = mustGetenv("DB_NAME") // e.g. 'my-database'
// 	)

// 	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)

// 	// [START_EXCLUDE]
// 	// [START cloud_sql_mysql_databasesql_sslcerts]
// 	// (OPTIONAL) Configure SSL certificates
// 	// For deployments that connect directly to a Cloud SQL instance without
// 	// using the Cloud SQL Proxy, configuring SSL certificates will ensure the
// 	// connection is encrypted. This step is entirely OPTIONAL.
// 	// dbRootCert := os.Getenv("DB_ROOT_CERT") // e.g., '/path/to/my/server-ca.pem'
// 	// if dbRootCert != "" {
// 	// 	var (
// 	// 		dbCert = mustGetenv("DB_CERT") // e.g. '/path/to/my/client-cert.pem'
// 	// 		dbKey  = mustGetenv("DB_KEY")  // e.g. '/path/to/my/client-key.pem'
// 	// 	)
// 	// 	pool := x509.NewCertPool()
// 	// 	pem, err := ioutil.ReadFile(dbRootCert)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	if ok := pool.AppendCertsFromPEM(pem); !ok {
// 	// 		return nil, err
// 	// 	}
// 	// 	cert, err := tls.LoadX509KeyPair(dbCert, dbKey)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	mysql.RegisterTLSConfig("cloudsql", &tls.Config{
// 	// 		RootCAs:               pool,
// 	// 		Certificates:          []tls.Certificate{cert},
// 	// 		InsecureSkipVerify:    true,
// 	// 		VerifyPeerCertificate: verifyPeerCertFunc(pool),
// 	// 	})
// 	// 	dbURI += "&tls=cloudsql"
// 	// }
// 	// [END cloud_sql_mysql_databasesql_sslcerts]
// 	// [END_EXCLUDE]

// 	// dbPool is the pool of database connections.
// 	dbPool, err := sql.Open("postgres", dbURI)
// 	if err != nil {
// 		return nil, fmt.Errorf("sql.Open: %v", err)
// 	}

// 	// [START_EXCLUDE]
// 	configureConnectionPool(dbPool)
// 	// [END_EXCLUDE]

// 	return dbPool, nil
// 	// [END cloud_sql_mysql_databasesql_create_tcp]
// }
// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of SQL Server.
func initSocketConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_mysql_databasesql_create_socket]
	var (
		dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
		dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
		dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	)
	log.Println("initSocketConnectionPool")

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	// if !isSet {
	// 	socketDir = "/cloudsql"
	// }

	// dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s??sslmode=disable", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	// log.Printf("dbURI: %v", dbURI)
	// // dbPool is the pool of database connections.
	// dbPool, err := sql.Open("postgres", dbURI)
	// if err != nil {
	// 	return nil, fmt.Errorf("sql.Open: %v", err)
	// }

	// // [START_EXCLUDE]
	// configureConnectionPool(dbPool)
	// // [END_EXCLUDE]

	// socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
			socketDir = "/cloudsql"
	}

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	// dbPool is the pool of database connections.
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
        return nil, fmt.Errorf("sql.Open: %v", err)
	}


	return db, nil
	// [END cloud_sql_mysql_databasesql_create_socket]
}

func createConnection() *sql.DB {
	var db *sql.DB

	// printFiles()
	//load env file
	err := godotenv.Load("env_file", "/app/env_file")
	if err != nil {
		log.Printf("Error loading env_file file. %v", err)
	}

	// //db, err = Open("mysql", "gorm:gorm@/gorm?charset=utf8&parseTime=True")
	if os.Getenv("POSTGRES_URL") != "" {
		db, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	} else {
		db, err = initSocketConnectionPool()
		if err != nil {
			log.Fatalf("initSocketConnectionPool: unable to connect: %v", err)
		}
	}
	log.Println("initSocketConnectionPool connected ")

	// Check the connection
	// err = db.Ping()
	// if err != nil {

	// 	panic(err)

	// }

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
