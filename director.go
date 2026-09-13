package main

// PizzaDirector знает несколько "стандартных" рецептов и не заботится
// о том, какой конкретно билдер ему передали — только бы он
// реализовывал PizzaBuilder. Это удобно для готовых пунктов меню,
// а не мешает собрать что-то своё вручную через сам билдер.
type PizzaDirector struct{}

func (d PizzaDirector) BuildMargherita(b PizzaBuilder) (Pizza, error) {
	return b.
		SetSize("M").
		SetSauce("томатный").
		SetCheese("моцарелла").
		AddTopping("базилик").
		Build()
}

func (d PizzaDirector) BuildPepperoniFeast(b PizzaBuilder) (Pizza, error) {
	return b.
		SetSize("L").
		SetSauce("острый томатный").
		SetCheese("двойная моцарелла").
		AddTopping("пепперони").
		AddTopping("халапеньо").
		Build()
}

func (d PizzaDirector) BuildVeganGarden(b PizzaBuilder) (Pizza, error) {
	return b.
		SetSize("M").
		AddTopping("шампиньоны").
		AddTopping("болгарский перец").
		AddTopping("оливки").
		Build()
}
