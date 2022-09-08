package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	GENGO = `
//go:build tools

package %[2]s

import (
	_ "github.com/edwarnicke/vpplink/cmd"
	_ "%[1]s"
)

// Run using go generate -tags tools ./...
//go:generate go run github.com/edwarnicke/vpplink/cmd --binapi-package "%[1]s"
`
	DOCGO = `package %[1]s`
)

func TestGen(t *testing.T) {
	testdatadir := "./testdata"
	testpkg := "testpkg"
	binapipkg := "go.fd.io/govpp/binapi"

	if err := os.RemoveAll(testdatadir); err != nil && err != os.ErrNotExist {
		t.Errorf("failed to rm %s dir: %+v", testdatadir, err)
	}

	err := os.Mkdir(testdatadir, 0777)
	if err != nil {
		t.Errorf("failed to create %s dir: %+v", testdatadir, err)
	}
	err = os.Chdir(testdatadir)
	if err != nil {
		t.Errorf("failed to cd into %s dir: %+v", testdatadir, err)
	}
	goModInitOutput, err := exec.Command("go", "mod", "init", testpkg).CombinedOutput()
	if err != nil {
		t.Errorf("failed to run \"go mod init %s\" %+v\n %s", testpkg, err, string(goModInitOutput))
	}
	fmt.Printf("%s", goModInitOutput)

	goWorkInitOutput, err := exec.Command("go", "work", "init", ".", "../..").CombinedOutput()
	if err != nil {
		t.Errorf("failed to run \"go work init . ../..\" %+v\n %s", err, string(goWorkInitOutput))
	}
	fmt.Printf("%s", goWorkInitOutput)

	genGo := fmt.Sprintf(GENGO, binapipkg, testpkg)
	genGoFilename := "gen.go"
	if err := os.WriteFile(genGoFilename, []byte(genGo), 0600); err != nil {
		t.Errorf("failed to write %s: %+v", genGoFilename, err)
	}

	//docGoFilename := "doc.go"
	//if err := os.WriteFile(docGoFilename, []byte(DOCGO), 0600); err != nil {
	//	t.Errorf("failed to write %s: %+v", docGoFilename, err)
	//}
	//
	goGetOutput, err := exec.Command("go", "get", "go.fd.io/govpp@v0.6.0-alpha").CombinedOutput()
	if err != nil {
		t.Errorf("failed to run \"go get go.fd.io/govpp@v0.6.0-alpha\" %+v\n %s", err, string(goGetOutput))
	}
	fmt.Printf("%s", goGetOutput)

	goGenOutput, err := exec.Command("go", "generate", "-tags", "tools", "./...").CombinedOutput()
	if err != nil {
		t.Errorf("failed to run \"go generate -tags tools ./...\" %+v\n %s", err, string(goGenOutput))
	}
	fmt.Printf("%s", goGenOutput)

	goModTidyOutput, err := exec.Command("go", "mod", "tidy").CombinedOutput()
	if err != nil {
		t.Errorf("failed to run \"go mod tidy\" %+v\n %s", err, string(goModTidyOutput))
	}
	fmt.Printf("%s", goModTidyOutput)

	goBuildOutput, err := exec.Command("go", "build", "./...").CombinedOutput()
	if err != nil {
		t.Errorf("failed to run \"go build ./...\" %+v\n %s", err, string(goBuildOutput))
	}
	fmt.Printf("%s", goBuildOutput)
}
