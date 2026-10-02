AI-05 — DATABASE + POSTGRESQL

TASK

---

PHASE 1 — UNDERSTAND EXISTING SYSTEM

Sebelum coding:

1. baca AI-01;
2. baca AI-02;
3. baca AI-03;
4. baca AI-04;
5. identifikasi Article model;
6. identifikasi repository interface;
7. identifikasi API contract;
8. identifikasi persistence requirements.

Buat catatan:

Existing Article model:
Existing ArticleRepository:
Search persistence requirement:
Cache requirement:
ID strategy:

---

PHASE 2 — DATABASE DESIGN

Rancang PostgreSQL schema.

Minimal:

articles
searches
search_results

---

TASK 2.1 — ARTICLES TABLE

Implement field yang dibutuhkan backend.

Minimal:

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

Sesuaikan dengan model AI-04.

---

TASK 2.2 — AUTHORS

Evaluasi dua pendekatan:

Option A

Authors disimpan sebagai JSONB.

Option B

Relational:

authors
article_authors

Pilih berdasarkan kebutuhan aktual.

Documentasikan:

Decision:
Reason:
Trade-off:

Jangan membuat relational model hanya demi kompleksitas.

---

TASK 2.3 — SEARCHES

Implement:

id
query
created_at

Optional:

normalized_query
result_count
duration_ms
status

---

TASK 2.4 — SEARCH_RESULTS

Implement:

id
search_id
article_id
rank
score
created_at

Pastikan:

search_id → searches.id
article_id → articles.id

---

TASK 2.5 — CONSTRAINTS

Implement appropriate constraints.

Examples:

NOT NULL
FOREIGN KEY
CHECK
UNIQUE

Jangan menambahkan NOT NULL pada metadata yang memang optional.

---

TASK 2.6 — UNIQUE IDENTIFIERS

Pastikan duplicate article tidak mudah masuk.

Pertimbangkan:

DOI
PMID
PMCID

Gunakan partial unique indexes apabila identifier optional.

---

TASK 2.7 — INDEXES

Buat index sesuai query backend.

Minimal evaluasi:

doi
pmid
pmcid
publication_date
search_id
article_id

Untuk setiap index:

Purpose:
Query:
Expected benefit:

---

PHASE 3 — MIGRATIONS

Buat migration pertama.

Contoh struktur:

migrations/
├── 000001_initial_schema.up.sql
└── 000001_initial_schema.down.sql

Nama mengikuti migration tool yang digunakan.

---

MIGRATION RULES

Jangan:

- membuat schema manual;
- mengubah production database secara langsung;
- mengedit migration lama setelah digunakan production.

---

PHASE 4 — DATABASE DRIVER

Pilih PostgreSQL driver/library yang sesuai dengan existing Go project.

Contoh:

pgx

Jika project sudah menggunakan library tertentu:

ikuti existing architecture.

Jangan mengganti driver hanya demi preferensi pribadi.

---

PHASE 5 — CONNECTION

Implement database connection layer.

Harus mendukung:

context
pool
timeouts
graceful close

---

PHASE 6 — REPOSITORY

Implement repository sesuai interface AI-04.

Minimal:

type ArticleRepository interface {
    GetByID(ctx context.Context, id string) (*Article, error)
    Save(ctx context.Context, article *Article) error
}

Interface aktual harus mengikuti AI-04.

---

TASK 6.1 — ARTICLE CREATE/UPSERT

Implement persistence artikel.

Behavior duplicate:

existing DOI
existing PMID
existing PMCID

harus mengikuti deduplication policy AI-04.

Jangan membuat duplicate article hanya karena datang dari source berbeda.

---

TASK 6.2 — ARTICLE GET

Implement:

GetByID

Query harus:

- parameterized;
- context-aware;
- indexed.

---

TASK 6.3 — ARTICLE UPDATE

Jika dibutuhkan oleh integration:

implement update/upsert.

Jangan overwrite valid metadata dengan empty value dari source yang lebih lemah.

