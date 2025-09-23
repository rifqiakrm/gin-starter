# Frequently Asked Questions (FAQ)

### General

**Q: What is this project about?**
A: This project is a backend service structured with modules, supporting migrations, testing, and deployment. It provides a framework for scalable, maintainable development.

**Q: Where can I find the code structure overview?**
A: Refer to the [Code Map](CODE_MAP.md) which explains folders, packages, and their purposes.

**Q: What language and frameworks are used?**
A: The project is written in Go, with support tools for migrations, testing, and containerized deployment.

---

### Running the Application

**Q: How do I run the application?**
A: You can run it manually by creating an `.env` file, generating RSA keys, and running `go run main.go`. Alternatively, you can use Docker and Docker Compose. See [How to Run](HOW_TO_RUN.md).

**Q: Do I need Docker to run the project?**
A: No, you can run it manually, but Docker Compose is recommended for convenience.

**Q: How do I run the project with Docker?**
A: Use:

```bash
make tidy
make compile-server
make docker-build-server
docker-compose up
```

---

### Modules

**Q: What is a module?**
A: A module is a single business vertical focus, located in the `/modules` directory.

**Q: How do I add a new module?**
A: Create a new folder inside `/modules`, optionally with versioning (e.g., `/modules/example/v1`). Then add your code following the [Code Map](CODE_MAP.md). For details, see [How to Add Module](HOW_TO_ADD_MODULE.md).

**Q: Should modules be versioned?**
A: Yes, versioning (e.g., `v1`, `v2`) ensures backward compatibility and smoother upgrades.

---

### Database

**Q: How do I create a database migration?**
A: Use `make migration name=<migration-name> module=<module-name>`. Follow the workflow in [Database Migration](DATABASE_MIGRATION.md).

**Q: Why do I need to run UP, DOWN, and UP migrations?**
A: To ensure both migration directions work correctly. It validates that rollback and re-apply are successful.

**Q: What tool is used for migrations?**
A: [golang-migrate](https://github.com/golang-migrate/migrate).

**Q: How do I handle a dirty migration state?**
A: Run `make force-migrate url=<db_url> module=<module> version=<latest-clean-version>`.

**Q: Do I need a separate schema for each module?**
A: Yes, schemas must match the module name to ensure isolation and consistency.

---

### Development

**Q: How do I generate mocks for interfaces?**
A: Run `make mockgen`. The mocks will be generated in `/test/mock` following the original structure. See [Create Mock](CREATE_MOCK.md).

**Q: What is the recommended workflow for making changes?**
A: Clone → Migration → Entity → Module code → Update code/tests/docs → Format/Lint → Commit/MR. See [Making Changes](MAKING_CHANGES.md).

**Q: How do I ensure code quality?**
A: Run `make pretty` before committing. It tidies, formats, and lints code.

**Q: How do I run tests?**
A: Use `make test.unit` (unit tests), `make test.integration` (integration tests), or `make cover` (with coverage).

**Q: Where should I place unit tests?**
A: Unit test files live in the same directory as the files they test, following Go best practices.

---

### Making Changes

**Q: What is the safe workflow to introduce changes?**
A: Follow this order:

1. Create migration → `make migration`
2. Generate entity → `make entity`
3. Generate module code → `make module-all`
4. Update code, tests, and docs
5. Run quality checks → `make pretty`
6. Commit using Conventional Commits and open an MR

For full details, see [Making Changes](MAKING_CHANGES.md).

**Q: Can I generate only part of the module (handler, repository, service)?**
A: Yes, run `make handler`, `make repository`, or `make service` with the right flags.

**Q: How do I handle formatting and linting?**
A: Run `make format`, `make check.import`, and `make lint`. Or use `make pretty` for all-in-one.

**Q: How do I run tests before pushing changes?**
A: Run `make test.unit`, `make test.integration`, and `make cover` for coverage.

**Q: Do I need to update documentation after making changes?**
A: Yes, update docs like [Database Migration](DATABASE_MIGRATION.md) or [Create Mock](CREATE_MOCK.md) when relevant.

---

### Deployment

**Q: How do I deploy the application?**
A: Deployment documentation is currently TBD. For now, you may rely on Kubernetes manifests under the `deployment/` folder.

**Q: Does this project support Kubernetes?**
A: Yes, Kubernetes manifests are included in the `deployment/` folder.

---

### Utilities

**Q: How do I generate RSA keys for JWT?**
A: Run `make key.generate`.

**Q: How do I seed the database?**
A: Run `make seed`.

**Q: How do I validate migrations?**
A: Run `make validate-migration`.

**Q: Can I generate only specific components (e.g., handler, repository, service)?**
A: Yes, use `make handler`, `make repository`, or `make service` with appropriate flags.

---

### Git & Collaboration

**Q: What commit message convention should I follow?**
A: Use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

**Q: How do I open a merge request?**
A: Push your branch, open a Merge Request (MR), describe the goal, request reviews, and merge once approved and CI passes.

**Q: What is the branch naming convention?**
A: Typically `feature/<name>`, `fix/<name>`, or `chore/<name>`, aligned with commit type.

---

### Best Practices

**Q: How should I name migrations?**
A: Use descriptive names like `create_table_user` or `add_index_order`.

**Q: Should SQL queries always reference schemas?**
A: Yes, always prefix queries with schema (e.g., `auth.users`) to avoid ambiguity.

**Q: How do I keep dependencies clean?**
A: Run `make tidy` regularly and check for unused packages.

**Q: Should I write tests for every change?**
A: Yes, unit tests for logic, integration tests for module/db interactions.

---

### Troubleshooting

**Q: I get `dirty database` errors after migration. What should I do?**
A: Use `make force-migrate` or `make force-schema` with the latest clean version.

**Q: My Docker build fails with `no such file or directory` for the binary.**
A: Ensure `make compile-server` runs before `make docker-build-server`, and that your `cmd/server` binary exists.

**Q: I cannot connect to the database.**
A: Check your `.env` values, network access, and ensure Postgres is running.

**Q: CI pipeline fails due to lint errors.**
A: Run `make pretty` locally before committing.

**Q: Why are my tests failing locally but passing on CI?**
A: Ensure you’re using the same Go version and dependencies (`make tidy` before running tests).

---