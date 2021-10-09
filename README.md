
# Instagram API

In the project, MongoDB was used locally. All functions have been
implemented using standard go library

The password of a user was encrypted and stored in its encrypted form in the database, when a user makes a GET reuqest for a user, the password returned in JSON is encrypted.

The project is divied into 3 folders, the handler file implements all the API endpoint functions and has method to connect to database. The model file implements the data struture for the API and the encrypt file has functions to encrypt the password. main.go file starts the server on port 8000.

```
Attributes of Users :
Id
Name
Email
password

Attributes of posts :
Id
Caption
Image_URL
Timestamp
UserId
```
## API Reference

#### Create users 
```http
  POST /users
```

#### Get users By ID 

```http
  GET /users/{id}
```
#### Create posts 

```http
  POST /posts
```
#### Get post By ID 

```http
  GET /posts/{id}
```
#### Get all post by a user

```http
  GET /posts/users/{id}
```

## Screenshots

Create a User 
![CreateUser](https://user-images.githubusercontent.com/62784600/136660575-a5582a18-d045-4dc8-b015-2dc0784536a4.png)

Get a User by ID 
Password is encrypted
![GetUserByID](https://user-images.githubusercontent.com/62784600/136660597-86cce7cc-9f33-491f-a278-57edf6502e94.png)

Create a Post
![App Screenshot](https://imgur.com/Q2eCZHm)

Get a Post by ID
![App Screenshot](https://imgur.com/JiPslos)

Get all post made by a User
![App Screenshot](https://imgur.com/65pcWHk)
![App Screenshot](https://imgur.com/yiL177Q)
![App Screenshot](https://imgur.com/l8jOX9C)

Data stored in mongo
![App Screenshot](https://imgur.com/0wAUxnp)

