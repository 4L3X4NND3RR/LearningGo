package helper

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	Packages []Package
	Data     string `gorm:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	ShipmentID uint
	Data string `gorm:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// GORM creates tables with plural names.
// Use this to suppress it
func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func InitDB() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=jdeleon password=jdeleon123 dbname=mydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Shipment{}, &Package{})
	return db, nil
}
