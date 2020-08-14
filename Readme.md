## Installation

```bash
make build-docker
make run-docker
```

## Usage
8090 - порт по умолчанию, задаётся переменной окружения PORT

*localhost:8090/generateKey/{eventId}* - создаёт новый ключ, возвращает токен (keyName)
*localhost:8090/access/{eventId}/{clientId}/{keyName}* - возвращает access key, если токен(keyName) валиден,
можно задать другой URL через переменную окружения ACCESS_KEY_URL