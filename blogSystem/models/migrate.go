package models

func Migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		Migrator().AutoMigrate(&User{}, &Blog{}, &BlogComment{})
}
