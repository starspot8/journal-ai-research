HANDOFF AI-01

Journal AI Research

1. STATUS

COMPLETED

AI-01 has defined the product scope, user flow, source strategy, architecture, ranking and deduplication principles, full-text policy, PDF policy, API contract, conceptual database design, security requirements, performance target, testing strategy, and acceptance criteria. This handoff is intended to be implementation-ready for AI-02 UX/UI.

---

2. OBJECTIVE

AI-01 establishes the product definition and architecture baseline for V1 of Journal AI Research.

The goal is to ensure that AI-02 can design the user experience without redefining the product or broadening the scope. The product must remain aligned to the following core concept:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

The system is a scientific literature search and reading experience, not an editorial recommendation engine and not a medical decision support system.

---

3. PRODUCT DEFINITION

Product name
- Journal AI Research

Product purpose
- Help a user quickly find recent scientific journal articles relevant to a topic or research question.
- Keep the interface simple enough that a user does not need to understand Boolean syntax or source-specific search behavior.
- Return a short, readable list of recent, relevant scientific articles and allow the user to read the article metadata and legally accessible content.

Primary user
- Researcher, doctor, healthcare professional, student, or general user who wants to locate current scientific literature quickly.

Main problem solved
- Searching across multiple scientific sources is fragmented, time-consuming, and difficult for non-expert users.
- The product reduces the friction between a natural-language query and discoverable, recent scientific metadata.

Product principle
- Scientific metadata comes from scientific sources.
- AI may assist with query understanding or interpretation, but it does not replace the underlying bibliographic sources.
- The user should always see the concept: 10 Latest Relevant Journals.

Required product wording
- Use: 10 Latest Relevant Journals
- Do not use: Best Journals, Top Journals, Number One Journal

---

4. SCOPE

IN SCOPE (V1)
- One search box
- Natural-language and simple keyword search
- Search by Enter or Search button
- Integration with PubMed
- Integration with Europe PMC
- Integration with Crossref
- Query processing and source orchestration
- Normalization of metadata from different sources
- Deduplication across sources
- Relevance + recency + metadata completeness scoring
- Display of up to 10 results
- Article detail page
- Display of abstract when available and legally permitted
- Display of full text only when legally available
- Original article/source link
- PDF research report export
- Error handling and graceful fallback
- Search history storage at system level
- Logging and basic observability
- Basic security and rate limiting
- Responsive web interface

OUT OF SCOPE (V1)
- User accounts and authentication
- Social login or profiles
- Citation manager
- Collaboration or sharing features
- Systematic review automation
- Meta-analysis automation
- Diagnosis or treatment recommendation
- Native mobile app
- Recommendation feed
- Paid subscription model
- Unauthorized article downloading or scraping
- Paywall bypass
- AI-generated medical conclusions
- Database implementation details beyond conceptual architecture

Scope rationale
- V1 is intentionally limited to search, read, and export of scientific literature with legal access controls.
- The product is designed to remain safe, legal, and implementation-ready for the next engineering stages.

---

5. USER FLOW

The product user flow is:

1. Open app
2. See a single primary search box
3. Enter a topic or natural-language question
4. Press Search or Enter
5. Search configured scientific sources in parallel
6. Normalize results from all sources
7. Deduplicate overlapping records
8. Score and rank results
9. Display 10 Latest Relevant Journals
10. Select a result to read more
11. View metadata, abstract, and legal full-text content when available
12. Open original source if full text is not available
13. Export a PDF research report

Error states
- No results found
  - Show: "Tidak ditemukan jurnal yang sesuai. Coba kata kunci lain."
- One source fails
  - Continue with remaining sources and return available results
- All sources fail
  - Show: "Pencarian jurnal sedang mengalami gangguan. Silakan coba lagi."
- Article missing metadata or unavailable
  - Show available metadata and original article link
- Empty or malformed query
  - Validate input and return a user-friendly error

---

6. ARCHITECTURE

Conceptual architecture

Browser / User
  ↓
Next.js + TypeScript frontend
  ↓
Go API layer
  ↓
Search orchestration service
  ├── PubMed client
  ├── Europe PMC client
  └── Crossref client
  ↓
