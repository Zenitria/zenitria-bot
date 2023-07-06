# Zenitria Bot

Discord bot maded for Get XNO Discord server.

## Configuring

1. Copy `.env` example file

```
cp .env.example .env
```

2. Set `TOKEN` and `MONGODB` in `.env` file

## Build and run

### With Docker

1. Build container image

```
docker build -t zenitria-bot .
```

2. Run builded image

```
docker run -d --network=host zenitria-bot
```

### Without Docker

1. Compile source code

```
go build
```

2. Run compiled bot

```
./zenitria-bot
```
