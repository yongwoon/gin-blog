# Set develop

- env file

```bash
cp ./api/.env.sample ./api/.env
```

- image build

```bash
docker-compose build --no-cache
```

- docker run

```bash
docker-compose up
```

- db create

```bash
docker-compose exec api sh
```

```bash
soda create -e development
soda create -e test
```

- db migrate

```bash
soda migrate
```

- Access to http://localhost:3001/api/v1/ping

以下の response が帰ってきたら設定完了

```json
{
    "message": "ping"
}
```
