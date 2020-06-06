# installed packages

- default

```bash
docker-compose run --rm api go get github.com/gin-gonic/gin
docker-compose run --rm api go get github.com/go-sql-driver/mysql
docker-compose run --rm api go get github.com/gin-contrib/cors
docker-compose run --rm api go get github.com/joho/godotenv
```

- database

```bash
docker-compose run --rm api go get github.com/jinzhu/gorm
docker-compose run --rm api go get github.com/gobuffalo/pop/...
```

- logger

```bash
docker-compose run --rm api go get github.com/sirupsen/logrus
docker-compose run --rm api go get github.com/lestrrat-go/file-rotatelogs
docker-compose run --rm api go get github.com/rifflock/lfshook
```

- test

```bash
docker-compose run --rm api go get github.com/stretchr/testify
```

## Postpone

- [viper](https://github.com/spf13/viper)
