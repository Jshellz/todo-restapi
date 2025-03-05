## Тестовое задание REST API на Go с Fiber, PostgreSQL(package pgx)

### Структура проекта
1. database -> соедение с бд
2. handlers -> котроллеры запросов CRUD
3. migrations -> сама миграция для бд pgx
4. models -> структура тестового задания TODO
5. docker-compose.yml -> контейнер для локальной разработки, образ базы данных
6. go.mod & go.sum -> установочные зависемости
7. main.go -> отправная точка для запуска проекта на go

С unit тестами для каждой функции(кроме main)

![alt text](https://github.com/Jshellz/todo-restapi/test_work/photo)