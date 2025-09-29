# fiberplate

**fiberplate** is a clean, scalable boilerplate for building RESTful APIs in Go using [Fiber v3](https://github.com/gofiber/fiber), [sqlx](https://github.com/jmoiron/sqlx) for database access, and modular architecture (handler, service, repository, middleware, responder, etc).

---

## 📁 Project Structure

```
fiberplate/
├── internal/
│   ├── app/
│   │   ├── router.go
│   │   ├── user_route.go
│   │   └── account_route.go
│   ├── handler/
│   │   └── rest/
│   │       └── v1/
│   │           ├── user/
│   │           │   ├── handler.go
│   │           │   ├── request.go
│   │           │   └── response.go
│   │           ├── account/
│   │           │   ├── handler.go
│   │           │   ├── request.go
│   │           │   └── response.go
│   │           └── ...
│   ├── service/
│   │   ├── user_service.go
│   │   └── account_service.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   └── account_repository.go
│   ├── entity/
│   │   ├── user.go
│   │   └── account.go
│   └── middleware/
│       ├── requestid.go
│       └── recover.go
│
├── pkg/
│   ├── responder/
│   │   ├── responder.go
│   │   └── error.go
│   └── logger/
│       └── logger.go
│
├── go.mod
├── README.md
└── ...
```

---

## 🚀 Quick Start

1. **Clone the repo:**
    ```bash
    git clone https://github.com/n0l3r/fiberplate.git
    cd fiberplate
    ```

2. **Install dependencies:**
    ```bash
    go mod tidy
    ```

3. **Run the app:**
    ```bash
    go run main.go
    ```

4. **Test endpoints:**
    ```
    GET    /ping
    ```

---

## 🧑‍💻 How to Extend

- **Add new entity**: Create repository, service, handler, route file for your feature.
- **Add new middleware**: Place in `internal/middleware`, register in Fiber app.
- **Change DB**: Switch sqlx driver or config.
- **Customize response**: Edit or extend `pkg/responder`.

---

## 📝 License

MIT

---

## 💡 Credits

- [Fiber](https://github.com/gofiber/fiber)
- [sqlx](https://github.com/jmoiron/sqlx)
- [zerolog](https://github.com/rs/zerolog)