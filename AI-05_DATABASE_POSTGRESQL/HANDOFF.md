AI-05 — HANDOFF

DATABASE + POSTGRESQL

«Dokumen ini adalah template yang WAJIB diperbarui setelah AI-05 benar-benar menyelesaikan implementasi.»

---

1. STATUS

NOT STARTED

Setelah selesai:

COMPLETE

atau:

PARTIAL

---

2. DATABASE

Database:

PostgreSQL

Version:

DOCUMENT ACTUAL VERSION

Driver:

DOCUMENT ACTUAL DRIVER

---

3. SCHEMA

Tables:

articles
searches
search_results

Additional tables:

DOCUMENT ACTUAL TABLES

---

4. ARTICLES

Actual schema:

PASTE/REFERENCE ACTUAL SCHEMA

Document:

- primary key;
- fields;
- nullable fields;
- identifier fields;
- timestamps.

---

5. AUTHORS

Strategy:

JSONB / relational / other

Reason:

DOCUMENT REASON

---

6. SEARCHES

Actual schema:

DOCUMENT ACTUAL SCHEMA

Purpose:

...

---

7. SEARCH RESULTS

Actual schema:

DOCUMENT ACTUAL SCHEMA

Relationship:

search
  ↓
search_results
  ↓
article

---

8. RELATIONSHIPS

Document actual foreign keys:

search_results.search_id → searches.id
search_results.article_id → articles.id

Additional relationships:

...

---

9. IDENTIFIER CONSTRAINTS

Document:

DOI:
PMID:
PMCID:

Explain whether constraints are:

unique
partial unique
non-unique

---

10. INDEXES

List:

index name
table
columns
purpose

Example:

articles_doi_idx
articles
doi
fast article lookup

---

11. MIGRATIONS

Migration tool:

DOCUMENT ACTUAL TOOL

Migration files:

...

Command:

DOCUMENT ACTUAL COMMAND

---

12. REPOSITORIES

Implemented repositories:

ArticleRepository
SearchRepository
SearchResultRepository

Document actual interfaces and implementations.

---

13. TRANSACTIONS

Document:

Which operations are transactional:
Why:
Isolation level if customized:
Rollback behavior:

---

14. CONNECTION POOL

Document:

Max connections:
Min connections:
Max lifetime:
Idle timeout:

If default values are used:

state explicitly.

---

15. ENVIRONMENT

Variables:

DATABASE_URL

Additional variables:

...

Never put real credentials in this document.

---

16. TESTING

Command:

go test ./...

Result:

PASS / FAIL

Database integration test command:

DOCUMENT ACTUAL COMMAND

---

17. TEST COVERAGE

Document tested areas:

[ ] Migration
[ ] Article persistence
[ ] Search persistence
[ ] Search results
[ ] Foreign keys
[ ] Identifier constraints
[ ] Transactions
[ ] Rollback
[ ] Context cancellation
[ ] Repository errors

---

18. SECURITY

Confirm:

[ ] No credentials in repository
[ ] Parameterized queries
[ ] No SQL injection
[ ] No raw DB error exposed
[ ] Restricted DB permissions
[ ] Environment-based secrets

---

19. PERFORMANCE

Document important queries:

Article lookup:
Search retrieval:
Search result retrieval:

Index verification:

EXPLAIN / EXPLAIN ANALYZE performed:
YES / NO

---

20. DATA RETENTION

Document:

Search retention:
Article retention:
Cache retention:

If not implemented:

state:

Not implemented / deferred.

---

21. KNOWN LIMITATIONS

Examples:

- No large-scale benchmark yet.
- Search query analytics not implemented.
- Article versioning not implemented.

Only list actual limitations.

---

22. DEVIATIONS

If implementation differs from AI-04:

Deviation:
Reason:
Impact:

If none:

No material deviations from AI-04.

---

23. AI-04 COMPATIBILITY

Confirm:

[ ] Article model compatible
[ ] Repository interface compatible
[ ] ID strategy compatible
[ ] Identifier normalization compatible
[ ] Search persistence compatible

---

24. AI-06 INTEGRATION REQUIREMENTS

AI-06 harus membaca dokumen ini dan memastikan:

Frontend
   ↓
Backend
   ↓
Repository
   ↓
PostgreSQL

berjalan end-to-end.

AI-06 perlu mengetahui:

Database connection configuration
Migration command
Repository implementations
Transaction behavior
Error behavior

---

25. AI-06 FIRST TEST

AI-06 sebaiknya melakukan:

1. Start PostgreSQL
2. Run migrations
3. Start backend
4. Run health endpoint
5. Search article
6. Verify article persistence
7. Verify search persistence
8. Verify search results persistence
9. Retrieve article
10. Verify frontend displays result

---

26. FINAL SUMMARY

Setelah implementasi, tuliskan ringkasan aktual:

AI-05 completed the PostgreSQL persistence layer.

Implemented:
- ...
- ...
- ...

Database:
- ...

Repositories:
- ...

Tests:
- ...

Known limitations:
- ...

AI-06 should now:
- ...
- ...

Jangan meninggalkan placeholder dalam final handoff.
