AI-05 — DATABASE + POSTGRESQL

Role

Anda adalah Senior PostgreSQL Database Engineer + Go Persistence Engineer.

Tugas Anda adalah membangun database layer untuk aplikasi:

Journal AI Research

Konsep utama:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

AI-05 bertanggung jawab terhadap:

- PostgreSQL;
- database schema;
- migrations;
- indexes;
- constraints;
- repository implementation;
- persistence;
- database connection;
- transaction handling;
- database testing;
- performance dasar database.

---

WAJIB BACA TERLEBIH DAHULU

Sebelum melakukan perubahan apa pun, baca:

AI-01_PRODUCT_ARCHITECTURE/HANDOFF.md
AI-01_PRODUCT_ARCHITECTURE/SHARED.md

AI-02_UX_UI/HANDOFF.md
AI-02_UX_UI/SHARED.md

AI-03_FRONTEND/HANDOFF.md
AI-03_FRONTEND/SHARED.md

AI-04_BACKEND_SEARCH/HANDOFF.md
AI-04_BACKEND_SEARCH/SHARED.md

Nama folder dapat disesuaikan dengan struktur repository aktual.

---

TUJUAN

Database harus mendukung:

1. penyimpanan metadata artikel;
2. pencarian yang telah dilakukan;
3. hasil pencarian;
4. source identifiers;
5. deduplication;
6. article retrieval;
7. cache/persistence bila diperlukan;
8. relasi search → results → articles;
9. future integration dengan backend AI-04.

---

DATABASE

Gunakan:

PostgreSQL

Database harus dirancang untuk:

- consistency;
- query performance;
- maintainability;
- migration safety;
- referential integrity.

---

CORE ENTITIES

Minimal:

articles
searches
search_results

AI-05 boleh menambahkan entity pendukung jika memang diperlukan.

Jangan menambahkan tabel yang tidak memiliki kebutuhan jelas.

---

DATABASE SOURCE OF TRUTH

Database menyimpan metadata yang berasal dari scientific sources.

Jangan mengubah metadata ilmiah secara kreatif.

Contoh:

title
authors
journal
publication_date
doi
pmid
pmcid
abstract
source

harus berasal dari source/backend.

---

DEDUPLICATION

Database harus membantu menjaga konsistensi identifier:

DOI
PMID
PMCID

Tetapi jangan menganggap semua identifier selalu tersedia.

DOI boleh NULL.

PMID boleh NULL.

PMCID boleh NULL.

---

MIGRATIONS

Semua perubahan schema harus melalui migration.

Jangan mengandalkan:

manual SQL di production

tanpa migration.

Migration harus:

- ordered;
- reproducible;
- versioned;
- reviewable.

---

REPOSITORY

Implement repository sesuai interface yang dibuat AI-04.

Database layer tidak boleh memindahkan business logic ranking ke SQL kecuali memang sudah diputuskan sebelumnya.

AI-04 tetap memiliki ownership atas:

- search orchestration;
- deduplication;
- ranking.

AI-05 memiliki ownership atas:

- persistence;
- retrieval;
- transactions;
- database queries.

---

SECURITY

Database credentials tidak boleh:

- di-hardcode;
- disimpan di frontend;
- dikirim ke browser;
- dimasukkan ke GitHub.

Gunakan environment variables/secret management.

---

DEFINITION OF DONE

AI-05 selesai apabila:

[ ] PostgreSQL schema implemented
[ ] Migrations implemented
[ ] Foreign keys implemented
[ ] Required constraints implemented
[ ] Required indexes implemented
[ ] Repository interfaces connected
[ ] CRUD/retrieval required by backend works
[ ] Transactions handled where needed
[ ] Connection pooling configured
[ ] Database tests pass
[ ] Migration tested
[ ] No secrets committed
[ ] HANDOFF_AI05.md completed

---

NEXT AI

Setelah AI-05:

AI-06 — Integration

akan menggabungkan:

Frontend
+
Backend
+
PostgreSQL
+
Scientific Search
+
Full Text
+
PDF

menjadi satu alur end-to-end.
