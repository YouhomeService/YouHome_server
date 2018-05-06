先定义了一些初次迭代用得到的API，以后要添加新功能再加吧

url中的信息，只是表明了post要传递什么参数

 

 

1、返回用户id=xx，场景=yy的设备信息。yy=“all”时返回所有信息。

GET/all?id=xx&scene=yy

{

{

type:light

data:{

device_id:xx

name:xx

state:xx

scene:xx

...

}

}

...

}

 

2、返回设备的详细信息（状态等）

GET/device?id=xx&device_id==xx

{

device_id: xx

data:{

...

}

}

 

3、更改设备状态：

POST /device?state=xx

{

state:xx

data:{

 ...

}//如果有必要的话

}

 

4、获取用户信息：

GET /user?id=xx

{

 username: xx

 ...

}

 

5、更改用户信息，info1，info2是要更改的属性：

POST  /user?id=xx&info1=xx&info2=xx

{

change:succeed/failed  //成功或者失败

data:{

...

}//如果有必要的话

}

 

6、获取该设备历史数据：

GET/history?id=xx&device_id=xx

{

data{

...

}

}//传回一个数组

 

 