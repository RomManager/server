package models

type Rom struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:"name"`
	Filepath string `json:"filepath"`
	Emulator string `json:"emulator"`
}

func (r *Rom) SaveRom() (*Rom, error) {
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
