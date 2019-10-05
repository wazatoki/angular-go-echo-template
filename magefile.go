// +build mage

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// BuildJs build javascript
func BuildJs() error {

	os.Chdir("./frontend")
	defer os.Chdir("../")

	cmd := exec.Command("npx", "ng", "build")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {

	mg.Deps(Clean)

	defer resetEnv()

	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOARCH", "amd64")

	fmt.Println("Building for linux")
	os.Setenv("GOOS", "linux")
	cmdLinux := exec.Command("go", "build", "-o", "build/project_linux_amd64.bin", ".")
	if err := cmdLinux.Run(); err != nil {
		return err
	}

	fmt.Println("Building for windows")
	os.Setenv("GOOS", "windows")
	cmdWindows := exec.Command("go", "build", "-o", "build/project_windows_amd64.exe", ".")
	if err := cmdWindows.Run(); err != nil {
		return err
	}

	fmt.Println("Building for darwin")
	os.Setenv("GOOS", "darwin")
	cmdDarwin := exec.Command("go", "build", "-o", "build/project_darwin_amd64.app", ".")
	if err := cmdDarwin.Run(); err != nil {
		return err
	}

	copyResources()

	fmt.Println("build finished !")

	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("build")
}

func resetEnv() {
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOARCH", "amd64")
	os.Setenv("GOOS", "linux")
}

func copyResources() error {

	separator := string(os.PathSeparator)
	root := "." + separator + "resources"
	targetDir := "." + separator + "build" + separator + "resources"
	files := []string{}

	// ディレクトリ内のファイルを取得
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		rel, err := filepath.Rel(root, path)
		files = append(files, rel)
		return nil
	})

	// ディレクトリを作成
	for _, file := range files {
		if fInfo, _ := os.Stat(root + separator + file); fInfo.IsDir() == true {
			os.MkdirAll(targetDir+separator+file, 0777)
		}
	}

	// コピー実行
	for _, file := range files {
		if fInfo, _ := os.Stat(root + separator + file); fInfo.IsDir() == false {
			src, err := os.Open(root + separator + file)
			if err != nil {
				panic(err)
			}

			dst, err := os.Create(targetDir + separator + file)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(dst, src)
			if err != nil {
				panic(err)
			}
			src.Close()
			dst.Close()
		}
	}

	return err

}
