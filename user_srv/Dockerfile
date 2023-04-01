FROM golang:1.20
ENV BASE_PATH=/data
# 安装vim插件
RUN yum install -y vim
# 声明当前容器中服务的端口
EXPOSE 8088
# 指定在容器后，终端登录进来的默认目录，一个落脚点,路径不存在则会自动创建
WORKDIR $BASE_PATH/project
ADD main.sh .
CMD ./main.sh
