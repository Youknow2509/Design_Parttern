/**
 * @author: Ly Tran Vinh
 * @contact: lytranvinh.work@gmail.com
 * @content: Mediator Pattern
 */
package main

import (
	"mediator_pattern/concretes"
)

func main() {
	mediator := concretes.NewChatMediatorImpl()

	user1 := concretes.NewUserImpl(mediator, "John")
	user2 := concretes.NewUserImpl(mediator, "Jane")
	user3 := concretes.NewUserImpl(mediator, "Doe")

	mediator.AddUser(user1)
	mediator.AddUser(user2)
	mediator.AddUser(user3)

	user1.Send("Hello, everyone!")
}