Normalizer
  ↓
Deduplicator
  ↓
Ranking engine
  ↓
PostgreSQL / cache / search metadata store
  ↓
API response to frontend
  ↓
UI display: results and article detail

Architecture notes
- The frontend is not responsible for API credentials or source-specific logic.
- The backend owns orchestration, validation, external API calls, scoring, and deduplication.
- Source adapters are independent; failures must be isolated and non-blocking.
- The database stores normalized article records and search history, but the source databases remain authoritative for underlying metadata.

---

7. SCIENTIFIC SOURCE STRATEGY

Primary sources
- PubMed
- Europe PMC
- Crossref

Source responsibilities

PubMed
- Primary biomedical literature discovery
- Metadata source for article identification, title, authors, journal, dates, and identifiers
- Intended for biomedical and research-topic discovery

Europe PMC
- Biomedical literature retrieval and open-access/full-text availability checks where supported
- Useful for legal access validation and cross-linking to original articles
- Must not be used as a bypass mechanism for restricted access

Crossref
- DOI resolution and scholarly metadata enrichment
- Useful for metadata normalization and duplicate detection
- Supports source-level cross-checking of scholarly records

Source integration requirements
- Requests must be performed in parallel where feasible
- Each source must have independent timeout and error handling
- Partial results are acceptable when some sources fail
- Search results must be normalized into a common article schema before ranking

Failure behavior
- A single source outage must not crash the whole search
- The system should continue with other available sources
- If all sources fail, return a search disruption message

Source-specific notes
- No source may be treated as the sole truth for authorship, metadata, or access legality
- Metadata must be traceable to the source record used
- Legal access decisions are based on accessible terms and source permissions, not on bypassing restrictions

---

8. SEARCH STRATEGY

Query processing
- Preserve the user's original query as the first-class intent
- Accept natural-language queries, keywords, and Indonesian/English phrasing
- Use AI support only as a controlled layer for understanding or expansion, not as a replacement for scientific metadata
- Do not silently reinterpret the user’s medical or research intent in a way that changes the actual topic being searched

Query handling principles
- Maintain user intent as far as possible
- Query expansion or source-specific transformation may occur in the backend, but must not alter the user’s core topic beyond a documented and traceable process
- Validation should block empty, abusive, and malformed input

Parallel source model
- Fire PubMed, Europe PMC, and Crossref requests concurrently when available
- Use timeout and cancellation logic
- Merge all successful results into a common normalized article set

Result pipeline
- Search query
- Source search
- Fetch and parse results
- Normalize metadata
- Deduplicate articles
- Score and rank
- Select top 10
- Return to frontend

---

9. NORMALIZATION

Normalization is required because every source formats article metadata differently.

Normalization goals
- Convert source-specific identifiers into canonical fields
- Standardize title, author, journal, date, and source-url fields
- Map missing values consistently
- Preserve raw source metadata internally for traceability

Normalization requirements
- Titles: trim whitespace, normalize spacing, unify casing when needed for matching but retain original display title
- Authors: normalize naming and list formatting
- Publication date: convert to a canonical publication_date field
- Identifiers: map DOI, PMID, PMCID into dedicated fields
- Journal: map source journal names consistently where possible
- URLs: keep the original source URL and any legal full-text/open-access URL when available
- Access status: classify as open access / restricted / unknown based on available metadata and legal capabilities

Normalization must not invent data
- If the source does not provide a field, it remains missing
- No fabricated DOI, PMID, PMCID, or abstract
- No undocumented AI-generated metadata may replace source data

---

10. DEDUPLICATION STRATEGY

Deduplication priority
1. DOI
2. PMID
3. PMCID
4. Normalized title + author + year

Rules
- If two records share the same DOI, treat them as the same article
- If DOI is missing but PMID matches, treat as same article
- If PMID is missing but PMCID matches, treat as same article
- If all identifiers are unavailable, use normalized title + author + publication year as a fallback

Implementation intent
- Keep a single canonical record per article
- Retain multiple source references internally when desired, but only display one article record in the result list
- Preserve a provenance trail or a source-reference list for debugging and traceability

