package main

import "fmt"

func main() {
	director := PizzaDirector{}

	fmt.Println("=== 1. Сборка через директора ===")
	if margherita, err := director.BuildMargherita(NewClassicPizzaBuilder()); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(margherita)
	}

	if pepperoni, err := director.BuildPepperoniFeast(NewClassicPizzaBuilder()); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(pepperoni)
	}

	if veganGarden, err := director.BuildVeganGarden(NewVeganPizzaBuilder()); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(veganGarden)
	}

	fmt.Println("\n=== 2. Ручная сборка через fluent API ===")
	custom, err := NewClassicPizzaBuilder().
		SetSize("L").
		SetDough("тонкое").
		SetSauce("белый соус").
		AddTopping("грибы").
		AddTopping("курица").
		Build()
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(custom)
	}

	fmt.Println("\n=== 3. Проверка валидации ===")

	_, err = NewClassicPizzaBuilder().SetSauce("томатный").Build()
	if err != nil {
		fmt.Println("Поймали ошибку (нет размера):", err)
	}

	_, err = NewVeganPizzaBuilder().
		SetSize("M").
		AddTopping("пепперони").
		Build()
	if err != nil {
		fmt.Println("Поймали ошибку (запрещённая начинка):", err)
	}

	_, err = NewVeganPizzaBuilder().
		SetSize("M").
		SetCheese("моцарелла").
		Build()
	if err != nil {
		fmt.Println("Поймали ошибку (не веганский сыр):", err)
	}
}
