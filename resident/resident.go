package resident

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func (r *Resident) AddAge() {
	r.Age++
}

func (r *Resident) ChangeMarried() {
	r.Married = !r.Married
	if r.Married {
		r.Events = append(r.Events, "Заключил брак")
	} else {
		r.Events = append(r.Events, "Развод")
	}
}

func (r *Resident) Die() {
	r.Alive = false
	r.Events = append(r.Events, "Житель умер")
}

func (r *Resident) Update() {
	if !r.Alive {
		return
	}

	r.AddAge()

	if rand.IntN(100) < 30 {
		r.ChangeMarried()
	}

	if rand.IntN(100) < 40 {
		r.Die()
	}
}

func (r *Resident) FlushInfo() string {
	events := ""

	if len(r.Events) == 0 {
		events = "Не было событий за год"
	} else {
		events = strings.Join(r.Events, ",\n")
	}

	isMarried := ""
	if r.Married {
		isMarried = "Состоит в браке"
	} else {
		isMarried = "Холост"
	}

	isAlive := ""
	if r.Alive {
		isAlive = "Да"
	} else {
		isAlive = "Нет"
	}

	info := fmt.Sprintf("Имя: %s:\n Возраст: %d; Статус брака: %s; Жив ли: %s; События за год: %s.\n\n", r.Name, r.Age, isMarried, isAlive, events)

	r.Events = []string{}

	return info
}
