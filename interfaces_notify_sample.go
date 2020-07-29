package main

import "fmt"

type user struct {
	name, email string
}

type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	userAdriano := user{"adriano", "adrianomsg@gmail.com"}
	sendNotification(&userAdriano)
}
