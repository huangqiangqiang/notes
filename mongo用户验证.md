MongoDB默认是无密码直接连接的

# docker-compose.yml
```
version: '2'
services:
  mongodb:
    image: mongo:latest
    environment:
      - MONGO_INITDB_ROOT_USERNAME=${username}
      - MONGO_INITDB_ROOT_PASSWORD=${password}
      - MONGO_INITDB_DATABASE=${database_name}
    ports:
      - "27017:27017"
    volumes:
      - ${PWD}/db:/data/db
```

# mongodb connect
```
mongodb://${username}:${password}@SERVER_IP:27017/${database_name}?authSource=admin
```
`authSource`表示用哪个数据库来验证用户名和密码，默认是`admin`