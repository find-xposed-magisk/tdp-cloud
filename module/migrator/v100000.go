package migrator

import (
	"tdp-cloud/module/dborm"
)

func v100000() error {

	return dborm.Db.AutoMigrate(
		&dborm.Config{},
		&dborm.Domain{},
		&dborm.Machine{},
		&dborm.Session{},
		&dborm.Sshkey{},
		&dborm.TaskHistory{},
		&dborm.TaskScript{},
		&dborm.User{},
		&dborm.Vendor{},
	)

}