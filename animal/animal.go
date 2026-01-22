package animal

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

func (a *Animal) AddAge() {
	a.Age++
}

func (a *Animal) Add_child() {
	if a.Alive {
		a.Events = append(a.Events, "Животное дало потомство")
	}
}

func (a *Animal) Die() {
	a.Alive = false
	a.Events = append(a.Events, "Животное умерло")
}

func (a *Animal) Update() {
	if !a.Alive {
		return
	}

	a.AddAge()

	if rand.IntN(100) < 30 {
		a.Add_child()
	}

	if rand.IntN(100) < 60 {
		a.Die()
	}
}

func (a *Animal) FlushInfo() string {
	events := ""

	if len(a.Events) == 0 {
		events = "Не было событий за год"
	} else {
		events = strings.Join(a.Events, ",\n")
	}

	info := fmt.Sprintf("Имя: %s, Возраст: %d, Вид животного: %s, Жив ли: %t, События за год: %s\n", a.Name, a.Age, a.Type, a.Alive, events)

	a.Events = []string{}

	return info
}
