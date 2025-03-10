package auto_migrate

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	err := db.AutoMigrate(&ExampleEntity{})
	if err != nil {
		return err
	}
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	return nil
}
