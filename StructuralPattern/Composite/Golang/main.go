/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Bridge Pattern
 */
package main

import (
	"test/Concretes"
)

func main() {
	folder := Concrete.NewFolder("Folder 1").(*Concrete.Folder)
	file1 := Concrete.NewFile("File 1")
	file2 := Concrete.NewFile("File 2")

	folder.Add(file1)
	folder.Add(file2)

	folder.ShowDetails()
}