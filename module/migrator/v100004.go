package migrator

import (
	"github.com/opentdp/go-helper/dborm"
)

func v100004() error {

	if isMigrated("v100004") {
		return nil
	}

	if err := v100004AddVendor(); err != nil {
		return err
	}

	return addMigration("v100004", "添加默认 Vendor")

}

func v100004AddVendor() error {

	// 添加默认 Vendor
	sql := "INSERT INTO `vendor` VALUES (0, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL)"

	return dborm.Db.Exec(sql).Error

}
