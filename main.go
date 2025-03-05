package main

import (
	_ "embed"
	"fmt"
	"objective-service/api/routes"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
)

//go:embed app-logo.txt
var appLogo []byte

func main() {

	fmt.Printf("%s\nMode: %s\nVersion: %s\nSha: %s\nBuildDate: %s\n", string(appLogo), core.Mode, core.BuildVersion, core.SHA, core.BuildDate)

	core.Invoke(routes.NewRouter)

	core.Start()
}
