/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Iterator Pattern
 */

package main

import (
	"v/Concretes"
)

func main() {
	name_collection := concretes.NewNameCollection()
    iterator := name_collection.CreateIterator()

	for iterator.HasNext() {
		name := iterator.Next().(string)
		println(name)
	}
}