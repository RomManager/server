package roms

type Emulator struct {
	Name       string
	FolderName string
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
