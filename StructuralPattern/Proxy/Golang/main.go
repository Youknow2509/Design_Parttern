/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Froxy Pattern
 */

package main

import (
	"fmt"
	"v/Concretes"
	"v/Interfaces"
)

/**
 * Class ProxyImage implements for Image
 */
type ProxyImage struct {
	fileName string
	realImage interfaces.Image
}

// Constructor
func NewProxyImage(fileName string) interfaces.Image {
	return &ProxyImage{fileName: fileName}
}

// Function Display override interface Image
func (proxyImage *ProxyImage) Display() {
	if proxyImage.realImage == nil {
		proxyImage.realImage = concretes.NewRealImage(proxyImage.fileName)
	}
	proxyImage.realImage.Display()
}

func main() {
	proxyImage := NewProxyImage("test_10mb.jpg")
    // Image will be loaded from disk
    proxyImage.Display()
    fmt.Println("")
	// Image will not be loaded from disk
    proxyImage.Display()
}