package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"
)

type WhoAmI struct {
	Hostname string    `json:"hostname"`
	Time     time.Time `json:"time"`
	Timezone string    `json:"timezone"`
	MOTD     string    `json:"motd,omitempty"`
	User     user.User `json:"user"`
}

func NewWhoAmi() WhoAmI {

	hostName, _ := os.Hostname()
	currentTime := time.Now()
	timezone, offset := currentTime.Zone()
	motd := os.Getenv("MOTD")
	current, _ := user.Current()

	return WhoAmI{
		Hostname: hostName,
		Time:     currentTime,
		Timezone: fmt.Sprintf("%s %d", timezone, offset),
		MOTD:     motd,
		User:     *current,
	}
}

func main() {

	// Return a new WhoAmI payload with server and OS information
	http.HandleFunc("/whoami", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling /whoami")
		payload := NewWhoAmi()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(payload)
		if err != nil {
			return
		}
	})

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
