package main

import (
	"encoding/json"
	"firstGo/modules/basic"
	"firstGo/modules/channel"
	"firstGo/modules/listen"
	"firstGo/modules/model"
	"firstGo/modules/speaker"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"time"
)

func main() {
	is_run := false
	if is_run {
		run()
		// basic
		basic.Run()
		// channel
		channel.Run()
	}
	// listen
	is_listen := false
	if is_listen {
		listen.Run()
	}
	// time
	runTime()
}

func runTime() {
	// time
	now := time.Now()
	// 1月2日 3時4分5秒 2006年 +007
	log.Println(now.Format("2006-01-02 15:04:05"))
	log.Println(now.Format(time.DateOnly))
	// duration
	d, err := time.ParseDuration("5s")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(d)
	now2 := now.Add(d)
	dd := now2.Sub(now)
	log.Println(now)
	log.Println(now2)
	log.Println(dd)

	// filepath
	p := "/Users/takashifuruya/Documents/work_go/first_go/main.go"
	println(filepath.Base(p))
	println(filepath.Dir(p))
	pd := "/Users/takashifuruya/..Documents/../first_go/main.go"
	if filepath.IsAbs(pd) == false {
		println(filepath.Clean(pd))
	}
	absolute, err := filepath.Rel("/test", "/test/ok/hello.txt")
	if err == nil {
		println(absolute)
	}
	files := []string{}
	filepath.WalkDir(
		".",
		func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			files = append(files, path)
			return nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}

func run() {
	// struct
	var user1 model.User
	user1.Name = "たろう"
	user1.Age = 22
	fmt.Println(user1)
	user1.UpdateName("じろう")
	fmt.Println(user1)
	// constructer
	user2 := *model.NewUser("David", 35)
	fmt.Println(user2)

	// interface
	dog := speaker.Dog{}
	speaker.DoSpeak(&dog) // &; pointer of a value
	cat := speaker.Cat{}
	speaker.DoSpeak(&cat)

	// json
	var content2 = `
	{
		"species": "ハト",
		"description": "岩にとまるのが好き",
		"dimensions": {
			"height": 24,
			"width": 10
		}
	}
	`
	var data model.Data
	err := json.Unmarshal([]byte(content2), &data)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	fmt.Println(data)
}
