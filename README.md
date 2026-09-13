# Pizza Builder — Assignment #1 (Software Design Patterns, ShP-2216)

Реализация паттерна **Builder** на Go. Продукт — `Pizza`, у которой есть два
принципиально разных представления: **Classic** (мягкие правила) и **Vegan**
(своя валидация: без мяса/рыбы, без обычного сыра).

## Структура

- `pizza.go` — продукт (`Pizza`)
- `builder.go` — интерфейс `PizzaBuilder`
- `validation.go` — общие правила валидации (не дублируются между билдерами)
- `classic_builder.go` — `ClassicPizzaBuilder`
- `vegan_builder.go` — `VeganPizzaBuilder`
- `director.go` — `PizzaDirector`, готовые рецепты (Margherita, Pepperoni Feast, Vegan Garden)
- `main.go` — демонстрация: сборка через директора, ручная сборка через fluent API, проверка валидации

## Как собрать каждое представление

Через директора (готовый рецепт):

```go
director := PizzaDirector{}
pizza, err := director.BuildMargherita(NewClassicPizzaBuilder())
```

Вручную через fluent API:

```go
pizza, err := NewClassicPizzaBuilder().
    SetSize("L").
    SetSauce("белый соус").
    AddTopping("грибы").
    Build()
```

```go
pizza, err := NewVeganPizzaBuilder().
    SetSize("M").
    AddTopping("шампиньоны").
    Build()
```

`Build()` возвращает `(Pizza, error)` — если обязательное поле не задано
или нарушено правило конкретного стиля пиццы, вернётся ошибка вместо паники.

## Как запустить

```
go run .
```