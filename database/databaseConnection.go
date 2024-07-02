package database

import(
	"fmt";
	"log";
	"os";
	"database/sql"
	"github.com/joho/godotenv";
	_ "github.com/lib/pq";
)

func ConnectDB() *sql.DB{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		dbname = os.Getenv("DB_NAME")
	)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", connectionString)
	
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

var Client *sql.DB = ConnectDB()