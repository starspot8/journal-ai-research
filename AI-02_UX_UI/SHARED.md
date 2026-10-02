SHARED PROJECT SPECIFICATION

Journal AI Research

1. PRODUCT

Journal AI Research

Core concept:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

2. PRODUCT PURPOSE

Allow users to search scientific literature using a simple search box and receive:

10 Latest Relevant Journals

---

3. V1 SOURCES

- PubMed
- Europe PMC
- Crossref

---

4. RANKING

Initial ranking:

Relevance              60%
Recency                30%
Metadata completeness  10%

This does not mean scientific quality ranking.

---

5. DEDUPLICATION

Priority:

DOI
PMID
PMCID
Title + Author + Year

---

6. FULL-TEXT POLICY

Legal access only.

The application must not:

- bypass paywalls
- bypass authentication
- scrape restricted content
- create unauthorized copies

When unavailable:

FULL TEXT NOT AVAILABLE

[ OPEN ORIGINAL ARTICLE ]

---

7. PDF

V1 PDF is a generated research report.

It may contain:

- search query
- search date
- selected articles
- metadata
- permitted abstract
- source links

---

8. CORE PAGES

V1 UX:

1. Search
2. Search Results
3. Article Detail / Reader
4. PDF Export

---

9. SEARCH

One primary search box.

Examples:

latest cataract surgery research

jurnal terbaru diabetic retinopathy

multifocal IOL elderly patients

The user does not need Boolean syntax.

---

10. RESULTS

Header:

10 Latest Relevant Journals

Each card:

- title
- authors
- journal
- publication date
- DOI when available
- source
- article action

---

11. ARTICLE

Article detail includes:

- title
- authors
- journal
- date
- DOI
- PMID/PMCID
- abstract
- full-text status
- source
- original link

---

12. UX PRINCIPLES

The application must be:

- simple
- professional
- readable
- responsive
- accessible
- uncluttered

---

13. MOBILE

Must support:

- phones
- tablets
- desktops

Mobile must not require horizontal scrolling for primary workflows.

---

14. ACCESSIBILITY

Must support:

- keyboard navigation
- focus states
- readable contrast
- semantic structure
- accessible labels
- touch-friendly controls
- meaningful errors

---

15. STACK

Frontend:

Next.js + TypeScript

Backend:

Go

Database:

PostgreSQL

---

16. HANDOFF RULE

Every AI must:

1. Read SHARED.md.
2. Read the previous AI handoff.
3. Follow frozen decisions.
4. Avoid unnecessary scope expansion.
5. Document changes.
6. Create the next handoff.

---

17. AI CHAIN

AI-01 Product + Architecture
↓
AI-02 UX/UI
↓
AI-03 Frontend
↓
AI-04 Backend + Search
↓
AI-05 Database
↓
AI-06 Integration
↓
AI-07 Testing + Security
↓
AI-08 Deployment

---

18. IMPORTANT

The application is a scientific literature search tool.

It is not:

- a medical diagnostic tool
- a treatment recommendation engine
- an automatic systematic review
- a replacement for scientific judgment
