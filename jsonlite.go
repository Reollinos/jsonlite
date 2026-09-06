package jsonlite

import (
	"encoding/json"
	"fmt"
	"os"
)

type header struct {
	JsonFileContent []byte
	JsonPath string
	Struct any
}


type internPreferences struct {
	Indent string
	Prefix string
	Perm os.FileMode
}

var Preferences = internPreferences{
	Indent: "    ",
	Prefix: "",
	Perm: 0644,
}

func (hd *header) Structure(structAddress any) error {
	hd.Struct = structAddress
	return json.Unmarshal(hd.JsonFileContent, structAddress)
}

func (hd *header) ReloadJson() error {
	data, err := json.MarshalIndent(hd.Struct, Preferences.Prefix, Preferences.Indent)
	if err != nil {
		return fmt.Errorf(
			"Unable to change the content of `%s`",
			hd.JsonPath,
		)
	}

	if err := os.WriteFile(hd.JsonPath, data, Preferences.Perm); err != nil {
		return nil
	}

	hd.JsonFileContent = data
	return nil
}

func Load(path string) (*header, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"Unable to access path `%s`. Check if the .json file actually exists! This may solve your problem.",
			path,
		)
	}

	return &header{
		JsonFileContent: data,
		JsonPath: path,
	}, nil
}