# 🚀 Gin Starter

This repository provides a **scalable starter** for building modular backend services in Go.
It includes best practices for **code structure, migrations, testing, and deployment**.

---

## 📂 Project Structure

See [Code Map](doc/CODE_MAP.md) for full details.

```
.
├── app/              # Routes
├── bin/              # Helper binaries (e.g., mock generator script)
├── cmd/              # main.go, and code generation executor
├── config/           # Configurations
├── db/
│   ├── migrations/   # Module-specific migrations
│   ├── schemas/      # Database schemas
│   └── seeders/      # Database seeders
├── deployment/       # Kubernetes manifests
├── doc/              # Documentation
├── middleware/       # Route middleware
├── entity/           # Domain entities
├── modules/          # Business logic modules
├── resource/         # API request/response structs
├── sdk/              # SDK integrations (pub/sub, gcs, etc.)
├── template/         # Templates (emails, etc.)
└── test/             # Unit tests, fixtures, mocks
```

---

## 🛠️ Getting Started

### Run Locally

See [How to Run](doc/HOW_TO_RUN.md) for details.

```bash
# Copy environment file
cp env.sample .env

# Generate EdDSA key
make key.generate

# Run the app
go run main.go
```

### Run with Docker

```bash
make tidy
make compile-server
make docker-build-server

docker run -d --env-file .env -p 8080:8080 --name gin-starter-container gin-backend-server:latest
```

---

## 🧩 Development Guide


* [Development Guide](doc/DEVELOPMENT_GUIDE.md)
* [Database Migration](doc/DATABASE_MIGRATION.md)
* [How to Add a Module](doc/HOW_TO_ADD_MODULE.md)
* [Create Mocks](doc/CREATE_MOCK.md)

---

## 📦 Deployment

See [Deployment](doc/DEPLOYMENT.md) (WIP).

---

## ❓ FAQs

See [FAQs](doc/FAQS.md).

---

## ✅ Contributing

1. Create a new branch:

   ```bash
   git checkout -b feature/your-change
   ```
2. Make your changes following the [development guide](doc/DEVELOPMENT_GUIDE.md).
3. Run tests:

   ```bash
   go test ./...
   ```
4. Commit with conventional message:

   ```bash
   git commit -m "feat: add user service"
   ```
5. Open a Pull Request.

---

## 📜 License

MIT License. Feel free to use and modify.
