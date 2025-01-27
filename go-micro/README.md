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

## Code explain

Broker-service

- Giống như API gateway, sẽ nhận mọi request từ UI
- Nó dùng 1 endpoint duy nhất để điều hướng request, đó là `/handler`
- UI muốn gọi tới chẳng hạn service auth, sẽ gọi tới endpoint `/handler` và kèm theo `action=auth`
- Update: có thêm endpoint khác là `/ping` và `/log-grpc` để handle ping-pong và gRPC request. Chi tiết xem file [routes.go](./broker-service/cmd/api/routes.go)

Authentication flow:

- Click `Test Auth` in UI
- UI gọi tới Broker:
  ```bash
  curl 'http://localhost:8080/handle' \
    --data-raw '{"action":"auth","auth":{"email":"admin@example.com","password":"verysecret"}}'
  ```
- Broker gọi sang auth-service để xác thực: `http://authentication-service/authenticate`
- Auth-service xác thực và gọi sang log-service để ghi log

Log flow:

- Click `Test Log` in UI
- UI gọi tới Broker:
  ```bash
  curl 'http://localhost:8080/handle' \
    --data-raw '{"action":"log","log":{"name":"event","data":"Some kind of data"}}'
  ```
- Broker gọi sang log-service thông qua RPC:
  - Server address: `logger-service:5001`
  - Server method: `client.Call("RPCServer.LogInfo", rpcPayload, &result)`
- Log-service nhận data và ghi vào mongoDB
- Ưu điểm:
  - Không cần Marshal/Unmarshal data (parse JSON)
  - Client chỉ cần chỉ rõ tên method của RPC server + param là có thể gọi được
- Requirement:
  - Cả 2 phải define cùng kiểu data cho param --> duplicate code
  - Cả 2 phải viết bằng Go (thấy tutorial bảo thế), và đều phải implement RPC
- Nếu muốn 2 phía có thể dùng 2 ngôn ngữ khác nhau --> dùng gRPC

Gen code gRPC

- Broker và logger có dùng cả gRPC để giao tiếp, do đó cả 2 sẽ phải define file proto giống hệt nhau
- Từ file proto sẽ gen code Go, xem thêm ở file [README](./logger-service/README.md) của 2 service đó

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
