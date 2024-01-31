# DeepRead 企业微信回调服务

## 1.部署运行

项目部署采用容器部署,命令如下

```shell
# 生成镜像
docker build -t ptonlix/go-deepdread:v0.0.5 -f Dockerfile .

# 或者直接从docker hub下载
docker pull  ptonlix/go-deepdread:v0.0.5

# 生成网络
docker network create test-network

# 运行Redis Redis可以自行部署，修改app.conf中配置项即可
docker pull redis

docker run -p 6379:6379 --name redis \
--network test-network \
--network-alias redis \
-v /e/Docker/redis/conf/redis.conf:/etc/redis/redis.conf \
-v /e/Docker/redis/data:/data \
-d redis redis-server /etc/redis/redis.conf \
--appendonly yes

# 运行容器
docker run -itd -p 8800:8800 --restart always -v /root/go-deepread/conf:/app/conf ptonlix/go-deepdread:v0.0.5
docker run -itd -p 8800:8800 --restart always -v ./conf:/app/conf -v ./log:/app/log ptonlix/go-deepdread:v0.0.5

# 查看日志
docker logs -f 6b05e1c81380
```
