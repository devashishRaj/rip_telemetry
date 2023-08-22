package postgresDB

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// struct to store json data
type SystemInfo struct {
	HardwareID  string  `json:"HardwareID"`
	CPUuserLoad float64 `json:"CPUuserLoad"`
	CPUidle     float64 `json:"CPUidle"`
	TotalMemory int64   `json:"TotalMemory"`
	FreeMemory  int64   `json:"FreeMemory"`
	IP          string  `json:"IP"`
	Temperature string  `json:"Temperature"`
	TimeStamp   string  `json:"TimeStamp"`
}

func ReceiveJSON(w http.ResponseWriter, r *http.Request) {
	var jsonData SystemInfo
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received Info: %+v\n", jsonData)
	w.WriteHeader(http.StatusOK)
	// UpdateDB in insert.go
	DBerr := InsertInDB(jsonData)
	if DBerr != nil {
		return
	}
}