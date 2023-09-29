package main

import (
	"fmt"
)

// Interfaces
type Observer interface {
	getNotification(string)
	getName() string
}

type Subject interface {
	register(o Observer)
	deregister(o Observer)
	notifyAll()
}

// Student object
type Reader struct {
	name string
}

func (s *Reader) getName() string {
	return s.name
}

func (s *Reader) getNotification(newsTheme string) {
	fmt.Printf("Hey, %s, there's new news about '%s'\n", s.name, newsTheme)
}

// News object
type Newsletter struct {
	subscribers []Observer
}

func (n *Newsletter) register(o Observer) {
	n.subscribers = append(n.subscribers, o)
}

func (n *Newsletter) deregister(o Observer) {
	lastInd := len(n.subscribers) - 1
	for i, obs := range n.subscribers {
		if obs.getName() == o.getName() {
			n.subscribers[lastInd], n.subscribers[i] = n.subscribers[i], n.subscribers[lastInd]
			n.subscribers = n.subscribers[:lastInd]
		}
	}
}

func (n *Newsletter) notifyAll(theme string) {
	for _, obs := range n.subscribers {
		obs.getNotification(theme)
	}
}

func (n *Newsletter) createArticle(theme string) {
	fmt.Printf("New news today: %s\n", theme)
	n.notifyAll(theme)
}

func main() {
	news := &Newsletter{}

	reader1 := &Reader{name: "Islam"}
	reader2 := &Reader{name: "Dos"}
	reader3 := &Reader{name: "Alim"}

	news.register(reader1)
	news.register(reader2)
	news.register(reader3)

	news.createArticle("Games are killing our kids")

	news.deregister(reader3)

	news.createArticle("Maybe games are not killing our kids")

}
