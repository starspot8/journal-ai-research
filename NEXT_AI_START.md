# Journal AI Research — Master Briefing for Next AI

## 1. PURPOSE

You are the next AI agent continuing development of the **Journal AI Research** project.

Your responsibility is to continue implementation from the repository's actual state.

**IMPORTANT: Do not assume that a specification, START.md, TASK.md, or HANDOFF.md means the corresponding implementation already exists.**

The repository has been audited. The current repository contains primarily project specifications, handoffs, and a minimal frontend package definition.

Your first responsibility is to inspect the repository and establish the actual implementation state before writing code.

---

## 2. CURRENT REPOSITORY STATE

Repository:

`journal-ai-research`

Current branch:

`main`

The working tree has been verified clean.

Current high-level structure:

```text
journal-ai-research/
├── AI-01_PRODUCT_ARCHITECTURE/
├── AI-02_UX_UI/
├── AI-03_FRONTEND/
├── AI-03 — Frontend/
├── AI-04_BACKEND_SEARCH/
├── AI-05_DATABASE_POSTGRESQL/
├── AI-06_INTEGRATION/
├── AI-07_TESTING_SECURITY/
├── AI-08_DEPLOYMENT/
├── docs/
├── frontend/
│   └── package.json
└── README.md
```

---

## 3. CRITICAL IMPLEMENTATION STATUS

At the time this document was created:

### AI-01 — Product & Architecture

Specification and handoff documentation exists.

**Implementation status: NOT VERIFIED / DO NOT ASSUME IMPLEMENTED.**

Read:

```text
AI-01_PRODUCT_ARCHITECTURE/START.md
AI-01_PRODUCT_ARCHITECTURE/TASK.md
AI-01_PRODUCT_ARCHITECTURE/SHARED.md
AI-01_PRODUCT_ARCHITECTURE/HANDOFF.md
```

### AI-02 — UX/UI

Specification and handoff documentation exists.

**Implementation status: NOT VERIFIED / DO NOT ASSUME IMPLEMENTED.**

Read:

```text
AI-02_UX_UI/START.md
AI-02_UX_UI/TASK.md
AI-02_UX_UI/SHARED.md
AI-02_UX_UI/HANDOFF.md
```

### AI-03 — Frontend

Frontend documentation exists.

The repository currently contains only:

```text
frontend/package.json
```

The complete frontend implementation has **not** been verified as present.

**Implementation status: NOT COMPLETE.**

Read both relevant documentation areas:

```text
AI-03_FRONTEND/START.md
AI-03_FRONTEND/SHARED.md

AI-03 — Frontend/START.md
AI-03 — Frontend/TASK.md
AI-03 — Frontend/SHARED.md
AI-03 — Frontend/HANDOFF.md
```

Do not create duplicate frontend architecture without first reconciling these documents.

### AI-04 — Backend Search

Documentation exists:

```text
AI-04_BACKEND_SEARCH/START.md
AI-04_BACKEND_SEARCH/TASK.md
AI-04_BACKEND_SEARCH/SHARED.md
AI-04_BACKEND_SEARCH/HANDOFF.md
```

The repository currently does **not** contain a verified backend implementation.

There is currently no verified:

```text
backend/
go.mod
```

for this project.

**Implementation status: NOT IMPLEMENTED / NOT VERIFIED.**

### AI-05 — PostgreSQL Database

Documentation exists:

```text
AI-05_DATABASE_POSTGRESQL/START.md
AI-05_DATABASE_POSTGRESQL/TASK.md
AI-05_DATABASE_POSTGRESQL/SHARED.md
AI-05_DATABASE_POSTGRESQL/HANDOFF.md
```

The repository currently does **not** contain a verified database implementation.

There is currently no verified:

```text
database/
database/migrations/
docker-compose.yml
```

for this project.

**Implementation status: NOT IMPLEMENTED / NOT VERIFIED.**

### AI-06 — Integration

Documentation exists, but implementation depends on the actual completion of the frontend, backend, and database layers.

Do not treat AI-06 as completed merely because its handoff documents exist.

### AI-07 — Testing & Security

Documentation exists.

Actual test/security implementation must be verified against the source code once implementation exists.

### AI-08 — Deployment

Documentation exists.

Deployment configuration must be verified against the actual implementation before claiming deployment readiness.

---

## 4. REQUIRED FIRST ACTION

Before implementing anything:

### Step 1 — Read the project status

Read:

```text
README.md
docs/PROJECT_STATUS.md
```

### Step 2 — Read the relevant handoffs

For the current implementation phase, read:

```text
AI-01_PRODUCT_ARCHITECTURE/HANDOFF.md
AI-02_UX_UI/HANDOFF.md
AI-03 — Frontend/HANDOFF.md
AI-04_BACKEND_SEARCH/HANDOFF.md
AI-05_DATABASE_POSTGRESQL/HANDOFF.md
```

Also inspect their corresponding:

```text
START.md
TASK.md
SHARED.md
```

where available.

### Step 3 — Audit the actual filesystem

Run an actual repository inspection.

Do not infer implementation from filenames or commit messages.

Verify:

```text
frontend/
backend/
database/
docker-compose.yml
```

and inspect their actual contents.

### Step 4 — Establish an implementation matrix

Create a clear table such as:

| Area                 | Specification | Source Code  | Tests        | Verified |
| -------------------- | ------------- | ------------ | ------------ | -------- |
| Product Architecture | Yes           | N/A          | N/A          | Yes      |
| UX/UI                | Yes           | Not verified | Not verified | No       |
| Frontend             | Yes           | Incomplete   | Not verified | No       |
| Backend Search       | Yes           | Missing      | Missing      | No       |
| PostgreSQL           | Yes           | Missing      | Missing      | No       |
| Integration          | Yes           | Not verified | Not verified | No       |
| Testing/Security     | Yes           | Not verified | Not verified | No       |
| Deployment           | Yes           | Not verified | Not verified | No       |

