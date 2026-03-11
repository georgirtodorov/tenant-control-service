# Universal Tenant Control Plane

Това е универсална gRPC услуга за управление на клиенти (tenants) в мултитенантна архитектура. Проектът служи като "Control Plane" за управление на инфраструктурата, изолацията на бази данни и сигурността.

## 🎯 Архитектурни цели
- **Изолация на данните:** Стратегия "База данни за всеки клиент" (Level 2/3 SaaS Architecture).
- **Сигурност:** Zero-Trust модел, защита срещу инжекции и неаутентикиран достъп.
- **Универсалност:** Независим от бизнес логиката "Blueprint", подходящ за всякакви системи.

---

## 🗺️ Roadmap & Implementation Phases

### Phase 0: Инфраструктура и Основи (Current)
- [ ] Настройка на `docker-compose.yml` (Control Plane, Data Plane, Redis, Vault).
- [ ] Инициализация на Go структурата в `cmd/tenant-control-service/`.
- [ ] Дефиниране на gRPC договора (`api/proto/tenant.proto`).

### Phase 1: Tenant Registry, Security & Testing
- [ ] **Security Interceptors:** Имплементиране на JWT аутентикация за защита на gRPC ендпоинтите.
- [ ] **Integration Tests:** "Black-box" тестове през реална мрежа в `internal/registry/integration_test.go`.
- [ ] **Testcontainers:** Автоматично вдигане на ефимерни Docker контейнери (Postgres/Redis) за тестовете.
- [ ] **Seeding & Migrations:** Автоматично попълване на тестови данни и мигриране на схемата преди тестове.
- [ ] **Registry Service:** CRUD за метаданни с валидация срещу SQL инжекции.
- [ ] **Caching Layer:** Redis интеграция за бърза резолюция.

### Phase 2: Tenant Provisioning (Automation)
- [ ] **Provisioner Module:** Автоматично създаване на PostgreSQL бази и потребители.
- [ ] **Vault Integration:** Динамични креденшъли чрез `CredentialsProvider` интерфейс.
- [ ] **Migration Orchestrator:** Автоматично пускане на схеми върху новите бази.

### Phase 3: Resolver & Context Propagation
- [ ] Middleware за резолюция (Header/Metadata базирана).
- [ ] Инжектиране на активни DB конекции в Go `context`.

### Phase 4: Monitoring & Health Checks
- [ ] Dashboard за мониторинг на здравето на тенантните бази.
- [ ] Одит логване на достъпа в Vault.

---

## 🧪 Тестване и Качество
Проектът използва модерна методология за интеграционно тестване:
1. **End-to-End Flow:** Тестване на пълния gRPC път – от клиента, през мрежата, до базата данни.
2. **Testcontainers:** Използва се `testcontainers-go` за стартиране на истински Postgres/Redis в Docker контейнери по време на тест.
3. **Database Seeding:** Всеки тест започва с прясно "сийдната" база данни за 100% възпроизводимост.



---

## 🛠️ Изисквания (Prerequisites)

Преди да започнете, уверете се, че имате:

1. **Go (1.18+):** [Инсталация](https://go.dev/doc/install)
2. **Protocol Buffer Compiler (protoc):** [Releases](https://github.com/protocolbuffers/protobuf/releases)
3. **Go плъгини за protoc:**
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest