## Code Map

Code Map explains how the codes are structured in this project. This document will list down all folders or packages and their purpose.

---

### `app`

This folder any routes that available in the project.

---

### `bin`

This folder contains any executable binary to support the project.
For example, there is `generate-mock.sh`. It is a shell script file used to generate mocks for all interfaces available in this project.

---

### `cmd`

This folder contains the **main entry points** and **command-line tools** for the project.

- `cmd/main.go` → the main application entry point for running the backend server.
- `cmd/entitygen` → CLI tool to generate entity structs from database tables.
- `cmd/modulegen` → CLI tool to generate boilerplate code for modules (handler, repository, service).

Each subfolder under `cmd/` can be built into a separate binary.

---

### `config`

This folder contains configuration .env

---

### `db/migrations`

This folder contains all database migration files. It has many subdirectories. Each subdirectory represents a single module.

---

### `db/migrations/<module-name>`

This folder contains all module's migration files. Each migration has exactly two files: UP and DOWN. 

---

### `db/schemas`

This folder contains all database schema migration files.

---

### `db/seeders`

This folder contains all needed seeder for the database.

---

### `deployment`

This folder contains Kubernetes manifest related to deployment.

---

### `doc`

This folder contains all documents related to the project.

---

### `middleware`

This folder contains routes middleware.

---

### `entity`

This folder contains the domain of the module.
Mostly, this folder contains only structs, constants, global variables, enumerations, or functions with simple logic related to the core domain of the module (not a business logic).

---

### `resource`

This folder contains all structs related to API requests.

---

### `modules`

This folder contains all vertical businesses logic.

---

### `modules/<module-name>/<module-version>/service`

This folder contains the main business logic of the module. Almost all interfaces and all the business logic flows are defined here.
If someone wants to know the flow of the module, they better start to open this folder.

---

### `modules/<module-name>/<module-version>`

All APIs/codes in the folder (and all if its subfolders) are designed to [not be able to be imported](https://golang.org/doc/go1.4packages).
This folder contains all detail implementation specified in the `modules/<module-name>/<module-version>/service` folder.

---

### `modules/<module-name>/<module-version>/builder`

This folder contains the [builder design pattern](https://sourcemaking.com/design_patterns/builder).
It composes all codes needed to build a full usecase.

---

### `modules/<module-name>/<module-version>/handler`

This folder contains the HTTP/2 gRPC handlers.
Codes in this folder implement gRPC server interface.

---

### `modules/<module-name>/<module-version>/repository`

This folder contains codes that connect to the repository, such as database.
Repository is not limited to databases. Anything that can be a repository can be put here.

---

### `sdk`

This folder contains related SDK needed for the project such as pub/sub and gcs.

---

### `template`

This folder contains template or related stuff.
For example, you can put email template here in this folder.

---

### `test`

This folder contains test related stuffs.
For the case of unit test, the unit test files are put in the same directory as the files that they test. It is one of the Go best practice, so we follow.

---

### `test/fixture`

This folder contains a well defined support for test.

---

### `test/mock`

This folder contains mock for testing.

---