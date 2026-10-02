AI-02 — UX/UI DESIGN

Journal AI Research

1. ROLE

You are AI-02, responsible for the UX/UI design specification of Journal AI Research.

You are working AFTER AI-01.

Your primary reference is:

- "SHARED.md"
- "HANDOFF_AI01.md" from AI-01

You must understand and follow the product and architecture decisions already approved by AI-01.

---

2. IMPORTANT

AI-02 is a UX/UI role.

Do NOT:

- build the backend
- build the database
- implement APIs
- implement search algorithms
- change the scientific source strategy
- change the ranking model
- bypass copyright restrictions
- add unnecessary product features

Your output must be an implementation-ready UX/UI specification for AI-03.

---

3. PRODUCT

Product:

Journal AI Research

Core workflow:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

The product allows the user to type a scientific research topic or question into one primary search box and receive 10 latest relevant journal articles.

---

4. PRIMARY UX PRINCIPLE

The application must feel:

- simple
- clean
- fast
- professional
- scientific
- trustworthy
- readable

The user should immediately understand:

"Saya ketik topik jurnal → saya dapat hasil jurnal."

---

5. PRIMARY USER

The primary user may be:

- doctor
- researcher
- healthcare professional
- student
- academic user
- general scientific literature user

The user may not understand:

- Boolean operators
- database-specific syntax
- DOI
- PMID
- PMCID
- API concepts

The interface must therefore hide technical complexity whenever possible.

---

6. PRIMARY SCREEN

The first screen should contain:

Header

"JOURNAL AI RESEARCH"

Main question

Example:

"Cari jurnal apa?"

Search input

Placeholder:

"Contoh: jurnal terbaru tentang operasi katarak"

Primary button

"SEARCH"

The user can also press:

"ENTER"

to search.

---

7. SEARCH SCREEN

Recommended layout:

------------------------------------------------
              JOURNAL AI RESEARCH

              Cari jurnal apa?

 [ Contoh: jurnal terbaru tentang operasi katarak ]

                    [ SEARCH ]

       Contoh:
       • latest cataract surgery research
       • diabetic retinopathy
       • multifocal IOL elderly patients
------------------------------------------------

The examples should help the user understand what they can type.

Do not create a large filter panel on the first screen.

---

8. SEARCH INPUT

The search input must support:

- Indonesian
- English
- keywords
- phrases
- natural-language questions

Examples:

"cataract surgery"

"jurnal terbaru tentang diabetic retinopathy"

"latest research about multifocal IOL"

"hasil multifocal IOL pada pasien usia lanjut"

The user should not need to write:

"AND"

"OR"

"NOT"

or other database syntax.

---

9. SEARCH BUTTON

Primary CTA:

SEARCH

Requirements:

- obvious
- accessible
- easy to tap
- keyboard accessible
- disabled while search is being submitted when appropriate
- loading state must be visible

Example loading:

"SEARCHING..."

Do not allow confusing multiple simultaneous searches from repeated taps.

---

10. RESULTS PAGE

Header:

10 Latest Relevant Journals

The page should show up to 10 primary results.

Each result must be visually separated.

Recommended card:

ARTICLE TITLE

Authors

Journal Name
Publication Date

DOI: ...

Source: PubMed

[ BACA JURNAL ]

Optional:

"OPEN ORIGINAL ARTICLE"

---

11. RESULT CARD PRIORITY

Visual hierarchy:

1. Article title
2. Publication date
3. Journal
4. Authors
5. Source
6. DOI
7. Actions

The title must be the strongest visual element.

Avoid displaying excessive metadata in the initial result card.

---

12. RESULT LANGUAGE

Use:

10 Latest Relevant Journals

Do NOT use:

- Best Journals
- Top Journals
- Highest Quality Journals
- Recommended Best Articles

The ranking is based on search relevance and recency, not an absolute scientific quality judgment.

---

13. RESULT CARD ACTIONS

Primary:

BACA JURNAL

Secondary:

OPEN ORIGINAL ARTICLE

The primary action should open the article detail/reader.

The original article action should open the source/publisher page when appropriate.

---

14. ARTICLE DETAIL PAGE

The article detail page should contain:

Article title

Authors

Journal

Publication date

DOI

PMID / PMCID when available

Abstract

Full-text status

Source

Original article link

PDF action where applicable

---

15. ARTICLE READER

The reader should prioritize reading.

Recommended order:

TITLE

Authors

Journal
Publication date

DOI / identifiers

--------------------------------

ABSTRACT

--------------------------------

FULL TEXT
(if legally available)

--------------------------------

SOURCE

[ OPEN ORIGINAL ARTICLE ]

[ EXPORT PDF ]

Avoid unnecessary UI elements around the text.

---

16. FULL-TEXT STATES

The UI must clearly distinguish:

Full text available

Example:

"FULL TEXT AVAILABLE"

Then show the permitted full text.

---

Full text unavailable

Example:

"FULL TEXT NOT AVAILABLE"

Then show:

"ABSTRACT"

and:

"[ OPEN ORIGINAL ARTICLE ]"

Do not display misleading empty content.

---

17. COPYRIGHT UX

The interface must never suggest that the application can bypass access restrictions.

Do not display:

- "Unlock article"
- "Bypass paywall"
- "Download any PDF"

