// +build dev

// build.go automates proper versioning of caddy binaries.
// Use it like:   go run build.go
// You can customize the build with the -goos, -goarch, and
// -goarm CLI options:   go run build.go -goos=windows
//
// To get proper version information, this program must be
// run from the directory of this file, and the source code
// must be a working git repository, since it needs to know
// if the source is in a clean state.
//
// This program is NOT required to build Caddy from source
// since it is go-gettable. (You can run plain `go build`
// in this directory to get a binary.) However, issues filed
// without version information will likely be closed.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/WedgeServer/builds"
)

var goos, goarch, goarm, gopath string
var multiple bool

type PlatformTriple struct {
	OperatingSystem string
	Architecture string
	ArmVersion string
}

func init() {
	flag.StringVar(&goos, "goos", "", "GOOS for which to build")
	flag.StringVar(&goarch, "goarch", "", "GOARCH for which to build")
	flag.StringVar(&goarm, "goarm", "", "GOARM for which to build")
	flag.BoolVar(&multiple, "multiple", false, "Builds all supported platforms, ignoring all other flags")
}

const binaryName = "wedge"

func main() {
	flag.Parse()

	gopath = os.Getenv("GOPATH")

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ldflags, err := builds.MakeLdFlags(filepath.Join(pwd, ".."))
	if err != nil {
		log.Fatal(err)
	}

	if !multiple {
		platform := PlatformTriple{goos, goarch, goarm}
		runBuild(ldflags, &platform, binaryName)
		return
	}

	platforms := []PlatformTriple{
		PlatformTriple{"windows", "amd64", ""},
		PlatformTriple{"windows", "386", ""},
		PlatformTriple{"linux", "amd64", ""},
		PlatformTriple{"linux", "386", ""},
		PlatformTriple{"darwin", "amd64", ""},
		PlatformTriple{"linux", "arm", "6"},
		PlatformTriple{"linux", "arm", "7"},
	}

	for _, platform := range platforms {
		name := binaryName + "_" + platform.OperatingSystem + "_" + platform.Architecture

		if (platform.ArmVersion != "") {
			name += "_" + platform.ArmVersion // e.g. linux_arm_7
		}

		if (platform.OperatingSystem == "windows") {
			name += ".exe"
		}

		runBuild(ldflags, &platform, name)
	}
}

func runBuild(ldflags string, triple *PlatformTriple, binaryName string) {
	args := []string{"build", "-o", binaryName, "-ldflags", ldflags}
	args = append(args, "-asmflags", fmt.Sprintf("-trimpath=%s", gopath))
	args = append(args, "-gcflags", fmt.Sprintf("-trimpath=%s", gopath))
	cmd := exec.Command("go", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()
	for _, env := range []string{
		"CGO_ENABLED=0",
		"GOOS=" + triple.OperatingSystem,
		"GOARCH=" + triple.Architecture,
		"GOARM=" + triple.ArmVersion,
	} {
		cmd.Env = append(cmd.Env, env)
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
