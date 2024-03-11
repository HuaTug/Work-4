package model

func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).AutoMigrate(&Comment{}).AutoMigrate(&Favorite{}).AutoMigrate(&Video{}).
		AutoMigrate(&Relation{})
}
