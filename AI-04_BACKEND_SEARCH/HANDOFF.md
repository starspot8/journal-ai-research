AI-04 — HANDOFF

BACKEND + SCIENTIFIC SEARCH

«Dokumen ini awalnya merupakan template handoff.
Setelah AI-04 selesai bekerja, dokumen ini WAJIB diperbarui menjadi laporan implementasi aktual.»

---

1. IMPLEMENTATION STATUS

Status:

NOT STARTED

Setelah implementasi:

COMPLETE

atau:

PARTIAL

Jika partial, jelaskan bagian yang belum selesai.

---

2. PROJECT STRUCTURE

Dokumentasikan struktur backend aktual.

Contoh:

backend/
├── cmd/
├── internal/
│   ├── search/
│   ├── ranking/
│   ├── articles/
│   ├── pubmed/
│   ├── europepmc/
│   ├── crossref/
│   ├── fulltext/
│   └── ...
├── tests/
├── go.mod
└── README.md

Jangan menulis struktur hipotetis setelah implementasi.

---

3. SERVER

Dokumentasikan:

Go version:
HTTP framework/router:
Server port:
Graceful shutdown:
Request timeout:

---

4. API IMPLEMENTED

Health

GET /api/v1/health

Status:

IMPLEMENTED / NOT IMPLEMENTED

---

Search

POST /api/v1/search

Status:

IMPLEMENTED / NOT IMPLEMENTED

Request:

{
  "query": "..."
}

Response aktual:

PASTE ACTUAL CONTRACT HERE

---

Article

GET /api/v1/articles/:id

Status:

IMPLEMENTED / NOT IMPLEMENTED

---

Full Text

GET /api/v1/articles/:id/fulltext

Status:

IMPLEMENTED / NOT IMPLEMENTED

---

5. SCIENTIFIC PROVIDERS

PubMed

Status:

IMPLEMENTED / NOT IMPLEMENTED

Implementation:

...

---

Europe PMC

Status:

IMPLEMENTED / NOT IMPLEMENTED

Implementation:

...

---

Crossref

Status:

IMPLEMENTED / NOT IMPLEMENTED

Implementation:

...

---

6. SEARCH PIPELINE

Actual pipeline:

Request
  ↓
Validation
  ↓
Query processing
  ↓
Parallel providers
  ↓
Normalization
  ↓
Merge
  ↓
Deduplication
  ↓
Ranking
  ↓
Limit 10
  ↓
Response

Jika berbeda dari architecture awal, jelaskan:

Deviation:
Reason:
Impact:

---

7. PROVIDER FAILURE

Dokumentasikan behavior ketika:

PubMed fails
Europe PMC fails
Crossref fails
All providers fail

Contoh:

PubMed timeout:
Europe PMC + Crossref results tetap digunakan.

---

8. NORMALIZED ARTICLE MODEL

Dokumentasikan model aktual.

Contoh:

ID
Title
Abstract
Authors
Journal
PublicationDate
DOI
PMID
PMCID
URL
PDFURL
Source
IsOpenAccess

Jika terdapat field tambahan:

jelaskan alasannya.

---

9. DEDUPLICATION

Urutan actual:

1. DOI
2. PMID
3. PMCID
4. Title + Author + Year

Dokumentasikan:

- DOI normalization;
- title normalization;
- merge policy;
- duplicate conflict resolution.

---

10. RANKING

Dokumentasikan formula aktual.

Target awal:

Relevance       60%
Recency         30%
Metadata        10%

Jika formula berubah:

Old:
New:
Reason:

Jangan mengubah ranking tanpa alasan yang terdokumentasi.

---

11. RELEVANCE

Dokumentasikan bagaimana relevance dihitung.

Contoh:

Title match
Abstract match
Term coverage
Keyword match

Jelaskan apakah menggunakan:

keyword scoring
BM25
embedding
LLM
other

Jika masih menggunakan simple scoring:

nyatakan secara eksplisit.

---

12. RECENCY

Dokumentasikan:

- date field yang digunakan;
- scoring formula;
- handling missing date;
- handling invalid/future date.

---

13. METADATA COMPLETENESS

Dokumentasikan field yang digunakan untuk completeness score.

Contoh:

Title
Authors
Journal
Publication Date
DOI
Abstract
URL

---

14. FULL TEXT

Dokumentasikan:

Legal access policy:
Open-access sources:
HTML support:
PDF support:
Fallback behavior:

Pastikan tidak ada bypass paywall.

---

15. SECURITY

Checklist:

