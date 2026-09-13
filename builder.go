package main

// PizzaBuilder — общий контракт для всех билдеров пиццы.
// Каждый сеттер возвращает PizzaBuilder (а не свой конкретный тип) —
// именно это и даёт fluent-цепочку вызовов:
//
//	builder.SetSize("L").SetSauce("томатный").AddTopping("грибы").Build()
//
// Важный момент про Go: здесь нет наследования и нет java-шного "this",
// которое само знает, что оно на самом деле ClassicPizzaBuilder. Поэтому
// каждый конкретный билдер реализует все методы сам и в конце пишет
// return b — возвращает сам себя как интерфейс. Дублирования логики
// это не создаёт (сеттеры — однострочники), а вот реальные правила
// (валидация, дефолты) вынесены в общие функции в validation.go.
type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetDough(dough string) PizzaBuilder
	SetSauce(sauce string) PizzaBuilder
	SetCheese(cheese string) PizzaBuilder
	AddTopping(topping string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	Build() (Pizza, error)
}
