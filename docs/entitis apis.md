POST  /v1/users{userCodeuserName}{userID}
POST /v1/scenes{userIDsceneName}{sceneID}
GET /v1/scenes/{userID}{[]sceneID}
DELETE /v1/scenes/{userID}
POST /v1/devices{userIDsceneIDdeviceType}{[{entityID}]}
GET /v1/devices      (maybe you can /v1/device?scenesID="XXX"){[{deviceIDdeviceNameentityID}]}
GET /v1/devices/{devideID}/states{deviceName lightSwitch (just for example)}
POST /v1/devices/{deviceID}/states{lightSwitch }{lightSwitch}
DELETE /v1/devices/{deviceID}