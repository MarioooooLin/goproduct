package main

import (
	"goproduct/config"
	"goproduct/routers"
	"log"
)

func main() {
	// 初始化資料庫（MySQL + GORM）
	db := config.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 啟動 Router（使用 Gin + Channel 模型支援高併發）
	r := routers.SetupRouter(db)
	log.Println("🚀 Server is running at http://localhost:8080")
	r.Run(":8080")
}
