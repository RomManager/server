package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Rom struct {
	ID            uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name          string    `json:"name"`
	Filepath      string    `gorm:"unique" json:"filepath"`
	FileURL       string    `json:"file_url"`
	Emulator      string    `json:"emulator"`       // Given emulator as emulator ID (Foldername)
	SteamGridDBID int       `json:"steamgriddb_id"` // If 0 no steamgrid connection
	ReleaseDate   time.Time `json:"release_date"`   // When not known its a 0 timestamp
	GridFilepath  string    `json:"grid_filepath"`
	GridURL       string    `json:"grid_url"`
}

func (r *Rom) SaveRom() (*Rom, error) {

	strID := strconv.FormatUint(uint64(r.ID), 10)
	r.FileURL = "/rom/" + strID + "/file"
	r.GridURL = "/rom/" + strID + "/grid"

	fmt.Println(r.ID)
	fmt.Println(strID)
	fmt.Println(r)

	err := DB.Create(&r).Error
	if err != nil {
		return &Rom{}, err
	}
	return r, nil
}

func GetAllRoms() (*[]Rom, error) {
	var roms = []Rom{}

	err := DB.Find(&roms).Error

	if err != nil {
		return &[]Rom{}, err
	}

	return &roms, err
}

func GetRomByID(rid uint) (Rom, error) {
	var r Rom

	if err := DB.First(&r, rid).Error; err != nil {
		return r, errors.New("rom not found")
	}

	return r, nil
}

// Check if rom exists by given filepath -> bool
func CheckRomExistsByFilepath(path string) bool {
	var roms []Rom

	result := DB.Where(&Rom{Filepath: path}, "filepath").Find(&roms)

	return result.RowsAffected != 0
}
