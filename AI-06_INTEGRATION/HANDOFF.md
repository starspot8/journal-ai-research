AI-06 — HANDOFF

INTEGRATION

«Dokumen ini wajib diisi berdasarkan implementasi aktual sebelum diserahkan kepada AI-07.»

---

1. STATUS

NOT STARTED

Final:

COMPLETE / PARTIAL

---

2. INTEGRATED ARCHITECTURE

Actual architecture:

DOCUMENT ACTUAL ARCHITECTURE

Expected:

Frontend
   ↓
Backend
   ↓
Search
   ↓
Scientific Sources
   ↓
Normalization
   ↓
Deduplication
   ↓
Ranking
   ↓
PostgreSQL

---

3. FRONTEND

Framework:

DOCUMENT ACTUAL

API base URL:

DOCUMENT ENV VARIABLE NAME

Pages integrated:

Search
Results
Article
PDF

---

4. BACKEND

Server:

DOCUMENT ACTUAL

Endpoints:

POST /api/v1/search
GET /api/v1/articles/:id
GET /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET /api/v1/health

Actual status:

DOCUMENT ACTUAL STATUS

---

5. DATABASE

Database:

PostgreSQL

Migration tool:

DOCUMENT ACTUAL

Persistence verified:

[ ] articles
[ ] searches
[ ] search_results

---

6. SEARCH FLOW

Actual flow:

QUERY
 ↓
BACKEND
 ↓
PUBMED
EUROPE PMC
CROSSREF
 ↓
NORMALIZE
 ↓
DEDUPLICATE
 ↓
RANK
 ↓
DATABASE
 ↓
FRONTEND

---

7. API COMPATIBILITY

Frontend/backend contract:

COMPATIBLE / CHANGES REQUIRED

If changed:

Change:
Reason:
Frontend impact:
Backend impact:

---

8. SEARCH TEST

Test query:

latest journal about cataract surgery

Result:

PASS / FAIL

Number returned:

DOCUMENT ACTUAL

---

9. PARTIAL FAILURE TEST

Scenario:

PubMed timeout
Europe PMC success
Crossref success

Result:

PASS / FAIL

Observed behavior:

DOCUMENT ACTUAL

---

10. ALL SOURCE FAILURE

Scenario:

PubMed fail
Europe PMC fail
Crossref fail

Result:

PASS / FAIL

Frontend behavior:

DOCUMENT ACTUAL

---

11. ARTICLE FLOW

Test:

Search
 ↓
Article card
 ↓
Article detail

Result:

PASS / FAIL

---

12. FULL TEXT

Available:

PASS / FAIL

Unavailable:

PASS / FAIL

Legal access policy preserved:

YES / NO

---

13. ORIGINAL ARTICLE

External original article link:

PASS / FAIL

URL validation:

PASS / FAIL

---

14. PDF

PDF export:

PASS / FAIL

Verified contents:

[ ] Query
[ ] Date
[ ] Article metadata
[ ] Abstract where permitted
[ ] Original/source link

---

15. DOCKER

Command:

docker compose up --build

Result:

PASS / FAIL

Services:

frontend
backend
postgres

Actual service names:

DOCUMENT ACTUAL

---

16. ENVIRONMENT

Required variables:

DOCUMENT ACTUAL VARIABLES

Secrets:

[ ] Not committed
[ ] Backend only

---

17. SECURITY

Checklist:

[ ] Input validation
[ ] Rate limiting
[ ] CORS
[ ] SSRF protection
[ ] URL validation
[ ] Timeout
[ ] Error sanitization
[ ] No frontend secrets
[ ] No database credentials in frontend

---

18. TEST RESULTS

Frontend:

npm run lint
npm run typecheck
npm test
npm run build

Results:

DOCUMENT ACTUAL

Backend:

go test ./...

Result:

DOCUMENT ACTUAL

Database:

DOCUMENT ACTUAL

E2E:

PASS / FAIL

---

19. PERFORMANCE

Document actual observations:

Frontend:
Backend:
Database:
External providers:
PDF:

Do not claim guaranteed external API latency.

---

20. ACCESSIBILITY

Verified:

[ ] Keyboard navigation
[ ] Search Enter key
[ ] Focus states
[ ] Semantic headings
[ ] Button labels
[ ] Screen-reader relevant labels

---

21. RESPONSIVE

Verified:

[ ] Mobile
[ ] Tablet
[ ] Desktop

---

22. KNOWN ISSUES

List actual issues:

1.
2.
3.

---

23. DEVIATIONS

From AI-01:

None / DOCUMENT

From AI-02:

None / DOCUMENT

From AI-03:

None / DOCUMENT

From AI-04:

None / DOCUMENT

From AI-05:

None / DOCUMENT

---

24. TECHNICAL DECISIONS

Document:

Decision:
Reason:
Impact:

---

25. AI-07 TEST INSTRUCTIONS

AI-07 harus memulai dengan:

1. Read all previous handoffs
2. Start PostgreSQL
3. Run migrations
4. Start backend
5. Start frontend
6. Execute E2E search
7. Test source failures
8. Test API errors
9. Test security
10. Test accessibility
11. Test responsive behavior
12. Test PDF
13. Run complete test suite

---

26. AI-07 FOCUS

AI-07 harus fokus pada:

Functional testing
Regression testing
Security testing
API testing
Database testing
Accessibility testing
Performance verification
Failure testing

AI-07 tidak boleh melakukan redesign besar.

---

27. FINAL SUMMARY

Setelah implementasi:

AI-06 successfully integrated:

Frontend:
...

Backend:
...

Scientific Search:
...

PostgreSQL:
...

Full Text:
...

PDF:
...

Testing:
...

Known Issues:
...

AI-07 should now focus on:
...

Jangan meninggalkan placeholder dalam final handoff.
