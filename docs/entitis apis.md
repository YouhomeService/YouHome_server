POST /v1/users

{

​	userCode

​	userName

}

{

​	userID

}

 

POST /v1/scenes

{

​	userID

​	sceneName

}

{

​	sceneID

}

 

GET /v1/scenes/{userID}

{

​	[]sceneID

}

 

DELETE /v1/scenes/{userID}

 

POST /v1/devices

{

​	userID

​	sceneID

​	deviceType

}

{

​	[

​		{

​		entityID

​		}

​	]

}

 

GET /v1/devices     (maybe you can/v1/device?scenesID="XXX")

{

​	[

​		{

​			deviceID

​			deviceName

​			entityID

​		}

​	]

}

 

GET /v1/devices/{devideID}/states

{

​	deviceName

​	 lightSwitch (just for example)

}

 

POST /v1/devices/{deviceID}/states

{

​	lightSwitch 

}

{

​	lightSwitch

}

 

DELETE /v1/devices/{deviceID}

 

 