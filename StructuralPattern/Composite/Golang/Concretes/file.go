package Concrete

import (
	"fmt"
	"test/Interfaces"
)

type File struct {
	name string
}

// Constructor
func NewFile(name string) Interface.FileSystemConponent {
	return &File{name: name}
}

func (f *File) ShowDetails() {
	fmt.Println("File name: " + f.name)
}