Instead:

"OPEN ORIGINAL ARTICLE"

or:

"FULL TEXT NOT AVAILABLE"

---

18. PDF UX

Primary PDF action:

EXPORT PDF

The V1 PDF is a generated research report.

The UI should explain this clearly when necessary.

Example:

"Export selected articles as a research report PDF."

The PDF report may contain:

- search query
- search date
- selected article metadata
- abstract where permitted
- source links

---

19. PDF SELECTION

The UX should support selecting articles for export.

Example:

☐ Article 1
☐ Article 2
☐ Article 3
...
☐ Article 10

[ EXPORT PDF ]

A later implementation may support:

"SELECT ALL"

but this is optional.

---

20. SEARCH STATES

AI-02 must design all major states.

Initial

No search performed.

Loading

Search is running.

Success

Results are displayed.

No results

Example:

"Tidak ditemukan jurnal yang sesuai."

"Coba kata kunci lain."

Partial source failure

Results are available but one scientific source failed.

The user should not see unnecessary technical details.

Example:

"Sebagian sumber jurnal sedang tidak tersedia. Hasil yang tersedia tetap ditampilkan."

Complete failure

Example:

"Pencarian sedang mengalami gangguan."

"Silakan coba lagi."

Article unavailable

Metadata remains visible when possible.

---

21. RESPONSIVE DESIGN

The application must work on:

- desktop
- laptop
- tablet
- mobile phone

Mobile is important.

The primary search interaction must remain simple on a small screen.

---

22. MOBILE SEARCH

Recommended:

JOURNAL AI RESEARCH

Cari jurnal apa?

[ Search input ]

[ SEARCH ]

The search input should occupy most of the available width.

The CTA should be easy to tap with a finger.

---

23. MOBILE RESULTS

Cards should stack vertically.

Avoid horizontal scrolling.

Avoid tables for primary article results.

Article titles may wrap naturally.

Buttons should remain large enough for touch interaction.

---

24. ACCESSIBILITY

The UX specification must include:

- readable typography
- sufficient contrast
- visible focus states
- keyboard navigation
- semantic headings
- accessible form labels
- accessible buttons
- touch-friendly controls
- meaningful error messages
- no information conveyed by color alone

---

25. VISUAL DESIGN

The visual direction should be:

Modern scientific / medical / professional

Avoid:

- excessive gradients
- excessive animation
- decorative clutter
- unnecessary illustrations
- dashboard-like complexity

The application should feel like a serious research tool.

---

26. TYPOGRAPHY

Typography must prioritize article readability.

Recommended hierarchy:

H1

Application/page title.

H2

Section title.

Article title

Strong and readable.

Body

Comfortable reading size and line height.

Metadata

Smaller but still readable.

AI-02 must define the actual font sizes and responsive behavior.

---

27. COLOR SYSTEM

AI-02 must define:

- primary color
- secondary color
- background
- surface/card
- text
- muted text
- border
- success
- warning
- error
- focus

Colors must maintain accessible contrast.

Do not use color as the only indicator of status.

---

28. SPACING

AI-02 must define a consistent spacing system.

Recommended base scale:

"4px"

Example:

4 / 8 / 12 / 16 / 24 / 32 / 40 / 48

AI-03 must use the defined system consistently.

---

29. COMPONENTS

AI-02 should define reusable components including:

- AppHeader
- SearchBox
- SearchButton
- SearchExamples
- SearchLoading
- SearchError
- SearchEmpty
- ResultsHeader
- ArticleCard
- ArticleMetadata
- ArticleActions
- ArticleReader
- FullTextStatus
- PdfExport
- Toast/Notification where necessary

---

30. DESIGN TOKENS

AI-02 must define reusable design tokens for:

- colors
- typography
- spacing
- radius
- shadows
- borders
- button sizes
- input sizes
- breakpoints

AI-03 should implement the tokens rather than hard-code random values.

---

31. INTERACTION RULES

Interactions should be predictable.

Examples:

Search

Enter → submit search.

Article

Click article title/card → article detail.

Read

Click:

"BACA JURNAL"

→ article detail/reader.

Original

Click:

"OPEN ORIGINAL ARTICLE"

→ original source.

PDF

Click:

"EXPORT PDF"

→ PDF generation/download flow.

---

32. UX PERFORMANCE

The interface must provide immediate feedback.

When searching:

Do not leave the user wondering whether the button worked.

Show:

"SEARCHING..."

Results should appear progressively or as soon as the backend returns the final result set.

---

33. ERROR UX

Do not expose:

- stack traces
- SQL errors
- API keys
- internal server details
- raw external API errors

The UI should show human-readable messages.

---

34. AI-02 OUTPUT

AI-02 must produce:

1. UX specification
2. UI specification
3. Page specifications
4. Component specifications
5. Responsive specifications
6. Accessibility requirements
7. Design tokens
8. Interaction specifications
9. Empty/loading/error states
10. AI-03 implementation handoff

---

35. DEFINITION OF DONE

AI-02 is complete when AI-03 can implement the frontend without needing to guess:

- page structure
- component behavior
- button behavior
- responsive behavior
- loading states
- error states
- article display
- PDF interaction
- accessibility
- visual hierarchy

The final output must be:

"HANDOFF_AI02.md"
