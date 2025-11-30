# seri-go-logger-loki

Пакет для интеграции seri-go-logger с Grafana Loki - горизонтально-масштабируемой системой агрегации логов.

## Возможности

- 📊 **Поддержка всех уровней логирования:** Debug, Info, Warn, Error, Fatal
- 🏷️ **Автоматическое добавление лейблов:** job, app, level
- 🔧 **Гибкая конфигурация:** настраиваемые уровни логирования и параметры подключения
- 🚀 **Асинхронная отправка:** не блокирует основное приложение
- 🔍 **Интеграция с контекстом:** поддержка trace_id и других метаданных

## Установка

```bash
go get github.com/SergeiKhanlarov/seri-go-logger-loki
```

## Быстрый старт

Базовая настройка

```go
package main

import (
    "context"
    
	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	sgloki "github.com/SergeiKhanlarov/seri-go-logger-loki"
)

func main() {
    ctx := context.Background()
    
    // Конфигурация Loki клиента
    lokiConfig := clients.LokiConfig{
        lokiUrl: "http://localhost:3100",
        job:     "user-service",
        app:     "authentication",
    }
    
    // Создание клиента Loki
    lokiClient := clients.NewLokiClient(lokiConfig)
    
    // Конфигурация провайдера
    providerConfig := sgloggerloki.ProviderConfig{
        LoggerConfig: sglogger.LoggerConfig{},
        Level: sglogger.LevelInfo,
    }
    
    // Создание провайдера
    provider := sgloggerloki.NewLokiProvider(providerConfig, lokiClient)
    
    // Создание логгера
    logger := sglogger.NewLogger(
		sglogger.LoggerConfig{}, 
		sglogger.NewFieldsHandler(),
		provider)
    
    // Использование логгера
    logger.Info(ctx, "Приложение запущено", sglogger.Fields{
        "version": "1.0.0",
        "port":    8080,
    })
}
```

### Структура пакета

Компоненты

1. LokiClient

Интерфейс для отправки логов в Loki:

```go
type LokiClient interface {
    SendLog(level, message string, params map[string]interface{}) error
}
```
2. LokiProvider

Провайдер для интеграции с seri-go-logger:

```go
type LoggerProvider interface {
    Write(ctx context.Context, level Level, message string, fields Fields) error
    ShouldLog(ctx context.Context, level Level) bool
    Close(ctx context.Context) error
}
```

## Конфигурация

### Уровни логирования

LevelDebug (0) - Отладочные сообщения<br>
LevelInfo (1) - Информационные сообщения<br>
LevelWarn (2) - Предупреждения<br>
LevelError (3) - Ошибки<br>
LevelFatal (4) - Критические ошибки<br>

## 📄 Лицензия

MIT License - смотрите файл [LICENSE](LICENSE) для деталей.

Copyright (c) 2025 Ханларов Сергей