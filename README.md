# USERS
### 1) userGroup.Post("/", controllers.CreateUser)
```console
POST http://localhost:3000/users
{
    "name": "Naushad",
    "email": "naushad@example.com",
    "balance": 100
}

```

### 2) userGroup.Get("/", controllers.GetAllUsers)
```console
GET http://localhost:8080/users
```

# POSTS
### 1) postRoutes.POST("/", controllers.CreatePost)
```console
POST http://localhost:3000/posts
{
    "title": "My First Post",
    "content": "This is the content of the post",
    "user_id": 2
}
```

### 2)  postRoutes.GET("/", controllers.GetAllPosts)
```console
GET http://localhost:3000/posts
```


