package pg

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dsn := "host=localhost user=admin password=password dbname=smart_bin port=5432 sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return db

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