Example
- PubMed and Europe PMC both return the same DOI for the same article
- Only one normalized result should appear in the top 10 list

---

11. RANKING STRATEGY

Baseline ranking model
- Relevance: 60%
- Recency: 30%
- Metadata completeness: 10%

Interpretation
- These are baseline product weights, not a claim of scientific superiority
- Ranking indicates how well the article matches the user query and how current it is, not whether the paper is medically “better” or “top quality”
- Metadata completeness rewards records with sufficient bibliographic data to be useful and trustworthy

Ranking principle
- A higher score means a better match within the search context, not a better study in an editorial or scientific sense

Selection policy
- After deduplication and scoring, select the top 10 results for display
- If fewer than 10 valid records are available, return the available records
- If results are tied, prefer more recent records, then metadata completeness, then source reliability and traceability

---

12. ARTICLE DATA MODEL

The normalized article object must at minimum include:

- id
- doi
- pmid
- pmcid
- title
- abstract
- authors
- journal
- publication_date
- source
- source_url
- fulltext_url
- is_open_access
- created_at
- updated_at

Recommended additional fields (only if justified)
- article_type
- year
- venue
- language
- access_status
- source_references
- raw_metadata
- has_abstract
- has_full_text

Article model intent
- Each article should be canonicalized and consistent across all scientific sources
- The object must support both list/result views and article detail views
- The object must support legal full-text decision logic without exposing illegal access paths

---

13. FULL-TEXT LEGAL ACCESS POLICY

The application may display full text only when access is legally available through an authorized source.

Allowed conditions
- Open-access article
- Publicly accessible full text
- Licensed or authorized source access
- System clearly respects source access rights

Disallowed actions
- Bypass paywalls
- Bypass authentication
- Circumvent access controls
- Scrape restricted content
- Reproduce copyrighted content without authorization
- Create unauthorized copies of publisher PDFs

Display behavior
- If full text is legally available:
  - show the full text or accessible content, if the source permits it
  - display a clear source reference
- If full text is not legally available:
  - show metadata
  - show abstract if available and legally permitted
  - show original article link
  - show “FULL TEXT NOT AVAILABLE” when appropriate

Legal rule
- Source permissions and provider terms take precedence over convenience
- If access is uncertain, the system must err on the side of not displaying the content

---

14. PDF POLICY

V1 PDF export is primarily a research report, not a copy of a restricted publisher article.

The generated research report may contain:
- Search query
- Search date and time
- Selected articles
- Article title
- Authors
- Journal
- Publication date
- DOI
- PMID / PMCID
- Abstract where legally permitted
- Source and original article link
- Summary text describing the article selection

The PDF must clearly distinguish:
- Generated research report
- Original publisher article

Important rules
- The application must not generate unauthorized copies of paywalled journal PDFs
- If a source PDF is legally accessible directly from the publisher or repository, the application may provide a link to that source, but not a copied or bypassed file
- The report must remain a summary/export artifact, not a disguised article PDF copy

---

15. API CONTRACT (INITIAL)

Base path
- /api/v1

Endpoints

1) POST /api/v1/search
Purpose
- Search scientific literature using a user query

Request
- application/json
- body:
  {
    "query": "latest journal about cataract surgery"
  }

Response
- JSON object with normalized results, metadata, and error status if needed
- Response should include up to 10 primary results

Validation
- Reject empty or malformed input
- Validate query length and content constraints
- Reject abusive input patterns

Errors
- 400 invalid query
- 422 unsupported payload shape
- 429 too many requests
- 503 external source unavailable
- 500 internal error

2) GET /api/v1/articles/:id
Purpose
- Return article details for a selected article

Request
- path parameter: article id

Response
- article metadata and available legal access information

Validation
- Ensure article id is valid and exists

Errors
- 404 article not found
- 400 invalid id
- 500 database or retrieval failure

3) GET /api/v1/articles/:id/fulltext
Purpose
- Return legal full-text availability or content if authorized and available

Request
- path parameter: article id

Response
- either:
  - full text content under legal access
  - availability status
  - original article URL
  - metadata-only result

Validation
- Do not expose restricted content unless explicitly legally available

Errors
- 404 article not found
- 403 restricted / not authorized
- 500 processing error

