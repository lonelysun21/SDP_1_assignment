package main

import (
	"fmt"
	"strings"
)

// requireSize — правило "размер обязателен", общее для обоих билдеров.
// Вынесено в отдельную функцию, чтобы это правило не пришлось
// копировать в каждый Build().
func requireSize(size string) error {
	if strings.TrimSpace(size) == "" {
		return fmt.Errorf("размер пиццы обязателен: укажи S, M или L")
	}
	return nil
}

// findForbidden возвращает те начинки из toppings, в которых встречается
// хотя бы одно из запрещённых слов. Используется веганским билдером,
// но сама функция ни к чему конкретному не привязана.
func findForbidden(toppings []string, forbidden []string) []string {
	var found []string
	for _, t := range toppings {
		lower := strings.ToLower(t)
		for _, f := range forbidden {
			if strings.Contains(lower, f) {
				found = append(found, t)
				break
			}
		}
	}
	return found
}
