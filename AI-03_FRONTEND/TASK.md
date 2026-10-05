AI-03 TASK

FRONTEND IMPLEMENTATION

Journal AI Research

OBJECTIVE

Implement the Journal AI Research V1 frontend based on the specifications produced by AI-01 and AI-02.

---

PHASE 1 — PRE-IMPLEMENTATION

Read:

SHARED.md
HANDOFF_AI01.md
HANDOFF_AI02.md

Confirm:

- product scope
- UX/UI decisions
- component requirements
- responsive requirements
- accessibility requirements

Do not begin coding before understanding them.

---

PHASE 2 — PROJECT SETUP

Create or verify:

Next.js
TypeScript
React

Configure:

- linting
- formatting if already approved
- TypeScript strictness
- testing
- build scripts

Avoid unnecessary dependencies.

---

PHASE 3 — APPLICATION SHELL

Implement:

- root layout
- application header
- global styles
- design tokens
- responsive container
- typography

---

PHASE 4 — SEARCH PAGE

Implement:

- search input
- search button
- examples
- validation
- Enter key
- loading
- error

Required empty query message:

"Masukkan topik jurnal terlebih dahulu."

---

PHASE 5 — SEARCH API ABSTRACTION

Create a typed API client.

Example:

searchArticles(query: string)

Do not put external scientific source logic in the frontend.

The frontend should communicate only with the application's backend API.

---

PHASE 6 — MOCK API

If backend is not available:

Create a mock adapter.

The mock must:

- clearly identify mock mode
- use obviously synthetic/test data
- never be enabled accidentally in production

The interface must be compatible with the future backend.

---

PHASE 7 — RESULTS

Implement:

- results header
- article list
- article cards
- metadata
- actions

Header:

10 Latest Relevant Journals

---

PHASE 8 — ARTICLE DETAIL

Implement route:

/articles/[id]

Display:

- title
- authors
- journal
- date
- DOI
- PMID
- PMCID
- abstract
- full-text status
- source
- original link

---

PHASE 9 — FULL-TEXT READER

Implement UI for:

Available

Show permitted content.

Unavailable

Show:

"FULL TEXT NOT AVAILABLE"

and:

"OPEN ORIGINAL ARTICLE"

Ensure external links are safe.

---

PHASE 10 — PDF

Implement:

- article selection
- export button
- loading
- success
- error

The backend PDF implementation comes later.

---

PHASE 11 — RESPONSIVE

Test at:

- 320px
- 375px
- 390px
- 768px
- 1024px
- 1280px+

Ensure:

- no horizontal scrolling
- no clipped text
- buttons usable
- article cards readable

---

PHASE 12 — ACCESSIBILITY

Verify:

- keyboard navigation
- focus
- form labels
- buttons
- links
- semantic headings
- contrast
- screen-reader-friendly states

---

PHASE 13 — ERROR HANDLING

Implement:

Empty

"Masukkan topik jurnal terlebih dahulu."

No results

"Tidak ditemukan jurnal yang sesuai."

Partial failure

"Sebagian sumber jurnal sedang tidak tersedia. Hasil yang tersedia tetap ditampilkan."

Complete failure

"Pencarian sedang mengalami gangguan. Silakan coba lagi."

PDF failure

"PDF tidak dapat dibuat. Silakan coba lagi."

---

PHASE 14 — TESTS

Create tests for:

Search

- valid query
- empty query
- Enter
- loading
- error

Results

- zero results
- one result
- ten results
- missing metadata

Article

- full text
- no full text
- missing DOI
- missing PMID

PDF

- idle
- loading
- success
- failure

Accessibility

- keyboard
- form labels
- buttons
- links

---

PHASE 15 — QUALITY CHECK

Run:

npm run lint
npm run typecheck
npm test
npm run build

Use the project's actual commands if they differ.

Fix all blocking errors.

---

PHASE 16 — FRONTEND SECURITY REVIEW

Check:

- no secrets in frontend
- no unsafe HTML
- no unsanitized rich text
- safe external links
- input validation
- safe URL handling
- no unnecessary third-party scripts

---

PHASE 17 — HANDOFF

Create:

"HANDOFF_AI03.md"

It must document:

1. implementation status
2. pages created
3. components created
4. routes
5. API interfaces
6. expected backend response
7. mock behavior
8. environment variables
9. tests
10. known limitations
11. technical decisions
12. AI-04 integration instructions

---

HARD RULES

Do NOT:

- implement PubMed
- implement Europe PMC
- implement Crossref
- create PostgreSQL tables
- create database migrations
- implement backend ranking
- implement backend deduplication
- bypass paywalls
- add login
- add unnecessary features
- change the product scope

---

DEFINITION OF DONE

AI-03 is complete only when:

[ ] frontend runs

[ ] search page works

[ ] results page works

[ ] article page works

[ ] full-text states work

[ ] PDF UI works

[ ] loading works

[ ] empty states work

[ ] error states work

[ ] mobile works

[ ] desktop works

[ ] accessibility implemented

[ ] tests pass

[ ] build passes

[ ] HANDOFF_AI03.md complete
