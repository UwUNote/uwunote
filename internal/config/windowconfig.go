package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/UwUNote/uwunote/internal/errors"
)

var (
	//Configuration contains the positions and sizes for all notes
	windowConfiguration = WindowConfig{}
)

func getWindowConfigPath() string {
	return filepath.Join(ConfigPath, "windows.json")
}

//WindowConfig contains a map of WindowData entries in a map
type WindowConfig struct {
	Data map[string]WindowData
}

//WindowData contains the sosition (x and y) and size of a Window
type WindowData struct {
	X int
	Y int

	Width  int
	Height int
}

//LoadWindowConfiguration loads the window configuration from its path.
func LoadWindowConfiguration() {
	log.Println("Loading window configuration")
	file, openError := os.Open(getWindowConfigPath())
	if openError == nil || os.IsExist(openError) {
		defer file.Close()
		decoder := json.NewDecoder(file)
		windowConfigLoadError := decoder.Decode(&windowConfiguration)

		if windowConfigLoadError != nil && windowConfigLoadError != io.EOF {
			errors.ShowErrorDialog(windowConfigLoadError)
		}
	}

	if windowConfiguration.Data == nil {
		//Creating an empty map to prevent nil pointer references
		windowConfiguration.Data = make(map[string]WindowData)
	}
}

//PersistWindowConfiguration saves the current window configuration to its iven path
func PersistWindowConfiguration() {
	windowConfigurationJSON, _ := json.Marshal(&windowConfiguration)
	writeError := ioutil.WriteFile(getWindowConfigPath(), windowConfigurationJSON, 0755)

	errors.ShowErrorDialogOnError(writeError)
}

//GetWindowDataForFile retrieves the window-config entry for the given file
func GetWindowDataForFile(file string) (WindowData, bool) {
	data, exists := windowConfiguration.Data[file]
	return data, exists
}

//DeleteWindowDataForFile removes an entry from the window data map.
func DeleteWindowDataForFile(file string) {
	delete(windowConfiguration.Data, file)
}

//SetWindowDataForFile sets coordinates and size in the window-config for the given file
func SetWindowDataForFile(file string, x, y, width, height int) {
	windowConfiguration.Data[file] = WindowData{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
