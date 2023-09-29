package main

import (
	"fmt"
)

// Language Interface
type Language interface {
	greeting()
}

// Language strategies
type English struct{}

func (l *English) greeting() {
	fmt.Print("Hello!\n")
}

type Russian struct{}

func (l *Russian) greeting() {
	fmt.Print("Привет!\n")
}

type Italian struct{}

func (l *Italian) greeting() {
	fmt.Print("Buongiorno!\n")
}

// Context
type userMenu struct {
	lang Language
}

func (u *userMenu) setLanguage(l Language) {
	u.lang = l
}

func (u *userMenu) greeting() {
	u.lang.greeting()
}

func main() {
	menu := userMenu{}

	langEng := &English{}
	menu.setLanguage(langEng)
	menu.greeting()

	langRus := &Russian{}
	menu.setLanguage(langRus)
	menu.greeting()

	langIta := &Italian{}
	menu.setLanguage(langIta)
	menu.greeting()
}
