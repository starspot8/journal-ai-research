AI-07 — TESTING + SECURITY

TASK

---

PHASE 1 — SYSTEM AUDIT

Read:

AI-01
AI-02
AI-03
AI-04
AI-05
AI-06

Record:

implemented feature
expected behavior
actual behavior
test status

---

PHASE 2 — BUILD VERIFICATION

Run actual commands.

Frontend:

npm run lint
npm run typecheck
npm test
npm run build

Backend:

go test ./...
go vet ./...

Use actual project commands if different.

---

PHASE 3 — DATABASE VERIFICATION

Verify:

migration
schema
constraints
indexes
repository
transactions

Run integration tests using isolated test database.

---

PHASE 4 — HEALTH CHECK

Test:

GET /api/v1/health

Verify:

- HTTP status;
- response format;
- database dependency behavior if health includes it;
- no secret exposure.

---

PHASE 5 — SEARCH FUNCTIONAL TEST

Run:

latest journal about cataract surgery

Verify:

request succeeds
results exist
maximum 10
metadata valid
ranking returned

---

PHASE 6 — LANGUAGE TEST

Run:

jurnal terbaru tentang diabetic retinopathy

and:

latest research about multifocal IOL

Verify both work.

---

PHASE 7 — LONG NATURAL QUERY

Run:

Saya ingin penelitian terbaru tentang hasil multifocal IOL pada pasien usia lanjut

Verify:

request accepted
results returned if available
no timeout caused by query length

---

PHASE 8 — EMPTY QUERY

Test:

""
"   "

Expected:

validation error

---

PHASE 9 — LARGE QUERY

Test an unusually large request.

Expected:

controlled validation error

not:

server crash
database overload

---

PHASE 10 — SPECIAL CHARACTERS

Test:

quotes
apostrophes
parentheses
slashes
Unicode
emoji

Verify:

no crash
no SQL injection

---

PHASE 11 — DUPLICATES

Force multiple sources to return the same article.

Verify:

one logical article

---

PHASE 12 — PARTIAL PROVIDER FAILURE

Simulate:

PubMed timeout

Verify:

other sources continue
results still render

Repeat for:

Europe PMC
Crossref

---

PHASE 13 — ALL PROVIDERS FAILURE

Simulate all providers unavailable.

Verify:

backend controlled response
frontend controlled error
no infinite loading

---

PHASE 14 — ARTICLE DETAIL

Test valid ID.

Verify:

title
authors
journal
date
identifiers
abstract
links

---

PHASE 15 — INVALID ARTICLE ID

Test:

invalid ID
unknown ID
malformed ID

Expected:

controlled 404/validation response

---

PHASE 16 — FULL TEXT

Test:

available
unavailable
provider failure

Verify correct UI state.

---

PHASE 17 — ORIGINAL ARTICLE

Verify external link:

valid
safe
correct destination

Test unsafe schemes.

---

PHASE 18 — PDF

Test:

one article
multiple articles
empty selection
invalid article

Verify PDF:

opens
contains correct metadata
contains query/date where required
does not expose secrets

---

PHASE 19 — API SECURITY

Test:

malformed JSON
missing body
wrong content type
oversized request
unexpected fields
rapid requests

---

PHASE 20 — SQL INJECTION

Use harmless test inputs.

Examples conceptually:

' OR '1'='1

Expected:

treated as data

No SQL execution manipulation.

---

PHASE 21 — XSS

Test metadata/rendering with strings resembling HTML/script payloads.

Verify:

not executed

---

PHASE 22 — SSRF

Audit every backend feature that fetches URLs.

Identify:

user-controlled URL
source-controlled URL
redirect
PDF URL
fulltext URL

Verify server does not blindly fetch arbitrary internal addresses.

---

PHASE 23 — CORS

Test:

allowed origin
unknown origin

Verify expected behavior.

---

PHASE 24 — RATE LIMITING

Send repeated search requests.

Record:

threshold
response
retry behavior

Do not run destructive traffic.

---

PHASE 25 — TIMEOUT

Simulate slow provider.

Verify:

bounded request
controlled failure

---

PHASE 26 — DATABASE FAILURE

Stop PostgreSQL.

Then test:

search
article
health

Verify:

controlled behavior

No raw database error.

---

PHASE 27 — BACKEND FAILURE

Stop backend.

Open frontend.

Verify:

user-friendly error

No infinite spinner.

---

PHASE 28 — ACCESSIBILITY

Test with keyboard only:

Tab
Shift+Tab
Enter
Escape where relevant

Verify search can be completed without mouse.

---

PHASE 29 — RESPONSIVE

Test:

mobile
tablet
desktop

Verify:

search
results
article
PDF

---

PHASE 30 — PERFORMANCE

Measure actual:

search
article detail
fulltext
PDF

Record environment and conditions.

Do not claim universal performance.

---

PHASE 31 — LOG REVIEW

Inspect logs for:

secrets
passwords
tokens
full API credentials
private filesystem paths
unnecessary sensitive query data

---

PHASE 32 — REPOSITORY SECRET SCAN

Search codebase for common credential patterns.

Check:

.env
.env.*
source
config
Docker
CI
frontend bundle

---

PHASE 33 — DEPENDENCY REVIEW

Check dependencies for:

- known security issues;
- abandoned critical packages;
- suspicious packages;
- unnecessary packages.

Use available package tooling.

---

PHASE 34 — FIX

For each issue:

identify
reproduce
fix
test
regression test
document

Do not mark an issue fixed without reproducing the original failure when practical.

---

PHASE 35 — REGRESSION

After fixes, repeat:

Search
Results
Article
Fulltext
Original
PDF

---

PHASE 36 — FINAL TEST MATRIX

Create:

Test ID
Feature
Scenario
Expected
Actual
Status
Severity if failed

---

PHASE 37 — RELEASE BLOCKERS

Mark blockers:

CRITICAL security
HIGH security
core flow broken
data corruption
production startup failure

---

PHASE 38 — DOCUMENTATION

Update:

README.md

if testing/setup instructions changed.

Create:

HANDOFF_AI07.md

---

PHASE 39 — FINAL HANDOFF

Include:

Total tests
Passed
Failed
Blocked
Security findings
Fixed issues
Remaining issues
Deployment blockers
AI-08 instructions

---

ACCEPTANCE CRITERIA

[ ] Build passes
[ ] Unit tests pass
[ ] Integration tests pass
[ ] E2E core flow passes
[ ] API validation passes
[ ] Security tests performed
[ ] SSRF reviewed/tested
[ ] SQL injection tested
[ ] XSS tested
[ ] CORS tested
[ ] Rate limit tested
[ ] Secrets checked
[ ] Database failure tested
[ ] External API failure tested
[ ] Accessibility tested
[ ] Responsive tested
[ ] PDF tested
[ ] Regression completed
[ ] All findings documented
[ ] No unresolved critical security issue
[ ] HANDOFF_AI07.md completed
