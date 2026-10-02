JOURNAL AI RESEARCH

SHARED PROJECT RULES

Dokumen ini berlaku untuk AI-04 dan AI berikutnya.

---

1. PRODUCT

Nama proyek:

Journal AI Research

Konsep:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

Tujuan:

Membantu pengguna menemukan artikel ilmiah terbaru dan relevan melalui satu search box sederhana.

---

2. USER FLOW

USER
  ↓
SEARCH BOX
  ↓
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
10 LATEST RELEVANT JOURNALS
  ↓
ARTICLE DETAIL
  ↓
FULL TEXT / ABSTRACT
  ↓
ORIGINAL ARTICLE
  ↓
PDF REPORT

---

3. SEARCH SOURCES

Primary sources:

1. PubMed
2. Europe PMC
3. Crossref

Source digunakan untuk mendapatkan metadata artikel ilmiah.

Backend harus memperlakukan source secara independent.

Jika satu source gagal:

PubMed       SUCCESS
Europe PMC   FAILED
Crossref     SUCCESS

search tetap menghasilkan response berdasarkan source yang berhasil.

---

4. SEARCH PRINCIPLE

User dapat memasukkan query natural language atau keyword sederhana.

Contoh:

latest journal about cataract surgery

jurnal terbaru tentang diabetic retinopathy

latest research about multifocal IOL

Saya ingin penelitian terbaru tentang hasil multifocal IOL pada pasien usia lanjut

Backend harus dapat memproses query tersebut tanpa mengharuskan user memahami syntax database.

---

5. RANKING

Initial ranking model:

Relevance       60%
Recency         30%
Metadata        10%

Formula konseptual:

final_score =
    relevance_score * 0.60
  + recency_score   * 0.30
  + metadata_score  * 0.10

Ranking bukan penilaian kualitas ilmiah.

Jangan menyebut hasil:

- Best Journal
- Best Article
- Top Journal

Gunakan:

10 Latest Relevant Journals

---

6. DEDUPLICATION

Urutan identifikasi duplicate:

1. DOI
2. PMID
3. PMCID
4. normalized title + author + year

DOI harus dinormalisasi sebelum comparison.

Contoh:

https://doi.org/10.xxxx/ABC
doi:10.xxxx/abc
10.xxxx/ABC

dapat dianggap DOI yang sama setelah normalisasi case dan prefix.

---

7. ARTICLE MODEL

Minimal normalized article:

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

Jika metadata tidak tersedia:

gunakan null / empty value sesuai contract.

Jangan mengarang data.

---

8. ARTICLE ID

Internal article ID harus stabil dan tidak bergantung hanya pada array index.

Contoh:

article:<hash>

atau UUID.

Namun source identifiers tetap harus disimpan:

DOI
PMID
PMCID

---

9. EXTERNAL SOURCE FAILURE

Contoh:

PubMed timeout
Europe PMC success
Crossref success

Response tetap valid.

Response sebaiknya dapat memberikan informasi source status jika diperlukan oleh frontend.

Jangan mengembalikan HTTP 500 hanya karena satu external source gagal.

HTTP 500 digunakan apabila backend sendiri mengalami failure yang membuat request tidak dapat diproses.

---

10. TIMEOUT

Semua external HTTP request harus memiliki timeout.

Jangan menggunakan:

http.Get(...)

tanpa context timeout yang sesuai.

Gunakan context-aware request.

Timeout harus configurable melalui environment/config.

---

11. CONCURRENCY

External source search dapat dijalankan secara parallel.

Contoh:

             ┌── PubMed
QUERY ───────┼── Europe PMC
             └── Crossref

Satu source lambat tidak boleh secara tidak perlu menahan source lain.

Gunakan:

- context;
- goroutine;
- errgroup atau equivalent;
- bounded concurrency jika diperlukan.

---

12. HTTP API

Minimum API:

POST /api/v1/search
GET /api/v1/articles/:id
GET /api/v1/articles/:id/fulltext
GET /api/v1/health

