AI-06 — INTEGRATION

TASK

---

PHASE 1 — AUDIT

Baca semua handoff.

Buat tabel internal:

Component | Owner | Actual Status | Integration Requirement

Minimal:

Frontend
Backend
Search
PostgreSQL
Full Text
PDF

---

PHASE 2 — VERIFY API CONTRACT

Bandingkan:

AI-03 frontend types
VS
AI-04 backend response

Periksa:

- request body;
- response;
- article fields;
- errors;
- HTTP status;
- nullability;
- dates;
- IDs.

---

PHASE 3 — FRONTEND CONFIGURATION

Frontend harus memiliki API base URL melalui environment.

Contoh:

NEXT_PUBLIC_API_BASE_URL=http://localhost:8080

Gunakan actual variable jika AI-03 sudah menetapkan nama berbeda.

Jangan hardcode production URL.

---

PHASE 4 — SEARCH INTEGRATION

Implement:

Search page
   ↓
POST /api/v1/search
   ↓
Results page

Test query:

latest journal about cataract surgery

---

PHASE 5 — SEARCH STATES

Implement/verify:

idle
loading
success
empty
error
partial

---

PHASE 6 — RESULTS

Pastikan:

maximum 10

article cards tampil dengan data aktual.

Tidak boleh ada mock data ketika production API tersedia.

---

PHASE 7 — MOCK DATA CONTROL

Jika AI-03 memiliki mock API:

pastikan mock dapat:

development only

dan tidak aktif secara default pada production.

Jangan menyembunyikan fallback mock apabila backend gagal.

Backend failure harus tampil sebagai error.

---

PHASE 8 — ARTICLE DETAIL

Connect:

GET /api/v1/articles/:id

Test:

search result
   ↓
click article
   ↓
article detail

---

PHASE 9 — FULL TEXT

Connect:

GET /api/v1/articles/:id/fulltext

Test:

available
not available
error

---

PHASE 10 — ORIGINAL ARTICLE

Verify:

OPEN ORIGINAL ARTICLE

mengarah ke URL asli yang disediakan backend.

Tidak boleh menggunakan arbitrary URL dari user.

---

PHASE 11 — DATABASE INTEGRATION

Start PostgreSQL.

Run migrations.

Start backend.

Verify:

search
 ↓
article persistence
 ↓
search persistence
 ↓
search results persistence

---

PHASE 12 — SEARCH PERSISTENCE

Test:

User search
   ↓
Backend
   ↓
PostgreSQL

Periksa:

- search row;
- article rows;
- search_result rows.

---

PHASE 13 — DUPLICATE ARTICLE

Simulasikan:

PubMed article
Europe PMC same article
Crossref same article

Verify:

one logical article

---

PHASE 14 — PARTIAL SOURCE FAILURE

Mock:

PubMed timeout

Expected:

Europe PMC + Crossref
        ↓
results

Frontend harus tetap usable.

---

PHASE 15 — ALL SOURCE FAILURE

Simulasikan:

PubMed fail
Europe PMC fail
Crossref fail

Verify backend returns meaningful error.

Frontend displays user-friendly message.

---

PHASE 16 — ARTICLE NOT FOUND

Request invalid ID.

Expected:

404

Frontend should display appropriate state.

---

PHASE 17 — PDF EXPORT

Connect:

POST /api/v1/export/pdf

Request should follow actual contract.

Test:

1 article
multiple articles
empty selection
invalid article ID

---

PHASE 18 — PDF CONTENT

Verify generated report contains:

Journal AI Research
Search query
Date
Article metadata
Abstract if permitted
Original/source URL

Do not copy restricted article content without authorization.

---

PHASE 19 — DOCKER COMPOSE

Verify:

docker compose up --build

Expected services:

frontend
backend
postgres

Actual service names may follow existing project.

---

PHASE 20 — ENVIRONMENT

Create/update:

.env.example

Document:

frontend API URL
backend port
database URL
timeouts
rate limit
other required configuration

No secrets.

---

PHASE 21 — CORS

Test:

frontend origin
   ↓
backend

Verify allowed origin.

Test rejected origin if practical.

---

PHASE 22 — SECURITY CHECK

Verify:

[ ] No database credentials in frontend
[ ] No scientific API secret in frontend
[ ] CORS configured
[ ] URL validation
[ ] SSRF protection
[ ] Rate limiting
[ ] Timeout
[ ] Input validation
[ ] Error sanitization

---

PHASE 23 — NETWORK FAILURE

Test:

backend unavailable
database unavailable
external API unavailable

Frontend must not hang indefinitely.

---

PHASE 24 — TIMEOUT

Verify:

external provider timeout
database timeout
frontend request timeout

No indefinite loading state.

---

PHASE 25 — RESPONSIVE

Test:

mobile
tablet
desktop

Verify:

- search;
- results;
- article reader;
- PDF controls.

---

PHASE 26 — ACCESSIBILITY

Verify:

keyboard only
tab navigation
Enter search
focus states
button labels
ARIA where needed
semantic headings

---

PHASE 27 — TESTING

Run actual project commands.

Frontend examples:

npm run lint
npm run typecheck
npm test
npm run build

Backend:

go test ./...

Database:

migration tests
repository tests

Use actual commands if project differs.

---

PHASE 28 — END-TO-END TEST

Run:

1. Start PostgreSQL
2. Run migrations
3. Start backend
4. Start frontend
5. Open Search
6. Enter query
7. Press Enter
8. Receive results
9. Open article
10. Read abstract/full text
11. Open original article
12. Export PDF
13. Verify PDF

---

PHASE 29 — REGRESSION

Verify existing functionality from AI-03 still works.

Do not accept:

backend works
but frontend broken

or:

frontend works
but database persistence broken

---

PHASE 30 — PERFORMANCE

Measure approximately:

frontend request time
backend processing
provider calls
database
PDF generation

Do not treat third-party provider latency as fully controllable.

---

PHASE 31 — DOCUMENTATION

Update:

README.md

Include:

architecture
local setup
environment
Docker
database
frontend
backend
testing

---

PHASE 32 — ACCEPTANCE CRITERIA

[ ] Search works end-to-end
[ ] Results work
[ ] Maximum 10
[ ] Partial source failure works
[ ] Article detail works
[ ] Full text works
[ ] Original link works
[ ] PDF works
[ ] Database persistence works
[ ] Docker works
[ ] Environment documented
[ ] Errors handled
[ ] Loading handled
[ ] Empty handled
[ ] Accessibility preserved
[ ] Responsive layout preserved
[ ] Frontend tests pass
[ ] Backend tests pass
[ ] Database tests pass
[ ] E2E test passes
[ ] HANDOFF_AI06.md completed

---

PHASE 33 — DO NOT ADD

Jangan implement:

login
subscription
social
citation manager
systematic review
meta-analysis
medical recommendation

Jika ada ide:

documentasikan sebagai future scope.
