
package concretes
import (
	"fmt"
	"v/Interfaces"
)

/**
 * Class RealImage implementation for interface image
 */
type RealImage struct {
	fileName string
}

// Constructor
func NewRealImage(fileName string) interfaces.Image {
	rI := &RealImage{fileName: fileName}
	rI.LoadFromDisk()
	return rI
}

// Function LoadFromDisk 
func (realImage *RealImage) LoadFromDisk() {
	fmt.Println("Loading " + realImage.fileName)
}

// Deloy method for Image
func (realImage *RealImage) Display() {
	fmt.Println("Displaying " + realImage.fileName)
}
