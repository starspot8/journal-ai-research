JOURNAL AI RESEARCH

AI-07 TESTING + SECURITY SHARED RULES

---

1. PROJECT

Journal AI Research

Core:

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

2. QUALITY PRINCIPLE

Kode yang berhasil build belum berarti aplikasi siap.

AI-07 harus memverifikasi:

Correctness
Reliability
Security
Accessibility
Performance
Compatibility

---

3. SOURCE POLICY

Scientific sources:

PubMed
Europe PMC
Crossref

AI-07 harus memastikan source handling sesuai policy.

Tidak boleh:

paywall bypass
unauthorized scraping
unauthorized article copying

---

4. API

Minimal:

POST /api/v1/search
GET /api/v1/articles/:id
GET /api/v1/articles/:id/fulltext
POST /api/v1/export/pdf
GET /api/v1/health

Test:

- valid requests;
- invalid requests;
- malformed JSON;
- missing fields;
- oversized input;
- unexpected values;
- repeated requests.

---

5. SEARCH TESTING

Test:

latest journal about cataract surgery

jurnal terbaru tentang diabetic retinopathy

latest research about multifocal IOL

Saya ingin penelitian terbaru tentang hasil multifocal IOL pada pasien usia lanjut

Expected:

maximum 10 results

---

6. SEARCH EDGE CASES

Test:

empty
whitespace
very long query
special characters
Unicode
Indonesian
English
mixed language
duplicate query

---

7. RESULTS

Verify:

title
authors
journal
publication date
DOI
PMID
PMCID
source

Do not require optional metadata that the scientific source does not provide.

---

8. DEDUPLICATION

Test same article returned from multiple providers.

Expected:

one logical article

Priority:

DOI
PMID
PMCID
normalized title + author + year

---

9. RANKING

Expected initial formula:

relevance     60%
recency       30%
completeness  10%

AI-07 verifies implementation consistency.

Do not redesign the ranking formula.

---

10. PARTIAL FAILURE

Test:

PubMed       FAIL
Europe PMC   SUCCESS
Crossref     SUCCESS

Expected:

results still available

---

11. ALL PROVIDERS FAILURE

Test all scientific providers unavailable.

Expected:

controlled backend error
+
user-friendly frontend message

No infinite spinner.

---

12. TIMEOUT

Test provider timeout.

Expected:

request eventually completes
or returns controlled error

Never:

infinite waiting

---

13. RATE LIMIT

Verify search endpoint cannot be abused by unlimited requests.

Test:

rapid repeated requests

Expected:

rate limit eventually activates

Exact policy follows implementation.

---

14. INPUT VALIDATION

Backend must validate even if frontend validates.

Test:

missing query
empty query
oversized query
malformed JSON
unexpected JSON types

---

15. SQL INJECTION

Test repository/API with harmless security test strings.

Goal:

input remains data

Never construct SQL from raw user input.

---

16. XSS

Test article metadata and user-controlled fields for HTML/script injection.

If rich text is rendered:

sanitize

Do not execute arbitrary HTML.

---

17. SSRF

This is a major test area.

Potentially dangerous fields:

external URL
PDF URL
article URL
source URL

Backend must not blindly fetch arbitrary user-provided URLs.

Test internal/private destinations where technically safe in test environment.

Expected:

blocked / validated

---

18. URL SECURITY

External links must be validated.

Avoid unsafe schemes such as:

javascript:
data:
file:

unless explicitly required and safely handled.

Expected article links should use appropriate web schemes.

---

19. CORS

Verify:

allowed frontend origin

and rejected unauthorized origins where applicable.

Production should not unnecessarily allow:

*

---

20. SECRET EXPOSURE

Search repository for:

API keys
database passwords
tokens
private keys
credentials

Check:

frontend bundle
source code
logs
error responses
Git history where practical

---

21. ERROR HANDLING

Errors shown to users must not expose:

SQL
stack trace
filesystem path
internal hostname
credentials
provider secrets

---

22. DATABASE

Test:

foreign keys
constraints
duplicate identifiers
transactions
rollback
connection failure

---

23. PDF SECURITY

Verify PDF does not unintentionally expose:

database credentials
API credentials
internal URLs
server filesystem paths
private debug information

---

24. FULL TEXT SECURITY

Verify:

legal full text available

is different from:

source says article exists

Do not treat metadata availability as authorization to download restricted full text.

---

25. ACCESSIBILITY

Verify:

keyboard
focus
labels
headings
buttons
links
forms
reader
errors

Search must work with:

Enter

---

26. RESPONSIVE

Verify:

320px-ish mobile
tablet
desktop

No critical functionality should disappear because of viewport size.

---

27. PERFORMANCE

Measure:

search request
article request
fulltext request
PDF generation
database query

Identify bottlenecks.

Do not fabricate benchmark numbers.

---

28. REGRESSION

After fixing a bug:

re-run related tests

and:

re-run core E2E flow

---

29. TEST REPORT

Every finding:

ID
Severity
Component
Description
Steps
Expected
Actual
Impact
Recommendation
Status

---

30. QUALITY GATE

BLOCK RELEASE

if unresolved:

CRITICAL

or major unresolved:

HIGH security vulnerability

or:

core user flow unusable

---

31. DOCUMENTATION

AI-07 must produce:

HANDOFF_AI07.md

with:

- tests executed;
- results;
- vulnerabilities;
- fixes;
- remaining issues;
- release recommendation based on explicit technical criteria.

Do not provide political/electoral-style ranking or unrelated evaluative conclusions; this is a technical quality gate based on documented test criteria.

---

32. NEXT STAGE

AI-08 uses AI-07 handoff to determine deployment blockers.

Deployment should not proceed past unresolved critical security issues.
