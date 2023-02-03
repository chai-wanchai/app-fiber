package driver

import (
	"iot/config"

	"gorm.io/gorm"
)

type Connections struct {
	SqlORM SqlORM
}
type SqlORM struct {
	MasterDB *gorm.DB
}

type connection struct {
	db     Connections
	config config.Config
}
type DriverConnection interface {
	GetConnection() Connections
}

func NewConnection(config config.Config) DriverConnection {
	conn := createConnections(config)
	return connection{
		db:     conn,
		config: config,
	}
}

func (c connection) GetConnection() Connections {
	return Connection
}

var Connection Connections

func createConnections(cfg config.Config) Connections {

	sql, err := NewRegistrySQLWithORM(SQLConfig{
		Username: cfg.SQLMasterDB.Username,
		Password: cfg.SQLMasterDB.Password,
		Host:     cfg.SQLMasterDB.Host,
		Port:     cfg.SQLMasterDB.Port,
		DB:       cfg.SQLMasterDB.Db,
	})
	if err != nil {
		panic("Err connect DB: " + err.Error())
	}
	var connections Connections = Connections{
		SqlORM: SqlORM{
			MasterDB: sql,
		},
	}
	Connection = connections
	return connections
}
func GetConnection() *Connections {
	return &Connection
}
