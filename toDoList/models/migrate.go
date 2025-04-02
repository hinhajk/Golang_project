package models

// 简历表格，将其迁移到数据库中
func Migration() {
	//err := DB.Migrator().AutoMigrate(Tasks{}, User{})
	////建立一张teacher表
	//if err != nil {
	//	return
	//}
	DB.Set("gorm:table_options", "charset=utf8mb4").
		Migrator().AutoMigrate(&User{}, &Tasks{})

	//建造多个表时在后面加即可
	//例如：DB.Migrator().AutoMigrate(Teacher{}, Course{})
}
