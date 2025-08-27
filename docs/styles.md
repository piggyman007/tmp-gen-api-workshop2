# Code Style Guide (System Prompt)

Use this guide as a system prompt to keep contributions consistent across the codebase.

## Language and Formatting
- Go 1.x, idiomatic Go.
- Always run `gofmt` and `go vet` before committing.
- Import groups: standard lib, third-party, internal (separated by blank lines).
- Keep functions small and cohesive; prefer early returns.
- No unused imports/variables; avoid overusing `interface{}`.

## Project Structure
- `handler/`: HTTP layer only (parse input, auth, map errors to HTTP, shape response).
- `service/`: Business logic with receiver-based services (e.g., `UserService`, `TransferService`). No framework imports here.
- `db/`: Persistence layer with receiver type (e.g., `TransferDB`) + interface (`TransferDBInterface`). Provide `mockdb.go` for tests.
- `model/`: Plain structs with GORM and JSON tags.
- `route/`: Wire dependencies (services, middleware) and register routes.

## Naming and JSON
- Exported: PascalCase; unexported: camelCase.
- JSON fields: snake_case via struct tags (e.g., `json:"receiver_code"`).
- Packages and files: lowercase, concise names.

## HTTP/Fiber Conventions
- Handler signature: `func(c *fiber.Ctx) error`.
- Parse body with `c.BodyParser(&req)` and validate required fields.
- Auth: read `user_id` from `c.Locals("user_id")` (set by middleware).
- Error mapping:
  - 400 invalid payload/validation.
  - 401 auth failure.
  - 404 not found.
  - 500 unexpected errors.
- Error JSON shape: `{ "error": "<message>" }`.

## Services (Business Logic)
- Receivers: `func (s *UserService) ...` with constructor (e.g., `NewUserService(db, jwtSecret)`).
- No HTTP or Fiber types inside services.
- Return domain values and errors; use sentinel errors (e.g., `gorm.ErrRecordNotFound`).
- Define domain-level errors like `ErrMissingFields`, `ErrInvalidCredentials`.

## DB Layer
- Receiver-based implementation: `type TransferDB struct { DB *gorm.DB }`.
- Interface: `TransferDBInterface` with methods:
  - `FindUserByCode(code interface{}) (model.User, error)`
  - `FindUserByID(id int) (model.User, error)`
  - `CreateTransfer(transfer *model.Transfer) error`
  - `FindTransfersByUser(userID int, limit int) []model.Transfer`
- Constructor returns interface: `NewTransferDB(*gorm.DB) TransferDBInterface`.
- Use GORM query builders; return `gorm.ErrRecordNotFound` when appropriate.

## Models
- Include GORM primary key/autoincrement tags as needed.
- JSON tags in snake_case.
- Example: `Transfer` includes `SenderID`, `ReceiverID`, `SenderCode`, `ReceiverCode`, `Points`, `CreatedAt (YYYY-MM-DD string)`.

## Errors
- Propagate and wrap as needed; handlers decide HTTP status.
- Do not expose raw DB errors to clients.

## Testing
- Services depend on `db.TransferDBInterface`.
- Use `db.NewMockDB()` in tests; seed `Users` and assert `Transfers`.
- Test happy paths and edge cases (not found, invalid input).

## Dates and Limits
- Keep `CreatedAt` string format `2006-01-02` unless a migration is planned.
- Default list limits to 10 at handler level; allow service to accept `limit`.

## Commits and Docs
- Conventional commits: `feat:`, `fix:`, `refactor:`, `docs:`, `chore:`.
- Update `README.md`/`docs/` when API shapes change (e.g., `sender_code`).
