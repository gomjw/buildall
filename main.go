package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gomjw/color"
)

var platforms = []string{
	"darwin/386",
	"darwin/amd64",
	"dragonfly/amd64",
	"freebsd/386",
	"freebsd/amd64",
	"freebsd/arm",
	"linux/386",
	"linux/amd64",
	"linux/arm",
	"linux/arm64",
	"linux/ppc64",
	"linux/ppc64le",
	"linux/mips",
	"linux/mipsle",
	"linux/mips64",
	"linux/mips64le",
	"linux/s390x",
	"nacl/386",
	"nacl/amd64p32",
	"nacl/arm",
	"netbsd/386",
	"netbsd/amd64",
	"netbsd/arm",
	"openbsd/386",
	"openbsd/amd64",
	"openbsd/arm",
	"plan9/386",
	"plan9/amd64",
	"plan9/arm",
	"solaris/amd64",
	"windows/386",
	"windows/amd64",
}

var wg sync.WaitGroup

func main() {
	_ = os.Mkdir("build", os.ModeDir)

	wg.Add(len(platforms))

	for index := range platforms {
		go build(index)
	}

	wg.Wait()

}

var blue = color.New(color.FgCyan).SprintFunc()
var magenta = color.New(color.FgMagenta).SprintFunc()

func build(n int) {
	dir, _ := os.Getwd()
	dir = filepath.Base(dir)

	ext := ""

	element := platforms[n]

	goos := strings.Split(element, "/")[0]
	goarch := strings.Split(element, "/")[1]

	fmt.Printf("%s: %s - %s\n", magenta("Building"), blue(goos), blue(goarch))

	switch goos {
	case "windows":
		ext = ".exe"
	}

	cmd := exec.Command("go", "build", "-o", "./build/"+dir+"_"+goos+"_"+goarch+ext)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)

	if err := cmd.Run(); err != nil {
		color.Red("Could not build: %s - %s.", goos, goarch)
		color.Red("You probably used code which is not compatible with this platform!")
		color.Red("Error: %s", err.Error())
	}

	wg.Done()
}
