package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"time"
)

var root = "/home/ionut/facultate/seminar/"

var cTemplate = `
#include <stdio.h>

int main(void) {
    printf("Hello, World!\n");
    return 0;
}
`
var cMakeTemplate = `
cmake_minimum_required(VERSION 3.29)
project(%s C)

set(CMAKE_C_STANDARD 11)

add_executable(%s main.c)
`

func getSeminarLocation() (location string) {
	now := time.Now()
	fmt.Println(now.Month())
	location = fmt.Sprintf("%d_%s_%d/", now.Day(), now.Month().String()[:3], now.Year())
	return
}

func ensureFolderExists(folderPath string) bool {
	_, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("No dir detected \nCreating one")
		err = os.Mkdir(folderPath, 0744)
		if err != nil {
			fmt.Printf("How the fuck we got here\nI cant create at this %s location", folderPath)
			log.Fatal(err)
		}
		return false
	}
	return true
}

// legacy
func folderInit() []fs.DirEntry {
	folder := root + getSeminarLocation()
	dirs, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("No dir detected for today\nCreating one")
		err = os.Mkdir(root+getSeminarLocation(), 0744)
		if err != nil {
			fmt.Printf("How the fuck we got here\nSomething is wrong in the %s and I cant create %s\n", root, getSeminarLocation())
			log.Fatal(err)
		}
		// i am to tired to understand why i do this extra step on an empty dir
		dirs, err = os.ReadDir(folder)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dirs
}

func main() {

	//initializarea folderului pentru ziua de seminar
	currentSeminarLocation := root + getSeminarLocation()
	ensureFolderExists(currentSeminarLocation)

	// ferificarea existentei parametrilor
	exerciseName := os.Args[1:2]
	fmt.Println(exerciseName)
	if len(exerciseName) != 1 {
		log.Fatal("No name for directory provided\nPlease enter the name for the directory")
	}
	cMakeTemplate = fmt.Sprintf(cMakeTemplate, exerciseName[0], exerciseName[0])

	// verificarea existentei folderului de exercitii, in cazul in care exista sa nu mai fie creat
	exerciseLocation := currentSeminarLocation + exerciseName[0]
	if ensureFolderExists(exerciseLocation) {
		fmt.Println("Open existing folder")
	} else {
		err := os.WriteFile(exerciseLocation+"/main.c", []byte(cTemplate), 0744)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(exerciseLocation+"/CMakeLists.txt", []byte(cMakeTemplate), 0744)
		if err != nil {
			log.Fatal(err)
		}
	}
	exec.Command("code", exerciseLocation).Run()

}
