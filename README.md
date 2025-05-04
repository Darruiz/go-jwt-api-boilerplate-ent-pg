🚀 Go Vertical Slice Boilerplate (Golang + JWT + PostgreSQL + Docker)

This repository is a production-ready boilerplate for building modern, secure, and scalable APIs using Golang, designed with the Vertical Slice Architecture, JWT authentication, Docker, and PostgreSQL integration.

📦 Why This Exists

We built this boilerplate to serve as a universal starting point for any backend project that demands:
	•	Robust and modular structure
	•	Clean separation of concerns
	•	Modern and explicit authentication flow with Access & Refresh Tokens
	•	Scalability from day one
	•	Developer-friendly local setup with Docker Compose

🧱 Architectural Pattern: Vertical Slice

Instead of organizing code by layers (like controller/service/repository), Vertical Slice Architecture groups files by feature. This results in better encapsulation, reduces coupling, and makes onboarding faster for new developers.

✅ Benefits:
	•	Easier to test and reason about
	•	Each feature is isolated
	•	Scales naturally as your app grows
	•	Avoids god-structs and huge service files

🧠 ENT Framework?

We’re using entgo in the subscription and user domains to enable:
	•	Code-first schema definitions
	•	Powerful query building with type safety
	•	Automatic migrations (optional)

It wasn’t applied to auth since auth has no persistent user model yet – it’s pure token validation.

🔐 Authentication (JWT)

Implements industry-standard Access Token and Refresh Token flows:
	•	POST /auth/login — receives user_id and password, returns tokens
	•	POST /auth/refresh — refreshes tokens securely
	•	GET /auth/me — protected route using a middleware that extracts userId from the token and adds it to context

Tokens:
	•	Access Token: short-lived, for API usage (15min)
	•	Refresh Token: long-lived, for session renewal (7 days)

🐳 Dockerized from the start

Everything is Dockerized:
	•	Go API container
	•	PostgreSQL 15 container

Just run:

make docker-build-nc  # build without cache
make docker-up         # start services
make docker-logs       # tail logs

🗂 Folder Structure (Vertical Slice + Internal Boundaries)

internal/
├── app/              # setup & routing logic
├── config/           # DB connection, ENV vars
├── features/         # feature-first vertical slice domains
│   ├── auth/         # stateless token-based login
│   ├── subscription/ # ent-based business domain
│   └── user/         # user schema & services
├── middleware/       # custom middlewares (e.g., auth)
├── migrations/       # db migrations (if needed)
pkg/                  # shared utils (e.g., logger, hash)

🛠 Stack
	•	Go 1.24 (alpine based)
	•	PostgreSQL 15
	•	Docker + Compose
	•	entgo for ORM
	•	github.com/golang-jwt/jwt/v5 for JWT
	•	godotenv for env loading in local

✅ Features
	•	🔐 Auth with Access + Refresh JWT
	•	🧠 Vertical Slice structure
	•	🐳 Dockerized infra
	•	💾 PostgreSQL with GORM and/or ENT ready
	•	🧪 Ready for testability and CI setup

📌 Future Suggestions
	•	Add CI pipeline (GitHub Actions)
	•	Swagger/OpenAPI generator
	•	Add user registration + hashed password
	•	Integrate ent migrations
	•	Rate limiting & audit logs

⸻

This boilerplate aims to be a blueprint for clean backend architectures in Go that balances modern DX, scalability, and security. Feel free to fork and adapt to your needs. Contributions welcome!