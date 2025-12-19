package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/digital-dream-labs/vector-cloud/internal/robot"
)

type ESNBlacklist struct {
	Blacklist []string `json:"blacklist"`
}

func CheckBlacklist() {
	fmt.Println("Running Amy's blacklist check!")
	var blacklist ESNBlacklist

	blacklistjson, err := http.Get("https://amymc.dev/blacklist")
	if err != nil {
		time.Sleep(30 * time.Second)
		blacklistjson, err = http.Get("https://amymc.dev/blacklist")
		if err != nil {
			if _, err := os.Stat("/data/vic-gateway/isBlacklisted"); err == nil {
				panic("ESN is blacklisted")
			}
			return
		}
	}
	defer blacklistjson.Body.Close()

	body, _ := io.ReadAll(blacklistjson.Body)
	err = json.Unmarshal(body, &blacklist)
	if err != nil {
		return
	}

	blacklistMap := make(map[string]bool)
	for _, esn := range blacklist.Blacklist {
		blacklistMap[esn] = true
	}

	esn, _ := robot.ReadESN()
	if blacklistMap[esn] {
		if _, err := os.Stat("/data/vic-gateway/isBlacklisted"); err != nil {
			f, _ := os.Create("/data/vic-gateway/isBlacklisted")
			f.Close()
		}
		panic("ESN is blacklisted")
	} else {
		if _, err := os.Stat("/data/vic-gateway/isBlacklisted"); err == nil {
			err = os.Remove("/data/vic-gateway/isBlacklisted")
		}
	}
}