Update this matrix based on actual evidence.

---

## 5. IMPLEMENTATION RULE

The most important rule:

> **Documentation is not implementation.**

Never state:

* "AI-04 is complete"
* "AI-05 is complete"
* "PostgreSQL is implemented"
* "Backend is implemented"
* "Frontend is complete"

unless the corresponding source code exists and has been verified.

Completion claims must be based on evidence such as:

* source files exist;
* dependencies are installed;
* application builds;
* migrations execute;
* database schema exists;
* API endpoints work;
* tests pass;
* integration has been verified.

---

## 6. EXPECTED TARGET ARCHITECTURE

Unless the existing specifications require otherwise, the target repository is expected to evolve toward:

```text
journal-ai-research/
├── frontend/
│   ├── package.json
│   ├── app/
│   ├── components/
│   ├── lib/
│   └── ...
│
├── backend/
│   ├── go.mod
│   ├── cmd/
│   ├── internal/
│   └── ...
│
├── database/
│   └── migrations/
│
├── docker-compose.yml
│
├── AI-01_PRODUCT_ARCHITECTURE/
├── AI-02_UX_UI/
├── AI-03_FRONTEND/
├── AI-04_BACKEND_SEARCH/
├── AI-05_DATABASE_POSTGRESQL/
├── AI-06_INTEGRATION/
├── AI-07_TESTING_SECURITY/
├── AI-08_DEPLOYMENT/
│
├── docs/
└── README.md
```

Do not blindly create this structure. Reconcile it with the existing specifications first.

---

## 7. ORDER OF WORK

The recommended implementation sequence is:

```text
1. Audit repository
       ↓
2. Reconcile AI-01 / AI-02 specifications
       ↓
3. Implement frontend according to AI-03
       ↓
4. Implement backend search according to AI-04
       ↓
5. Implement PostgreSQL persistence according to AI-05
       ↓
6. Integrate frontend + backend + database according to AI-06
       ↓
7. Run testing/security work according to AI-07
       ↓
8. Prepare deployment according to AI-08
```

However, if the handoff documents define a more appropriate dependency order, follow the documented dependency.

---

## 8. DATABASE REQUIREMENTS

When implementing AI-05:

Do not merely create SQL files.

Verify that:

1. PostgreSQL starts successfully.
2. Database migrations execute successfully.
3. Tables, indexes, constraints, and relationships exist.
4. The backend can persist data through the defined repository boundary.
5. Search/retrieval behavior works against PostgreSQL where required.
6. The implementation matches the AI-01 architecture and AI-05 specification.
7. Tests verify persistence behavior.

Do not claim PostgreSQL implementation is complete until it has been executed and verified.

---

## 9. BACKEND REQUIREMENTS

When implementing AI-04:

Verify:

* Go module configuration;
* package structure;
* API boundaries;
* request/response contracts;
* search behavior;
* database/repository boundary;
* validation;
* error handling;
* tests.

The backend must be implemented according to the project specifications rather than inventing an unrelated architecture.

---

## 10. FRONTEND REQUIREMENTS

When implementing AI-03:

Verify:

* Next.js configuration;
* TypeScript;
* application routes;
* components;
* journal/search UI;
* API integration;
* loading states;
* error states;
* empty states;
* responsive behavior;
* tests.

Do not replace the UX specification with a generic dashboard.

---

## 11. VERIFICATION STANDARD

Every major implementation phase must end with verification.

At minimum report:

```text
What was implemented
Files changed
Commands executed
Build result
Test result
Database result, if applicable
Known limitations
```

If something cannot be verified, explicitly say:

```text
NOT VERIFIED
```

Do not substitute assumptions for verification.

---

## 12. GIT DISCIPLINE

Keep commits understandable and incremental.

Recommended commit style:

```text
Implement frontend foundation
Implement backend search API
Implement PostgreSQL schema and migrations
Integrate frontend with backend
Add integration tests
Add security validation
Prepare deployment configuration
```

Before each commit:

```bash
git status
```

After implementation:

```bash
git diff
```

Before claiming completion:

```bash
git status
```

and run the relevant build/test commands.

---

## 13. DO NOT DO THIS

Do not:

* assume handoff documentation equals completed code;
* rewrite specifications without reason;
* delete existing handoff documentation;
* replace the architecture without checking AI-01;
* skip database verification;
* claim success without running tests/builds;
* create duplicate implementations because similarly named documentation exists;
* remove project files merely because they appear unused;
* overwrite existing configuration without inspecting it first.

---

## 14. FIRST RESPONSE EXPECTED FROM THE NEXT AI

Before coding, provide a concise audit containing:

1. Current repository structure.
2. Existing implementation found.
3. Missing implementation.
4. Relevant specifications discovered.
5. Dependency/blocker analysis.
6. Recommended first implementation step.

Then proceed with implementation.

---

## 15. PROJECT PRINCIPLE

The project must progress from:

```text
SPECIFICATION
      ↓
IMPLEMENTATION
      ↓
VERIFICATION
      ↓
INTEGRATION
      ↓
TESTING
      ↓
DEPLOYMENT
```

Never treat one stage as completed merely because the previous stage has documentation.

**Evidence over assumption.**

---

## 16. HANDOFF STATUS

This repository is a valid continuation point, but it is **not yet a fully implemented Journal AI application**.

The next AI should treat the repository as:

> **A specification-driven project with a minimal frontend package scaffold, requiring implementation of the application layers described by AI-01 through AI-08.**

Start by auditing the repository and then implement the highest-priority missing layer according to the documented dependencies.
