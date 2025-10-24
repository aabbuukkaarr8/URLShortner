## URL Shortener Service

Микросервис для сокращения ссылок:
- REST API: создание и редирект коротких URL
- PostgreSQL: хранение ссылок и статистики
- Redis: кэширование коротких ссылок
- Retry/Backoff для БД и Redis
- Middleware для логирования и метрик
- Конфигурация через TOML (`config.toml`)
- Docker Compose для локального запуска

---

### Архитектура
- `internal/handler`: Gin-обработчики (создание, редирект)
- `internal/service`: бизнес-логика (валидация, генерация short code)
- `internal/repository`: доступ к PostgreSQL
- `internal/cache`: работа с Redis
- `internal/config`: загрузка конфигурации
- `pkg/retry`: стратегия повторных попыток (экспоненциальный backoff)
- `cmd/api`: запуск HTTP-сервера

---

### Требования
- Go 1.23+
- Docker + Docker Compose
- PostgreSQL + Redis

---

### Быстрый старт

1) Поднять инфраструктуру:
```bash
docker compose up -d db redis
Запуск приложения:
go run cmd/api/main.go -config config.toml
Проверка здоровья:
curl -s http://localhost:8080/healthz
Конфигурация (config.toml)
[server]
http_port = ":8080"

[database.master]
host = "db"
port = "5432"
user = "postgres"
pass = ""
name = "shortener_db"
ssl_mode = "disable"

[redis]
addr = "redis:6379"
password = ""
db = 0

[shortener]
base_url = "http://localhost:8080"
code_length = 8
API
POST /api/shorten
Создание короткой ссылки.
Пример запроса

curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/long/url"}'
Пример ответа
{
  "short_url": "http://localhost:8080/Ab3dE7xK"
}
GET /:code
Редирект на оригинальный URL по короткому коду.
Пример

curl -v http://localhost:8080/Ab3dE7xK
→ 302 Found → Location: https://example.com/long/url
Путь выполнения кода
Handler.Shorten принимает JSON с исходным URL
Сервис проверяет формат, генерирует короткий код
Данные сохраняются в PostgreSQL
Ссылка кэшируется в Redis для ускорения редиректа
Handler.Redirect по коду ищет ссылку в Redis
При отсутствии — берёт из PostgreSQL, обновляет кэш
Структура проекта
.
├── cmd/
│   └── api/          # HTTP сервер
├── internal/
│   ├── handler/      # REST API
│   ├── service/      # Бизнес-логика
│   ├── repository/   # PostgreSQL слой
│   ├── cache/        # Redis клиент
│   ├── config/       # TOML-конфиг
├── pkg/
│   └── retry/        # Backoff стратегия
├── config.toml
└── go.mod
Модель данных (internal/model/url.go)
type URL struct {
	ID        int64     `db:"id"`
	Code      string    `db:"code"`
	Original  string    `db:"original"`
	CreatedAt time.Time `db:"created_at"`
}
Retry / Backoff
Реализовано в pkg/retry
Используется при сохранении в PostgreSQL и обращении к Redis
Экспоненциальный backoff с jitter
Пример Docker Compose
services:
  db:
    image: postgres:16
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: shortener_db
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  redis:
    image: redis:7
    ports:
      - "6379:6379"

volumes:
  pgdata:
Тестирование
go test ./...
Развитие
Добавить авторизацию (JWT)
Метрики Prometheus (/metrics)
TTL для временных ссылок
OpenAPI документация
Полезные ссылки
Gin: https://github.com/gin-gonic/gin
go-redis: https://github.com/redis/go-redis
TOML config: https://github.com/pelletier/go-toml