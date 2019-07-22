# King's Thursday Bootcamp on Microservices

The subject matter of this bootcamp are microservices. We will look at some basic
principles of microservices, and then we will build from scratch two microservices
and a client that consumes them.

## Agenda

- Theoretical overview of microservices (first 20 minutes or so)
- Coding session (rest of the time, up to two hours, or longer if desired)

## Prerequisites

For the coding session, depending on what technology you plan to use, please prepare
your environment (IDE, JDK, frameworks, debuggers, etc.) before the bootcamp. We don't
want to spend time during the workshop installing Java or Python.

Some tools and frameworks that can be used for building microservices:

* Python/Flask: https://palletsprojects.com/p/flask/
* Java/Spring Boot: https://spring.io/guides/gs/rest-service/
* Java/JAX-RS (GlassFish Server):  https://docs.oracle.com/javaee/7/tutorial/jaxrs.htm#GIEPU
* NodeJS/Express: https://expressjs.com/en/starter/hello-world.html
* .NET: https://docs.microsoft.com/en-us/aspnet/core/tutorials/first-web-api

There are many options available. Choose the one that you feel comfortable with. If
you feel like exploring, choose two different technologies, given that we will build
two microservices.

## Coding Session

Develop two microservices that can be used together for managing a list of users:
basic data like name, gender and year of birth, and profile pictures. One
microservice will manage the basic data, the other one will manage the profile
pictures. Use the programming language(s) of your choice for building the two microservices.
Feel free to even use different languages for the two.

Write a very simple client that orchestrates the two microservices and generates 
some form of output where users and their photos are processed together. Again, use the language
of your choice to build this minimalist client.

You have some sample data in the `samples/mock-data` directory.

You should keep the implementation of the microservices as simple as possible. For
example you can use an in-memory store (e.g. a map or dictionary), don't bother
storing things in a database. Focus on getting the REST interfaces right.

If everything works out well, feel free to do one more step. In your client implement
a transactional update of a user. In one transaction you should POST data to both 
microservices, and if the second call fails for a reason you make sure that no garbage
remains, which means you cleanup the data you sent to the first microservice. This 
might require adding some DELETE endpoints to the microservices. To test, you can 
shutdown one of the microservices when you run the client.

The specification for the REST API of the two microservices follows.

In the `samples/python` directory you can find a sample implementation of the two 
microservices in Python. It is a trivial implementation, basically without error
checking. In `samples/client.py` you can find a sample implementation of a client,
again in Python.

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
