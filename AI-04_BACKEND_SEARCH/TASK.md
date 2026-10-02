AI-04 — BACKEND + SCIENTIFIC SEARCH

TASK

1. PRIMARY OBJECTIVE

Implement backend untuk Journal AI Research menggunakan Go.

Backend harus menjadi penghubung antara frontend dan scientific data sources.

---

2. READ FIRST

Wajib membaca:

AI-01_PRODUCT_ARCHITECTURE/HANDOFF.md
AI-01_PRODUCT_ARCHITECTURE/SHARED.md

AI-02_UX_UI/HANDOFF.md
AI-02_UX_UI/SHARED.md

AI-03_FRONTEND/HANDOFF.md
AI-03_FRONTEND/SHARED.md

Setelah membaca, buat catatan:

Architecture assumptions
Frontend contract
Backend requirements
Known constraints

Jangan langsung coding sebelum memahami handoff.

---

3. BACKEND STRUCTURE

Gunakan struktur yang maintainable.

Target:

backend/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── search/
│   ├── ranking/
│   ├── articles/
│   ├── pubmed/
│   ├── europepmc/
│   ├── crossref/
│   ├── fulltext/
│   ├── cache/
│   ├── http/
│   └── config/
│
├── migrations/
│
├── tests/
│
├── go.mod
└── README.md

Jangan memaksakan struktur di atas apabila existing repository memiliki struktur yang lebih baik.

---

4. CONFIGURATION

Configuration harus berasal dari environment/configuration layer.

Minimal:

PORT
DATABASE_URL
HTTP_TIMEOUT
SEARCH_TIMEOUT
RATE_LIMIT
ENVIRONMENT

Jika source tertentu membutuhkan configuration tambahan, dokumentasikan dalam ".env.example".

Jangan menyimpan secret di source code.

---

5. HTTP SERVER

Implement HTTP server dengan:

- routing;
- JSON request;
- JSON response;
- graceful shutdown;
- request context;
- timeout;
- logging;
- recovery middleware.

Server harus berhenti secara graceful ketika menerima shutdown signal.

---

6. HEALTH ENDPOINT

Implement:

GET /api/v1/health

Response minimal:

{
  "status": "ok"
}

Health endpoint tidak boleh membutuhkan search API.

Jika database health check ditambahkan, bedakan:

application health
dependency health

Jangan membuat health endpoint lambat karena external scientific API.

---

7. SEARCH ENDPOINT

Implement:

POST /api/v1/search

Request:

{
  "query": "latest journal about cataract surgery"
}

Validation:

- query wajib;
- trim whitespace;
- reject empty query;
- enforce maximum length.

---

8. SEARCH SERVICE

Buat abstraction:

type SearchProvider interface {
    Search(ctx context.Context, query string) ([]Article, error)
}

Source:

PubMedProvider
EuropePMCProvider
CrossrefProvider

Search service:

Receive query
    ↓
Validate
    ↓
Parallel provider search
    ↓
Collect results
    ↓
Normalize
    ↓
Deduplicate
    ↓
Score
    ↓
Sort
    ↓
Take first 10
    ↓
Return

---

9. PUBMED

Implement adapter untuk PubMed official APIs.

Adapter harus:

- menerima context;
- memiliki timeout;
- menangani HTTP errors;
- menangani malformed response;
- parse metadata;
- map metadata ke normalized Article;
- tidak membuat asumsi metadata yang tidak tersedia.

Gunakan source identifier:

source = "pubmed"

Simpan:

PMID
DOI
PMCID jika tersedia

---

10. EUROPE PMC

Implement adapter untuk Europe PMC.

Map metadata ke model yang sama.

Gunakan:

source = "europepmc"

Jika Europe PMC memberikan informasi open access/full text, simpan informasi tersebut sesuai model.

---

11. CROSSREF

Implement adapter Crossref.

Crossref terutama digunakan sebagai complementary metadata source.

Gunakan:

source = "crossref"

Map:

title
authors
journal
publication date
DOI
URL

Jangan menganggap Crossref selalu memiliki abstract atau full text.

---

12. SOURCE INTERFACE

Semua provider harus mengikuti interface yang sama.

Contoh:

type SearchProvider interface {
    Name() string
    Search(ctx context.Context, query string) ([]Article, error)
}

Jika repository membutuhkan provider metadata terpisah, buat interface yang jelas.

---

13. PARALLEL SEARCH

Source calls harus dijalankan secara parallel.

Concept:

              ┌── PubMed
              │
QUERY ────────┼── Europe PMC
              │
              └── Crossref

Jika:

