# Migration CLI (buffalo/pop)

## Initialize

- create database.yml file

```bash
soda g config
```

## Creating Databases

Once the database.yml has been configured with the appropriate settings,
and the database server is running,
Soda can create all of the databases in the database.yml file with a simple command:

```bash
soda create -a
```

You can also create just one of the configured databases by using the -e flag and the name of the database:

```bash
soda create -e test
```

## Dropping Databases

Soda can drop all of your databases, should you want to, with one command:

```bash
soda drop -a
```

You can also drop just one of the configured databases by using the -e flag and the name of the database:

```bash
soda drop -e test
```

## Migration

### Create migration file

```bash
soda generate fizz name_of_migration

# ex
soda generate fizz posts
soda generate fizz post_comments
```

### Migrate

- migration

```bash
soda migrate
# or
soda migrate up
```

- rollback

```bash
soda migrate down
```

## Link

- [github buffalo/pop](https://github.com/gobuffalo/pop)
- [pop docs](https://gobuffalo.io/en/docs)
