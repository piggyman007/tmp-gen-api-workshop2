
# tmp-gen-api-workshop2

## API Documentation

### 1. Register User
**POST** `/register`

Request:
```
{
	"email": "user1@example.com",
	"password": "pass1",
	"firstname": "User1",
	"lastname": "One",
	"phone_number": "1111111111",
	"birthday": "1990-01-01"
}
```
Response:
```
true
```

### 2. Login
**POST** `/login`

Request:
```
{
	"email": "user1@example.com",
	"password": "pass1"
}
```
Response:
```
{
	"token": "<JWT_TOKEN>"
}
```

### 3. Get User Profile
**GET** `/me`

Headers:
```
Authorization: Bearer <JWT_TOKEN>
```
Response:
```
{
	"id": 1,
	"email": "user1@example.com",
	"password": "",
	"firstname": "User1",
	"lastname": "One",
	"phone_number": "1111111111",
	"birthday": "1990-01-01"
}
```

### 4. Transfer Points
**POST** `/transfer`

Headers:
```
Authorization: Bearer <JWT_TOKEN>
```
Request:
```
{
	"receiver_code": "user2@example.com",
	"points": 100
}
```
Response:
```
{
	"success": true
}
```


### 5. Get Point Histories
**GET** `/point-histories`

Headers:
```
Authorization: Bearer <JWT_TOKEN>
```
Response:
```
[
	{
		"from": "User1 One",
		"to": "User2 Two",
		"points": 100,
		"date": "2024-01-19",
		"sender_code": "user1@example.com"
	}
]
```

### 6. Hello World
**GET** `/`

Response:
```
{
	"message": "Hello World"
}
```
