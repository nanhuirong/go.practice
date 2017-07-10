package main

import "fmt"

type user struct {
	name string
	email string
	//ext int
	//privileged bool
}

type admin struct {
	name string
	email string
}

type notifier interface {
	notify()
}

func (u *user) notify()  {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

func (a *admin) notify()  {
	fmt.Printf("sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifier)  {
	n.notify()
}

func main()  {
	//var bill user
	//
	//lisa := user {
	//	name:"lisa",
	//	email:"lisa@163.com",
	//	ext:123,
	//	privileged:true,
	//}
	u := user{
		name:"Bill",
		email:"bill@email.com",
	}
	sendNotification(&u)
	a := admin{
		name:"admin",
		email:"email@163.com",
	}
	sendNotification(&a)
}


