SHARED PROJECT SPECIFICATION

Journal AI Research

1. PRODUCT

Journal AI Research

Core workflow:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

2. PRODUCT SCOPE

V1 allows users to:

1. enter a research topic
2. search scientific literature
3. receive 10 latest relevant journals
4. open an article
5. read legally available content
6. open the original source
7. export a research report PDF

---

3. SCIENTIFIC SOURCES

Backend sources:

- PubMed
- Europe PMC
- Crossref

AI-03 does NOT implement these integrations.

---

4. RANKING

Backend ranking:

Relevance              60%
Recency                30%
Metadata completeness  10%

The frontend must not change or reinterpret this ranking.

---

5. DEDUPLICATION

Backend:

1. DOI
2. PMID
3. PMCID
4. title + author + year

The frontend receives already normalized/deduplicated results.

---

6. RESULTS

UI terminology:

10 Latest Relevant Journals

Never display:

10 Best Journals

---

7. FULL-TEXT POLICY

Only legally accessible content may be displayed.

If unavailable:

FULL TEXT NOT AVAILABLE

[ OPEN ORIGINAL ARTICLE ]

No paywall bypass.

---

8. PDF

V1 PDF is a generated research report.

It may include:

- query
- search date
- selected articles
- metadata
- permitted abstracts
- source links

---

9. FRONTEND STACK

Required:

- Next.js
- TypeScript
- React

Use the project's approved version where already established.

---

10. FRONTEND PRINCIPLES

The UI must be:

- simple
- modern
- professional
- scientific
- readable
- responsive
- accessible

---

11. PRIMARY USER EXPERIENCE

The user should understand the product immediately:

Type topic
↓
Search
↓
See 10 journals
↓
Read
↓
Export PDF

---

12. API

Expected backend endpoints:

POST /api/v1/search
GET  /api/v1/articles/:id
GET  /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET  /api/v1/health

AI-03 should build an API abstraction around these endpoints.

---

13. SECURITY

Frontend must not contain:

- PubMed credentials
- Crossref credentials if any
- Europe PMC secrets if any
- database credentials
- server secrets

All server-side secrets belong to the backend.

---

14. ACCESSIBILITY

Required:

- keyboard navigation
- semantic HTML
- focus states
- readable contrast
- accessible forms
- accessible buttons
- touch-friendly controls

---

15. TESTING

Frontend must be tested before handoff.

Minimum:

- search
- article display
- loading
- empty
- errors
- PDF interaction
- responsive behavior
- accessibility

---

16. AI CHAIN

AI-01 Product + Architecture
↓
AI-02 UX/UI
↓
AI-03 Frontend
↓
AI-04 Backend + Search
↓
AI-05 Database
↓
AI-06 Integration
↓
AI-07 Testing + Security
↓
AI-08 Deployment

---

17. HANDOFF RULE

AI-03 must not leave undocumented technical assumptions for AI-04.

Anything AI-04 needs to know must be recorded in:

"HANDOFF_AI03.md"
