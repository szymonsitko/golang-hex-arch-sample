package persistence

import (
	"log"

	"github.com/ssitko/hex-domain/config"
	album "github.com/ssitko/hex-domain/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewPersistenceLayer() DB {
	// Establish mysql connection
	dsn := config.GetConfigValue(config.DB_USER) + ":" + config.GetConfigValue(config.DB_PASSWORD) + "@tcp(" + config.GetConfigValue(config.DB_HOST) + ":" + config.GetConfigValue(config.DB_PORT) + ")/" + config.GetConfigValue(config.DB_NAME) + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	// Perform DB migrations
	mysqlDB.AutoMigrate(&album.Album{})

	// Return gorm wrapper that implements DB interface type
	return NewGormDBWrapper(mysqlDB)
}

type DB interface {
	Create(value interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Save(value interface{}) error
	Delete(value interface{}, conds ...interface{}) error
}

type GormDBWrapper struct {
	db *gorm.DB
}

func NewGormDBWrapper(db *gorm.DB) *GormDBWrapper {
	return &GormDBWrapper{db: db}
}

// Implement the interface methods
func (g *GormDBWrapper) Create(value interface{}) error {
	return g.db.Create(value).Error
}

func (g *GormDBWrapper) Find(dest interface{}, conds ...interface{}) error {
	return g.db.Find(dest, conds...).Error
}

func (g *GormDBWrapper) First(dest interface{}, conds ...interface{}) error {
	return g.db.First(dest, conds...).Error
}

func (g *GormDBWrapper) Save(value interface{}) error {
	return g.db.Save(value).Error
}

func (g *GormDBWrapper) Delete(value interface{}, conds ...interface{}) error {
	return g.db.Delete(value, conds...).Error
}
