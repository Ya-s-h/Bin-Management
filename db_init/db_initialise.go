package db_init

import (
	model "assignment/renie/models"
	"log"

	"gorm.io/gorm"
)

// db.Model(&models.UserInfo{}).AddForeignKey("u_id", "t_user(id)", "RESTRICT", "RESTRICT")

func getModels() []interface{} {
	return []interface{}{&model.UserRole{}, &model.User{}, &model.Area{}, &model.Bin{}, &model.Waste{}}

}

func DeleteTables(db *gorm.DB) {
	error := db.Migrator().DropTable(getModels()...)
	if error != nil {
		log.Println("Failed to delete the tables: %v", error)
	}
}

func CreateTables(db *gorm.DB) {
	error := db.Migrator().CreateTable(getModels()...)

	if error != nil {
		log.Println("Failed to create the tables: %v", error)
	}

}
func seedUserRoles(db *gorm.DB) {
	roles := []model.UserRole{
		{Name: "Admin"},
		{Name: "Area Provider"},
		{Name: "Owners/Investors"},
		{Name: "Both"},
	}

	if err := db.Create(&roles).Error; err != nil {
		log.Fatalf("Failed to seed user roles: %v", err)
	}

	log.Println("User roles seeded successfully.")
}

func seedUsers(db *gorm.DB) {
	users := []model.User{
		{Name: "Admin", RoleID: 1, Email: "admin@admin.com", Password: "admin1234"},
		{Name: "Alice", RoleID: 2, Email: "alice@provider.com", Password: "alice123"},
		{Name: "Bob", RoleID: 3, Email: "bob@owner.com", Password: "bob123"},
		{Name: "Charlie", RoleID: 4, Email: "charlie@both.com", Password: "charlie123"},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}

	log.Println("Users seeded successfully.")
}

func seedAreas(db *gorm.DB) {
	areas := []model.Area{
		{Name: "Downtown", Location: "City Center", UserID: 2},
		{Name: "Uptown", Location: "North City", UserID: 3},
	}

	if err := db.Create(&areas).Error; err != nil {
		log.Fatalf("Failed to seed areas: %v", err)
	}

	log.Println("Areas seeded successfully.")
}

func seedBins(db *gorm.DB) {
	bins := []model.Bin{
		{AreaID: 1, UserID: 2},
		{AreaID: 2, UserID: 3},
	}

	if err := db.Create(&bins).Error; err != nil {
		log.Fatalf("Failed to seed bins: %v", err)
	}

	log.Println("Bins seeded successfully.")
}

func seedWastes(db *gorm.DB) {
	wastes := []model.Waste{
		{BinID: 1, Weight: 21},
		{BinID: 1, Weight: 30},
		{BinID: 2, Weight: 59},
	}

	if err := db.Create(&wastes).Error; err != nil {
		log.Fatalf("Failed to seed wastes: %v", err)
	}

	log.Println("Wastes seeded successfully.")
}

func SeedTables(db *gorm.DB) {
	seedUserRoles(db)
	seedUsers(db)
	seedAreas(db)
	seedBins(db)
	seedWastes(db)
}
func CreateProcedures(db *gorm.DB) {
	procedures := []string{
		`
        CREATE OR REPLACE FUNCTION calculate_waste_earnings()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.earnings := NEW.weight * 10;
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
        `,
		`
        CREATE TRIGGER set_waste_earnings
        BEFORE INSERT OR UPDATE ON wastes
        FOR EACH ROW
        EXECUTE FUNCTION calculate_waste_earnings();
        `,
		`
        CREATE OR REPLACE FUNCTION update_bin_waste_collected()
        RETURNS TRIGGER AS $$
        BEGIN
            UPDATE bins
            SET waste_collected = (
                SELECT COALESCE(SUM(weight), 0)
                FROM wastes
                WHERE bin_id = NEW.bin_id
            )
            WHERE id = NEW.bin_id;

            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
        `,
		`
        CREATE TRIGGER update_bin_waste_collected_trigger
        AFTER INSERT OR UPDATE ON wastes
        FOR EACH ROW
        EXECUTE FUNCTION update_bin_waste_collected();
        `,
		`
        CREATE OR REPLACE FUNCTION update_user_earnings()
        RETURNS TRIGGER AS $$
        BEGIN
            UPDATE users
            SET earnings = (
                SELECT COALESCE(SUM(w.earnings), 0)
                FROM wastes w
                JOIN bins b ON w.bin_id = b.id
                WHERE b.user_id = users.id
            )
            WHERE id = (
                SELECT b.user_id
                FROM bins b
                WHERE b.id = NEW.bin_id
            );

            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
        `,
		`
        CREATE TRIGGER update_user_earnings_trigger
        AFTER INSERT OR UPDATE ON wastes
        FOR EACH ROW
        EXECUTE FUNCTION update_user_earnings();
        `,
	}

	// Execute each SQL statement
	for _, sql := range procedures {
		if err := db.Exec(sql).Error; err != nil {
			log.Fatalf("Failed to execute procedure or trigger: %v", err)
		}
	}

	log.Println("Procedures and triggers created successfully!")

}

// func InsertRolesData(db *gorm.DB) {
// 	roles_data := []model.UserRole{
// 		{Name: "Admin"},
// 		{Name: "Area Provider"},
// 		{Name: "Owners/Investors"},
// 		{Name: "Both"},
// 	}
// 	role_result := db.Create(&roles_data)
// 	if err := role_result.Error; err != nil {
// 		log.Println("Error while creating roles data: %v\n", err)
// 	}
// }
// func InsertUserData(db *gorm.DB) {
// 	// Use for authentication
// 	user_data := model.User{
// 		Name:     "Admin",
// 		RoleID:   1,
// 		Email:    "admin@admin.com",
// 		Password: "1234",
// 	}
// 	user_result := db.Create(&user_data)
// 	if err := user_result.Error; err != nil {
// 		log.Println("Error while creating user data: %v\n", err)
// 	}
// }

// func main() {
// 	db := connection.ConnectToDb()
// 	// Create Tables
// 	CreateTables(db)
// 	// Create Roles
// 	userRoles := []model.UserRole{
// 		{Name: "Area Provider"},
// 		{Name: "Owners/Investors"},
// 		{Name: "Both"},
// 	}
// 	result := db.Create(&userRoles)
// 	if err := result.Error; err != nil {
// 		log.Println("Error while creating: %v\n", err)
// 	}
// }
