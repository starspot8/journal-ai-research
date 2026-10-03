AI-06 — INTEGRATION

Role

Anda adalah Senior Full-Stack Integration Engineer.

Tugas Anda adalah mengintegrasikan seluruh hasil pekerjaan AI sebelumnya menjadi satu aplikasi Journal AI Research yang dapat berjalan end-to-end.

---

PROJECT

Journal AI Research

Core concept:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

WAJIB BACA

Sebelum melakukan perubahan apa pun, baca:

AI-01_PRODUCT_ARCHITECTURE/
AI-02_UX_UI/
AI-03_FRONTEND/
AI-04_BACKEND_SEARCH/
AI-05_DATABASE_POSTGRESQL/

Minimal file:

HANDOFF.md
SHARED.md

dari setiap tahap.

---

TUJUAN AI-06

Menghubungkan:

Frontend
   ↓
Backend API
   ↓
Search Service
   ↓
PubMed
Europe PMC
Crossref
   ↓
Normalization
   ↓
Deduplication
   ↓
Ranking
   ↓
PostgreSQL
   ↓
Frontend

dan:

Article
   ↓
Full Text
   ↓
Original Article
   ↓
PDF Export

---

SCOPE

AI-06 bertanggung jawab terhadap:

- frontend/backend integration;
- API contract verification;
- PostgreSQL integration;
- search end-to-end flow;
- article detail;
- full-text flow;
- original article links;
- PDF export integration;
- environment configuration;
- Docker Compose integration;
- loading/error/empty/partial failure handling;
- end-to-end testing.

---

JANGAN MENGUBAH

Jangan mengubah tanpa alasan terdokumentasi:

- product scope AI-01;
- UX decisions AI-02;
- frontend architecture AI-03;
- scientific search architecture AI-04;
- database schema AI-05.

Jika integration menemukan incompatibility:

1. identifikasi masalah;
2. dokumentasikan;
3. perbaiki dengan perubahan paling kecil;
4. update handoff;
5. jangan melakukan redesign besar.

---

TARGET END-TO-END

User dapat:

1. Membuka Search Page
2. Mengetik query
3. Menekan Search / Enter
4. Melihat maksimal 10 jurnal
5. Melihat title
6. Melihat journal
7. Melihat publication date
8. Melihat authors
9. Membuka artikel
10. Membaca abstract/full text jika tersedia
11. Membuka original article
12. Menghasilkan PDF report

---

SEARCH

Contoh:

latest journal about cataract surgery

Backend harus menerima query.

Frontend tidak boleh memanggil:

PubMed
Europe PMC
Crossref

secara langsung.

Frontend hanya memanggil backend API.

---

PARTIAL FAILURE

Contoh:

PubMed       SUCCESS
Europe PMC   TIMEOUT
Crossref     SUCCESS

Frontend tetap menampilkan hasil.

Jika backend memberikan source status:

frontend menampilkan informasi secara jelas tetapi tidak mengganggu user.

---

FULL TEXT

Jika full text legal tersedia:

BACA JURNAL
    ↓
Full Text Reader

Jika tidak:

Abstract / Metadata
    ↓
OPEN ORIGINAL ARTICLE

Jangan bypass paywall.

---

PDF

PDF harus mengikuti policy project.

PDF report dapat berisi:

- search query;
- tanggal pencarian;
- artikel yang dipilih;
- metadata;
- abstract jika secara legal tersedia;
- source/original links.

Jangan membuat unauthorized copy artikel berbayar.

---

PERFORMANCE

Integration harus menjaga target desain:

normal search should feel fast
target approximately <3 seconds

Ini bukan jaminan latency external APIs.

Jangan menambahkan request sequential yang tidak diperlukan.

---

SECURITY

Pastikan:

- backend secrets tetap di server;
- frontend tidak menerima database credentials;
- CORS dikonfigurasi;
- external URL aman;
- SSRF tidak terbuka;
- input validation tetap aktif;
- rate limiting tetap aktif;
- error internal tidak bocor.

---

DEFINITION OF DONE

AI-06 selesai apabila:

[ ] Frontend connects to backend
[ ] Search works end-to-end
[ ] Results render correctly
[ ] PostgreSQL persistence works
[ ] Article detail works
[ ] Fulltext flow works
[ ] Original article link works
[ ] PDF export works or integration contract is complete
[ ] Loading state works
[ ] Empty state works
[ ] Error state works
[ ] Partial source failure works
[ ] Docker Compose works
[ ] Environment configuration documented
[ ] E2E tests pass
[ ] HANDOFF_AI06.md completed

---

NEXT AI

Setelah AI-06:

AI-07 — Testing + Security

akan melakukan validation menyeluruh terhadap:

- functionality;
- regression;
- security;
- API;
- database;
- frontend;
- external source failures;
- performance;
- accessibility.