[ ] Input validation
[ ] Request timeout
[ ] External timeout
[ ] Rate limiting
[ ] SSRF protection
[ ] Secure external URL handling
[ ] No frontend secrets
[ ] No API keys in source
[ ] No stack traces exposed
[ ] Structured logging

---

16. CONFIGURATION

Dokumentasikan environment variables aktual.

Contoh:

PORT=
DATABASE_URL=
HTTP_TIMEOUT=
SEARCH_TIMEOUT=
RATE_LIMIT=
ENVIRONMENT=

Jangan memasukkan nilai secret sebenarnya.

---

17. DATABASE INTERFACES

AI-04 harus mendokumentasikan interface yang dibutuhkan AI-05.

Contoh:

type ArticleRepository interface {
    GetByID(ctx context.Context, id string) (*Article, error)
    Save(ctx context.Context, article *Article) error
}

Dokumentasikan interface aktual.

---

18. DATABASE REQUIREMENTS FOR AI-05

AI-05 harus mengetahui:

Required entities

articles
searches
search_results

Jika terdapat entity tambahan, dokumentasikan.

Required identifiers

article_id
doi
pmid
pmcid

Required indexes

Dokumentasikan kebutuhan index aktual.

Contoh:

articles.doi
articles.pmid
articles.pmcid
articles.publication_date

Required uniqueness

Contoh:

normalized DOI
PMID
PMCID

Jangan menganggap semua field selalu unique jika source behavior tidak menjaminnya.

---

19. CACHE REQUIREMENTS

Dokumentasikan cache yang diperlukan:

search cache
article cache
source response cache

Untuk setiap cache:

key
TTL
invalidation

---

20. TESTS

Dokumentasikan command:

go test ./...

Status:

PASS / FAIL

Jika fail:

Test:
Reason:
Impact:

---

21. TEST COVERAGE

Dokumentasikan minimal coverage area:

Provider
Normalization
Deduplication
Ranking
Validation
Partial failure
Timeout
Article retrieval
Full text

---

22. PERFORMANCE

Dokumentasikan:

Average local search latency:
Provider timeout:
Search timeout:
Concurrency:
Cache:

Jangan membuat klaim latency external provider yang tidak dapat diverifikasi.

---

23. KNOWN LIMITATIONS

Tuliskan keterbatasan aktual.

Contoh:

- Crossref may not provide abstract.
- Some articles do not have open full text.
- External provider latency varies.
- Search relevance is currently lexical.

---

24. TECHNICAL DECISIONS

Catat keputusan penting:

Decision:
Reason:
Alternatives considered:
Impact:

Contoh:

Decision:
Use provider interface instead of direct source calls from handlers.

Reason:
Keeps source-specific implementation isolated.

Impact:
Easier testing and future provider additions.

---

25. DEVIATIONS FROM AI-01 / AI-02 / AI-03

Jika tidak ada:

No material deviations.

Jika ada:

Deviation:
Reason:
Impact:
Approved by:

Jangan menyembunyikan perubahan contract.

---

26. FRONTEND INTEGRATION NOTES

AI-03 perlu mengetahui:

Search endpoint:
Request shape:
Response shape:
Error shape:
Article endpoint:
Fulltext endpoint:
Loading expectations:
Partial failure behavior:

Jika frontend perlu perubahan:

jelaskan secara eksplisit.

---

27. AI-05 DATABASE HANDOFF

AI-05 harus membaca:

This HANDOFF

dan mengimplementasikan:

PostgreSQL schema
migrations
indexes
constraints
repository implementations
persistence

AI-05 tidak boleh mengubah search ranking tanpa alasan dan dokumentasi.

---

28. AI-06 INTEGRATION HANDOFF

AI-06 berikutnya harus fokus pada:

Frontend + Backend integration
Database integration
Full end-to-end flow
PDF integration
Environment configuration
Error state integration

---

29. FINAL STATUS

AI-04:

[ ] Backend implemented
[ ] Search implemented
[ ] PubMed implemented
[ ] Europe PMC implemented
[ ] Crossref implemented
[ ] Normalization implemented
[ ] Deduplication implemented
[ ] Ranking implemented
[ ] Article endpoint implemented
[ ] Fulltext endpoint implemented
[ ] Tests passing
[ ] Documentation complete
[ ] AI-05 handoff ready

---

30. FINAL MESSAGE TO NEXT AI

Setelah AI-04 selesai, tuliskan ringkasan nyata:

AI-04 completed the backend and scientific search layer.

The backend now:
- ...
- ...
- ...

AI-05 should now:
- ...
- ...
- ...

Known issues:
- ...

Jangan menggunakan placeholder setelah handoff final diserahkan.
