package sgloggerloki

import (
	"maps"
	"context"
	"log"

	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"github.com/SergeiKhanlarov/seri-go-logger-loki/clients"
)

// lokiProvider реализует интерфейс sglogger.LoggerProvider для отправки логов в Loki
type lokiProvider struct {
	client clients.LokiClient
	config ProviderConfig
}

// NewLokiProvider создает новый провайдер для отправки логов в Loki
//
// Параметры:
//   - config: конфигурация провайдера (уровень логирования и другие настройки)
//   - lokiClient: клиент для взаимодействия с Loki
//
// Возвращает:
//   - sglogger.LoggerProvider: интерфейс провайдера логирования
//
// Пример использования:
//
//	config := sgloggerloki.ProviderConfig{
//	    Level: sglogger.LevelInfo,
//	}
//	lokiConfig := clients.LokiConfig{
//	    LokiURL: "http://localhost:3100",
//	    Job:     "my-service",
//	    App:     "user-api",
//	}
//	client := clients.NewLokiClient(lokiConfig)
//	provider := sgloggerloki.NewLokiProvider(config, client)
//	logger := sglogger.NewLogger(provider)
func NewLokiProvider(config ProviderConfig, lokiClient clients.LokiClient) sglogger.LoggerProvider {
	return &lokiProvider{
		client: lokiClient,
		config: config,
	}
}

// Write отправляет лог-сообщение в Loki
// Реализует интерфейс sglogger.LoggerProvider
//
// Параметры:
//   - ctx: контекст для управления временем жизни операции
//   - level: уровень логирования
//   - message: текст сообщения
//   - fields: дополнительные поля для логирования
//
// Возвращает:
//   - error: ошибка, если таковая возникла при обработке лога
//
// Особенности:
//   - Проверяет уровень логирования через ShouldLog перед отправкой
//   - Преобразует уровни sglogger в строковые представления для Loki
//   - Отправляет логи асинхронно через горутину
//   - Копирует fields для избежания гонок данных
func (p *lokiProvider) Write(ctx context.Context, level sglogger.Level, message string, fields sglogger.Fields) error {
	// Проверяем, нужно ли логировать сообщение данного уровня
	if !p.ShouldLog(ctx, level) {
		return nil
	}

	// Создаем копию fields для безопасного использования в горутине
	lokiFields := make(map[string]interface{})
	maps.Copy(lokiFields, fields)

	// Преобразуем уровень логирования в строковое представление для Loki
	var levelStr string
	switch level {
	case sglogger.LevelDebug:
		levelStr = "debug"
	case sglogger.LevelInfo:
		levelStr = "info"
	case sglogger.LevelWarn:
		levelStr = "warning"
	case sglogger.LevelError:
		levelStr = "error"
	case sglogger.LevelFatal:
		levelStr = "critical"
	default:
		levelStr = "unknown"
	}

	// Асинхронная отправка в Loki
	go func() {
		if err := p.client.SendLog(levelStr, message, lokiFields); err != nil {
			// Логируем ошибку отправки в стандартный логгер
			// В продакшн можно добавить более сложную обработку ошибок
			log.Printf("Failed to send log to Loki: %v", err)
		}
	}()

	return nil
}

// ShouldLog определяет, нужно ли логировать сообщение данного уровня
// Реализует интерфейс sglogger.LoggerProvider
//
// Параметры:
//   - ctx: контекст выполнения
//   - level: уровень логирования сообщения
//
// Возвращает:
//   - bool: true если сообщение должно быть залогировано, false в противном случае
//
// Логика:
//   - Сообщение логируется если его уровень >= установленному в конфигурации
//   - Например, при LevelInfo логируются Info, Warn, Error, Fatal
func (p *lokiProvider) ShouldLog(ctx context.Context, level sglogger.Level) bool {
	return level >= p.config.Level
}

// Close закрывает провайдер и освобождает ресурсы
// Реализует интерфейс sglogger.LoggerProvider
//
// Параметры:
//   - ctx: контекст для управления временем жизни операции
//
// Возвращает:
//   - error: ошибка, если таковая возникла при закрытии
//
// В текущей реализации не требует закрытия ресурсов,
// но оставлен для будущего расширения функциональности
func (p *lokiProvider) Close(ctx context.Context) error {
	// Здесь можно добавить закрытие соединений, flush буферов и т.д.
	return nil
}