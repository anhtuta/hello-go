# Working with Microservices in Go (Golang)

Build highly available, scalable, resilient distributed applications using Go

Rate: 4.5 (2502 ratings) 23,668 students

https://cbtwapac.udemy.com/course/working-with-microservices-in-go/

## How to run on local

All commands are executed in `project` folder

- Start BE: `make up_build`
- Init data for postgres: check below
- Start FE: `make start`
- Go to FE page: http://localhost:8090/, click to each button to test it
  - Check data in DB: use PgAdmin and MongoDB compass
  - Check mail result: http://localhost:8025/

## Init data

Run [users.sql](./users.sql) in for `users` database after starting the containers. This will create a new `users` table inside `users` database, and create a sample record for it

After that, try to login to it and see the data (we can use PgAdmin for easier)

```bash
psql -h localhost -p 5433 -U postgres

postgres=# \dt
Did not find any relations.
postgres=# \l
                                                List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    | ICU Locale | Locale Provider |   Access privileges
-----------+----------+----------+------------+------------+------------+-----------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            |
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | =c/postgres          +
           |          |          |            |            |            |                 | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | =c/postgres          +
           |          |          |            |            |            |                 | postgres=CTc/postgres
 users     | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            |
(4 rows)

postgres=# \c users
psql (15.6 (Homebrew), server 14.0 (Debian 14.0-1.pgdg110+1))
You are now connected to database "users" as user "postgres".
users=# SELECT * FROM users;

 id |       email       | first_name | last_name |                           password                           | user_active |     created_at      |     updated_at
----+-------------------+------------+-----------+--------------------------------------------------------------+-------------+---------------------+---------------------
  2 | admin@example.com | Admin      | User      | $2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe |           1 | 2022-03-14 00:00:00 | 2022-03-14 00:00:00
(1 row)
```

## Postgres error

Dùng PgAdmin connect tới `localhost:5432` bị lỗi sau:

```bash
➜  project git:(main) ✗ psql -U postgres
psql: error: connection to server on socket "/tmp/.s.PGSQL.5432" failed: FATAL:  role "postgres" does not exist
```

Có thể nó đang conflict với thằng postgres chạy ở local (trước t cài postgres = brew mà éo nhớ :v)

Solution: chuyển sang port `5433` cho đỡ bị conflict

## Should not use `localhost` to connect to internal service

- Container `authentication-service` sẽ kết nối tới postgresql qua host `postgres`, chính là tên của service này trong file `docker-compose`, chứ không dùng `localhost` nhé, vì chúng dùng chung network
- Tương tự, service `broker` sẽ gọi sang `authentication-service` qua và cũng không dùng `localhost` (check file [handlers.go](./broker-service/cmd/api/handlers.go)):
  ```go
  request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
  ```

## Fix warning `orphans`

Nếu gặp warning này khi start docker container:

```
Fix warning: WARN[0003] Found orphan containers ([project-adminer-1 project-pgadmin-1]) for this project.
If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up.
```

Thì thêm `--remove-orphans` vào `docker-compose` là được
