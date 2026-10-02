HANDOFF AI-02

UX/UI

Journal AI Research

1. STATUS

"NOT STARTED"

Change to:

"COMPLETED"

only when AI-02 has completed and reviewed the UX/UI specification.

---

2. OBJECTIVE

AI-02 converts the AI-01 product and architecture specification into an implementation-ready UX/UI specification.

---

3. PRODUCT

Journal AI Research

Core workflow:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

4. PRIMARY SCREEN

The primary screen contains:

JOURNAL AI RESEARCH

Cari jurnal apa?

[ Contoh: jurnal terbaru tentang operasi katarak ]

[ SEARCH ]

Examples

The user should understand the application immediately.

---

5. SEARCH

Search supports:

- Indonesian
- English
- keywords
- natural language
- questions

Search can be submitted by:

- clicking SEARCH
- pressing Enter

---

6. RESULTS

Page title:

10 Latest Relevant Journals

Each article card contains:

- title
- authors
- journal
- publication date
- DOI when available
- source
- BACA JURNAL
- OPEN ORIGINAL ARTICLE where applicable

---

7. ARTICLE DETAIL

Article page includes:

- title
- authors
- journal
- publication date
- DOI
- PMID/PMCID
- abstract
- full-text status
- source
- original article
- export PDF

---

8. FULL-TEXT STATES

Available

Display legally accessible full text.

Unavailable

Display:

"FULL TEXT NOT AVAILABLE"

and:

"OPEN ORIGINAL ARTICLE"

---

9. PDF

Primary action:

EXPORT PDF

The PDF is a generated research report.

It may include:

- search query
- date
- selected articles
- metadata
- permitted abstract
- source links

---

10. RESPONSIVE DESIGN

Required:

- mobile
- tablet
- desktop

Mobile must be optimized for touch.

Avoid:

- horizontal scrolling
- tiny buttons
- dense metadata
- complicated navigation

---

11. ACCESSIBILITY

Required:

- keyboard navigation
- semantic headings
- accessible labels
- visible focus
- sufficient contrast
- touch-friendly targets
- understandable error messages

---

12. DESIGN SYSTEM

AI-02 must define:

Colors

- primary
- secondary
- background
- surface
- text
- muted
- border
- success
- warning
- error
- focus

Typography

- H1
- H2
- article title
- body
- metadata
- caption

Spacing

Use consistent spacing tokens.

Components

Use reusable components rather than one-off styles.

---

13. REQUIRED STATES

AI-03 must implement:

Initial
Loading
Success
Empty
Partial Failure
Complete Failure
Article Unavailable
PDF Loading
PDF Success
PDF Failure

---

14. UX LANGUAGE

Preferred language:

SEARCH
BACA JURNAL
OPEN ORIGINAL ARTICLE
EXPORT PDF
SEARCHING...
FULL TEXT AVAILABLE
FULL TEXT NOT AVAILABLE

Error language should be simple and human-readable.

---

15. IMPLEMENTATION RULE

AI-03 must implement the approved UX/UI specification.

AI-03 must not independently redesign the product.

If AI-03 finds a technical conflict:

1. Document it.
2. Explain the conflict.
3. Propose a minimal change.
4. Preserve the original UX intention.

---

16. ACCEPTANCE CRITERIA

AI-02 is accepted when:

- Search UI is defined.
- Results UI is defined.
- Article UI is defined.
- Reader UI is defined.
- PDF UX is defined.
- Responsive rules are defined.
- Accessibility is defined.
- Design tokens are defined.
- Component behavior is defined.
- All major states are defined.
- AI-03 can implement without guessing.

---

17. AI-03 INSTRUCTIONS

Next AI:

AI-03 — FRONTEND

AI-03 must:

1. Read "SHARED.md".
2. Read "HANDOFF_AI01.md".
3. Read this "HANDOFF_AI02.md".
4. Implement the approved UX/UI.
5. Use Next.js + TypeScript.
6. Create reusable components.
7. Implement responsive behavior.
8. Implement loading/error/empty states.
9. Implement accessibility.
10. Do not implement backend search logic.
11. Do not modify the database.
12. Do not invent new product features.

Expected output:

Working frontend
+
Tests
+
HANDOFF_AI03.md

---

18. FINAL AI-02 CHECK

Before marking COMPLETED, verify:

[ ] Search page complete

[ ] Search input defined

[ ] Search button defined

[ ] Results page complete

[ ] Article card defined

[ ] Article detail defined

[ ] Full-text states defined

[ ] PDF export defined

[ ] Loading states defined

[ ] Empty states defined

[ ] Error states defined

[ ] Mobile defined

[ ] Tablet defined

[ ] Desktop defined

[ ] Accessibility defined

[ ] Design tokens defined

[ ] Components defined

[ ] AI-03 instructions complete
