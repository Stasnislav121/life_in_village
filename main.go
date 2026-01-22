package main

import (
	"fmt"
	resident "life_in_village/Resident"

	// animal "life_in_village/animal"
	village "life_in_village/village"
)

func main() {
	village := village.Village{}

	// Создаем жителей деревни
	resident1 := &resident.Resident{Name: "Алиса", Age: 30, Married: false, Alive: true, Events: []string{}}
	resident2 := &resident.Resident{Name: "Борис", Age: 40, Married: true, Alive: true, Events: []string{}}

	// Создаем животных
	// animal1 := &animal.Animal{Name: "Бобик", Age: 5, Type: "собака", Alive: true, Events: []string{}}
	// animal2 := &animal.Animal{Name: "Мурка", Age: 3, Type: "кошка", Alive: true, Events: []string{}}

	// Добавляем элементы в деревню
	village.AddElement(resident1)
	village.AddElement(resident2)
	// village.AddElement(animal1)
	// village.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		village.UpdateAll()
		fmt.Println(village.ShowAllInfo())
	}
}
