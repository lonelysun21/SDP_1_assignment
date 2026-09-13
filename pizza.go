package main

import (
	"fmt"
	"strings"
)

// Pizza — это "продукт": итоговый объект, который в конце концов
// выдаёт билдер. Сама структура ничего не проверяет и не решает —
// она просто хранит готовое состояние. Вся логика (дефолты, правила,
// валидация) живёт в билдерах, а не здесь.
type Pizza struct {
	Style    string // "Classic" или "Vegan" — проставляется билдером в конце Build()
	Size     string
	Dough    string
	Sauce    string
	Cheese   string
	Crust    string
	Toppings []string
}

// String делает Pizza красиво печатаемой через fmt.Println.
func (p Pizza) String() string {
	toppings := "нет"
	if len(p.Toppings) > 0 {
		toppings = strings.Join(p.Toppings, ", ")
	}
	return fmt.Sprintf(
		"%s пицца [%s, тесто: %s, борт: %s]\n  Соус: %s\n  Сыр: %s\n  Начинка: %s",
		p.Style, p.Size, p.Dough, p.Crust, p.Sauce, p.Cheese, toppings,
	)
}
