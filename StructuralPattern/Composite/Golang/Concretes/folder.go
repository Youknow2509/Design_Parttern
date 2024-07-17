package Concrete

import (
	"fmt"
	"test/Interfaces"
)

type Folder struct {
	name string
	files []Interface.FileSystemConponent
}

// Constructor
func NewFolder(name string) Interface.FileSystemConponent {
	return &Folder{name: name}
}

func(f *Folder) Add(file Interface.FileSystemConponent) {
	f.files = append(f.files, file)
}

func (f *Folder) ShowDetails() {
	fmt.Println("Folder name: " + f.name)
	for _, file := range f.files {
		file.ShowDetails()
	}
}