PDF export dapat memiliki interface yang disiapkan untuk tahap berikutnya jika implementasi final berada pada integration stage.

---

13. SEARCH REQUEST

Conceptual request:

{
  "query": "latest journal about cataract surgery"
}

---

14. SEARCH RESPONSE

Conceptual response:

{
  "query": "latest journal about cataract surgery",
  "results": [
    {
      "id": "article:...",
      "title": "...",
      "abstract": "...",
      "authors": [],
      "journal": "...",
      "publication_date": "...",
      "doi": "...",
      "pmid": "...",
      "pmcid": "...",
      "url": "...",
      "pdf_url": "...",
      "source": "...",
      "is_open_access": true
    }
  ],
  "total": 10,
  "sources": {
    "pubmed": "success",
    "europepmc": "success",
    "crossref": "success"
  }
}

Contract boleh disesuaikan dengan hasil AI-03 jika handoff frontend sudah menetapkan format berbeda.

---

15. ERROR FORMAT

Gunakan format error yang konsisten.

Contoh:

{
  "error": {
    "code": "INVALID_QUERY",
    "message": "Search query is required."
  }
}

Jangan mengirim stack trace ke client.

---

16. VALIDATION

Minimal:

- query wajib;
- query tidak boleh kosong;
- panjang query memiliki batas;
- malformed request ditolak;
- content type diperiksa.

Jangan menerima input tanpa batas.

---

17. SECURITY

Backend harus:

- menggunakan HTTPS di production;
- melakukan input validation;
- menggunakan timeout;
- menggunakan secure headers jika relevan;
- tidak menyimpan secret di source code;
- menggunakan environment variables;
- melakukan rate limiting;
- tidak mengekspos internal errors;
- melakukan structured logging;
- memvalidasi external URL sebelum digunakan;
- menghindari SSRF;
- tidak mengeksekusi URL user secara bebas.

---

18. API KEYS

Secret/API credentials jika dibutuhkan:

BACKEND ONLY

Jangan:

NEXT_PUBLIC_...

Jangan pernah memasukkan secret:

- GitHub repository;
- frontend bundle;
- browser;
- JSON response.

---

19. DATABASE

AI-04 tidak memiliki ownership atas physical database schema.

AI-05 akan membuat:

PostgreSQL
migrations
indexes
constraints
repository implementation

AI-04 boleh membuat interface:

type ArticleRepository interface {
    GetByID(...)
    Save(...)
}

Tetapi jangan membuat migration PostgreSQL kecuali diminta secara eksplisit.

---

20. PERFORMANCE

Target awal:

Search harus terasa cepat pada kondisi normal.

Target desain:

~ < 3 seconds

Namun latency external scientific API tidak dapat dijamin.

Gunakan:

- parallel source calls;
- timeout;
- connection reuse;
- caching bila sesuai;
- minimal payload;
- bounded concurrency.

---

21. LOGGING

Gunakan structured logging.

Log informasi seperti:

request_id
query_id
source
duration
status
result_count
error_code

Jangan log:

- secrets;
- API keys;
- authorization tokens;
- sensitive user data.

---

22. TESTING

Minimal test:

source adapter
normalization
DOI normalization
deduplication
ranking
query validation
timeout
partial source failure
empty result
successful search
article retrieval
fulltext availability

External API tests sebaiknya menggunakan mock server.

Jangan membuat unit test yang selalu bergantung pada internet.

---

23. NO SCOPE CREEP

Jangan menambahkan:

- login;
- social features;
- subscription;
- citation manager;
- systematic review;
- meta-analysis;
- diagnosis;
- treatment recommendation;
- medical decision support.

Tetap fokus pada scientific article search.

---

24. HANDOFF RULE

AI-04 wajib membuat:

HANDOFF_AI04.md

Dokumen harus menjelaskan:

- backend structure;
- API;
- source adapters;
- normalized model;
- ranking;
- deduplication;
- configuration;
- tests;
- repository interfaces;
- known limitations;
- database requirements untuk AI-05.
