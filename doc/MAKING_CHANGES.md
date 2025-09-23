# Making Changes

This document explains how to safely make changes in the project and use the provided tools.

---

## 🚀 Workflow Overview

1. Clone the project

   ```bash
   git clone <repo-url>
   cd <repo>
   ```
2. Create a migration → generate entity → generate module code.
3. Update code, tests, and documentation as needed.
4. Format, lint, and run tests before committing.
5. Push changes and open a Merge Request (MR).

---

## 🔨 Code Generation Workflow

When creating new functionality, **follow this order**:

1. **Create migration** for your module

   ```bash
   make migration name=<migration_name> module=<module>
   ```
2. **Create entity** based on your migration

   ```bash
   make entity table=<table_name> module=<module>
   ```
3. **Generate module code** (handler, repository, service)

   ```bash
   make module-all module=<name> version=<version> entity=<entity>
   ```

---

## 🔨 Code Generation (Individual)

* **Generate only a handler:**

  ```bash
  make handler module=<name> version=<version> entity=<entity>
  ```
* **Generate only a repository:**

  ```bash
  make repository module=<name> version=<version> entity=<entity>
  ```
* **Generate only a service:**

  ```bash
  make service module=<name> version=<version> entity=<entity>
  ```

---

## 🧹 Code Quality

* **Format code & clean imports:**

  ```bash
  make format
  make check.import
  ```
* **Run linter:**

  ```bash
  make lint
  ```
* **Tidy dependencies:**

  ```bash
  make tidy
  ```
* **Run full pre-push check (tidy + format + lint):**

  ```bash
  make pretty
  ```
* Always run `make pretty` **before committing**, as this will also be checked in CI pipelines.

---

## ✅ Testing

* **Run unit tests:**

  ```bash
  make test.unit
  ```
* **Run integration tests:**

  ```bash
  make test.integration
  ```
* **Run unit tests with coverage:**

  ```bash
  make cover
  ```
* **View coverage report in HTML:**

  ```bash
  make coverhtml
  ```

---

## 🗄️ Database

* **Create new schema migration:**

  ```bash
  make schema name=<schema_name>
  ```

* **Apply schema migration:**

  ```bash
  make migrate-schema url=<db_url>
  ```

* **Rollback schema migration:**

  ```bash
  make rollback-schema url=<db_url>
  ```

* **Create new module migration:**

  ```bash
  make migration name=<migration_name> module=<module>
  ```

* **Apply migration:**

  ```bash
  make migrate url=<db_url> module=<module>
  ```

* **Rollback last migration:**

  ```bash
  make rollback url=<db_url> module=<module>
  ```

* **Rollback all migrations:**

  ```bash
  make rollback-all url=<db_url> module=<module>
  ```

* **Force set migration version (recover from dirty state):**

  ```bash
  make force-migrate url=<db_url> module=<module> version=<version>
  ```

* **Seed database:**

  ```bash
  make seed
  ```

📌 **Important:** Always include the module's schema in SQL queries. Example:

```sql
-- ❌ Avoid
SELECT * FROM users;

-- ✅ Correct
SELECT * FROM auth.users;
```

Apply this to **all SQL operations** (SELECT, INSERT, UPDATE, DELETE, etc.).

---

## 🔑 Utilities

* **Generate EdDSA key for JWT:**

  ```bash
  make key.generate
  ```
* **Generate mocks for interfaces:**

  ```bash
  make mockgen
  ```
* **Validate migrations:**

  ```bash
  make validate-migration
  ```

---

## 📄 Documentation

* If your changes require any document to be updated, please update the relevant docs (e.g., [Database Migration](DATABASE_MIGRATION.md), [Create Mock](CREATE_MOCK.md)).

---

## 💾 Git & MR Process

1. Stage and commit your changes:

   ```bash
   git add .
   git commit -s -m "feat: add new user module"
   git push origin <your-branch>
   ```

    * Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for messages.

2. Open a **Merge Request (MR)**.

    * Clearly describe the goal of the MR and the changes introduced.

3. Ask contributors to review your MR.

4. Once approved and CI pipeline is green → **merge your MR**.
