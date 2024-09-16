package seeds

import (
	"github.com/deveshmishra34/groot/pkg/clients/logger"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Seeds *gormigrate.Gormigrate
var SeedsList = []*gormigrate.Migration{}

func Init(db *gorm.DB) {
	Seeds = gormigrate.New(db, gormigrate.DefaultOptions, SeedsList)
}

func AddSeed(seed *gormigrate.Migration) {
	SeedsList = append(SeedsList, seed)
}

func Apply() error {
	return Seeds.Migrate()
}

func logSuccess(id string) {
	logger.Info("Applied seed: %s", id)
}

func logFail(id string, err error) {
	logger.Error("Failed to apply seed: %s, error: %s", id, err)
}
