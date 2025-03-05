## Тестовое задание REST API на Go с Fiber, PostgreSQL(package pgx)

### Структура проекта
1. database -> соедение с бд
2. handlers -> котроллеры запросов CRUD
3. migrations -> сама миграция для бд pgx
4. models -> структура тестового задания TODO
5. docker-compose.yml -> контейнер для локальной разработки, образ базы данных
6. go.mod & go.sum -> установочные зависемости
7. main.go -> отправная точка для запуска проекта на go

## Для запуска сделать контейнер из образа postgres и через команду go run main.go запустить сервер с пробросом к бд 


![photo](https://github.com/Jshellz/todo-restapi/tree/main/photo)