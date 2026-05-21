# DevArena Architecture

DevArena is an educational backend project designed to learn Go and backend development through practical implementation.

The project starts as a console battle simulator and gradually evolves into a backend system with REST API, PostgreSQL, Redis, MongoDB, gRPC, RabbitMQ, Kafka, Docker, Kubernetes and CI/CD.

---

## Development principle

The project is developed as a learning map.

Every topic must go through this cycle:

1. Learn the concept
2. Apply it in DevArena
3. Explain why it is implemented this way
4. Commit the change
5. Mark the topic in LEARNING_CHECKLIST.md

The goal is not to build fast. The goal is to understand every layer.

---

## Evolution stages

```text
Console Go application
→ Go packages
→ In-memory REST API
→ Layered backend architecture
→ PostgreSQL persistence
→ Docker Compose environment
→ Redis mechanics
→ Tests
→ RabbitMQ workers
→ Kafka event streaming
→ MongoDB logs
→ gRPC internal services
→ Kubernetes deployment
→ CI/CD automation
```

---

## Initial structure

```text
devarena/
  cmd/
    arena/
      main.go
  README.md
  ARCHITECTURE.md
  LEARNING_CHECKLIST.md
  .gitignore
  go.mod
```

---

## Main domain entities

```text
User
Hero
HeroClass
Enemy
Battle
BattleRound
Item
Inventory
Reward
Rating
Tournament
BattleLog
AuditLog
GameEvent
```

---

## Future backend structure

```text
devarena/
  cmd/
    arena/
      main.go
    api/
      main.go
    notification-worker/
      main.go
    reward-worker/
      main.go
    analytics-worker/
      main.go

  internal/
    app/
    config/
    handler/
    service/
    repository/
    model/
    middleware/
    response/
    validator/
    infrastructure/

  proto/
  gen/
  migrations/
  tests/
  deploy/
  .github/
```

---

## Layered architecture

The main backend request flow:

```text
HTTP request
→ handler
→ service
→ repository
→ database/storage
```

### Handler layer

Responsible for:

```text
HTTP requests
JSON decoding
path/query parameters
basic validation
calling services
JSON responses
HTTP status codes
```

### Service layer

Responsible for:

```text
business logic
battle rules
hero leveling
reward calculation
cooldown checks
event publishing
transactions orchestration
```

### Repository layer

Responsible for:

```text
data storage
SQL queries
PostgreSQL
MongoDB
Redis
CRUD operations
```

### Infrastructure layer

Responsible for:

```text
PostgreSQL connection
Redis client
MongoDB client
Kafka producer/consumer
RabbitMQ producer/consumer
gRPC clients
```

---

## Core scenario

```text
1. User registers
2. User logs in
3. User creates Hero
4. Hero starts Battle
5. API checks JWT
6. BattleService checks cooldown in Redis
7. BattleService uses Redis distributed lock
8. BattleService loads Hero and Enemy from PostgreSQL
9. BattleService runs battle algorithm
10. Battle result is saved in PostgreSQL transaction
11. Battle rounds are saved in PostgreSQL
12. Full battle log is saved in MongoDB
13. battle.finished event is published to Kafka
14. reward.grant job is published to RabbitMQ
15. RewardWorker grants experience or item
16. RatingService updates leaderboard in Redis
17. API returns battle result as JSON
```

---

## Planned REST API

### Health

```http
GET /health
GET /health/live
GET /health/ready
```

### Auth

```http
POST /api/v1/auth/register
POST /api/v1/auth/login
POST /api/v1/auth/refresh
POST /api/v1/auth/logout
GET  /api/v1/users/me
```

### Heroes

```http
POST   /api/v1/heroes
GET    /api/v1/heroes
GET    /api/v1/heroes/{id}
PUT    /api/v1/heroes/{id}
PATCH  /api/v1/heroes/{id}
DELETE /api/v1/heroes/{id}
```

### Battles

```http
POST /api/v1/battles
GET  /api/v1/battles/{id}
GET  /api/v1/heroes/{id}/battles
GET  /api/v1/battles/{id}/rounds
GET  /api/v1/battles/{id}/log
```

### Inventory

```http
GET    /api/v1/heroes/{id}/inventory
POST   /api/v1/heroes/{id}/inventory/items
PATCH  /api/v1/heroes/{id}/inventory/items/{item_id}/equip
DELETE /api/v1/heroes/{id}/inventory/items/{item_id}
```

### Rating

```http
GET /api/v1/leaderboard
GET /api/v1/heroes/{id}/rating
```

---

## Planned PostgreSQL tables

```text
users
heroes
enemies
battles
battle_rounds
items
hero_inventory
ratings
refresh_tokens
```

---

## Planned Redis usage

```text
hero:{id}
leaderboard:global
battle:cooldown:hero:{id}
rate_limit:ip:{ip}
refresh_token:{token_id}
lock:battle:hero:{id}
arena:events
```

---

## Planned MongoDB collections

```text
battle_logs
audit_logs
event_debug_logs
```

---

## Planned RabbitMQ usage

```text
Exchanges:
- devarena.direct
- devarena.fanout
- devarena.topic
- devarena.dlx

Queues:
- notification.queue
- reward.queue
- levelup.queue
- dead_letter.queue

Routing keys:
- battle.finished
- hero.level_up
- reward.grant
- notification.send
```

---

## Planned Kafka topics

```text
devarena.hero.events
devarena.battle.events
devarena.inventory.events
devarena.rating.events
```

---

## Planned gRPC services

```text
HeroService
BattleService
RatingService
InventoryService
```

RPC types to learn:

```text
unary RPC
server streaming
client streaming
bidirectional streaming
```

---

## Planned Kubernetes resources

```text
namespace
deployment
service
ingress
configmap
secret
readiness probe
liveness probe
resource limits
horizontal pod autoscaler
```

---

## Planned CI/CD

```text
GitHub Actions:
- build
- test
- lint
- docker build
- docker push
- deploy
```