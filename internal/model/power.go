package model

type Power struct {
	ID          int `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"`
	ActivePower int `gorm:"column:active_power;type:int;" json:"active_power"`
	PowerInput  int `gorm:"column:power_input;type:int;" json:"power_input"`
}

func (Power) TableName() string {
	return "power"
}
