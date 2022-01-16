package main

import (
	"encoding/gob"
	"fmt"
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

/**
gob 是仅golang支持的，其他语言不支持
*/

// serialize 将内存中的数据序列化到磁盘上
// 不仅支持go中内置类型，也支持开发者自定义类型
func serialize() {
	// userMap := map[int]User{
	// 	1: User{Id: 1, Name: "keke", Birthday: time.Now(), Tel: "18487299603", Addr: "天骄北麓"},
	// 	2: User{Id: 2, Name: "keke1", Birthday: time.Now(), Tel: "18487299604", Addr: "天骄北麓1"},
	// 	3: User{Id: 3, Name: "keke2", Birthday: time.Now(), Tel: "18487299605", Addr: "天骄北麓2"},
	// 	4: User{Id: 4, Name: "keke3", Birthday: time.Now(), Tel: "18487299604", Addr: "天骄北麓3"},
	// }
	user := User{
		Id:       1,
		Name:     "murray",
		Birthday: time.Now(),
		Tel:      "124235364747",
		Addr:     "kunmings",
	}
	f, err := os.Create("user.gob")
	if err != nil {
		log.Fatalf("Failed to open file user.gob, error: %s\n", err)
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	encoder.Encode(&user)
}

// deserialize 将磁盘上的文件反序列化到内存中
func deserialize() {
	// users := map[int]User{}
	user := User{}

	f, err := os.Open("user.gob")
	if err != nil {
		log.Fatalf("Failed to open file user.gob, error: %s\n", err)
	}
	defer f.Close()
	decoder := gob.NewDecoder(f)
	decoder.Decode(&user)

	fmt.Println(user)
}

func main() {
	serialize()
	deserialize()
}
