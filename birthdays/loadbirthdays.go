package birthdays

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadBirthdays(path string) ([]Birthday, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Cannot find birthdays.json.")
	}
	defer file.Close()

	var birthdays []Birthday

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&birthdays)
	if err != nil {
		return nil, err
	}

	return birthdays, nil
}
