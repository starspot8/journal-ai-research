SHARED PROJECT SPECIFICATION

Journal AI Research

1. PROJECT PRINCIPLE

The product principle is:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

The application must be simple enough that a user can search scientific literature using one search box.

---

2. CORE REQUIREMENT

The user types a topic or question.

Example:

"latest journal about cataract surgery"

The system searches scientific sources and returns:

10 Latest Relevant Journals

---

3. SOURCES

V1 sources:

1. PubMed
2. Europe PMC
3. Crossref

Scientific metadata must be traceable to source data.

---

4. RANKING

Initial ranking:

- Relevance: 60%
- Recency: 30%
- Metadata completeness: 10%

Ranking means relevance/recency for the user's search.

Ranking does NOT mean:

- scientific quality ranking
- best article ranking
- clinical recommendation
- medical recommendation

---

5. DEDUPLICATION

Priority:

1. DOI
2. PMID
3. PMCID
4. normalized title + author + year

---

6. FULL TEXT

Only legally accessible full text may be displayed.

Never:

- bypass paywalls
- bypass authentication
- scrape restricted content
- reproduce copyrighted content without authorization

When full text is unavailable:

Show metadata + abstract where available + original article link.

---

7. PDF

V1 PDF is primarily a research report.

It may contain:

- search query
- date
- selected articles
- metadata
- abstract where permitted
- source links

Original publisher PDFs must remain subject to their access rights.

---

8. V1 UI

Main screen:

JOURNAL AI RESEARCH

[ What journal are you looking for? ]

[ SEARCH ]

Results:

10 Latest Relevant Journals

[Article 1]
Title
Authors
Journal
Date
DOI
[BACA JURNAL]

[Article 2]
...

---

9. STACK

Frontend:

Next.js + TypeScript

Backend:

Go

Database:

PostgreSQL

Deployment:

Docker + cloud

---

10. API

POST /api/v1/search
GET  /api/v1/articles/:id
GET  /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET  /api/v1/health

---

11. SECURITY

Required:

- HTTPS
- validation
- rate limiting
- timeout
- parameterized SQL
- secure headers
- secret management
- structured logging
- no API secrets in frontend

---

12. PERFORMANCE

Initial target:

Normal search should generally return within approximately 3 seconds when external sources respond normally.

Use:

- parallel source requests
- timeout
- caching
- connection reuse

---

13. FAILURE HANDLING

One source failing must not necessarily cause the entire search to fail.

Example:

PubMed      ERROR
Europe PMC  SUCCESS
Crossref    SUCCESS

The application should return available results.

---

14. PRODUCT LANGUAGE

Use:

10 Latest Relevant Journals

Do not use:

10 Best Journals

Do not imply that ranking equals scientific quality.

---

15. V1 EXCLUSIONS

No:

- login
- social account
- native mobile app
- citation manager
- systematic review automation
- meta-analysis automation
- unauthorized article downloading
- paywall bypass
- medical diagnosis
- medical treatment recommendation

---

16. AI ROLE

AI may assist with:

- query understanding
- query expansion
- natural-language interpretation
- future summarization

AI is not the authoritative source for bibliographic metadata.

Source databases remain authoritative for metadata.

---

17. DEVELOPMENT RULE

Every AI must:

1. Read this file.
2. Read the previous AI handoff.
3. Follow the frozen V1 scope.
4. Avoid unnecessary feature expansion.
5. Document important decisions.
6. Produce a handoff for the next AI.

---

18. HANDOFF CHAIN

AI-01
Product + Architecture
        ↓
AI-02
UX/UI
        ↓
AI-03
Frontend
        ↓
AI-04
Backend + Search
        ↓
AI-05
Database
        ↓
AI-06
Integration
        ↓
AI-07
Testing + Security
        ↓
AI-08
Deployment

Each AI must leave the project in a state that the next AI can continue safely.
