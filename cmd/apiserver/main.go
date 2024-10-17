package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gaara1994/gorm_learm/internal/config"
	"github.com/gaara1994/gorm_learm/router"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("%#v", cfg)

	//初始化数据库
	db, err := gorm.Open(postgres.Open(cfg.Database.ConnectionURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println(db.Statement.Vars...)
	// 自动迁移模式（可选）
	// db.AutoMigrate()

	//gin启动http服务
	// 创建 Gin 路由器
	r := router.SetupRouter()
	// 启动 HTTP 服务器
	log.Printf("Starting server on :%s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
