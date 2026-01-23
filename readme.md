
# API GO


Simple Authentication (Login, Register) Using JTW create token, refresh

Simple CRUD(Get user, Update user, Delete user by status) using Gorm


## Deployment

To deploy this project run

```bash
  Create database belajar_go

  fill .env like a sample
    JWT_SECRET=
    JWT_REFRESH_SECRET=
    JWT_ACCESS_EXPIRE_MINUTES=1400
    JWT_REFRESH_EXPIRE_DAYS=7


    DATABASE_USER=root
    DATABASE_PASSWORD=
    DATABASE_ADDRESS=127.0.0.1:3306
    DATABASE_NAME=belajar_go

  Migrate database 
  migrate -database "mysql://root:password@tcp(localhost:3306)/belajar_go" -path migrations up

  after database migrating you can run 
  go run main.go
```


## Endpoint, You can test by Bruno or Postman


To run tests, run the following command

```bash
  {
  "name": "Belajar GO",
  "version": "1",
  "items": [
    {
      "type": "http",
      "name": "Register",
      "filename": "Register.bru",
      "seq": 2,
      "settings": {
        "encodeUrl": false,
        "timeout": 0
      },
      "tags": [],
      "examples": [],
      "request": {
        "url": "http://localhost:8080/auth/register",
        "method": "POST",
        "headers": [
          {
            "name": "Content-Type",
            "value": "application/json",
            "enabled": true
          }
        ],
        "params": [],
        "body": {
          "mode": "json",
          "json": "{\n  \"nama\": \"Dicky\",\n  \"email\": \"dickydraknes@gmail.com\",\n  \"password\":\"admin\"\n}",
          "formUrlEncoded": [],
          "multipartForm": [],
          "file": []
        },
        "script": {},
        "vars": {},
        "assertions": [],
        "tests": "",
        "docs": "",
        "auth": {
          "mode": "inherit"
        }
      }
    },
    {
      "type": "http",
      "name": "Get user",
      "filename": "Get user.bru",
      "seq": 5,
      "settings": {
        "encodeUrl": false,
        "timeout": 0
      },
      "tags": [],
      "examples": [],
      "request": {
        "url": "http://localhost:8080/user",
        "method": "GET",
        "headers": [
          {
            "name": "Content-Type",
            "value": "application/json",
            "enabled": true
          }
        ],
        "params": [],
        "body": {
          "mode": "json",
          "json": "",
          "formUrlEncoded": [],
          "multipartForm": [],
          "file": []
        },
        "script": {},
        "vars": {},
        "assertions": [],
        "tests": "",
        "docs": "",
        "auth": {
          "mode": "inherit"
        }
      }
    },
    {
      "type": "http",
      "name": "Update User",
      "filename": "Update User.bru",
      "seq": 6,
      "settings": {
        "encodeUrl": false,
        "timeout": 0
      },
      "tags": [],
      "examples": [],
      "request": {
        "url": "http://localhost:8080/user/1",
        "method": "PUT",
        "headers": [
          {
            "name": "Content-Type",
            "value": "application/json",
            "enabled": true
          }
        ],
        "params": [],
        "body": {
          "mode": "json",
          "json": "{\n  \"nama\": \"Dickys\",\n  \"email\": \"dickydraknes@gmail.coms\"\n}",
          "formUrlEncoded": [],
          "multipartForm": [],
          "file": []
        },
        "script": {},
        "vars": {},
        "assertions": [],
        "tests": "",
        "docs": "",
        "auth": {
          "mode": "inherit"
        }
      }
    },
    {
      "type": "http",
      "name": "Delete user",
      "filename": "Delete user.bru",
      "seq": 7,
      "settings": {
        "encodeUrl": false,
        "timeout": 0
      },
      "tags": [],
      "examples": [],
      "request": {
        "url": "http://localhost:8080/user/2",
        "method": "DELETE",
        "headers": [
          {
            "name": "Content-Type",
            "value": "application/json",
            "enabled": true
          }
        ],
        "params": [],
        "body": {
          "mode": "json",
          "json": "",
          "formUrlEncoded": [],
          "multipartForm": [],
          "file": []
        },
        "script": {},
        "vars": {},
        "assertions": [],
        "tests": "",
        "docs": "",
        "auth": {
          "mode": "inherit"
        }
      }
    },
    {
      "type": "http",
      "name": "Login",
      "filename": "Login.bru",
      "seq": 3,
      "settings": {
        "encodeUrl": false,
        "timeout": 0
      },
      "tags": [],
      "examples": [],
      "request": {
        "url": "http://localhost:8080/auth/login",
        "method": "POST",
        "headers": [
          {
            "name": "Content-Type",
            "value": "application/json",
            "enabled": true
          }
        ],
        "params": [],
        "body": {
          "mode": "json",
          "json": "{\n  \"email\": \"dickydraknes@gmail.com\",\n  \"password\":\"admin\"\n}",
          "formUrlEncoded": [],
          "multipartForm": [],
          "file": []
        },
        "script": {},
        "vars": {},
        "assertions": [],
        "tests": "",
        "docs": "",
        "auth": {
          "mode": "inherit"
        }
      }
    },
    {
      "type": "http",
      "name": "Refresh token",
      "filename": "Refresh token.bru",
      "seq": 4,
      "settings": {
        "encodeUrl": false,
        "timeout": 0
      },
      "tags": [],
      "examples": [],
      "request": {
        "url": "http://localhost:8080/auth/login",
        "method": "POST",
        "headers": [
          {
            "name": "Content-Type",
            "value": "application/json",
            "enabled": true
          }
        ],
        "params": [],
        "body": {
          "mode": "json",
          "json": "{\n  \"refresh_token\": \"dickydraknes@gmail.com\"\n}",
          "formUrlEncoded": [],
          "multipartForm": [],
          "file": []
        },
        "script": {},
        "vars": {},
        "assertions": [],
        "tests": "",
        "docs": "",
        "auth": {
          "mode": "inherit"
        }
      }
    }
  ],
  "environments": [],
  "brunoConfig": {
    "version": "1",
    "name": "Belajar GO",
    "type": "collection",
    "ignore": [
      "node_modules",
      ".git"
    ],
    "size": 0.00026035308837890625,
    "filesCount": 1
  },
  "exportedAt": "2026-01-23T07:42:20.701Z",
  "exportedUsing": "Bruno/3.0.2"
}
```

