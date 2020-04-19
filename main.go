package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func createFolder(path string, name string) error {
	err := os.Mkdir(path+"/"+name, os.ModeDir)
	if err != nil {
		return err
	}
	return nil
}

func createFile(path, template, end string) error {
	m, err := assets.Open("/" + template + ".txt")
	if err != nil {
		return err
	}
	mText, err := ioutil.ReadAll(m)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path+"/"+template+end, os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(mText)
	if err != nil {
		return err
	}
	return nil
}

func createMod(projectPath, repo string) error {
	cmd := exec.Command("go", "mod", "init", repo)
	cmd.Dir = projectPath
	err := cmd.Run()
	if err != nil {
		return err
	}
	get := exec.Command("go", "get", "-u", "github.com/gorilla/mux")
	get.Dir = projectPath
	err = get.Run()
	if err != nil {
		return err
	}
	return nil
}

func gitInit(projectPath string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = projectPath
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func createTypescript(frontendPath string) error {
	err := createFolder(frontendPath, "src")
	if err != nil {
		return err
	}
	err = createFile(frontendPath, "src/index", ".tsx")
	if err != nil {
		return err
	}
	err = createFile(frontendPath, "tsconfig", ".json")
	if err != nil {
		return err
	}
	return nil
}

func createBuild(frontendPath string) error {
	err := createFolder(frontendPath, "dist")
	if err != nil {
		return err
	}
	err = createFile(frontendPath, "dist/index", ".html")
	if err != nil {
		return err
	}
	err = createFile(frontendPath, "webpack.config", ".js")
	if err != nil {
		return err
	}
	return nil
}

func fatalErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("provide a project name and mod file repo")
	}
	project := os.Args[1]
	modRepo := os.Args[2]
	err := createFolder("./", project)
	fatalErr(err)
	projectPath := "./" + project
	err = createFile(projectPath, "main", ".go")
	fatalErr(err)
	err = createFolder(projectPath, "frontend")
	fatalErr(err)
	frontendPath := projectPath + "/frontend"
	err = createFile(frontendPath, "package", ".json")
	fatalErr(err)
	err = createTypescript(frontendPath)
	fatalErr(err)
	err = createBuild(frontendPath)
	fatalErr(err)
	err = createMod(projectPath, modRepo)
	fatalErr(err)
	err = gitInit(projectPath)
	fatalErr(err)

	log.Println(project, "created! Now go into the frontend folder run \"npm install\"")
}
