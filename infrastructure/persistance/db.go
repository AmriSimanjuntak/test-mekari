package persistance

import (
	"Test_Mek/domain/repository"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Repositories struct {
	Db           *gorm.DB
	EmployeeRepo repository.IEmployee
}

func NewRepositories() (*Repositories, error) {
	db, err := initPostgres("",
		"localhost", "postgres",
		"kmb123", "postgres",
		"5432", "disable")
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Db:           db,
		EmployeeRepo: NewEmployeeRepo(db),
	}, nil
}
func initPostgres(env, host, user, password, name, port, sslMode string, cfg ...gorm.Config) (*gorm.DB, error) {
	// init connection postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		host, user, password, name, port, sslMode)
	gormCfg := new(gorm.Config)
	if len(cfg) > 0 {
		gormCfg = &cfg[0]
	}
	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}
	db.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	return db, nil
}
