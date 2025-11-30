package sgloggerloki

import sglogger "github.com/SergeiKhanlarov/seri-go-logger"

// ProviderConfig расширяет базовую конфигурацию логгера специфичными для провайдера настройками.
// Встраивает общую конфигурацию логгера и добавляет параметры, специфичные для конкретного провайдера.
//
// Структура объединяет:
//   - sglogger.LoggerConfig - базовая конфигурация логгера (уровни логирования, формат, выходные потоки)
//   - level - специфичный для провайдера уровень логирования, который может отличаться от общего уровня
//
// Пример использования:
//
//	config := &ProviderConfig{
//	    LoggerConfig: sglogger.LoggerConfig{...},
//	    level:        sglogger.LevelDebug,
//	}
type ProviderConfig struct {
	sglogger.LoggerConfig        // Встроенная базовая конфигурация логгера
	level       sglogger.Level   // Специфичный для провайдера уровень логирования
}

// GetLevel возвращает строковое представление уровня логирования провайдера.
// Используется для совместимости с системами, которые ожидают строковые значения уровней логирования.
//
// Возвращаемые значения:
//   - "debug"    - для уровня отладки (LevelDebug)
//   - "info"     - для информационного уровня (LevelInfo) 
//   - "warning"  - для предупреждений (LevelWarn)
//   - "error"    - для ошибок (LevelError)
//   - "critical" - для критических ошибок (LevelFatal)
//   - "info"     - значение по умолчанию, если уровень не распознан
//
// Пример:
//
//	level := config.GetLevel() // возвращает "debug", "info", etc.
func (p *ProviderConfig) GetLevel() string {
	switch p.level {
	case sglogger.LevelDebug:
		return "debug"
	case sglogger.LevelInfo:
		return "info"
	case sglogger.LevelWarn:
		return "warning"
	case sglogger.LevelError:
		return "error"
	case sglogger.LevelFatal:
		return "critical"
	}

	return "info" // Значение по умолчанию
}