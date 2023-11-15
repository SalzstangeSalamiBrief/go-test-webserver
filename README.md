# Test GO Webserver

## Routes

| method | route          | description      | body                                                        |
| ------ | -------------- | ---------------- | ----------------------------------------------------------- |
| GET    | /users         | get all users    |                                                             |
| GET    | /users/:int    | get a user by id |                                                             |
| DELETE | /users/:ind    | delete a user    |                                                             |
| POST   | /users         | create a user    | "FirstName": string, LastName:string, Age:Number            |
| PUT    | /users         | update a user    | "Id": int, "FirstName": string, LastName:string, Age:Number |
| GET    | /files/:string | get a file       |                                                             |
| POST   | /files         | upload a file    | form-data with the field "file"                             |
