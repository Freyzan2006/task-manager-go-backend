# task-manager-go-backend

![task-manager-go-backend](./public/intro.png)

# Описание 
task-manager-go-backend - это REST API для управления задачами. 
Позволяет создавать, обновлять, удалять и получать задачи.


# Установка & Запуск(Native)

1. Скачайте репозиторий
```bash
git clone https://github.com/Freyzan2006/task-manager-go-backend.git
cd task-manager-go-backend
```
2. Установите зависимости
```bash
go get
```

3. Просмотре флагов:
```bash
go run cmd/main.go -h
```

3. Сборка(Можете в конце указать нужные флаги)
```bash
go build -o build/task-manager ./cmd/main.go 
```

4. Запуск
```bash
./build/task-manager
```

# Установка & Запуск(Docker)

1. Скачайте репозиторий 
```bash
git clone https://github.com/Freyzan2006/task-manager-go-backend.git
cd task-manager-go-backend
```

2. Скачайте Docker
```bash
https://docs.docker.com/get-started/get-docker/
```

3. Запустите Docker
```bash
docker build -t task-manager-go-backend .
docker run -p 8080:8080 -d --name task-manager-go-backend freyzan2006/task-manager-go-backend
```


----------------------------------