Merge policy harus mengikuti AI-04.

---

TASK 6.4 — SEARCH PERSISTENCE

Implement persistence untuk search:

create search
save results
retrieve results

---

TASK 6.5 — TRANSACTION

Jika search + results harus atomic:

gunakan transaction.

Concept:

BEGIN
 ↓
INSERT search
 ↓
INSERT search results
 ↓
COMMIT

Jika gagal:

ROLLBACK

---

PHASE 7 — ERROR HANDLING

Repository errors harus dapat dibedakan.

Minimal:

not found
duplicate
constraint violation
connection failure
query failure
transaction failure

Jangan expose raw database errors ke browser.

---

PHASE 8 — TESTING

Implement database tests.

---

TEST 1 — MIGRATION

Test:

up migration
schema exists
tables exist
indexes exist
foreign keys exist

---

TEST 2 — ARTICLE

Test:

insert article
get article

---

TEST 3 — IDENTIFIER

Test:

same DOI
same PMID
same PMCID

Pastikan constraint behavior benar.

---

TEST 4 — SEARCH

Test:

create search

---

TEST 5 — SEARCH RESULTS

Test:

search → article relation

---

TEST 6 — FOREIGN KEY

Test invalid:

search_result.article_id

harus ditolak jika article tidak ada.

---

TEST 7 — TRANSACTION

Simulasikan failure:

insert search succeeds
insert result fails

Pastikan rollback terjadi jika operation harus atomic.

---

PHASE 9 — PERFORMANCE

Gunakan:

EXPLAIN
EXPLAIN ANALYZE

untuk query penting.

Jangan melakukan benchmark dengan data kosong saja.

Jika belum ada dataset besar:

dokumentasikan keterbatasannya.

---

PHASE 10 — SECURITY

Pastikan:

[ ] parameterized SQL
[ ] no hardcoded password
[ ] no committed .env
[ ] restricted DB credentials
[ ] context-aware queries
[ ] no raw SQL error to client

---

PHASE 11 — ENVIRONMENT

Update:

.env.example

Contoh:

DATABASE_URL=postgres://user:password@localhost:5432/journal_ai

Gunakan placeholder saja.

Jangan commit credential asli.

---

PHASE 12 — DOCKER

Jika project menggunakan Docker:

pastikan PostgreSQL dapat dijalankan melalui:

docker compose up

Jangan mengubah deployment architecture tanpa alasan.

---

PHASE 13 — DOCUMENTATION

Update:

backend/README.md

Tambahkan:

- database setup;
- migration command;
- test command;
- environment variables;
- repository architecture.

---

PHASE 14 — INTEGRATION CHECK

Setelah repository selesai:

test:

Backend
  ↓
Repository
  ↓
PostgreSQL

Pastikan backend dapat:

save article
retrieve article
save search
save search results
retrieve article

---

PHASE 15 — ACCEPTANCE CRITERIA

AI-05 diterima jika:

[ ] Schema created
[ ] Migration works
[ ] Rollback tested if supported
[ ] Articles table works
[ ] Searches table works
[ ] Search results table works
[ ] Foreign keys work
[ ] Identifier constraints work
[ ] Indexes exist
[ ] Repository works
[ ] Transactions work
[ ] Connection pool works
[ ] Context cancellation works
[ ] SQL injection protections verified
[ ] Tests pass
[ ] Documentation updated
[ ] HANDOFF_AI05.md completed

---

PHASE 16 — DO NOT CHANGE

AI-05 tidak boleh mengubah tanpa alasan dan dokumentasi:

search ranking
relevance formula
UX
frontend behavior
scientific source policy
full-text legal policy

---

PHASE 17 — FINAL HANDOFF

Sebelum selesai:

buat:

HANDOFF_AI05.md

Isi minimal:

Schema
Migrations
Repositories
Indexes
Constraints
Transactions
Configuration
Tests
Known limitations
AI-06 integration instructions
