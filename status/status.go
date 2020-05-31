package status

import (
	"fmt"
	"strconv"

	"github.com/pelletier/go-toml"
)

// Status describes a HTTP status
type Status struct {
	Code        int    `toml:"code" json:"code"`
	Phrase      string `toml:"phrase" json:"phrase"`
	Class       string `toml:"class" json:"class"`
	Description string `toml:"description" json:"description"`
	Reference   string `toml:"reference" json:"reference"`
}

var allStatus map[int]Status

// Init initializes status data using the file at the given path.
func Init(path string) error {
	loaded, err := toml.LoadFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	allStatus = make(map[int]Status)
	strKeyData := make(map[string]Status)
	loaded.Unmarshal(&strKeyData)

	for k, v := range strKeyData {
		code, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		allStatus[code] = v
	}

	return nil
}

// Get returns the status with the given code.
func Get(code int) (Status, error) {
	status, exists := allStatus[code]
	if !exists {
		return status, fmt.Errorf("no status with code %d", code)
	}

	return status, nil
}

// All returns the list of all status.
func All() ([]Status, error) {
	return Filter(func(s Status) bool {
		return true
	})
}

// Filter returns a list of status matching the given filter.
func Filter(filter func(Status) bool) ([]Status, error) {
	statusList := make([]Status, 0, len(allStatus))
	for _, s := range allStatus {
		if filter(s) {
			statusList = append(statusList, s)
		}
	}

	return statusList, nil
}
