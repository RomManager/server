package models

type Rom struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:"name"`
	Filepath string `gorm:"unique" json:"filepath"`
	Emulator string `json:"emulator"` // Given emulator as emulator ID (Foldername)
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

// Check if rom exists by given filepath -> bool
func CheckRomExistsByFilepath(path string) bool {
	var roms []Rom

	result := DB.Where(&Rom{Filepath: path}, "filepath").Find(&roms)

	return result.RowsAffected != 0
}
