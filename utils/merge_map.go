package utils

import (
	"fmt"
	"strconv"
)

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
//	// Результат: map[string]string{"a": "1", "b": "world", "c": "3.14"}
//
// Примечания:
//   - Если передать nil или пустые карты, они будут проигнорированы
//   - Исходные карты не изменяются в процессе объединения
//   - Для вложенных структур рекомендуется использовать глубокое копирование
func MergeMaps(maps ...map[string]interface{}) map[string]string {
    result := make(map[string]string)
    
    for _, m := range maps {
        for key, value := range m {
            switch v := value.(type) {
            case string:
                result[key] = v
            case bool:
                result[key] = strconv.FormatBool(v)
            case int:
                result[key] = strconv.Itoa(v)
            case int64:
                result[key] = strconv.FormatInt(v, 10)
            case float64:
                result[key] = strconv.FormatFloat(v, 'f', -1, 64)
            default:
                result[key] = fmt.Sprint(v)
            }
        }
    }
    
    return result
}