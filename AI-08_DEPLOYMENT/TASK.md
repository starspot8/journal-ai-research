AI-08 — DEPLOYMENT

TASK

---

PHASE 1 — READ HANDOFFS

Read:

AI-01_PRODUCT_ARCHITECTURE/HANDOFF.md
AI-02_UX_UI/HANDOFF.md
AI-03_FRONTEND/HANDOFF.md
AI-04_BACKEND_SEARCH/HANDOFF.md
AI-05_DATABASE_POSTGRESQL/HANDOFF.md
AI-06_INTEGRATION/HANDOFF.md
AI-07_TESTING_SECURITY/HANDOFF.md

Extract:

architecture
environment requirements
known issues
deployment blockers
security requirements
database requirements

Do not deploy if AI-07 has unresolved critical blockers.

---

PHASE 2 — AUDIT REPOSITORY

Verify:

frontend/
backend/
migrations/
Dockerfile
docker-compose
.env.example
README

Identify actual production entrypoints.

---

PHASE 3 — ENVIRONMENT CONFIGURATION

Create/update:

.env.example

Document required variables.

Rules:

no real secret
no production credential
no private key

---

PHASE 4 — DOCKER

Verify frontend container.

Verify backend container.

Verify PostgreSQL configuration.

Use production-appropriate images/configuration.

---

PHASE 5 — CONTAINER SECURITY

Check:

non-root user where practical
minimal image
no secret in image
no unnecessary package
health check
restart behavior

---

PHASE 6 — NETWORK

Design:

Internet
 ↓
HTTPS
 ↓
Reverse Proxy
 ↓
Frontend / Backend
 ↓
Private PostgreSQL

Do not expose PostgreSQL unnecessarily.

---

PHASE 7 — DOMAIN

Configure production domain.

Verify:

DNS
TLS certificate
HTTPS
redirect

Do not mark complete without verification.

---

PHASE 8 — CORS

Set production frontend origin.

Test:

allowed origin
unauthorized origin

---

PHASE 9 — DATABASE

Provision PostgreSQL.

Configure:

database
user
password
connection
pool
network

Do not hardcode credentials.

---

PHASE 10 — MIGRATIONS

Before production migration:

backup if database contains existing data

Then:

run migration
verify schema
verify indexes
verify constraints

---

PHASE 11 — DATABASE BACKUP

Configure automated backup.

Document:

frequency
retention
location
restore method

---

PHASE 12 — RESTORE TEST

If possible:

backup
 ↓
restore to isolated environment
 ↓
verify tables
 ↓
verify representative data

Record result.

---

PHASE 13 — BACKEND CONFIGURATION

Configure:

PORT
DATABASE
CORS
provider settings
timeouts
rate limits
logging

Use actual implementation names.

---

PHASE 14 — EXTERNAL PROVIDERS

Verify backend connectivity to:

PubMed
Europe PMC
Crossref

Test normal operation.

Do not bypass source access restrictions.

---

PHASE 15 — FRONTEND CONFIGURATION

Verify frontend points to production backend.

Never point production frontend directly to:

PubMed
Europe PMC
Crossref

Frontend communicates with backend.

---

PHASE 16 — BUILD

Run actual project commands.

Example:

npm run lint
npm run typecheck
npm test
npm run build

Backend:

go test ./...
go vet ./...
go build ./...

Use project-specific commands when different.

---

PHASE 17 — IMAGE BUILD

Build versioned images.

Example concept:

journal-ai-frontend:<version>
journal-ai-backend:<version>

Actual registry/tag strategy depends on infrastructure.

---

PHASE 18 — STAGING DEPLOYMENT

Deploy first to staging where available.

Run:

health
search
article
fulltext
PDF

---

PHASE 19 — STAGING E2E

Test:

KETIK
 ↓
CARI
 ↓
10 RESULTS
 ↓
BACA
 ↓
PDF

---

PHASE 20 — PRODUCTION DEPLOYMENT

Deploy only after staging verification.

Record:

version
commit
timestamp
environment

---

PHASE 21 — MIGRATION

Execute production migration safely.

Verify:

tables
indexes
constraints
application compatibility

---

PHASE 22 — HEALTH CHECK

Test:

GET /api/v1/health

Expected operational response.

Do not expose internal secrets.

---

PHASE 23 — PRODUCTION SMOKE TEST

Perform:

Open homepage
Search "latest research about multifocal IOL"
Open result
Read abstract/fulltext
Open original article
Export PDF

---

PHASE 24 — FAILURE TEST

Verify controlled behavior for:

provider failure
database failure
backend restart
frontend/backend communication failure

Do not perform destructive tests against production.

Use staging for destructive/failure simulation.

---

PHASE 25 — LOGGING

Verify production logs are generated.

Check:

errors
latency
request IDs where available
provider errors
database errors

Verify no secrets appear.

---

PHASE 26 — MONITORING

Configure monitoring for:

uptime
5xx
latency
CPU
memory
database
storage

---

PHASE 27 — ALERTING

Configure meaningful alerts.

At minimum consider:

backend unavailable
frontend unavailable
database unavailable
high 5xx
resource exhaustion

---

PHASE 28 — ROLLBACK TEST

Document and, where practical, test:

previous version
deployment rollback
health verification

Database rollback must be handled separately.

---

PHASE 29 — INCIDENT PROCEDURE

Document:

Who/what detects incident
How to inspect logs
How to rollback
How to restore DB
How to verify recovery

---

PHASE 30 — README

Update:

README.md

with:

local setup
production configuration
deployment
migration
backup
restore
rollback
monitoring
troubleshooting

Never put actual production secrets in README.

---

PHASE 31 — FINAL VERIFICATION

Verify:

[ ] HTTPS
[ ] Domain
[ ] Frontend
[ ] Backend
[ ] PostgreSQL
[ ] Migrations
[ ] External providers
[ ] CORS
[ ] Rate limiting
[ ] Health
[ ] Logging
[ ] Monitoring
[ ] Backup
[ ] Restore procedure
[ ] Rollback
[ ] PDF
[ ] Fulltext legal policy

---

PHASE 32 — HANDOFF

Complete:

HANDOFF.md

with actual deployment information.

Do not leave:

DOCUMENT
TBD
PASS / FAIL

in the final handoff.

---

ACCEPTANCE CRITERIA

[ ] Production build succeeds
[ ] Production containers start
[ ] Database migration succeeds
[ ] HTTPS works
[ ] Domain works
[ ] Backend health works
[ ] Frontend connects to backend
[ ] Search works
[ ] Article detail works
[ ] Fulltext behavior works
[ ] Original article links work
[ ] PDF works
[ ] Provider failure handling works
[ ] Secrets protected
[ ] CORS restricted
[ ] Rate limiting active
[ ] Logs available
[ ] Monitoring available
[ ] Backup configured
[ ] Restore procedure documented
[ ] Rollback documented
[ ] Production smoke test passes
[ ] HANDOFF_AI08.md completed
