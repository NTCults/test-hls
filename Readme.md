## Installation

```bash
make build-docker
make run-docker
```

## Usage

*localhost:8090/generateKey/{eventId}* - создаёт новый ключ, возвращает токен (keyName).

*localhost:8090/access/{eventId}/{clientId}/{keyName}* - возвращает access key, если токен(keyName) валиден