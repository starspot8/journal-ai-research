HANDOFF AI-01

Product & Architecture

Journal AI Research

1. STATUS

"NOT STARTED"

AI-01 must change this to:

"COMPLETED"

only after all AI-01 requirements have been reviewed and finalized.

---

2. OBJECTIVE

AI-01 defines the product requirements and architecture for Journal AI Research.

The goal is to provide AI-02 with a stable product specification for UX/UI implementation.

---

3. PRODUCT

Product:

Journal AI Research

Core workflow:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

Primary function:

The user enters a journal/research topic in a single search box.

The system searches scientific sources and returns:

10 Latest Relevant Journals

---

4. V1 SCOPE

Required:

- one search box
- natural-language/simple search
- PubMed
- Europe PMC
- Crossref
- normalization
- deduplication
- relevance ranking
- recency ranking
- top 10 results
- article detail
- abstract
- legal full-text access
- original article link
- PDF report
- security
- error handling
- responsive web interface

Not included:

- login
- user profile
- social login
- native mobile app
- citation manager
- automatic systematic review
- automatic meta-analysis
- paywall bypass
- unauthorized full-text scraping
- medical diagnosis
- treatment recommendation

---

5. USER FLOW

USER
↓
SEARCH BOX
↓
QUERY
↓
PUBMED + EUROPE PMC + CROSSREF
↓
NORMALIZE
↓
DEDUPLICATE
↓
RANK
↓
TOP 10
↓
ARTICLE DETAIL
↓
FULL TEXT IF LEGALLY AVAILABLE
↓
PDF REPORT

---

6. SOURCE STRATEGY

Initial sources:

1. PubMed
2. Europe PMC
3. Crossref

The application should not depend on one source.

One source failing should not automatically cause the entire search to fail.

---

7. RANKING

Initial model:

Relevance              60%
Recency                30%
Metadata completeness  10%

Ranking represents relevance/recency for the query.

It must not be presented as scientific quality ranking.

UI terminology:

10 Latest Relevant Journals

---

8. DEDUPLICATION

Use:

1. DOI
2. PMID
3. PMCID
4. normalized title + author + year

Multiple source records representing the same article must be merged.

---

9. ARTICLE MODEL

Minimum normalized article:

id
doi
pmid
pmcid
title
abstract
authors
journal
publication_date
source
source_url
fulltext_url
is_open_access
created_at
updated_at

---

10. FULL-TEXT POLICY

The application may display full text only when access is legally available.

If unavailable:

FULL TEXT NOT AVAILABLE

[OPEN ORIGINAL ARTICLE]

The application must never:

- bypass a paywall
- bypass authentication
- scrape restricted content
- create unauthorized copies

---

11. PDF POLICY

V1 PDF means a generated research report.

It may include:

- query
- date
- selected articles
- metadata
- abstract where permitted
- source links

A publisher's original article PDF remains subject to its access rights.

---

12. API

POST /api/v1/search
GET  /api/v1/articles/:id
GET  /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET  /api/v1/health

---

13. DATABASE

Conceptual tables:

articles
searches
search_results

AI-05 will finalize PostgreSQL implementation.

---

14. SECURITY

Required:

- HTTPS
- input validation
- rate limiting
- timeouts
- parameterized SQL
- secret management
- secure headers
- structured logging
- no frontend API secrets

---

15. PERFORMANCE

Initial target:

Approximately:

< 3 seconds

under normal external API conditions.

Implementation should use:

- parallel source requests
- timeouts
- caching
- connection reuse

Third-party API latency cannot be guaranteed.

---

16. TESTING

Must test:

- Indonesian queries
- English queries
- natural-language queries
- empty query
- long query
- duplicate articles
- missing metadata
- source failures
- full-text availability
- PDF generation
- malformed input
- rate limiting
- injection attempts

---

17. ACCEPTANCE CRITERIA

AI-01 product specification is accepted when:

- V1 scope is defined.
- User flow is defined.
- Sources are defined.
- Ranking is defined.
- Deduplication is defined.
- Full-text policy is defined.
- PDF policy is defined.
- API requirements are defined.
- Database model is defined.
- Security requirements are defined.
- Testing requirements are defined.
- AI-02 can design the UI without redefining the product.

---

18. DECISIONS

Record final decisions here.

Example:

DECISION-001
V1 uses one primary search box.

DECISION-002
Initial scientific sources are PubMed, Europe PMC and Crossref.

DECISION-003
Results are described as "10 Latest Relevant Journals".

DECISION-004
Paywall bypass is prohibited.

DECISION-005
V1 PDF is a generated research report.

Add new decisions only when required.

---

19. OPEN QUESTIONS

Only genuine unresolved questions should be listed here.

For each question:

Question:
Why it matters:
Suggested owner:
Required decision before:

Do not use this section to postpone decisions that can already be made from the specification.

---

20. INSTRUCTIONS FOR AI-02

AI-02 is responsible for UX/UI.

AI-02 must:

1. Read "SHARED.md".
2. Read this "HANDOFF_AI01.md".
3. Follow the frozen V1 scope.
4. Design the simplest possible search experience.
5. Design the results page.
6. Design the article reader.
7. Design the PDF/export interaction.
8. Design loading/error/empty states.
9. Keep the primary workflow simple.
10. Avoid adding features outside V1.

AI-02 must NOT redefine:

- scientific sources
- ranking model
- legal full-text policy
- product scope

unless a genuine conflict is discovered.

Any proposed change must be documented before implementation.

---

21. NEXT AI

Next:

AI-02 — UX/UI

Expected output:

START.md
SHARED.md
TASK.md
HANDOFF_AI02.md

AI-02 should produce an implementation-ready UX/UI specification for AI-03 Frontend.
