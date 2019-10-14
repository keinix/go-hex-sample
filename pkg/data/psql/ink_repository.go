package psql

import (
	"fmt"
	"go-hex-sample/pkg/ink"
)

type inkRepository struct {
}

func NewInkRepository() ink.Repository {
	return &inkRepository{}
}

func (r *inkRepository) AddInk(ink ink.Ink) error {
	db, err := openDb()
	if err != nil {
		return fmt.Errorf("error connecting to db: %v", err)
	}
	db.Save(&ink)
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func (r *inkRepository) GetInk(id int64) (*ink.Ink, error) {
	db, err := openDb()
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}
	var result ink.Ink
	db.First(&result, id)
	if result.Id == 0 {
		return nil, fmt.Errorf("could not find ink with id: %d", id)
	}
	if err := db.Close(); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *inkRepository) GetAllInks() (*[]ink.Ink, error) {
	db, err := openDb()
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}
	var result []ink.Ink
	db.Find(&result)
	if err := db.Close(); err != nil {
		return nil, err
	}
	return &result, nil
}
