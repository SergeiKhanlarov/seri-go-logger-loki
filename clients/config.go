package clients

// LokiConfig содержит конфигурационные настройки для отправки логов в Loki
type LokiConfig struct {
    // lokiUrl - URL endpoint сервера Loki для отправки логов
    LokiUrl string
    
    // job - идентификатор job/сервиса для группировки логов в Loki
    Job string
    
    // app - название приложения, используемое как лейбл в Loki
    App string
}