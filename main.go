package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var projectName = flag.String("projectName", "", "the project name to create")

func main() {

	flag.Parse()

	if *projectName == "" {
		fmt.Printf("supply a name using the projectName flag.")
	}

	fmt.Printf("setting up Go project %s\n", *projectName)


	projectRoot := *projectName
	cmdDir := fmt.Sprintf("%s/cmd", projectRoot)
	mainDir := fmt.Sprintf("%s/cmd/%s", projectRoot, projectRoot)
	pkgDir := fmt.Sprintf("%s/pkg", projectRoot)
	mainFilePath := fmt.Sprintf("%s/cmd/%s/main.go", projectRoot, projectRoot)

	if err := os.Mkdir(projectRoot, 0700); err != nil {
		log.Fatalf("error creating parent directory /%s : %+v", projectRoot, err)
	}

	if err := os.Mkdir(cmdDir, 0700); err != nil {
		log.Fatalf("error creating directory /%s : %+v", cmdDir, err)
	}

	if err := os.Mkdir(mainDir, 0700); err != nil {
		log.Fatalf("error creating directory /%s : %+v", mainDir, err)
	}

	if err := os.Mkdir(pkgDir, 0700); err != nil {
		log.Fatalf("error creating directory /%s : %+v", pkgDir, err)
	}

	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		log.Fatalf("error creating file /%s : %+v", mainFilePath, err)
	}
	defer func() {
		if err := mainFile.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := mainFile.WriteString("package main \n\nfunc main() { \n\n}"); err != nil {
		log.Fatalf("error writing to file %s : %+v", mainFilePath, err)
	}

	if err = exec.Command("git", fmt.Sprintf("init %s", projectRoot)).Run(); err != nil {
		log.Fatalf("error initializing git repository : %+v", err)
	}

}
