Описание проекта:

Test_API - это проект на Go, который использует PostgreSQL в качестве базы данных и Docker для контейнеризации. 
Проект предоставляет REST API для работы с данными из Binance, позволяя добавлять новые тикеры и получать информацию о них.

Стек:
- Go
- PostgreSQL
- Docker
- Docker Compose

Установка и запуск:

1)Клонируйте репозиторий:
  git clone "https://github.com/theborzet/test_task-binance-"
  
2)Создайте файл .env:
  В корне проекта найдите файл .env и добавьте в него свои переменные:
  
3)Запустите Docker Compose:
  docker-compose up --build
  
4)Проверьте, что контейнеры запущены:
  Убедитесь, что контейнеры app и db запущены и работают корректно.
  
Использование:

1)Добавление нового тикера
  Для добавления нового тикера используйте похожий запрос:
        Invoke-WebRequest -Uri "http://<IP-адрес>:3000/add_ticker" -Method POST -Headers @{ "Content-Type" = "application/json" } -Body '{ "ticker": "ETH" }'
        
2)Получение данных о тикере
  Для получения данных о тикере используйте похожий запрос:
        Invoke-WebRequest -Uri "http://<IP-адрес>:3000/fetch?ticker=ETH&date_from=2024-06-18%2009:09:30&date_to=2024-06-18%2009:15:32" -Method GET
