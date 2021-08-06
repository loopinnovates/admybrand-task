# ADmyBRAND-task for Backend Engineer

>This task is given to check my understanding and skill set of Golang.


### Technologies Used
**Backend**: Go\
**Database**: MongoDB

## Steps to follow for task execution

1. Clone this repo using following command in your Go workplace:

```bash
git clone https://github.com/rajput9piyush/admybrand-task.git
```
2. Now, Go to "admybrand-task"
```bash
cd admybrand-task
```
3. Create Go module
```go
go mod init admybrand
```
4. Get dependencies using "**go get**"
* Below dependency is required to use **Gin Web framework**
```go
go get github.com/gin-gonic/gin
```
* Below dependencies are required to work with **MongoDB**
```go
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/bson/primitive
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```
* Below dependency is used for **assertion** 
```go
go get github.com/stretchr/testify/assert
``` 
5. Building the Project
```go
go build
```
6. Run Project
```bash
./admybrand
```
## APIs Info
> This application is working on port number _"8080"_ on _localhost_\
> And MongoDB is running on port number _"27017"_ also on _localhost_

**1. GET /users**
> This api is used to return information of all the users present in the system.

###### Example: 
> URL: _"http://localhost:8080/users"_

**2. GET /user/:ID**
> This api is used to return information of specific user whose ID is given.
###### Example:
> URL: _"http://localhost:8080/user/610c959fc2937d8ae49d12a5"_ \
> Here, "610c959fc2937d8ae49d12a5" is an user id.

**3. POST /adduser**
> This api is used to create an user in the system.\
> To execute this api, following conditions should be meet:
> 1. "Content-Type": "application/json"
> 2. Following field are required:
>    * name
>    * Dob (yyyy-mm-dd)
>    * Address
>    * Description
###### Example:
```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "dummy", "Dob": "2000-01-30", "Address": "dummy Address", "Description": "Dummy Description"}' \
    http://localhost:8080/adduser
```

**4. PUT /updateuser**
> This api is used to modify information of user, present in the system.\
> To execute this api, certain conditions to be meet and certain data should be passed:
> 1. "Content-Type": "application/json"
> 2. Following field is required:
>    * id
> 3. User can change following fields:
>    * name
>    * Dob (yyyy-mm-dd)
>    * Address
>    * Description
###### Example:
```bash
curl -X PUT -H "Content-Type: application/json" \
    -d '{"id":"610c959fc2937d8ae49d12a5", "name": "dummyUpdated", "Dob": "2000-01-30", "Address": "dummy Address", "Description": "Dummy Description"}' \
    http://localhost:8080/updateuser
```




**5. DELETE /deleteuser/:ID**
> This api will delete an user from system using it's ID
###### Example
```bash
curl -X DELETE http://localhost:8080/deleteuser/610c959fc2937d8ae49d12a5
```

# Testing
To run tests use following syntax:
```go
go test -run NameOfTest
```
**NameOfTest** can be replaced by:
  1. TestGetAllUsers
  2. TestGetUserByID
  3. TestCreateUser
  4. TestDeleteAnUser
  5. TestUpdateUser
#
That's all from my side, if you think we can grow together, then click on below links to connect with me
## __________
[LinkedIn](https://www.linkedin.com/in/iampkr/) |  [Email](mailto:rajput9piyush@gmail.com) | [Phone](tel:+918347365590)