JOURNAL AI RESEARCH

AI-06 INTEGRATION SHARED RULES

---

1. SYSTEM

Frontend
   ↓
Go Backend
   ↓
Search Service
   ├── PubMed
   ├── Europe PMC
   └── Crossref
   ↓
Normalize
   ↓
Deduplicate
   ↓
Rank
   ↓
PostgreSQL

Article:

Article
   ↓
Full Text Service
   ↓
Reader

PDF:

Article metadata
   ↓
PDF Generator
   ↓
Research Report

---

2. SOURCE OF TRUTH

Scientific metadata:

PubMed
Europe PMC
Crossref

Database:

Persistence
Cache
Search history
Article retrieval

Frontend:

Presentation
Interaction
User experience

AI:

Optional query understanding

AI bukan source of truth untuk scientific metadata.

---

3. FRONTEND API

Frontend harus menggunakan backend API.

Minimal:

POST /api/v1/search
GET /api/v1/articles/:id
GET /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET /api/v1/health

Contract final harus mengikuti HANDOFF AI-03 dan AI-04.

---

4. API ERROR

Error format harus konsisten.

Concept:

{
  "error": {
    "code": "INVALID_QUERY",
    "message": "Search query is required."
  }
}

Frontend harus memetakan error menjadi user-friendly message.

Jangan menampilkan:

- stack trace;
- SQL error;
- internal path;
- API credentials.

---

5. SEARCH RESPONSE

Expected:

query
results[]
total
sources

Article:

id
title
abstract
authors
journal
publication_date
doi
pmid
pmcid
url
pdf_url
source
is_open_access

Gunakan actual backend contract jika berbeda.

---

6. MAXIMUM RESULTS

Frontend harus menampilkan maksimal:

10

artikel.

Backend tetap menjadi authority untuk result limit.

Frontend tidak boleh mengambil 1000 artikel lalu hanya menyembunyikan sisanya.

---

7. PARTIAL FAILURE

Jika satu scientific provider gagal:

continue

Contoh:

PubMed      success
Europe PMC  timeout
Crossref    success

Frontend tetap menampilkan hasil.

---

8. EMPTY RESULTS

Jika tidak ada artikel:

tampilkan:

No relevant journals found.

atau copy yang sudah ditentukan AI-02.

Jangan membuat dummy article.

---

9. INVALID QUERY

Jika query kosong:

frontend melakukan basic validation.

Backend tetap melakukan validation.

Defense in depth.

---

10. ARTICLE DETAIL

Article detail harus menampilkan:

- title;
- authors;
- journal;
- publication date;
- DOI jika ada;
- PMID jika ada;
- abstract;
- full text state;
- original article link.

---

11. FULL TEXT STATES

Minimal:

AVAILABLE
NOT_AVAILABLE
LOADING
ERROR

Jika unavailable:

jangan membuat reader kosong.

Tampilkan:

Abstract / metadata
+
OPEN ORIGINAL ARTICLE

---

12. ORIGINAL ARTICLE

External link harus:

- valid;
- safe;
- source-aware;
- membuka halaman asli.

Jangan mengubah URL menjadi URL proxy ilegal.

---

13. PDF

PDF export harus menjadi research report.

Minimal:

Journal AI Research

Search Query:
Date:

Selected Articles:
1. Title
   Authors
   Journal
   Date
   DOI
   PMID
   Abstract
   Source

Jika original article URL tersedia:

sertakan link.

---

14. DATABASE

AI-06 harus menggunakan PostgreSQL melalui repository/backend.

Frontend tidak pernah mengakses database secara langsung.

---

15. ENVIRONMENT

Development:

Frontend
Backend
PostgreSQL

dapat dijalankan secara konsisten.

Gunakan:

.env.example

Jangan commit ".env".

---

16. CORS

Development dapat menggunakan allowed origin tertentu.

Production:

jangan menggunakan:

Access-Control-Allow-Origin: *

jika tidak diperlukan.

---

17. NETWORK

Docker Compose harus memungkinkan:

frontend → backend
backend → postgres
backend → scientific APIs

Pastikan service names dan ports konsisten.

---

18. LOGGING

Request flow harus dapat ditelusuri menggunakan request ID jika tersedia.

Jangan log secret.

---

19. TIMEOUT

Timeout harus diterapkan:

browser/API request
backend
external provider
database
PDF generation

sesuai kebutuhan.

---

20. RETRIES

Jangan melakukan unlimited retry.

Retry hanya untuk error yang memang retryable.

Gunakan bounded retry dengan backoff bila diperlukan.

---

21. CACHE

Jika cache telah dibuat AI-04:

integrasikan tanpa mengubah semantics.

Jangan membuat second cache layer tanpa alasan.

---

22. DATABASE PERSISTENCE

AI-06 harus memverifikasi:

article saved
search saved
search results saved
article retrieved

---

23. NO DUPLICATE LOGIC

Jangan membuat ranking kedua di frontend.

Jangan membuat deduplication kedua yang berbeda di frontend.

Backend adalah authority.

---

24. ACCESSIBILITY

Integration tidak boleh merusak:

- keyboard navigation;
- focus state;
- semantic HTML;
- screen reader labels;
- contrast;
- responsive layout.

---

25. RESPONSIVE

Test:

mobile
tablet
desktop

---

26. SECURITY BOUNDARY

Frontend:

NO DATABASE CREDENTIAL
NO SCIENTIFIC API SECRET

Backend:

VALIDATION
RATE LIMIT
TIMEOUT
SAFE URL

Database:

RESTRICTED ACCESS

---

27. NO SCOPE CREEP

Jangan menambahkan:

- login;
- subscription;
- social features;
- citation manager;
- systematic review;
- diagnosis;
- treatment recommendation.

---

28. INTEGRATION PRIORITY

Prioritas:

Correctness
↓
Security
↓
Reliability
↓
Compatibility
↓
Performance
↓
UX polish

---

29. HANDOFF

AI-06 harus membuat:

HANDOFF_AI06.md

Dokumen harus berisi:

- integrated architecture;
- API contract;
- environment;
- Docker;
- end-to-end flow;
- tests;
- known issues;
- AI-07 test instructions.
