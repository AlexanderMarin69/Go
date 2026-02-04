# Go Commands Quick Reference

## Dependency Management

### `go mod download`
Downloads all dependencies listed in `go.mod` to local cache.
```bash
go mod download
```
**Use**: Before building in CI/CD or fresh clone.

### `go mod tidy`
Removes unused dependencies and adds missing ones. Updates `go.mod` and `go.sum`.
```bash
go mod tidy
```
**Use**: After adding/removing imports or when dependencies are messy. **Do this before commits.**

### `go get <package>`
Add or update a specific dependency.
```bash
go get github.com/go-chi/chi/v5
go get -u github.com/go-chi/chi/v5  # Update to latest
```
**Use**: When adding new libraries to your project.

## Running & Building

### `go run main.go`
Compile and run directly (doesn't create binary).
```bash
go run main.go
```
**Use**: Development, quick testing. Code changes require re-running.

### `go build`
Compile into binary (executable).
```bash
go build                    # Creates ./gowebapp (macOS/Linux)
go build -o myapp          # Custom output name
```
**Use**: Creating deployable artifacts, production releases.

### `go build -o dist/app && ./dist/app`
Build and run binary.
```bash
go build -o gowebapp && ./gowebapp
```
**Use**: Testing compiled version before deployment.

## Best Practices

| Task | Command | Why |
|------|---------|-----|
| **Fresh setup** | `go mod download` | Ensures clean dependency state |
| **Before commit** | `go mod tidy` | Keeps go.mod/go.sum clean |
| **Development** | `go run main.go` | Fast iteration, no compile wait |
| **Deployment** | `go build` | Single executable, fast startup |
| **Add dependency** | `go get package` | Updates go.mod automatically |
| **Clean build** | `go clean` | Removes build cache |

## Workflow Example

```bash
# 1. Clone or fresh start
go mod download

# 2. Development loop
go run main.go

# 3. Before committing
go mod tidy
git add go.mod go.sum
git commit -m "Update dependencies"

# 4. Before deployment
go build
./gowebapp  # Run the binary
```

## Useful Flags

- `go run -race main.go` - Detect race conditions
- `go build -v` - Verbose output
- `go mod verify` - Verify module integrity
- `go list -u -m all` - List available updates

---

**Remember**: `go mod tidy` before every commit to keep dependencies clean!
