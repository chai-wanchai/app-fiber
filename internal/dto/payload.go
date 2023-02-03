package dto

type QueryPayload struct {
	ActivePower int `query:"active_power" json:"active_power"`
	PowerInput  int `query:"power_input" json:"power_input"`
}
type SumData struct {
	ActivePower int `json:"active_power" gorm:"column:active_power"`
	PowerInput  int `json:"power_input" gorm:"column:power_input"`
}
