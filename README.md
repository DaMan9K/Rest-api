# rest-api-client

# user-service

# REST API

GET /users -- list of users -- 200,404,500
GET /users/:id -- list by id -- 200,404,500
POST /users/:id -- create user -- 204,4xx,Header Location: url
PUT /users/:id -- fully update user -- 204/200,404,400,500
PATCH /users/:id -- partially update user -- 204/200,404,400,500
DELETE /users/:id -- remove user by id -- 204,404,400
