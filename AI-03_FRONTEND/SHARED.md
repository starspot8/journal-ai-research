SHARED PROJECT SPECIFICATION

Journal AI Research Frontend

1. CORE WORKFLOW

«KETIK → CARI → 10 JURNAL → BACA → PDF»

Search box → Search results → Article detail → PDF export

---

2. PAGES

V1 Frontend has 3 core pages:

1. Search Page (/)
   - Primary landing page
   - Single search box
   - Examples
   - Loading state

2. Results Page (/results)
   - Query parameter: ?q=search+term
   - Up to 10 article cards
   - States: loading, results, empty, error, partial-failure

3. Article Detail Page (/articles/[id])
   - Full article metadata
   - Abstract
   - Full text if available
   - Actions: OPEN ORIGINAL, EXPORT PDF

---

3. KEY TERMINOLOGY

Use:
- "10 Latest Relevant Journals"
- "BACA JURNAL" (Read Article)
- "OPEN ORIGINAL ARTICLE"
- "EXPORT PDF"
- "FULL TEXT AVAILABLE"
- "FULL TEXT NOT AVAILABLE"

Do NOT use:
- "Best Journals", "Top Journals", "Number One Journal"
- "Fetch Article Content", "Redirect to Source"

---

4. DESIGN SYSTEM (from AI-02)

Colors:
- Primary: #1E40AF (blue)
- Secondary: #0EA5E9
- Success: #10B981
- Warning: #F59E0B
- Error: #EF4444
- Background: #FFFFFF
- Surface: #F9FAFB
- Text: #111827

Typography:
- H1: 32px (desktop), 24px (mobile), weight 700
- H2: 24px (desktop), 20px (mobile), weight 600
- Body: 16px, weight 400, line-height 1.6
- Metadata: 14px, weight 400

Spacing: 4px base unit (4 / 8 / 12 / 16 / 24 / 32 / 40 / 48 / 64)

Breakpoints:
- Mobile: 320-767px
- Tablet: 768-1023px
- Desktop: 1024px+

---

5. API ABSTRACTION

Frontend must define:

```typescript
interface SearchResponse {
  articles: Article[];
  totalCount: number;
  query: string;
  hasPartialFailure?: boolean;
}

interface Article {
  id: string;
  doi?: string;
  pmid?: string;
  pmcid?: string;
  title: string;
  abstract?: string;
  authors?: string[];
  journal?: string;
  publication_date?: string;
  source: string;
  source_url: string;
  fulltext_url?: string;
  is_open_access?: boolean;
}

// Functions
async function searchArticles(query: string): Promise<SearchResponse>
async function getArticle(id: string): Promise<Article>
async function getArticleFullText(id: string): Promise<{ fulltext: string | null; status: string }>
async function exportPdf(articleIds: string[]): Promise<Blob>
```

---

6. MOCK DATA

If backend unavailable, provide mock implementation:

- Use development-only mock data
- Clearly mark as mock (comments, separate files)
- Never silently fallback from real API to mock
- Must switch to real backend without code changes

Example mock article:

```typescript
{
  id: "mock-1",
  title: "Sample Cataract Surgery Article",
  authors: ["Smith J.", "Johnson K."],
  journal: "The Lancet",
  publication_date: "2024-01-15",
  abstract: "This is a sample abstract...",
  source: "PubMed",
  source_url: "https://pubmed.ncbi.nlm.nih.gov/...",
  is_open_access: true
}
```

---

7. RESPONSIVE BEHAVIOR

Mobile (320-767px):
- Full-width, single column
- 16px padding
- Stack buttons vertically
- 44px minimum touch targets

Tablet (768-1023px):
- Max-width 800px, centered
- 20-24px padding
- Inline buttons

Desktop (1024px+):
- Max-width 900px, centered
- 32px padding
- Optional 2-column results layout

---

8. ACCESSIBILITY

Minimum WCAG AA:

- Semantic HTML (<header>, <main>, <button>, <a>, proper heading hierarchy)
- Labels for all inputs (htmlFor linking)
- Focus indicators (3px #1E40AF outline)
- Keyboard navigation (Tab, Shift+Tab, Enter, Escape)
- Screen reader support (aria-live, aria-label, aria-describedby)
- Color contrast 4.5:1 (normal text), 3:1 (large text)
- Touch targets 44x44px minimum
- No autofocus unless necessary
- Error messages clear and actionable

---

9. FULL-TEXT POLICY

Frontend displays:

Full text available:
- Show permitted full text
- Clear source reference

Full text unavailable:
- Show: "FULL TEXT NOT AVAILABLE"
- Show abstract if available
- Provide: "OPEN ORIGINAL ARTICLE" link

Never:
- Bypass paywalls
- Circumvent access controls
- Show misleading "unlock" or "download" options

---

10. PDF FEATURE

PDF is a research report, not article copies:

- Modal/interface to select articles
- "Select articles to include in your research report PDF"
- Export button triggers PDF generation
- Backend handles PDF generation
- Frontend receives PDF blob and triggers download

---

11. ERROR HANDLING

Do NOT expose:
- Stack traces
- SQL errors
- API keys
- Internal server details
- Raw external API errors

Show user-friendly messages:
- "Tidak ditemukan jurnal yang sesuai." (no results)
- "Pencarian sedang mengalami gangguan." (search error)
- "Silakan coba lagi." (retry prompt)

---

12. STATE MANAGEMENT

Use React Hooks + Context API:

Context examples:
- SearchContext (query, results, loading, error)
- ArticleDetailContext (article, fulltext, loading)

Alternative: Simple useState in pages (for V1 simplicity)

---

13. ENVIRONMENT VARIABLES

.env.local or .env.development:

```
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
NEXT_PUBLIC_API_VERSION=v1
```

No secrets in frontend!

---

14. TESTING STRATEGY

Unit tests:
- Components (SearchBox, ArticleCard, etc.)
- Utilities
- Custom hooks

Integration tests:
- Page flows (search → results → detail)
- API calls (with mocks)

Accessibility tests:
- Focus management
- Keyboard navigation
- Contrast checks

Example Jest + React Testing Library.

---

15. PRODUCTION READINESS

Before handoff:

✓ npm run lint → no errors
✓ npm run typecheck → no errors
✓ npm test → all pass
✓ npm run build → succeeds
✓ No console errors/warnings in development
✓ Performance: Lighthouse score acceptable
✓ Mobile responsive works
✓ Accessibility audit passes

---

16. NEXT AI STAGE

After AI-03 Frontend:

AI-04 Backend + Scientific Search

Will implement:
- Go API layer
- PubMed integration
- Europe PMC integration
- Crossref integration
- Normalization
- Deduplication
- Ranking
- PDF generation backend

Frontend must integrate with AI-04's API.
