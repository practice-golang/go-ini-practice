package main // import "set-ini"

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ini/ini"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Path: " + wd)

	// giteaRootEnv := os.Getenv("GITEAROOT")
	// giteaRoot := ""

	// if len(giteaRootEnv) == 0 {
	// 	giteaRoot = wd
	// } else {
	// 	giteaRoot = giteaRootEnv
	// }

	// giteaRoot = strings.ReplaceAll(giteaRoot, "\\", "/")
	// userProfile := strings.ReplaceAll(os.Getenv("UserProfile"), "\\", "/")

	// cfg, err := ini.Load(giteaRoot + "/custom/conf/app.ini")
	// if err != nil {
	// 	fmt.Printf("Fail to read file: %v", err)
	// 	os.Exit(1)
	// }

	// // Print items to console
	// // fmt.Println("APP_NAME:", cfg.Section("").Key("APP_NAME").String())
	// // cfg.Section("").Key("APP_NAME").SetValue("Gitea")
	// // fmt.Println("APP_NAME:", cfg.Section("").Key("APP_NAME").String())
	// // fmt.Println("LFS_CONTENT_PATH:", cfg.Section("server").Key("LFS_CONTENT_PATH").String())
	// // fmt.Println("DB PATH:", cfg.Section("database").Key("PATH").String())
	// // fmt.Println("REPO ROOT:", cfg.Section("repository").Key("ROOT").String())
	// // fmt.Println("LOG PATH:", cfg.Section("log").Key("ROOT_PATH").String())

	// cfg.Section("server").Key("LFS_CONTENT_PATH").SetValue(giteaRoot + "/data/lfs")
	// cfg.Section("database").Key("PATH").SetValue(giteaRoot + "/data/gitea.db")
	// cfg.Section("repository").Key("ROOT").SetValue(userProfile + "/data/gitea.db")
	// cfg.Section("log").Key("ROOT_PATH").SetValue(giteaRoot + "/log")

	// cfg.SaveTo(giteaRoot + "/custom/conf/app.ini")

	klipperRoot := "../"
	cfg, err := ini.Load(klipperRoot + "/custom/conf/klipper-example.cfg")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	sections := cfg.Sections()
	for i, s := range sections {
		log.Println(i, s.Name())
		keys := s.Keys()
		for j, k := range keys {
			log.Println(j, k.Name(), k.String())
		}
	}
}
