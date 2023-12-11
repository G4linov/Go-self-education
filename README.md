# Go selfeducation.
## План обучения:
### Подготовка окружения. Основы языка (Basics)
1. [x] [Переменные, типы данных](https://github.com/G4linov/Go-self-education/blob/main/src/Basics/les1/main.go)
2. [x] (les2) Указатели в Go   
3. [x] (les3) Слайсы  
4. [x] (les4) Map’ы в Go  
5. [x] (les5) Конструкции языка и функции 
6. [x] (les6) defer - обработка выхода из функции  
7. [x] (les7) Panic и их обработка  
8. [x] (les8) Обработка ошибок  
### Модули и пакеты (Modules_and_packages)
1. [x] (les1) Области видимости, инициализация через init() 
2. [x] (les2) Работа с зависимостями, go mod 
3. [x] (les3) Создание модулей и их версионирование 
4. [x] (les4) layout проекта (структура проекта) 
### Структуры и интерфейсы (Structures_and_interfaces)
1. [x] (les1) Структуры в Go 
2. [x] (les2) Методы структур 
3. [x] (les3) Интерфейсы и утиная типизация 
4. [x] (les4) Пустой интерфейс 
5. [x] (les5) Композитное наследование  
[Structures repo](https://github.com/G4linov/Go-struct)
### Асинхронность (Async)
1. [ ] (les1) Goroutines 
2. [ ] (les2) Go sheduler 
3. [ ] (les3) Race condition 
4. [ ] (les4) Пакеты sync и atomic 
5. [ ] (les5) Пакет errgroup  
6. [ ] (les6) sync.Pool  
7. [ ] (les7) Каналы ч.1. Deadlocks 
8. [ ] (les8) Каналы ч.2. Context 
### Тестирование, бенчмарки и профилирование (Testing) 
1. [ ] (les1) Unit-тестирование в Go  
2. [ ] (les2) Моки, стабы и генерация через GoMock 
3. [ ] (les3) Table driven test vs closure driven tests 
4. [ ] (les4) Test coverage 
5. [ ] (les5) Benchmarks 
6. [ ] (les6) Профилирование с pprof 
### Кодогерация (Code_generation)
1. [ ] (les1) Рефлексия 
2. [ ] (les2) AST 
3. [ ] (les3) Templates 
4. [ ] (les4) Решение проблем рефлексии 
5. [ ] (les5) Враппинг
### Работа с БД
1. [ ] Конфигурирование подключения к БД (postgres)
2. [ ] Работа с БД (postgres)
3. [ ] Миграции (goose)
4. [ ] Работа с базой на примере использования GORM
5. [ ] Работа с NoSql (Mongo)
### Сервер на Go (обработка запросов, context, middleware)
1. [ ] Поднимаем сервер, роутинг, первый handler
2. [ ] Работа с параметрами
3. [ ] Роутер Gorilla
4. [ ] Роутер Chi
5. [ ] Middleware
6. [ ] Контекст запроса
7. [ ] fasthttp
8. [ ] WebSockets
### Низкоуровневость + продвинутая сборка
1. [ ] Продвинутая работа с модулями
2. [ ] Сборка с использованием Idflags
3. [ ] Сборка для разных ОС
4. [ ] Пакет unsafe
5. [ ] Cgo
### Микросервисная архитектура
1. [ ] Заворачиваем наш сервис в образ
2. [ ] Продвинутая сборку образа
3. [ ] Подтягиваем конфигурации из key-value store (consul)
### Межсервисное взаимодействие
1. [ ] Синхронное взаимодействие через REST
2. [ ] gRPC взаимодействие
3. [ ] Interceptors
4. [ ] Асинхронное взаимодействие (kafka) и pub-sub
5. [ ] Работа с GraphQL в Go
### Оптимизация
1. [ ] in-memory хранение
2. [ ] Redis
3. [ ] Самописный LRU cache
### Работа сервиса в кластере
1. [ ] Формат логов и уровни логирования
2. [ ] Пишем логи (Greylog)
3. [ ] Сквозное логирование
4. [ ] Метрики (Prometheus + Grafana)
5. [ ] graceful обработка сигналов
6. [ ] healthcheck