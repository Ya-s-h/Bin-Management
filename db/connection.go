package main
import (
    "context"
    "log"
    "github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDb() *pgxpool.Pool{
	connectionURL := "postgres://admin:password@localhost:5432/smart_bin"
	conn, err := pgxpool.New(context.Background(), connectionURL)
	if err != nil {
		log.Fatalf("Unable to Connect to database: %v\n", err)
	}
	log.Println("Connected to the PostgreSQL database!")
    return conn
}


// Testing Postgres Connection
// func main(){
// 	db := ConnectToDb() // Establishes a connection to the database
//     defer db.Close()    // Ensures the pool is closed when the program exits

//     var result string
//     err := db.QueryRow(context.Background(), "SELECT 'Hello, PostgreSQL!'").Scan(&result)
//     if err != nil {
//         log.Fatalf("Query failed: %v\n", err)
//     }
// 	if err := db.Ping(context.Background()); err != nil {
// 		log.Fatalf("Ping to the database failed: %v\n", err)
// 	}

//     log.Println(result)
// }