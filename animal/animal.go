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

	isAlive := ""
	if a.Alive {
		isAlive = "Да"
	} else {
		isAlive = "Нет"
	}

	info := fmt.Sprintf("Имя: %s:\n Возраст: %d; Вид животного: %s; Жив ли: %s; События за год: %s.\n\n", a.Name, a.Age, a.Type, isAlive, events)

	a.Events = []string{}

	return info
}