4) POST /api/v1/export/pdf
Purpose
- Generate a PDF research report from selected search results

Request
- article ids and/or search context

Response
- PDF binary or generated file reference

Validation
- Require valid selected item list
- Empty selection should fail gracefully or return a user-friendly message

Errors
- 400 invalid request
- 422 no valid article ids
- 500 PDF generation failure

5) GET /api/v1/health
Purpose
- Health check for the API and dependencies

Response
- service status, dependency status, timestamp

Validation
- No user input required

Errors
- 503 degraded service

---

16. CONCEPTUAL DATABASE DESIGN

Conceptual tables

articles
- id
- doi
- pmid
- pmcid
- title
- abstract
- authors
- journal
- publication_date
- source
- source_url
- fulltext_url
- is_open_access
- created_at
- updated_at

searches
- id
- query
- created_at
- result_count

search_results
- id
- search_id
- article_id
- rank
- relevance_score
- recency_score
- final_score

Design principles
- Keep the article table canonicalized to a single source-of-truth record per article
- Search history is system-level metadata only; it does not require a user account in V1
- search_results stores ranking decisions for traceability and later performance analysis
- AI-05 will finalise the exact PostgreSQL physical schema, indexes, constraints, and migrations

---

17. SECURITY REQUIREMENTS

Baseline security
- HTTPS in production
- Input validation on all API endpoints
- Rate limiting
- Request timeouts
- Secure HTTP headers
- Parameterized SQL and safe query patterns
- Secret management through environment variables or secret stores
- No API keys or secrets in frontend code
- No secrets committed to Git repositories
- Structured logging
- Error messages must not expose secrets
- Dependency updates and vulnerability reviewing
- Basic audit logging for source access and PDF export actions

Special security rules
- API credentials from scientific sources must remain server-side only
- User queries may be logged for technical and product analysis only in accordance with privacy policy
- No sensitive user data is required in V1
- No user personal data like email, phone, address, or medical record should be collected unless a future requirement explicitly introduces it

---

18. PERFORMANCE REQUIREMENTS

Target behavior
- A normal search should typically return within approximately 3 seconds when external scientific APIs behave normally
- This is an operational target, not a guaranteed SLA for third-party services

Performance strategies
- Parallel source requests
- Source-specific timeouts
- Request concurrency control
- Response caching for repeated searches when allowed by source terms
- Connection reuse
- Efficient database queries and moderate result counts
- Failure isolation so one service outage does not block all search results

Performance principles
- The system must be responsive enough for a simple search flow
- Third-party API latency is variable and cannot be fully controlled by the application
- Cache strategy must not violate source terms or legal access restrictions

---

19. TESTING REQUIREMENTS

Functional tests
- Search with natural-language queries
- Search with keywords
- Indonesian and English queries
- Empty input
- Long input and special characters
- Result display
- Article detail view
- Full-text decision logic
- PDF generation

Source failure tests
- One source unavailable
- Multiple sources unavailable
- All sources unavailable
- Partial result success

Deduplication tests
- Same DOI
- Same PMID
- Same PMCID
- Same title fallback

Ranking tests
- Relevance weighting
- Recency weighting
- Metadata completeness handling
- Results with missing metadata

Security tests
- Malformed input
- SQL injection attempts
- Rate limiting checks
- Unauthorized API access
- Secret leakage prevention

PDF tests
- Valid PDF generation
- Empty selection
- Long titles and abstracts
- Multiple selected articles

---

20. ACCEPTANCE CRITERIA

AI-01 is complete when all of the following are true:

1. User can search using one input field
2. System queries configured scientific sources
3. Results are normalized across data sources
4. Duplicates are removed according to the documented priority
5. Results are ranked using the agreed baseline weights
6. Maximum 10 primary results are shown
7. Article detail pages are available
8. Legal full text is displayed when available
9. Restricted full text is not bypassed
10. Original article links are provided when appropriate
11. PDF research report export is defined and supported
12. Source failure does not unnecessarily collapse the entire search
13. Security requirements are defined for later engineering stages
14. AI-02 can start UX/UI work without redefining the product

---

