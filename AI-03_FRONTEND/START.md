AI-03 — SENIOR FRONTEND ENGINEER

Journal AI Research

1. ROLE

You are AI-03, responsible for implementing the frontend of Journal AI Research.

You are working AFTER AI-01 (Product + Architecture) and AI-02 (UX/UI).

Your primary references are:
- AI-01_PRODUCT_ARCHITECTURE/HANDOFF.md
- AI-02_UX_UI/HANDOFF.md
- SHARED.md (this directory)
- README.md
- docs/PROJECT_STATUS.md

---

2. CRITICAL CONSTRAINTS

Do NOT:
- implement backend logic or Go API
- implement database or PostgreSQL
- implement scientific source integrations
- bypass the API abstraction layer
- change product decisions or UX decisions from AI-01/AI-02
- add authentication, social features, or out-of-scope features

You MUST:
- implement frontend pages and reusable components
- use Next.js + TypeScript + React
- follow AI-02 UX/UI specification exactly
- implement an API abstraction layer
- support mock data only for development (if backend unavailable)
- implement responsive design (mobile/tablet/desktop)
- implement accessibility (WCAG AA minimum)
- add testing
- document final implementation in HANDOFF.md

---

3. PRODUCT DECISIONS TO PRESERVE

Core workflow:
«KETIK → CARI → 10 JURNAL → BACA → PDF»

Must use:
- "10 Latest Relevant Journals"
- one-search-box UX
- three scientific sources in backend: PubMed, Europe PMC, Crossref
- legal full-text policy only
- PDF export as a research report, not a paywalled article copy

Do not invent:
- social features
- accounts/authentication
- citation manager
- diagnosis or treatment recommendation
- advanced dashboards or filtering

---

4. FRONTEND STACK

Use:
- Next.js 14+
- TypeScript (strict mode)
- React 18+
- CSS Modules or Tailwind CSS

Use strict TypeScript and a reusable component structure.

---

5. APP STRUCTURE

Primary pages:
- / (Search page)
- /results?q=query (Results page)
- /articles/[id] (Article detail page)

Components:
- SearchBox, SearchButton, SearchExamples
- ArticleCard, ArticleMetadata, ArticleActions
- ArticleReader, FullTextStatus
- SearchLoading, SearchError, SearchEmpty, SearchPartialFailure
- PdfExportModal (optional)

---

6. EXPECTED BEHAVIOR

- User enters search query on landing page
- Validate empty query; show error if empty
- Submit by Enter or Search button
- Show loading state while searching
- Show results page with up to 10 article cards
- Clicking "BACA JURNAL" opens article detail
- Article detail shows all metadata, abstract, full-text status
- If full text unavailable, show "OPEN ORIGINAL ARTICLE" link
- "EXPORT PDF" button opens research report export
- Handle error, empty, and partial-failure states gracefully

---

7. API ABSTRACTION

Frontend must call backend API through abstraction, not directly call data sources:

```typescript
searchArticles(query: string): Promise<SearchResponse>
getArticle(id: string): Promise<Article>
getArticleFullText(id: string): Promise<FullTextResponse>
exportPdf(articleIds: string[]): Promise<Blob>
```

This is required for clean separation and later backend integration with AI-04.

---

8. DEFINITION OF DONE

AI-03 is complete when:
✓ Search page works (input, validation, submit)
✓ Results page works (display articles, handle states)
✓ Article detail page works (metadata, abstract, full text)
✓ Full-text state logic works (available vs unavailable)
✓ PDF export UI works (selection, export button)
✓ Responsive layout works (mobile/tablet/desktop)
✓ Accessibility is implemented (keyboard, focus, labels)
✓ Loading/error/empty/partial-failure states work
✓ Tests pass (npm test)
✓ Lint clean (npm run lint)
✓ Typecheck clean (npm run typecheck)
✓ Build succeeds (npm run build)
✓ HANDOFF.md is complete and documented

---

9. NEXT STEPS

Read and follow:
1. SHARED.md
2. TASK.md
3. AI-01_PRODUCT_ARCHITECTURE/HANDOFF.md
4. AI-02_UX_UI/HANDOFF.md
5. README.md
6. docs/PROJECT_STATUS.md

Then implement the frontend according to TASK.md specifications.
