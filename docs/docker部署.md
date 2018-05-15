> blog: https://blog.csdn.net/fong_613/article/details/80323279

##gateway

```
sudo docker run -it -p 8080:8080 -p 8088:8088 --name=gateway --network=my_net2 --ip 172.22.16.2 temp

/go-path/src/YouHome_server/services/Gateway/main


```



## scene

```
sudo docker run -it --name=scene --network=my_net2 --ip 172.22.16.3 temp

/go-path/src/YouHome_server/services/Scene/main

```



## devcie

```
sudo docker run -it --name=device --network=my_net2 --ip 172.22.16.4 temp

/go-path/src/YouHome_server/services/Device/main

```



## user

```
sudo docker run -it --name=user --network=my_net2 --ip 172.22.16.5 temp

/go-path/src/YouHome_server/services/User/main

```



## 进入数据库

```
sudo docker run -it --net host mysql:5.7 "sh"

mysql -h127.0.0.1 -P3306 -uroot -proot
```

