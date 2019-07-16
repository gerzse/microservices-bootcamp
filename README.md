# microservices-bootcamp
Bootcamp on Microservices

## Study Material

* https://martinfowler.com/articles/microservices.html
* https://palletsprojects.com/p/flask/

## Main Reasons for Using Microservices

* Services are independently deployable. In a monolith application, a change to
any component requires a full redeploy of the entire application. In a Microservices
application only affected services must be redeployed (most of the times).
* More explicit component interface. HTTP interfaces define unbreakable boundaries
between services. In a monolith application with shared codebase it is difficult to
forbid over-tight coupling between components.

## Coding Session

Develop two microservices that can be used together for managing a list of users:
basic data like name, gender and year of birth, and profile pictures. One
microservice will manage the user data, the other one will manage the profile
pictures.

Write a very simple client that orchestrates the two microservices and generates 
some form of output where users and their photos are processed together.

You have some sample data in the `mock-data` directory.

You should keep the implementation of the microservices as simple as possible. For
example you can use a quick-and-dirty in-memory store (e.g. a map or dictionary),
don't bother connecting to databases. Focus on getting the REST interfaces right.

### Users Data Microservice

#### Create a new user: `POST /users`

Create a new user by posting its details to `/users`. You receive back a UUID that
was assigned to the new user:

```
POST /users
Content-Type: application/json
{"gender": "male", "name": {"last": "Mandvi", "first": "Aasif"}, "born": 1966}
```

```
HTTP/1.0 200 OK
Content-Type: application/json
{"uuid": "d0ad2480-a80d-11e9-b948-181dea663bfa"}
```

#### Retrieve all users: `GET /users`

You receive back a list with all existing users. There is no need for search and
filter functionality, just use a small number of users for testing.

```
GET /users
```

```
HTTP/1.0 200 OK
Content-Type: application/json
{
  "users": [
    {
      "born": 2019, 
      "gender": "female", 
      "id": "047d5e6e-a817-11e9-9db1-181dea663bfa", 
      "name": {
        "first": "Adele", 
        "last": "Mara"
      }
    }, 
    ...
    {
      "born": 1926, 
      "gender": "male", 
      "id": "0478555e-a817-11e9-a1f3-181dea663bfa", 
      "name": {
        "first": "Abdus", 
        "last": "Salam"
      }
    }
  ]
}
```

#### Update a User: `POST /user/<uuid>`

Send the new data for the user:

```
POST /user/047d5e6e-a817-11e9-9db1-181dea663bfa
Content-Type: application/json
{
  "born": 2019,
  "id": "047d5e6e-a817-11e9-9db1-181dea663bfa",
  "name": {
    "last": "Mara",
    "first": "Adele"
  },
  "gender": "female"
}
```

```
HTTP/1.0 200 OK
```

#### Retrieve a User by its UUID: `GET /user/<uuid>`

```
GET /user/047d5e6e-a817-11e9-9db1-181dea663bfa
```

```
HTTP/1.0 200 OK
Content-Type: application/json
{
  "born": 2019, 
  "gender": "female", 
  "id": "047d5e6e-a817-11e9-9db1-181dea663bfa", 
  "name": {
    "first": "Adele", 
    "last": "Mara"
  }
}
```

### Photos Microservice

#### Upload a profile picture: `POST /photo/<uuid>`

Send the binary data of the picture as the request payload.

```
POST /photo/890e0a9e-a811-11e9-a970-181dea663bfa
Content-Type: application/octet-stream
<binary image data>
```

```
HTTP/1.0 200 OK
```

#### Retrieve a profile picture by the UUID of the user: `GET /photo/<uuid>`

```
GET /photo/bd64c15e-a814-11e9-8bbe-181dea663bfa
```

```
HTTP/1.0 200 OK
Content-Type: application/octet-stream
<binary image data>
```