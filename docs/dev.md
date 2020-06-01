# Set develop

- image build

```bash
docker-compose build --no-cache
```

- docker run

```bash
docker-compose up db
```

- db create

```bash
# db container に access
docker-compose exec db bash

# access to mysql(password 入力部分ではそのまま Enter を押す)
mysql -uroot -p
```

```sql
CREATE DATABASE blog_development CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

exit
```

```bash
exit
```

- docker stop

起動している docker-compose を止める

- docker run

```bash
docker-compose up
```
