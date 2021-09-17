#  IOT logger API in Golang
Simple back-end to log data from IOT devices (using http)



***
## Api Endpoints

|Method    |  URL               | Function
|---           |---                  |---|
|POST      |/new_user     | Create user |
|POST      |/u/user_name              |login user|
|GET        |/u/user_name             | user Data, and devices|
|DELETE | /u/user_name             | delete user
| | |
|POST  |/u/user_name/device      |  create a new device, for user|
|GET| /u/user_name/device/:id     | get device data and sensors|
|PUT | /user/device/:id    | modify a device |
|Delete | /user/device/:id | delete existing sensor device|
|||
|POST| /user/device/:id/sensor|  create a new sensor for a device
|PUT |/user/device/:dev_id/sensor/:sensor_id|modify properties for sensor
|GET |/user/device/:id/sensor/:sensor_id|    get data only from this sensor
|POST|/u/user_id/dev/dev_id/sensor/:sensor_id| Log data for sensor
|||
|other sensor types|...|


***
### API calls examples with CURL

Test if API is working, and/or wake up

    curl http://127.0.0.1:3000/hello

Create new user

    curl -X POST -d 'Username=pedro&Email=juana@gmail.com&Pswd=123'  http://127.0.0.1:3000/new_user

Login existing user

    curl -X POST -d"username=pedro&password=123" http://127.0.0.1:3000/login
    {"data":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ",
    "message":"Success login","status":"success"}

Get user Data

    curl -X GET -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ" http://127.0.0.1:3000/u/pedro

Delete user

    curl -X DELETE -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ" http://127.0.0.1:3000/u/pedro


Create a Device for user

    curl -X POST  -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ" -d "Name=ejemploi&Description=otroejemplo"  http://127.0.0.1:3000/u/pedro/device




Get device

    curl -X GET -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ" http://127.0.0.1:3000/u/pedro/device/ejemploi

Edit device

    curl -X PUT -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ" -d 'newDescription=ejemplo_nuevo' http://127.0.0.1:3000/u/pedro/device/ejemploi

Delete device:

    curl -X DELETE -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ"   http://127.0.0.1:3000/u/pedro/device/ejemploi

Create sensor:

    curl -X POST http://127.0.0.1:3000/u/new_user/device/2/sensor


get Devices for user:

  curl -GET -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwOTUyMjksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoicGVkcm8ifQ.2Fo3ecGxih0ZMq-OeOWmsX29srPegok98HoBGMDi4vQ"  http://127.0.0.1:3000/u/pedro/device/

***
## SQL  Database Schemas

![schemas](docs/schemas.png)

***
## Development
Run using Air for automatic compilation and load of saved code, in the root directory just run:

    $ air

Build will be saved on **go-iot-loger-api/tmp**  (gitignored),
check **.air.toml** and **.airl.linux.conf** for implementation details


***
## TO DO  (coming soon)
* OAUTH2 + google sign up
* redis blacklist for JWT
* TLS https
* email notification to admin
* frontend
  * login
  * add device, sensor ...
  * dashboard to fetch sensor data
