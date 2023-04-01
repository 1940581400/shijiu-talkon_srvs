FROM golang:1.20 AS budilder

ENV GORPOXY=https://goproxy.cn,direct
    GO111MODULE=on
    CGO_ENABLED=0
    GOOS=linux
    GOARCH=amd64


COPY ./user_srv .
WORKDIR ./user_srv
RUN go mod download
RUN go build -o usersrv .

FROM scratch AS deployer
COPY --from=budilder
# 声明当前容器中服务的端口
EXPOSE 8088
# 指定在容器后，终端登录进来的默认目录，一个落脚点,路径不存在则会自动创建

ADD
CMD ./main.sh
