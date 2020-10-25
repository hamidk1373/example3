package migrate

import (
	"hamid/example3/settings"
	"hamid/example3/users"
)

// MakeMigrations migrates all models to database tables.
func MakeMigrations() error {
	db, err := settings.GetDB()
	if err != nil {
		return err
	}
	return db.AutoMigrate(&users.User{})
}
