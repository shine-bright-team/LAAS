package initialize

import (
	"fmt"

	"github.com/shine-bright-team/LAAS/v2/db"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbSetUp() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", utils.GetEnv("DB_HOST"), utils.GetEnv("DB_USER"), utils.GetEnv("DB_PASSWORD"), utils.GetEnv("DB_POSTGRES"), utils.GetEnv("DB_PORT"))
	db_ref, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.DB = db_ref
	return err
}
