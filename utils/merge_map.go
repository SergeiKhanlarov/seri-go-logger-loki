package utils

import maps0 "maps"

// MergeMaps объединяет несколько карт map[string]interface{} в одну новую карту
// При наличии одинаковых ключей в исходных картах, значения из последующих карт
// перезаписывают значения из предыдущих
//
// Параметры:
//   - maps: вариативный параметр, принимающий произвольное количество карт для объединения
//
// Возвращает:
//   - map[string]interface{}: новая карта, содержащая все ключи и значения из переданных карт
//
// Пример использования:
//
//	map1 := map[string]interface{}{"a": 1, "b": "hello"}
//	map2 := map[string]interface{}{"b": "world", "c": 3.14}
//	merged := MergeMaps(map1, map2)
//	// Результат: map[string]interface{}{"a": 1, "b": "world", "c": 3.14}
//
// Примечания:
//   - Если передать nil или пустые карты, они будут проигнорированы
//   - Исходные карты не изменяются в процессе объединения
//   - Для вложенных структур рекомендуется использовать глубокое копирование
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		maps0.Copy(result, m)
	}
	return result
}