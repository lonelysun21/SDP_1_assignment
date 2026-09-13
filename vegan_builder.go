package main

import (
	"fmt"
	"strings"
)

// VeganPizzaBuilder собирает пиццу без мяса, рыбы и обычного сыра.
// Правила валидации у неё принципиально другие, чем у
// ClassicPizzaBuilder — это и есть то самое "meaningfully different
// representation", которое требуется по заданию.
type VeganPizzaBuilder struct {
	pizza Pizza
}

func NewVeganPizzaBuilder() *VeganPizzaBuilder {
	return &VeganPizzaBuilder{}
}

func (b *VeganPizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}

func (b *VeganPizzaBuilder) SetDough(dough string) PizzaBuilder {
	b.pizza.Dough = dough
	return b
}

func (b *VeganPizzaBuilder) SetSauce(sauce string) PizzaBuilder {
	b.pizza.Sauce = sauce
	return b
}

func (b *VeganPizzaBuilder) SetCheese(cheese string) PizzaBuilder {
	b.pizza.Cheese = cheese
	return b
}

func (b *VeganPizzaBuilder) AddTopping(topping string) PizzaBuilder {
	b.pizza.Toppings = append(b.pizza.Toppings, topping)
	return b
}

func (b *VeganPizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.pizza.Crust = crust
	return b
}

// forbiddenToppings — то, чего в веганской пицце быть не может в принципе.
// Список сознательно с русскими и английскими вариантами слов.
var forbiddenToppings = []string{
	"pepperoni", "пепперони",
	"bacon", "бекон",
	"ham", "ветчина",
	"sausage", "колбаса",
	"anchovy", "анчоус",
}

func (b *VeganPizzaBuilder) Build() (Pizza, error) {
	if err := requireSize(b.pizza.Size); err != nil {
		return Pizza{}, err
	}

	if bad := findForbidden(b.pizza.Toppings, forbiddenToppings); len(bad) > 0 {
		return Pizza{}, fmt.Errorf("веганская пицца не может содержать: %v", bad)
	}

	if b.pizza.Dough == "" {
		b.pizza.Dough = "цельнозерновое"
	}
	if b.pizza.Sauce == "" {
		b.pizza.Sauce = "томатный"
	}
	if b.pizza.Cheese == "" {
		b.pizza.Cheese = "веганский сыр"
	} else if !isVeganCheese(b.pizza.Cheese) {
		return Pizza{}, fmt.Errorf("сыр %q не веганский — уточни веганский вариант", b.pizza.Cheese)
	}
	if b.pizza.Crust == "" {
		b.pizza.Crust = "тонкий"
	}

	b.pizza.Style = "Vegan"
	return b.pizza, nil
}

// isVeganCheese — простая эвристика: считаем сыр веганским, если в
// названии явно написано "веган"/"vegan" или пользователь вообще
// отказался от сыра.
func isVeganCheese(cheese string) bool {
	lower := strings.ToLower(cheese)
	return strings.Contains(lower, "vegan") ||
		strings.Contains(lower, "веган") ||
		lower == "нет" || lower == "без сыра"
}
