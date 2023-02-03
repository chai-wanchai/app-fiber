package migration

import (
	"fmt"
	"iot/config"
	driver "iot/dependency"
	"iot/internal/model"
	"math/rand"
	"time"
)

func InsertData(cfg config.Config) error {
	dep := driver.NewConnection(cfg)
	sql := dep.GetConnection().SqlORM
	m := sql.MasterDB.Migrator()
	_ = m.AutoMigrate(&model.Power{})
	var count int64 = 0
	rowsCheck := sql.MasterDB.Model(&model.Power{}).Count(&count)
	if rowsCheck.Error != nil {
		return rowsCheck.Error
	}
	if count > 1000 {
		fmt.Println("Insert complete")
		return nil
	}
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		data := model.Power{
			ActivePower: rand.Intn(1000) + 1,
			PowerInput:  rand.Intn(1000) + 1,
		}
		tx := sql.MasterDB.Model(&model.Power{}).Create(&data)
		if tx.Error != nil {
			fmt.Errorf("Err insert : %s", tx.Error.Error())
			return tx.Error
		}
	}
	return nil
}
