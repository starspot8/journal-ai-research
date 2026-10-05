HANDOFF AI-03

FRONTEND IMPLEMENTATION

Journal AI Research

1. STATUS

"NOT STARTED"

Change to:

"COMPLETED"

only after the frontend has been implemented, tested, and built successfully.

---

2. OBJECTIVE

AI-03 implements the frontend based on AI-01 and AI-02.

The frontend provides the user interface for:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

3. TECHNOLOGY

Expected:

- Next.js
- TypeScript
- React

---

4. IMPLEMENTED PAGES

Document the final routes here.

Expected:

/
/articles/[id]

If a separate results route is implemented, document it here.

---

5. IMPLEMENTED COMPONENTS

Document final components.

Minimum expected:

AppHeader
SearchBox
SearchButton
SearchExamples
SearchLoading
SearchError
SearchEmpty
ResultsHeader
ArticleCard
ArticleMetadata
ArticleActions
ArticleReader
FullTextStatus
PdfExport

---

6. SEARCH API CONTRACT

Frontend expects:

POST /api/v1/search

Expected request:

{
  "query": "latest cataract surgery research"
}

Document the exact response schema implemented by the frontend.

Example conceptual response:

{
  "query": "latest cataract surgery research",
  "results": [],
  "total": 0,
  "warnings": []
}

---

7. ARTICLE API CONTRACT

Frontend expects:

GET /api/v1/articles/:id

Document the exact article schema used.

---

8. FULL-TEXT API

Frontend expects:

GET /api/v1/articles/:id/fulltext

Document:

- available state
- unavailable state
- content format
- source information

---

9. PDF API

Frontend expects:

POST /api/v1/export/pdf

Document:

- request body
- response
- content type
- error behavior

---

10. MOCK MODE

If mock mode was used, document:

- how it is enabled
- where mock data lives
- how to disable it
- how production prevents accidental mock usage

---

11. ENVIRONMENT VARIABLES

Document frontend environment variables.

Example:

NEXT_PUBLIC_API_BASE_URL

Do not place secrets in "NEXT_PUBLIC_*".

Document which variables are required and which are optional.

---

12. DESIGN SYSTEM

Document final:

- colors
- typography
- spacing
- breakpoints
- components
- states

Any deviation from AI-02 must be explained.

---

13. TESTING

Document:

- test framework
- test files
- unit tests
- component tests
- accessibility tests
- build validation

Record final results:

Lint: PASS/FAIL
Typecheck: PASS/FAIL
Tests: PASS/FAIL
Build: PASS/FAIL

---

14. RESPONSIVE VALIDATION

Record validation for:

320px
375px
390px
768px
1024px
1280px+

Document any known issue.

---

15. ACCESSIBILITY

Record:

- keyboard navigation
- focus
- form labels
- semantic HTML
- contrast
- touch targets

---

16. SECURITY

Confirm:

[ ] no API secrets in frontend

[ ] no database credentials

[ ] external links handled safely

[ ] rich text sanitized if applicable

[ ] user input handled safely

[ ] no unnecessary third-party scripts

---

17. KNOWN LIMITATIONS

Document anything intentionally incomplete.

Examples:

- real scientific search API not connected yet
- PDF generation not connected yet
- full-text backend not connected yet

Do not hide known limitations.

---

18. TECHNICAL DECISIONS

Document important frontend decisions.

Format:

DECISION-001
Decision:
Reason:

DECISION-002
Decision:
Reason:

---

19. AI-04 REQUIREMENTS

AI-04 is responsible for:

Backend + Search Engine

AI-04 must connect the frontend to:

- PubMed
- Europe PMC
- Crossref

AI-04 must preserve the frontend API contract where possible.

If the backend response must change:

1. document the change
2. update the shared contract
3. update the frontend accordingly
4. record the change in the handoff

---

20. AI-04 MUST IMPLEMENT

AI-04 will provide:

POST /api/v1/search
GET /api/v1/articles/:id
GET /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET /api/v1/health

AI-04 must not require the frontend to directly call PubMed, Europe PMC, or Crossref.

---

21. NEXT AI

Next:

AI-04 — BACKEND + SEARCH

AI-04 must read:

SHARED.md
HANDOFF_AI01.md
HANDOFF_AI02.md
HANDOFF_AI03.md

AI-04 then implements the real backend search pipeline:

User Query
↓
Go API
↓
PubMed
Europe PMC
Crossref
↓
Normalize
↓
Deduplicate
↓
Rank
↓
Top 10
↓
Frontend

---

22. FINAL CHECK

Before setting status to COMPLETED:

[ ] frontend works

[ ] build passes

[ ] tests pass

[ ] search UI works

[ ] results UI works

[ ] article UI works

[ ] PDF UI works

[ ] loading states work

[ ] empty states work

[ ] error states work

[ ] mobile works

[ ] accessibility implemented

[ ] API contract documented

[ ] AI-04 instructions complete
