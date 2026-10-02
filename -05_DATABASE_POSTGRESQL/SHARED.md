JOURNAL AI RESEARCH

AI-05 DATABASE SHARED RULES

---

1. PROJECT

Nama:

Journal AI Research

Core flow:

KETIK
 ↓
CARI
 ↓
10 JURNAL
 ↓
BACA
 ↓
PDF

---

2. DATABASE ROLE

Database bukan scientific search engine.

Database berfungsi sebagai:

Persistence Layer

Database menyimpan:

- article metadata;
- searches;
- search results;
- identifiers;
- timestamps;
- source information;
- persistence state.

---

3. ARCHITECTURE

Frontend
   ↓
Go Backend
   ↓
Search Service
   ├── PubMed
   ├── Europe PMC
   └── Crossref
   ↓
Normalization
   ↓
Deduplication
   ↓
Ranking
   ↓
Repository
   ↓
PostgreSQL

Database tidak boleh dipanggil langsung oleh frontend.

---

4. CORE TABLES

Minimal:

articles
searches
search_results

---

5. ARTICLES

Conceptual fields:

id
title
abstract
journal
publication_date
doi
pmid
pmcid
url
pdf_url
source
is_open_access
created_at
updated_at

Authors dapat disimpan sesuai keputusan schema.

AI-05 harus menentukan apakah authors:

JSONB
separate article_authors table

berdasarkan kebutuhan query dan maintainability.

Jangan memilih hanya karena lebih mudah.

---

6. SEARCHES

Minimal concept:

id
query
created_at

Optional:

normalized_query
result_count
duration_ms
status

Tambahkan hanya jika benar-benar diperlukan.

---

7. SEARCH RESULTS

Relationship:

search
  ↓
search_results
  ↓
article

Minimal:

search_id
article_id
rank
score
created_at

Search result harus menyimpan snapshot informasi ranking jika diperlukan untuk reproducibility.

---

8. PRIMARY KEYS

Gunakan ID yang stabil.

UUID dapat digunakan apabila konsisten dengan backend.

Jika AI-04 sudah menentukan format ID, ikuti keputusan tersebut.

Jangan mengganti ID strategy tanpa alasan.

---

9. IDENTIFIERS

Article identifiers:

DOI
PMID
PMCID

Identifier optional.

Unique constraint hanya diterapkan dengan benar untuk non-null values.

Contoh konsep:

CREATE UNIQUE INDEX ...
WHERE doi IS NOT NULL;

Jangan membuat:

UNIQUE(doi)

tanpa memahami behavior NULL dan kebutuhan aktual.

---

10. DOI

DOI harus disimpan dalam normalized form sesuai keputusan AI-04.

Contoh:

10.1234/ABC

bukan secara random dalam berbagai format.

Original URL dapat tetap tersedia dalam field URL bila diperlukan.

---

11. PMIDs

PMID sebaiknya disimpan sebagai value yang dapat dicari dengan cepat.

PMCID juga demikian.

---

12. INDEXING

Minimal pertimbangkan index pada:

doi
pmid
pmcid
publication_date
search_id
article_id

Index harus berdasarkan query nyata.

Jangan membuat puluhan index tanpa kebutuhan.

---

13. FOREIGN KEYS

Relasi:

search_results.search_id
    → searches.id

search_results.article_id
    → articles.id

Foreign key harus aktif.

Jangan menggunakan relation hanya berdasarkan convention tanpa database constraint.

---

14. DELETE POLICY

Tentukan behavior:

search deleted
article deleted
search_result deleted

Gunakan "ON DELETE" policy yang sesuai.

Untuk historical search data, hindari cascade yang dapat menghapus article metadata tanpa sengaja.

Documentasikan keputusan.

---

15. TRANSACTIONS

Gunakan transaction untuk operasi yang harus atomic.

Contoh:

create search
+
insert search results

Jika salah satu gagal dan atomicity diperlukan:

ROLLBACK

---

16. CONNECTION POOL

Backend harus menggunakan PostgreSQL connection pool.

Konfigurasi harus dapat mengatur:

max connections
min connections
connection lifetime
idle timeout

Nilai final disesuaikan deployment.

---

17. TIMEOUT

Database query harus menggunakan context.

Jangan menggunakan query tanpa timeout/context awareness.

---

18. SQL INJECTION

Semua parameter user harus menggunakan parameterized queries.

Jangan:

fmt.Sprintf("SELECT ... WHERE query = '%s'", userInput)

Gunakan:

prepared/parameterized query

atau equivalent dari library yang digunakan.

---

19. MIGRATION RULE

Migration harus:

ordered
idempotency-aware
reviewable
reproducible

Jangan mengedit migration lama yang sudah digunakan production.

Buat migration baru.

---

20. DATABASE TESTING

Minimal test:

migration up
migration down jika supported
insert article
get article
update article
search relation
search result relation
duplicate identifier
foreign key constraint
transaction rollback

---

21. NO BUSINESS LOGIC DUPLICATION

AI-05 tidak boleh membuat versi kedua dari ranking.

Jangan membuat:

SQL ranking

jika AI-04 sudah memiliki ranking engine.

Database hanya menyimpan dan mengambil data sesuai contract.

---

22. PERFORMANCE

Prioritas:

correct query
proper index
connection pooling
reasonable payload
pagination if required

Jangan melakukan premature optimization.

---

23. DATA RETENTION

AI-05 harus mendokumentasikan:

- apakah searches disimpan;
- berapa lama;
- apakah article metadata dapat di-cache;
- apakah cache memiliki TTL.

Jangan membuat retention policy yang bertentangan dengan product requirements.

---

24. PRIVACY

Query pengguna dapat berisi informasi sensitif.

Jangan menyimpan query lebih lama dari yang dibutuhkan tanpa alasan.

Jangan memasukkan query lengkap ke log jika tidak diperlukan.

---

25. NO SECRETS

".env":

DO NOT COMMIT

".env.example":

SAFE TO COMMIT

---

26. DATABASE SOURCE

Scientific metadata tetap berasal dari:

PubMed
Europe PMC
Crossref

PostgreSQL bukan sumber ilmiah utama.

Database adalah persistence/cache layer.

---

27. FUTURE COMPATIBILITY

Schema sebaiknya memungkinkan future:

- additional scientific source;
- improved search;
- caching;
- analytics;
- article versions.

Namun jangan membuat schema kompleks hanya untuk kemungkinan yang belum diperlukan.

---

28. HANDOFF

AI-05 harus menghasilkan:

HANDOFF_AI05.md

yang menjelaskan:

- schema;
- migrations;
- repository;
- indexes;
- constraints;
- environment;
- tests;
- limitations;
- integration instructions.
