AI-01 TASK

Product & Architecture

Journal AI Research

OBJECTIVE

Define and freeze the product and technical architecture for V1 of Journal AI Research.

AI-01 must create the specification required for subsequent AI agents to implement the product.

---

TASK 1 — PRODUCT DEFINITION

Define:

- product purpose
- primary user
- main problem
- core workflow
- V1 scope
- V1 exclusions

The central workflow must remain:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

TASK 2 — USER FLOW

Document the complete user flow:

Open App
↓
Enter Query
↓
Search
↓
Search Scientific Sources
↓
Normalize
↓
Deduplicate
↓
Rank
↓
Display 10 Results
↓
Open Article
↓
Read Available Content
↓
Export PDF

Include error states.

---

TASK 3 — SOURCE ARCHITECTURE

Define integration requirements for:

- PubMed
- Europe PMC
- Crossref

Specify:

- source purpose
- expected metadata
- failure behavior
- normalization requirements
- timeout requirements

Do not invent undocumented API behavior.

---

TASK 4 — SEARCH ENGINE

Define:

- query processing
- parallel source requests
- normalization
- deduplication
- ranking
- top 10 selection

Initial scoring:

Relevance              60%
Recency                30%
Metadata completeness  10%

Explain how these scores should be interpreted.

---

TASK 5 — ARTICLE MODEL

Define the normalized article model.

Minimum fields:

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

Additional fields may be proposed only when justified.

---

TASK 6 — FULL-TEXT POLICY

Define exactly what the application may display.

Must distinguish:

Metadata available
Abstract available
Full text available
Full text unavailable
Original article available

The application must not bypass access restrictions.

---

TASK 7 — PDF

Define PDF behavior.

V1 PDF should support a research report containing:

- search query
- search timestamp
- selected articles
- article metadata
- abstract where legally permitted
- source/original links

Define how full-text article PDFs are handled.

---

TASK 8 — API CONTRACT

Define initial API contracts:

POST /api/v1/search
GET  /api/v1/articles/:id
GET  /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET  /api/v1/health

For each endpoint define:

- purpose
- request
- response
- validation
- errors

Do not write production implementation code.

---

TASK 9 — DATABASE ARCHITECTURE

Define the conceptual data model for:

- articles
- searches
- search_results

Do not over-engineer.

AI-05 will produce the final physical PostgreSQL implementation.

---

TASK 10 — SECURITY

Define:

- input validation
- rate limiting
- timeout
- SQL injection prevention
- secret management
- API key protection
- HTTPS
- logging
- error handling

---

TASK 11 — PERFORMANCE

Define:

- expected search latency
- external API timeout
- parallel requests
- caching strategy
- failure isolation

Initial target:

Approximately <3 seconds under normal external API conditions.

Do not promise a guaranteed response time when third-party APIs are slow.

---

TASK 12 — TEST STRATEGY

Define test categories:

Functional

- search
- result display
- article detail
- full text
- PDF

Search

- Indonesian
- English
- natural language
- keywords
- empty input
- long input

Source failures

- one source unavailable
- multiple sources unavailable
- all sources unavailable

Deduplication

- DOI
- PMID
- PMCID
- title fallback

Security

- malformed input
- injection attempts
- rate limiting
- unauthorized access

---

TASK 13 — ACCEPTANCE CRITERIA

Define measurable V1 acceptance criteria.

Minimum:

1. User can search using one input.
2. System queries configured scientific sources.
3. Results are normalized.
4. Duplicates are removed.
5. Results are ranked.
6. Maximum 10 primary results are shown.
7. Article details are accessible.
8. Legal full text is displayed when available.
9. Restricted full text is not bypassed.
10. Original article links are provided.
11. PDF report can be generated.
12. Source failure does not unnecessarily destroy the entire search.
13. Security requirements are implemented by later engineering stages.

---

TASK 14 — OPEN ISSUES

Create a section called:

"OPEN QUESTIONS"

Only include questions that genuinely require later decisions.

Do not leave basic product requirements undecided.

---

TASK 15 — HANDOFF

Create:

"HANDOFF_AI01.md"

The handoff must be detailed enough for AI-02 to start UX/UI work without redefining the product.

It must contain:

STATUS
OBJECTIVE
COMPLETED
FINAL PRODUCT SPECIFICATION
USER FLOW
SOURCE STRATEGY
SEARCH STRATEGY
RANKING
DEDUPLICATION
ARTICLE MODEL
FULL-TEXT POLICY
PDF POLICY
API CONTRACT
SECURITY
PERFORMANCE
TEST REQUIREMENTS
ACCEPTANCE CRITERIA
DECISIONS
OPEN QUESTIONS
AI-02 INSTRUCTIONS

---

HARD RULES

AI-01 must NOT:

- start frontend implementation
- start backend implementation
- create production database migrations
- introduce unnecessary features
- bypass copyright restrictions
- scrape paywalled content
- claim that ranked articles are the "best"
- make medical recommendations
- redefine the project without documenting the change

---

DEFINITION OF DONE

AI-01 is complete when:

- Product specification is clear.
- V1 scope is frozen.
- User flow is documented.
- Scientific sources are defined.
- Search architecture is defined.
- Ranking is defined.
- Deduplication is defined.
- Full-text rules are defined.
- PDF rules are defined.
- API contracts are defined.
- Security requirements are defined.
- Testing requirements are defined.
- Acceptance criteria are measurable.
- HANDOFF_AI01.md is ready for AI-02.