PubMed      success
Europe PMC  timeout
Crossref    success

hasil dari PubMed + Crossref tetap dikembalikan.

---

14. SOURCE STATUS

Track status setiap provider:

success
failed
timeout

Jika contract frontend mendukung source status, expose status tersebut.

Jangan expose internal stack trace.

---

15. NORMALIZATION

Semua source harus menghasilkan:

Article

Model minimal:

type Article struct {
    ID              string
    Title           string
    Abstract        string
    Authors         []Author
    Journal         string
    PublicationDate *time.Time
    DOI             string
    PMID            string
    PMCID           string
    URL             string
    PDFURL          string
    Source          string
    IsOpenAccess    bool
}

Sesuaikan dengan contract AI-03 apabila diperlukan.

---

16. AUTHOR MODEL

Gunakan model terstruktur.

Contoh:

type Author struct {
    GivenName  string
    FamilyName string
    FullName   string
}

Jangan menyimpan author hanya sebagai string apabila source menyediakan struktur yang lebih baik.

---

17. DOI NORMALIZATION

Implement helper:

normalizeDOI()

Handle:

doi:10.xxxx/abc
DOI:10.xxxx/abc
https://doi.org/10.xxxx/abc
http://doi.org/10.xxxx/abc
10.xxxx/ABC

Tujuan:

same DOI → same deduplication key

---

18. TITLE NORMALIZATION

Implement:

normalizeTitle()

Minimal:

- lowercase;
- trim;
- normalize whitespace;
- remove irrelevant punctuation;
- Unicode normalization bila diperlukan.

Jangan mengubah title asli yang ditampilkan kepada user.

Simpan:

original title
normalized title

jika diperlukan.

---

19. DEDUPLICATION

Priority:

DOI
↓
PMID
↓
PMCID
↓
title + author + year

Jika dua artikel memiliki DOI sama:

duplicate

Jika tidak ada DOI tetapi PMID sama:

duplicate

Jika identifier tidak tersedia:

gunakan fallback:

normalized title
+
primary author
+
publication year

---

20. MERGING DUPLICATES

Jika duplicate berasal dari dua source, jangan membuang seluruh metadata.

Contoh:

PubMed:
PMID + abstract

Crossref:
DOI + publisher URL

Gabungkan field yang valid.

Prioritaskan source metadata yang lebih authoritative untuk field tertentu.

Documentasikan merge policy.

---

21. RANKING

Implement ranking service.

Initial weighting:

Relevance = 60%
Recency = 30%
Metadata completeness = 10%

Formula:

final =
    relevance * 0.60
    + recency * 0.30
    + completeness * 0.10

Semua score harus berada pada range yang konsisten, misalnya:

0.0 → 1.0

---

22. RELEVANCE

Initial relevance dapat menggunakan kombinasi:

title match
abstract match
keyword match
query term coverage

Jangan mengklaim bahwa relevance score adalah scientific quality score.

---

23. RECENCY

Gunakan publication date.

Artikel yang lebih baru mendapat recency score lebih tinggi.

Pastikan:

- tanggal invalid tidak membuat panic;
- missing publication date ditangani;
- future date ditangani defensively.

---

24. METADATA COMPLETENESS

Contoh field:

title
authors
journal
publication date
DOI
abstract
URL

Semakin lengkap metadata valid, semakin tinggi completeness score.

Jangan menganggap missing DOI berarti artikel tidak valid.

---

25. RESULT LIMIT

Backend mengembalikan maksimal:

10 articles

Jika hanya terdapat:

4 valid articles

kembalikan 4.

Jangan membuat artikel palsu untuk mencapai 10.

---

26. ARTICLE ENDPOINT

Implement:

GET /api/v1/articles/:id

Return normalized article.

Jika article tidak ditemukan:

404

Gunakan error contract yang konsisten.

---

27. FULL TEXT ENDPOINT

Implement:

GET /api/v1/articles/:id/fulltext

Endpoint harus:

1. menentukan apakah full text tersedia;
2. mengecek apakah akses legal;
3. mengembalikan full text jika tersedia secara sah;
4. jika tidak tersedia, mengembalikan metadata/availability state;
5. menyediakan original article URL bila ada.

Jangan bypass paywall.

---

28. FULL TEXT RESPONSE

Concept:

{
  "available": true,
  "format": "html",
  "content": "...",
  "source_url": "..."
}

Jika tidak tersedia:

{
  "available": false,
  "source_url": "...",
  "reason": "FULL_TEXT_NOT_AVAILABLE"
}

Contract final harus diselaraskan dengan AI-03.

---

