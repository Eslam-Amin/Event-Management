# Event Booking

Simple event booking REST API built with Go and Gin.

**Project Summary**: This repository implements a small event booking backend with user signup/login (JWT), event CRUD, and event registration using a SQLite database.

**Quick Start**

- Prerequisites: Go 1.25+ installed.
- Run locally:

```powershell
go run main.go
```

The server listens on `:8080` and will create a local SQLite file `api.db` in the project root.

**Technologies Used**

- **Language**: Go
- **Web framework**: `github.com/gin-gonic/gin`
- **DB**: SQLite via `modernc.org/sqlite`
- **Auth**: JWT using `github.com/golang-jwt/jwt/v5`
- **Password hashing**: `golang.org/x/crypto`

**File Structure**

- `main.go` : application entrypoint, registers routes and starts the server
- `go.mod` : Go module and dependencies
- `db/` : database initialization and table creation (`db.go`)
- `controllers/` : request handlers (`auth.go`, `events.go`, `users.go`)
- `dtos/` : request DTOs (`auth.go`)
- `middlewares/` : authentication/authorization middleware (`auth.go`)
- `models/` : domain models and DB interactions (`event.go`, `user.go`)
- `routes/` : route registration (`routes.go`)
- `utils/` : helpers for hashing and JWT (`hash.go`, `jwt.go`)

**API Endpoints**

- Health

  - `GET /ping` : returns `{ "message": "pong" }`

- Auth

  - `POST /auth/signup` : Create a new user.
    - Body: `{ "name": "Name", "email": "you@example.com", "password": "secret" }`
    - Response: created user (without password)
  - `POST /auth/login` : Login and receive a JWT token.
    - Body: `{ "email": "you@example.com", "password": "secret" }`
    - Response: `{ "data": { "token": "<JWT>", "user": { ... } } }`

- Events

  - `GET /events/` : Get all events (public)
  - `GET /events/:id` : Get a single event by id (public)
  - `POST /events/` : Create event (protected — requires auth)
    - Body: `{ "name": "Event", "description": "...", "location": "..", "event_date": "2025-12-31T19:00:00Z" }`
  - `PUT /events/:id` : Update event (protected — must be owner)
  - `DELETE /events/:id` : Delete event (protected — must be owner)
  - `POST /events/:id/register` : Register the authenticated user for the event (protected)
  - `DELETE /events/:id/register` : Cancel registration (protected)

- Users
  - `GET /users/` : Get all users (public)
  - `GET /users/:id` : Get user by id (public)
  - `GET /users/me/registrations` : Get current user's event registrations (protected)

**Authentication**

- The project uses JWT for authentication. The login endpoint returns a token.
- Protected endpoints expect the token in the `Authorization` header as the raw token value (for example: `Authorization: <TOKEN>`). The middleware reads `Authorization` header and validates the token.

**Database**

- SQLite DB file: `api.db` created automatically in the project root by `db.InitDB()`.
- Tables created: `users`, `events`, `event_registrations`.

**Examples (curl / PowerShell)**

- Signup

```powershell
curl -X POST http://localhost:8080/auth/signup -H "Content-Type: application/json" -d '{"name":"John","email":"john@example.com","password":"secret"}'
```

- Login

```powershell
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d '{"email":"john@example.com","password":"secret"}'
```

- Create Event (replace `<TOKEN>` with the token returned from login)

```powershell
curl -X POST http://localhost:8080/events/ -H "Content-Type: application/json" -H "Authorization: <TOKEN>" -d '{"name":"Party","description":"End of year","location":"City Hall","event_date":"2025-12-31T19:00:00Z"}'
```

**Next Steps**

- The JWT `secretKey` is set in `utils/jwt.go` as a constant `supersecretkey` — for production change it to a securely loaded secret.
- You can change the DB path in `db/db.go` if you want the database stored elsewhere.
- Consider adding request validation, pagination for lists, and tests.
