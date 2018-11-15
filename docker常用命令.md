
删除none镜像
```
docker rmi $(docker images -f "dangling=true" -q)
```
