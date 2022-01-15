package main

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)

type User struct {
	Id       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
}

func main() {
	userMap := map[int]User{
		1: User{Id: 1, Name: "keke", Birthday: time.Now(), Tel: "18487299603", Addr: "天骄北麓"},
		2: User{Id: 2, Name: "keke1", Birthday: time.Now(), Tel: "18487299604", Addr: "天骄北麓1"},
		3: User{Id: 3, Name: "keke2", Birthday: time.Now(), Tel: "18487299605", Addr: "天骄北麓2"},
		4: User{Id: 4, Name: "keke3", Birthday: time.Now(), Tel: "18487299604", Addr: "天骄北麓3"},
	}
	f, err := os.Create("user.gob")
	if err != nil {
		log.Fatalf("Failed to open file user.gob, error: %s\n", err)
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	encoder.Encode(&userMap)
}
