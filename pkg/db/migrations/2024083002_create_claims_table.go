package migrations

import (
	"github.com/deveshmishra34/groot/pkg/db/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	m := &gormigrate.Migration{}

	m.ID = "2024083001_create_claims_table"

	m.Migrate = func(db *gorm.DB) error {
		type Claim struct {
			models.Claim
		}

		return AutoMigrateAndLog(db, &Claim{}, m.ID)
	}

	m.Rollback = func(db *gorm.DB) error {
		if err := db.Migrator().DropTable("claims"); err != nil {
			logFail(m.ID, err, true)
		}
		logSuccess(m.ID, true)
		return nil
	}

	AddMigration(m)
}
