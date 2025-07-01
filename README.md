
# Toll Calculator

Toll Calculator — это система на основе микросервисов, предназначенная для расчета платы за проезд транспортных средств на основе данных об их перемещении. Проект симулирует устройства на борту автомобилей (On-Board Units, OBU), которые отправляют данные о местоположении, а затем эти данные обрабатываются для вычисления пройденного расстояния и формирования счета (инвойса) для оплаты.

## Описание проекта

Система состоит из нескольких микросервисов, которые взаимодействуют через WebSocket, Kafka и gRPC/HTTP. Основная задача — сбор геолокационных данных от OBUs, расчет расстояния, агрегация данных и генерация счетов. Kafka используется как брокер сообщений для передачи данных между сервисами.

## Архитектура

Проект включает следующие микросервисы:

- **obu**: Симулятор бортовых устройств (On-Board Units), который генерирует случайные координаты (широта и долгота) и отправляет их через WebSocket в `data_receiver`.

- **data_receiver**: Принимает данные от OBUs через WebSocket, логирует их и отправляет в Kafka-топик `obudata`.

- **distance_calculator**: Читает данные из Kafka, вычисляет расстояние, пройденное каждым автомобилем, и отправляет результаты в `aggregator` через gRPC.

- **aggregator**: Агрегирует данные о расстоянии для каждого автомобиля и вычисляет итоговый счет (инвойс). Поддерживает HTTP и gRPC интерфейсы для взаимодействия.

- **gateway**: Служит API-шлюзом для взаимодействия с сервисами.

- **types**: Пакет с общими типами данных и сгенерированными файлами Protobuf для gRPC.

### Поток данных

1. **OBU → Data Receiver**: OBU отправляет геолокационные данные через WebSocket.
2. **Data Receiver → Kafka**: Данные записываются в Kafka-топик `obudata`.
3. **Kafka → Distance Calculator**: Сервис читает данные, вычисляет расстояние и отправляет его в `aggregator`.
4. **Distance Calculator → Aggregator**: Передача данных через gRPC.
5. **Aggregator**: Накопление данных и генерация инвойсов, доступных через HTTP.

## Предварительные требования

Для работы с проектом вам понадобятся следующие инструменты:

- **Go**: Версия 1.16 или выше.
- **Docker**: Для запуска Kafka и Zookeeper.
- **protoc**: Компилятор Protocol Buffers.
- **protoc-gen-go** и **protoc-gen-go-grpc**: Плагины для генерации Go-кода из Protobuf.

### Установка protoc и плагинов для Go

#### Для Linux или WSL2

```bash
sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
```

#### Для macOS (через Homebrew)

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
```

#### Настройка PATH

Убедитесь, что директория `go/bin` добавлена в ваш PATH:

```bash
export PATH="${PATH}:${HOME}/go/bin"
```

#### Установка зависимостей Go

Установите необходимые пакеты:

```bash
go get google.golang.org/protobuf
go get google.golang.org/grpc
```

## Инструкции по установке

1. **Клонируйте репозиторий**:

   ```bash
   git clone <repository-url>
   cd toll-calculator
   ```

2. **Настройте Kafka и Zookeeper**:

   Используйте файл `docker-compose.yml` для запуска Kafka и Zookeeper:

   ```bash
   docker-compose up -d
   ```

   Это запустит Zookeeper и брокер Kafka, который будет доступен на порту `9092`.

   Альтернативно, для быстрого запуска Kafka без Zookeeper можно использовать:

   ```bash
   docker run --name kafka -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest
   ```

3. **Сгенерируйте Protobuf-код** (при необходимости):

   Если вы изменили `types/ptypes.proto`, выполните:

   ```bash
   protoc --go_out=. --go-grpc_out=. types/ptypes.proto
   ```

   Если сгенерированные файлы уже присутствуют, этот шаг можно пропустить.

## Запуск сервисов

```
//TODO (ПОКА ЗАПУСКАЮТСЯ ПО ОТДЕЛЬНОСТИ)
```
## Использование

После запуска всех сервисов система начнет работать следующим образом:

- **OBU Simulator** отправляет данные о местоположении в `data_receiver` через WebSocket.
- **Data Receiver** записывает данные в Kafka-топик `obudata`.
- **Distance Calculator** вычисляет расстояние и отправляет его в `aggregator` через gRPC.
- **Aggregator** накапливает данные и формирует инвойсы.

### Получение инвойса

Чтобы получить инвойс для конкретного OBU, используйте HTTP-запрос:

```bash
curl "http://localhost:8080/invoice?obu=<obuID>"
```

Замените `<obuID>` на идентификатор OBU (целое число), для которого нужен инвойс.

Пример ответа:

```json
{
  "obuID": 123,
  "value": 45.67,
  "totalamount": 143.8605
}
```

