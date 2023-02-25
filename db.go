package awfuldb

import (
	"encoding/json"
	"os"
)

type Database struct {
	DBPath string
}

func (db *Database) Set(key, value string) error {
	// get from file
	data, err := db.readData()
	if err != nil {
		return err
	}
	data[key] = value

	// write to file
	file, err := os.OpenFile(db.DBPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	new_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := file.Write(new_data); err != nil {
		return err
	}
	return nil
}

func (db *Database) Get(key string) (string, error) {
	data, err := db.readData()
	if err != nil {
		return "", err
	}
	return data[key], nil
}

func (db *Database) readData() (map[string]string, error) {
	data := map[string]string{}
	bytes, err := os.ReadFile(db.DBPath)
	if err != nil {
		return map[string]string{}, err
	}
	json.Unmarshal(bytes, &data)
	return data, nil
}