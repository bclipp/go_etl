#### Cockroach DB demo

### URLS
[Admin GUI](http://localhost:8080/)

### command line SQL

```
cockroach sql \
--url='postgres://<username>:<password>@<global host>:26257/<database>?sslmode=verify-full&sslrootcert=<path to the CA certificate>' \
--execute="CREATE TABLE accounts (id INT PRIMARY KEY, balance DECIMAL);"
```
### Docker

```
cd docker 
sudo docker-compose up
```

### ETL app

```
cd src/cockroach_etl/
go run main
```