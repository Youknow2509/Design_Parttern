/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Adapter Pattern
 */

/**
 * Adapter Pattern: Translate Vietnamese to English
 */

package main

import (
	"fmt"
)

/**
 * Interface for vietnamese translation
 */
type vietnameseTranslate interface {
	send() string
}

/**
 * Adaptee handle translation
 */
type Adaptee struct {
}

// Handle translation
func (a *Adaptee) send(context string) string {

	// TODO: use api handle translation

	// I handle for example
	if context == "xin chao" {
		context = "Hello"
	}
	if context == "tam biet" {
		context = "Goodbye"
	}

	return "Translated: " + context
}

/**
 * Adapter for translation
 */
type Adapter struct {
	adaptee *Adaptee
}

// Constructor
func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

// Handle send
func (a *Adapter) send(context string) string {
	return a.adaptee.send(context)
}

func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)

	// Test
	fmt.Println(adapter.send("xin chao"))
	fmt.Println(adapter.send("tam biet"))

	return
}
