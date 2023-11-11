# Install docker environment in centos

[follow this](https://docs.docker.com/engine/install/centos/#set-up-the-repository)

# start up

docker compose up -d

# mongodb setup

```shell
docker exec -it <container_id> bash
mongo
use admin
db.createUser({user:"admin",pwd:"123456",roles:["root"]})
db.auth("admin", "123456")
db.createUser({user: "root", pwd: "123456", roles: [{ role: "dbOwner", db: "test" }]})
```

# redis setup

no
