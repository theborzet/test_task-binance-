# Описание проекта

Test_API - это проект на Go, который использует PostgreSQL в качестве базы данных и Docker для контейнеризации. Проект предоставляет REST API для работы с данными из Binance, позволяя добавлять новые тикеры и получать информацию о них.

## Стек технологий

- Go
- PostgreSQL
- Docker
- Docker Compose

## Установка и запуск

1. Клонируйте репозиторий:
   git clone "https://github.com/theborzet/test_task-binance-"

2. Измените файлы configs/config.example.yaml и .env.example:

   Введите конфигурацию для БД в этих файлах.

   Уберите ".example" из названия файла

4. Запустите Docker Compose:

   docker-compose up --build

6. Проверьте, что контейнеры запущены:

   Убедитесь, что контейнеры app и db запущены и работают корректно.

## Использование

1. Добавление нового тикера:
   Для добавления нового тикера используйте следующий запрос:
   
   Invoke-WebRequest -Uri "http://<IP-адрес>:3000/add_ticker" -Method POST -Headers @{ "Content-Type" = "application/json" } -Body '{ "ticker": "ETH" }'

3. Получение данных о тикере:
   Для получения данных о тикере используйте следующий запрос:

   (Invoke-WebRequest -Uri "http://localhost:3000/fetch?ticker=ETH&date_from=18.06.24%2009:09:30&date_to=18.06.24%2009:17:09" -Method GET).Content