29. URL SECURITY

External URLs harus divalidasi.

Jangan menerima URL arbitrer dari client lalu server melakukan fetch.

Gunakan allowlist/domain policy untuk source yang dipercaya.

Perhatikan SSRF.

---

30. CACHING

Design backend agar caching dapat ditambahkan.

Potential cache:

search result
article metadata
source response

Cache harus memiliki TTL.

Jangan cache indefinitely.

---

31. RATE LIMITING

Search endpoint harus memiliki rate limiting.

Tujuan:

- melindungi backend;
- mencegah abuse;
- melindungi external API quota.

Configuration harus dapat diubah.

---

32. ERROR HANDLING

Gunakan typed/application errors jika sesuai.

Contoh:

INVALID_QUERY
ARTICLE_NOT_FOUND
SOURCE_TIMEOUT
SOURCE_UNAVAILABLE
FULL_TEXT_NOT_AVAILABLE
INTERNAL_ERROR

Client tidak boleh menerima:

database stack trace
Go panic
API key
internal filesystem path

---

33. TESTING

Buat unit test dan integration-style test menggunakan mock HTTP server.

Minimum:

Validation

empty query
whitespace query
too long query
valid query

Provider

success
HTTP 4xx
HTTP 5xx
timeout
malformed JSON/XML
empty result

Search service

all providers success
one provider failure
two providers failure
all providers failure

Deduplication

same DOI
same PMID
same PMCID
same title/author/year
different article

Ranking

relevance
recency
completeness
combined score

Article

found
not found

Full text

available
unavailable

---

34. EXTERNAL API TESTS

Jangan membuat seluruh test suite bergantung pada internet.

Gunakan mock server.

Contoh:

httptest.Server

Integration test terhadap live APIs boleh dipisahkan.

---

35. OBSERVABILITY

Implement structured logs minimal:

request_id
method
path
duration
status
provider
provider_duration
result_count
error

Jangan log query jika query dapat mengandung sensitive user information tanpa alasan yang jelas.

---

36. GRACEFUL SHUTDOWN

Server harus:

1. menerima SIGINT/SIGTERM;
2. stop accepting new requests;
3. menunggu request aktif dalam timeout tertentu;
4. menutup HTTP client/resource;
5. exit dengan cleanly.

---

37. DOCUMENTATION

Buat:

backend/README.md

Berisi:

- cara menjalankan;
- environment variables;
- API endpoints;
- request examples;
- response examples;
- testing;
- architecture overview.

---

38. DATABASE BOUNDARY

AI-04 tidak membuat:

CREATE TABLE
ALTER TABLE
migration
index migration

kecuali AI-01/AI-03 handoff secara eksplisit membutuhkan sesuatu yang sangat spesifik.

AI-04 hanya mendefinisikan repository interfaces.

AI-05 akan mengimplementasikan PostgreSQL.

---

39. FRONTEND COMPATIBILITY

AI-04 wajib memastikan response contract kompatibel dengan frontend AI-03.

Jika ada perbedaan:

1. cek "HANDOFF_AI03.md";
2. pertahankan contract frontend jika masuk akal;
3. jika perlu perubahan, dokumentasikan;
4. jangan diam-diam mengubah contract.

---

40. ACCEPTANCE CRITERIA

AI-04 diterima jika:

[ ] Backend starts
[ ] Health endpoint works
[ ] Search endpoint works
[ ] Query validation works
[ ] PubMed adapter works
[ ] Europe PMC adapter works
[ ] Crossref adapter works
[ ] Providers execute independently
[ ] Timeout works
[ ] Partial failure works
[ ] Normalization works
[ ] Deduplication works
[ ] Ranking works
[ ] Maximum 10 results
[ ] Article endpoint works
[ ] Fulltext availability works
[ ] Rate limiting exists
[ ] External URLs are validated
[ ] No secrets in source
[ ] Tests pass
[ ] Documentation exists
[ ] HANDOFF_AI04.md exists

---

41. HANDOFF TO AI-05

AI-04 harus menyerahkan kepada AI-05:

Article model
Search result model
ArticleRepository interface
SearchRepository requirements
Cache requirements
Database persistence requirements
Indexes needed
Unique constraints needed

AI-05 kemudian mengimplementasikan PostgreSQL.

---

42. FINAL RULE

Jangan memperluas produk.

Jika menemukan ide fitur baru:

DOCUMENT AS FUTURE IDEA

jangan implementasikan.

Prioritas:

CORRECTNESS
SECURITY
RELIABILITY
MAINTAINABILITY
TESTABILITY

bukan jumlah fitur.
