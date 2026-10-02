AI-02 TASK

UX/UI — Journal AI Research

OBJECTIVE

Create an implementation-ready UX/UI specification for Journal AI Research V1.

AI-02 receives the product architecture from AI-01 and converts it into a precise interface specification for AI-03.

---

TASK 1 — READ AI-01

Before doing any work:

Read:

SHARED.md
HANDOFF_AI01.md

Do not redesign the product architecture.

---

TASK 2 — SEARCH PAGE

Design the Search page.

Must define:

- page layout
- header
- search box
- placeholder
- search button
- examples
- loading state
- validation
- responsive layout

Primary action:

"SEARCH"

---

TASK 3 — RESULTS PAGE

Define:

- page structure
- results header
- article card
- metadata hierarchy
- article actions
- source display
- loading state
- empty state
- partial failure state
- complete failure state

Header must be:

10 Latest Relevant Journals

---

TASK 4 — ARTICLE PAGE

Define:

- title
- authors
- journal
- date
- identifiers
- abstract
- full-text status
- source
- original article
- PDF action

---

TASK 5 — READER

Define how legal full text is displayed.

Consider:

- typography
- line length
- headings
- spacing
- navigation
- mobile reading
- source attribution

Do not design a paywall bypass.

---

TASK 6 — PDF EXPORT

Define:

- selection behavior
- select all behavior if included
- export button
- loading state
- success state
- failure state
- mobile behavior

---

TASK 7 — COMPONENT SYSTEM

Define reusable components.

Minimum:

AppHeader
SearchBox
SearchButton
SearchExamples
SearchLoading
SearchError
SearchEmpty
ResultsHeader
ArticleCard
ArticleMetadata
ArticleActions
ArticleReader
FullTextStatus
PdfExport

---

TASK 8 — DESIGN SYSTEM

Define:

- typography
- colors
- spacing
- border radius
- shadows
- buttons
- inputs
- cards
- breakpoints
- focus states

Use reusable tokens.

---

TASK 9 — RESPONSIVE

Specify layouts for:

Mobile

approximately 320–767px

Tablet

approximately 768–1023px

Desktop

1024px+

The exact breakpoint values may be adjusted by AI-03 only if technically necessary and documented.

---

TASK 10 — ACCESSIBILITY

Define:

- keyboard navigation
- screen-reader labels
- focus states
- contrast
- touch targets
- semantic HTML
- error messaging
- form accessibility

---

TASK 11 — INTERACTION

Define:

- Enter key behavior
- button states
- hover
- focus
- pressed
- disabled
- loading
- navigation
- external links

---

TASK 12 — ERROR STATES

Design:

Empty query

Example:

"Masukkan topik jurnal terlebih dahulu."

No results

"Tidak ditemukan jurnal yang sesuai."

Partial source failure

"Sebagian sumber jurnal sedang tidak tersedia. Hasil yang tersedia tetap ditampilkan."

Complete failure

"Pencarian sedang mengalami gangguan. Silakan coba lagi."

Article unavailable

Show available metadata and source link.

---

TASK 13 — CONTENT DESIGN

Use simple and professional language.

Avoid unnecessary technical terms.

Preferred:

"BACA JURNAL"

instead of:

"FETCH ARTICLE CONTENT"

Preferred:

"OPEN ORIGINAL ARTICLE"

instead of:

"REDIRECT TO SOURCE URI"

---

TASK 14 — VISUAL HIERARCHY

Define exact hierarchy:

1. Application identity
2. Search action
3. Article title
4. Publication information
5. Article actions
6. Secondary metadata

Do not allow metadata to visually overpower the article title.

---

TASK 15 — DESIGN HANDOFF

Create a final UX/UI handoff containing:

- pages
- components
- dimensions
- spacing
- typography
- colors
- states
- interactions
- accessibility
- responsive rules

---

TASK 16 — DO NOT ADD

Do not add:

- login
- profile
- social features
- subscription
- recommendation feed
- complex filtering
- chat
- medical diagnosis
- treatment recommendations

unless explicitly approved later.

---

DEFINITION OF DONE

AI-02 is complete when:

1. Search page is specified.
2. Results page is specified.
3. Article reader is specified.
4. PDF export UX is specified.
5. Loading states are specified.
6. Empty states are specified.
7. Error states are specified.
8. Mobile layout is specified.
9. Desktop layout is specified.
10. Accessibility is specified.
11. Components are specified.
12. Design tokens are specified.
13. AI-03 handoff is complete.

Final required file:

"HANDOFF_AI02.md"
