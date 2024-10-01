# Music Library API

# Запуск сервиса

## Через Docker

```bash
make docker
```

Запускает 2 контейнера с PostgreSQL и самим сервисом. Конфигурации берутся через .env файл
После запуска в консоле выводится ссылка на Swagger эндпоинт


## Локально

```bash
make run
```

Запуск сервиса локально. Требуется запущенный PostgreSQL.


## Конфигурации

Реализованы через переменные окружения

### API

| Name             | Description                                                                                        | Required | Example                            |
| ---------------- | -------------------------------------------------------------------------------------------------- | -------- | ---------------------------------- |
| PORT             | Порт Сервера. Дефолтное значение '8080'                                                            | False    | 9000                               |
| HOST             | Хост Сервера. Дефолтное значение `localhost`                                                       | False    | `localhost`                        |
| ENV              | Среда разработки. Определяет конфигурацию логгера. Доступные значения `dev`, `prod`                | True     | `prod`                             |
| SERVICE_URL      | URL Info сервиса                                                                                   | True     | `http://localhost:2000`            |
| POSTGRES_DSN     | DSN подключения к базе данных                                                                      | True     | `postgres://:@:5432/db`            |


### PostgreSQL

Необходимы для запуска через docker-compose. Указать в соответствии с официальной документицией.


# (Пере)Генерация Swagger документации

```bash
make swagger
```

# Миграции

Реализованы при запуске сервиса.

Также миграциями можно управлять:

Up
```bash
make up
```

Down
```bash
make down
```

# TODO

- Покрыть код тестами
    - Unit
    - Integration