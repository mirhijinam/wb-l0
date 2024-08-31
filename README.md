# WB Tech: level 0

## Описание 
Демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе.

## Стек
- Golang
- PostgreSQL
- Nats-streaming
  
## Инструкция по запуску
Для запуска сервиса необходимо:

1. Склонировать репозиторий с проектом: ```git clone https://github.com/mirhijinam/wb-l0```
2. Войти с директорию с проектом: ```cd wb-l0```
3. Сбилдить окружение: ```make up```
4. Запустить приложение: ```go run main.go```

## Стресс-тесты с использованием 'wrk'

```wrk -t12 -c1000 -d10s http://localhost:8080/order/1```
```
Running 10s test @ http://localhost:8080/order/1
  12 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.89ms    1.97ms  26.03ms   79.68%
    Req/Sec     4.08k     1.03k    7.76k    71.58%
  486956 requests in 10.01s, 292.57MB read
  Socket errors: connect 757, read 104, write 0, timeout 0
  Non-2xx or 3xx responses: 486956
Requests/sec:  48658.26
Transfer/sec:     29.23MB
```
