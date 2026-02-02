package database

import (
    "fmt"
    "task-management-system/internal/config"
    "task-management-system/internal/model"
    
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func NewMySQLDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset,
    )
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    
    if err := db.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
        return nil, fmt.Errorf("failed to migrate database: %w", err)
    }
    
    return db, nil
}