package persistance

import (
	"Test_Mek/domain/repository"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Repositories struct {
	Db           *gorm.DB
	EmployeeRepo repository.IEmployee
}

func NewRepositories() (*Repositories, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SCHEMA")
	fmt.Println("pantrek", password)

	db, err := initPostgres("",
		host, user,
		password, dbname,
		port, sslMode)
	if err != nil {
		return nil, err
	}
	fmt.Println("pantrek", password)

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