21. DECISIONS MADE

DECISION-001
- V1 product scope is limited to simple scientific journal search and reading, not a broader research platform.

DECISION-002
- The core workflow remains: KETIK → CARI → 10 JURNAL → BACA → PDF.

DECISION-003
- The product uses PubMed, Europe PMC, and Crossref as the initial scientific data sources.

DECISION-004
- Results are labeled as 10 Latest Relevant Journals rather than Best Journals or Top Journals.

DECISION-005
- Ranking is a relevance/recency/metadata completeness model: 60/30/10 baseline.

DECISION-006
- Deduplication priority is DOI > PMID > PMCID > normalized title + author + year.

DECISION-007
- Full-text content is shown only if legally available; no paywall bypass or scraping is allowed.

DECISION-008
- The V1 PDF is a generated research report, not a copy of restricted publisher PDFs.

DECISION-009
- API contract for V1 is limited to search, article detail, fulltext availability, PDF export, and health.

DECISION-010
- Database design remains conceptual in AI-01; AI-05 will finalize the PostgreSQL implementation.

---

22. KNOWN LIMITATIONS

- Search quality depends on the quality of the underlying scientific APIs and the normalization layer.
- The system does not guarantee official scientific review or editorial curation beyond source metadata.
- Recency and relevance ranking are not claims about clinical quality or journal prestige.
- Legal access decisions must be handled conservatively because access terms vary by source and publisher.
- PDF export is report-oriented rather than article-copy production.
- The product intentionally excludes user authentication, social features, and advanced research workflows.

---

23. OPEN DECISIONS

OPEN DECISION-001
- Exact source-specific query transformation rules for PubMed/Europe PMC/Crossref are not yet finalized.
- Why it matters: search quality and result recall depend on query interpretation strategy.
- Suggested owner: AI-04 backend + search stage.
- Required decision before: implementation of search service.

OPEN DECISION-002
- Final PDF report layout and template are not yet specified beyond the report’s content requirements.
- Why it matters: UX and export experience depend on layout decisions and multi-article formatting.
- Suggested owner: AI-02 UX/UI + AI-03 frontend + AI-08 deployment.
- Required decision before: PDF export UX and front-end implementation.

OPEN DECISION-003
- Search history retention policy is not yet finalized.
- Why it matters: it affects compliance, privacy, and storage volume.
- Suggested owner: AI-05 database + AI-07 testing/security.
- Required decision before: production storage specification.

OPEN DECISION-004
- Final legal validation strategy for full-text access across all sources is not yet finalized in operational detail.
- Why it matters: it affects source adapter behavior and content-rendering rules.
- Suggested owner: AI-04 backend + AI-07 security/testing.
- Required decision before: production full-text API implementation.

---

24. EXPLICIT INSTRUCTIONS FOR AI-02

AI-02 is responsible for UX/UI design for V1.

AI-02 must:
1. Read SHARED.md and this handoff before designing the product experience.
2. Follow the frozen V1 scope and do not expand the feature set.
3. Design the simplest possible user flow: one search box → results → article detail → PDF export.
4. Use the exact product language: 10 Latest Relevant Journals.
5. Avoid using any language indicating editorial ranking or quality: Best Journals, Top Journals, Recommended Journals.
6. Design clear empty states, loading states, and error states.
7. Create a clean results page and article detail page that support metadata, abstract, source link, and legal full-text status.
8. Design the PDF export interaction as a research report export, not a paywalled article copy.
9. Keep the user experience focused on search, reading, and export without social or personalization features.
10. Produce a UX/UI handoff for AI-03 that remains aligned with the architectural and product decisions above.

AI-02 must not:
- redefine the product scope
- add login or account flows
- introduce social features
- bypass paywalls or copyrighted full texts
- redefine the scientific source strategy
- change the ranking interpretation from relevance/recency/metadata completeness
- label the ranking as a quality ranking

---

25. NEXT STAGE

Next AI stage:
- AI-02 — UX/UI

Expected handoff chain:
- AI-02 should produce a UX/UI specification for AI-03 Frontend based on the decisions in this handoff.

This handoff is intended to be the definitive product and architecture baseline for all subsequent stages.
