package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SergeiKhanlarov/seri-go-logger-loki/utils"
)

// LokiClient интерфейс для отправки логов в Loki
type LokiClient interface {
	// SendLog отправляет лог-сообщение в Loki
	// level: уровень логирования (error, warn, info, debug)
	// message: текст сообщения
	// params: дополнительные параметры для добавления в лейблы
	SendLog(level, message string, params map[string]interface{}) error
}

// lokiClient реализация клиента для отправки логов в Loki
type lokiClient struct {
	config LokiConfig
	stream map[string]interface{}
	client *http.Client
}

// LokiStream представляет поток логов в формате Loki
type LokiStream struct {
	// Stream лейблы потока для индексации и поиска
	Stream map[string]string `json:"stream"`
	// Values массив пар [timestamp, message]
	Values [][]string `json:"values"`
}

// LokiPayload основной payload для отправки в Loki API
type LokiPayload struct {
	// Streams массив потоков логов
	Streams []LokiStream `json:"streams"`
}

// NewLokiClient создает новый экземпляр клиента для Loki
// config: конфигурация клиента (URL, job, app)
func NewLokiClient(config LokiConfig) LokiClient {
	return &lokiClient{
		config: config,
		stream: map[string]interface{}{
			"job": config.Job,
			"app": config.App,
		},
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SendLog отправляет лог-сообщение в Loki
// Принимает уровень логирования, сообщение и дополнительные параметры
// Возвращает ошибку в случае неудачи
func (l *lokiClient) SendLog(level, message string, params map[string]interface{}) error {
	// Получаем текущее время в наносекундах для точного таймстампа
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())

	// Объединяем базовые лейблы с уровнем логирования и пользовательскими параметрами
	stream := utils.MergeMaps(l.stream, map[string]interface{}{"level": level}, params)
	
	// Формируем payload для Loki
	payload := LokiPayload{
		Streams: []LokiStream{
			{
				Stream: stream,
				Values: [][]string{
					{timestamp, message},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal Loki payload: %w", err)
	}

	resp, err := l.client.Post(l.config.LokiUrl+"/loki/api/v1/push", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send log to Loki: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("loki returned non-ok status: %d", resp.StatusCode)
	}

	return nil
}