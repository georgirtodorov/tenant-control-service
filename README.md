# HR Leave Management System - Multi-Tenant Control Plane

## 🎯 Project Overview
This is a high-scale, multi-tenant infrastructure designed to manage isolated HR environments (Leave Requests) for different companies. The project follows the **Level 2/3 SaaS Architecture** (Control Plane vs. Data Plane isolation).

### Key Architectural Pillars:
- **Tenant Isolation:** Database-per-tenant strategy.
- **Security:** Zero-Trust architecture using HashiCorp Vault for dynamic credentials.
- **Scalability:** Application-level sharding (Tenant Routing).
- **Technology Stack:** Go (Golang), gRPC, PostgreSQL, Redis, HashiCorp Vault, Docker.

---

## 🗺️ Roadmap & Implementation Phases

### Phase 0: Infrastructure & Foundation (Current)
- [ ] Setup `docker-compose.yml` with Control Plane DB, Data Plane DB, Redis, and Vault.
- [ ] Define the `tenant.proto` gRPC contract for the Registry service.
- [ ] Initialize the Go project structure (`/cmd`, `/internal`, `/api`).

### Phase 1: Tenant Registry (The Control Plane)
- [ ] Implement the `Registry` service: CRUD for tenant metadata.
- [ ] Add Redis caching layer for fast tenant resolution.
- [ ] Implement JWT Service-to-Service authentication.

### Phase 2: Tenant Provisioning (Automation)
- [ ] Build the `Provisioner` module: Automate PostgreSQL database & user creation.
- [ ] Integrate a Migration Orchestrator to run schemas on new tenant databases.
- [ ] Implement the `CredentialsProvider` interface for Vault simulation.

### Phase 3: Resolver & Context Propagation
- [ ] Create a Go Middleware for tenant resolution (Subdomain/Header based).
- [ ] Implement gRPC client communication between the HR App and the Tenant Manager.
- [ ] Inject active DB connections into the Request Context.

### Phase 4: Monitoring & Health Checks
- [ ] Implement a dashboard/service to monitor tenant database health.
- [ ] Setup automated logging and audit trails in Vault.

---

## 🤖 Instructions for AI Assistants (Gemini/IDX)
*When working on this project, please adhere to the following rules:*

1. **Strict Isolation:** Never mix Control Plane (Registry) logic with Data Plane (HR Business Logic).
2. **Security First:** Always use interfaces for credential retrieval (prepare for Vault integration). No hardcoded passwords.
3. **Context Matters:** Every database operation must be tenant-aware and scoped via Go `context`.
4. **Step-by-Step:** Do not generate massive codebases at once. Focus only on the current Phase as marked in the Roadmap.

---

## 🛠️ Getting Started
1. Clone the repository.
2. Run `docker-compose up -d` to start the infrastructure.
3. Follow the instructions in `Phase 0` to begin development.
