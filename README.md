# Task Manager
REST API для управления задачами

## Функционал
- Создание задачи
- Получение списка всех задач
- Получение задачи по ID
- Обновление задачи
- Удаление задачи

## Эндпоинты
`POST /todos`

**Тело запроса**
```json
{
    "title": "Название задачи",
    // (опционально)
    "description": "Описание задачи",
    // (опционально)
    "is_done": false
}
```
**Ответ (Пример)**
```json
{
    "id": 1,
    "title": "Название задачи",
    "description": "Описание задачи",
    "is_done": false
}
```

`GET /todos`

**Ответ (Пример)**
```json
[
    {
        "id": 1,
        "title": "Задача 1",
        "description": "Описание 1",
        "is_done": false
    },
    {
        "id": 2,
        "title": "Задача 2",
        "description": "Описание 2",
        "is_done": true
    }
]
```

`GET /todos/{id}`

**Ответ (Пример)**
```json
{
    "id": 1,
    "title": "Название задачи",
    "description": "Описание задачи",
    "is_done": false
}
```

`PUT /todos/{id}`

**Тело запроса**
```json
{
    "title": "Обновленное название",
    // (опционально)
    "description": "Обновленное описание",
    // (опционально)
    "is_done": true
}
```

`DELETE /todos/{id}`

## Запуск проекта
```bash
# Клонирование репозитория
git clone https://github.com/SANEKNAYMCHIK/task-manager.git
cd task-manager

# Запуск проекта
go run cmd/task-manager/main.go
```
**Сервер будет доступен по адресу: http://localhost:8080**
