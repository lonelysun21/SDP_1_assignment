package main

import "fmt"

// ClassicPizzaBuilder собирает "обычную" пиццу: правила мягкие,
// единственное реальное ограничение — не перегрузить пиццу начинками.
type ClassicPizzaBuilder struct {
	pizza Pizza
}

func NewClassicPizzaBuilder() *ClassicPizzaBuilder {
	return &ClassicPizzaBuilder{}
}

func (b *ClassicPizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}

func (b *ClassicPizzaBuilder) SetDough(dough string) PizzaBuilder {
	b.pizza.Dough = dough
	return b
}

func (b *ClassicPizzaBuilder) SetSauce(sauce string) PizzaBuilder {
	b.pizza.Sauce = sauce
	return b
}

func (b *ClassicPizzaBuilder) SetCheese(cheese string) PizzaBuilder {
	b.pizza.Cheese = cheese
	return b
}

func (b *ClassicPizzaBuilder) AddTopping(topping string) PizzaBuilder {
	b.pizza.Toppings = append(b.pizza.Toppings, topping)
	return b
}

func (b *ClassicPizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.pizza.Crust = crust
	return b
}

const maxClassicToppings = 6

// Build проверяет обязательные поля, подставляет дефолты на всё,
// что осталось пустым, и возвращает готовую Pizza. Если что-то
// нарушено — возвращает error вместо паники: это Go-шный аналог
// исключений из java-варианта у одногруппников, роль та же —
// не дать собрать невалидный продукт.
func (b *ClassicPizzaBuilder) Build() (Pizza, error) {
	if err := requireSize(b.pizza.Size); err != nil {
		return Pizza{}, err
	}
	if len(b.pizza.Toppings) > maxClassicToppings {
		return Pizza{}, fmt.Errorf(
			"слишком много начинок для классической пиццы: %d (максимум %d)",
			len(b.pizza.Toppings), maxClassicToppings,
		)
	}

	if b.pizza.Dough == "" {
		b.pizza.Dough = "традиционное"
	}
	if b.pizza.Sauce == "" {
		b.pizza.Sauce = "томатный"
	}
	if b.pizza.Cheese == "" {
		b.pizza.Cheese = "моцарелла"
	}
	if b.pizza.Crust == "" {
		b.pizza.Crust = "обычный"
	}

	b.pizza.Style = "Classic"
	return b.pizza, nil
}
