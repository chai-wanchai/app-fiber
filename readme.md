how to run with docker
```sh
docker compose -f docker-compose.yml up -d
```

Need Set Env:

MYSQL_USERNAME=admin

MYSQL_PASSWORD=password

MYSQL_HOST=localhost

MYSQL_PORT=3306

MYSQL_DATABASE=master


API :

GET /api/data
    - get all raw data

GET /api/sum
    - get sum data

GET /api/random
    - random data


CURL
```sh
    curl --location --request GET 'http://localhost:3002/api/sum'
```