package roms

type Emulator struct {
	Name       string
	FolderName string // Has to be unique
}

var EmulatorList []Emulator = []Emulator{
	{
		Name:       "Nintendo DS",
		FolderName: "nds",
	},
	{
		Name:       "Nintendo Wii",
		FolderName: "wii",
	},
}
