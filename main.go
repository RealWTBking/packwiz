package main

import (
	// Modules of packwiz
	"github.com/RealWTBking/packwiz/cmd"
	_ "github.com/RealWTBking/packwiz/curseforge"
	_ "github.com/RealWTBking/packwiz/github"
	_ "github.com/RealWTBking/packwiz/migrate"
	_ "github.com/RealWTBking/packwiz/modrinth"
	_ "github.com/RealWTBking/packwiz/settings"
	_ "github.com/RealWTBking/packwiz/url"
	_ "github.com/RealWTBking/packwiz/utils"
)

func main() {
	cmd.Execute()
}
