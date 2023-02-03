package service

import (
	"context"
	"fmt"
	driver "iot/dependency"
	"iot/internal/dto"
	"iot/internal/model"
	"math/rand"
	"time"
)

type IoTService interface {
	GetData(ctx context.Context, query dto.QueryPayload) ([]model.Power, error)
	GetSumData(ctx context.Context, query dto.QueryPayload) (dto.SumData, error)
	RandomData(ctx context.Context) error
}
type iotService struct {
	db driver.SqlORM
}

func NewIoTService(sql driver.SqlORM) IoTService {
	return iotService{
		db: sql,
	}
}

func (svc iotService) GetData(ctx context.Context, query dto.QueryPayload) ([]model.Power, error) {
	data := make([]model.Power, 0)
	tx := svc.db.MasterDB.Model(&model.Power{})
	tx.Find(&data)
	if tx.Error != nil {
		return data, tx.Error
	}
	return data, nil
}
func (svc iotService) GetSumData(ctx context.Context, query dto.QueryPayload) (dto.SumData, error) {
	data := dto.SumData{}
	modelInfo := model.Power{}
	tx := svc.db.MasterDB.Debug().Table(modelInfo.TableName())
	tx.Select("sum(power.active_power) as active_power,sum(power.power_input) as power_input")
	tx.Find(&data)
	if tx.Error != nil {
		return data, tx.Error
	}
	return data, nil
}
func (svc iotService) RandomData(ctx context.Context) error {
	modelInfo := model.Power{}
	tx := svc.db.MasterDB.Debug().Table(modelInfo.TableName())
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		data := model.Power{
			ActivePower: rand.Intn(1000) + 1,
			PowerInput:  rand.Intn(1000) + 1,
		}
		tx := tx.Where("id=?", i+1).Updates(&data)
		if tx.Error != nil {
			fmt.Errorf("Err insert : %s", tx.Error.Error())
			return tx.Error
		}
	}
	return nil

}
