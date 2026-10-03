HANDOFF AI-02

UX/UI Specification

Journal AI Research

1. STATUS

COMPLETED

AI-02 has finalized all UX/UI specifications for V1 Journal AI Research. The design is implementation-ready for AI-03 Frontend.

---

2. OBJECTIVE

AI-02 converts the AI-01 product and architecture specification into a complete, implementation-ready UX/UI specification.

This handoff enables AI-03 to build the frontend without guessing about page layout, component behavior, or interaction patterns.

---

3. CORE PRINCIPLE

The application must feel simple, professional, and trustworthy while maintaining scientific credibility.

Core workflow:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

4. DESIGN SYSTEM

4.1 Color Palette

Primary Color: #1E40AF (professional blue)
Secondary Color: #0EA5E9 (lighter blue)
Success Color: #10B981 (green)
Warning Color: #F59E0B (amber)
Error Color: #EF4444 (red)

Neutral Colors:
- Background: #FFFFFF
- Surface/Card: #F9FAFB
- Text Primary: #111827
- Text Secondary: #6B7280
- Text Muted: #9CA3AF
- Border: #E5E7EB
- Focus Ring: #1E40AF with 3px outline

4.2 Typography

Font: System stack or 'Inter', 'Segoe UI', sans-serif

H1 (Page Title)
- Size: 32px (desktop), 24px (mobile)
- Weight: 700
- Line-height: 1.2
- Letter-spacing: -0.01em
- Example: "10 Latest Relevant Journals"

H2 (Section Title)
- Size: 24px (desktop), 20px (mobile)
- Weight: 600
- Line-height: 1.35
- Example: "Search Results"

Article Title
- Size: 20px (desktop), 18px (mobile)
- Weight: 600
- Line-height: 1.4
- Color: #111827
- Applies to article cards and detail page

Body Text
- Size: 16px
- Weight: 400
- Line-height: 1.6
- Color: #111827
- Used for readable content, abstracts, and full text

Metadata
- Size: 14px
- Weight: 400
- Line-height: 1.5
- Color: #6B7280
- Used for authors, journal, date, source

Caption/Label
- Size: 12px
- Weight: 500
- Line-height: 1.5
- Color: #9CA3AF
- Used for field labels and helper text

4.3 Spacing Scale

Base unit: 4px

Scale: 4 / 8 / 12 / 16 / 24 / 32 / 40 / 48 / 64

Usage:
- 4/8px: small gaps, tight spacing
- 12/16px: component padding, small sections
- 24/32px: section spacing, large gaps
- 40/48px: page-level vertical spacing
- 64px: major section breaks

4.4 Border Radius

- Inputs, buttons: 6px
- Cards: 8px
- Modals: 12px

4.5 Shadows

Subtle: 0 1px 2px 0 rgba(0, 0, 0, 0.05)
Small: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)
Medium: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)
Large: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)

Cards use: small shadow

4.6 Breakpoints

Mobile: 320px - 767px
Tablet: 768px - 1023px
Desktop: 1024px+

---

5. SEARCH PAGE (PRIMARY LANDING PAGE)

5.1 Layout Structure

Mobile (320-767px)
- Full-screen, single column
- Header padding: 24px top/bottom, 16px left/right
- Search box: 100% width, max-width constrained by screen
- Button: 100% width

Tablet (768-1023px)
- Centered column, max-width 600px
- Padding: 32px top/bottom, 20px left/right

Desktop (1024px+)
- Centered column, max-width 700px
- Padding: 48px top/bottom, 32px left/right

5.2 Content Hierarchy

Component: AppHeader
- Content: "JOURNAL AI RESEARCH"
- Typography: H1
- Color: #1E40AF
- Alignment: center
- Margin-bottom: 32px

Component: SearchPrompt
- Content: "Cari jurnal apa?" (or "What journal are you looking for?")
- Typography: H2
- Color: #111827
- Margin-bottom: 16px

Component: SearchBox
- Input field with placeholder text
- Placeholder: "Contoh: jurnal terbaru tentang operasi katarak"
- Placeholder color: #9CA3AF
- Size: 16px
- Width: 100%
- Padding: 12px 16px
- Border: 1px solid #E5E7EB
- Border-radius: 6px
- Margin-bottom: 16px
- Focus state: 3px solid #1E40AF outline
- Background on focus: #FFFFFF

Component: SearchButton
- Text: "SEARCH"
- Size: 16px, weight: 600
- Padding: 12px 32px
- Width: 100% (mobile), auto (tablet/desktop)
- Background: #1E40AF
- Color: #FFFFFF
- Border-radius: 6px
- Border: none
- Cursor: pointer
- Hover state: background #1E3A8A (darker blue)
- Active state: background #1E3A8A
- Focus state: 3px solid #1E40AF outline, offset 2px
- Disabled state: background #D1D5DB, cursor not-allowed, color #9CA3AF
- Loading state: display "SEARCHING..." text with spinner

Component: SearchExamples
- Margin-top: 24px
- Heading: "Contoh:" (size 14px, weight 500, color #6B7280)
- List of examples as clickable text links or simple text
- Example 1: "latest cataract surgery research"
- Example 2: "jurnal terbaru diabetic retinopathy"
- Example 3: "multifocal IOL elderly patients"
- Spacing between examples: 8px
- Color: #0EA5E9
- Hover (if clickable): underline, color #1E40AF

5.3 Search Input Behavior

- Accepts text input (no character restrictions beyond length)
- Supports Indonesian and English characters
- Supports natural-language questions and keywords
- Maximum length: 500 characters (can be adjusted by backend if needed)
- No special filtering; query passed as-is to backend
- On Enter key: submit search
- On click Search button: submit search
- Empty query: show validation error "Masukkan topik jurnal terlebih dahulu." before submission

---

6. RESULTS PAGE

6.1 Page Layout

Header Section
- Back navigation (mobile/tablet only): "<" button or "← Search"
- Page title: "10 Latest Relevant Journals"
- Typography: H1
- Alignment: left
- Padding: 24px (mobile), 32px (tablet/desktop)
- Margin-bottom: 24px

Results Container
- Single column on mobile/tablet
- Up to 2 columns on desktop (optional, can be 1 column for simplicity)
- Gap between columns: 24px (desktop only)
- Padding: 0 16px (mobile), 0 20px (tablet), 0 32px (desktop)

6.2 Article Card Component

Card Structure (ArticleCard)

Container
- Background: #F9FAFB
- Border: 1px solid #E5E7EB
- Border-radius: 8px
- Padding: 16px
- Margin-bottom: 16px
- Shadow: small
- Hover state (desktop): shadow medium, scale 1.02 (optional)

Article Title (ArticleCardTitle)
- Typography: Article Title (20px, weight 600)
- Line-height: 1.4
- Margin-bottom: 12px
- Color: #111827
- Text truncation: show full title, allow wrapping

Authors (ArticleCardAuthors)
- Typography: Metadata (14px)
- Color: #6B7280
- Margin-bottom: 8px
- Format: "Author 1, Author 2, Author 3, ..."
- Truncate after 3 authors if more exist: "Author 1, Author 2, Author 3, et al."

Journal & Date Row (ArticleCardMetadata)
- Typography: Metadata (14px)
- Color: #6B7280
- Margin-bottom: 8px
- Format: "Journal Name · YYYY-MM-DD" or "Journal Name · Published YYYY"
- Example: "The Lancet · 2024-01-15"

DOI & Source Row (ArticleCardIds)
- Typography: Metadata (14px)
- Color: #6B7280
- Margin-bottom: 12px
- Display DOI if available: "DOI: 10.xxxx/xxxxx" (clickable link to https://doi.org/...)
- Display source: "Source: PubMed" or "Source: Europe PMC" or "Source: Crossref"
- If multiple sources: "Sources: PubMed, Europe PMC"

Actions Row (ArticleCardActions)
- Spacing: 8px between buttons
- Button 1: "BACA JURNAL" (primary action)
  - Style: solid background #1E40AF, white text
  - Size: 14px, weight: 600
  - Padding: 10px 16px
  - On click: navigate to /articles/:id
- Button 2: "OPEN ORIGINAL" (secondary action, optional if available)
  - Style: outline border #E5E7EB, text #1E40AF
  - Size: 14px, weight: 600
  - Padding: 10px 16px
  - On click: open source URL in new tab

6.3 Results Page States

Loading State (SearchLoading)
- Display spinner or animated loading indicator
- Message: "Mencari jurnal..."
- Position: center of page
- Spacing: 48px from top

Empty Results State (SearchEmpty)
- Display message: "Tidak ditemukan jurnal yang sesuai."
- Subtext: "Coba kata kunci lain atau cari topik yang berbeda."
- Button: "KEMBALI" or "← SEARCH AGAIN"
- Position: center of page

Partial Source Failure State (SearchPartialFailure)
- Display available results
- Show warning banner at top:
  - Message: "Sebagian sumber jurnal sedang tidak tersedia. Hasil yang tersedia tetap ditampilkan."
  - Background: #FEF3C7
  - Text color: #92400E
  - Padding: 12px 16px
  - Border-left: 4px solid #F59E0B

Complete Failure State (SearchError)
- Display message: "Pencarian sedang mengalami gangguan."
- Subtext: "Silakan coba lagi nanti."
- Button: "KEMBALI" or "← SEARCH AGAIN"
- Position: center of page
- Do not expose technical error details

---

7. ARTICLE DETAIL / READER PAGE

7.1 Page Structure

Header Navigation
- Back button: "<" or "← Results" (mobile/tablet)
- Page title: Article title or "Article" (desktop)
- Actions bar (desktop only): "EXPORT PDF" button top-right

Article Header Section (ArticleHeader)

Title (ArticleTitle)
- Typography: 24px (desktop), 20px (mobile), weight 600
- Color: #111827
- Line-height: 1.4
- Margin-bottom: 16px

Metadata Group 1 (ArticleMetadataTop)
- Authors
  - Typography: Metadata
  - Format: "Author 1, Author 2, Author 3, et al."
  - Margin-bottom: 8px
- Journal
  - Typography: Metadata
  - Format: "Journal Name"
  - Margin-bottom: 8px
- Publication Date
  - Typography: Metadata
  - Format: "Published: YYYY-MM-DD" or "Published: Month Year"
  - Margin-bottom: 8px

Metadata Group 2 (ArticleIdentifiers)
- DOI
  - Display as link: "DOI: 10.xxxx/xxxxx"
  - Links to https://doi.org/...
  - Margin-bottom: 4px
- PMID
  - Display as text if available: "PMID: 12345678"
  - Margin-bottom: 4px
- PMCID
  - Display as text if available: "PMCID: PMC12345678"
  - Margin-bottom: 8px

Source Information (ArticleSource)
- Format: "Source: PubMed" or "Source: Europe PMC"
- If multiple sources: list all
- Margin-bottom: 16px
- Typography: Metadata
- Color: #6B7280

---

7.2 Divider

Horizontal rule (1px solid #E5E7EB) with margin: 24px 0

---

7.3 Abstract Section (AbstractSection)

Heading
- Typography: H2 (20px, weight 600)
- Text: "Abstract"
- Margin-bottom: 12px
- Color: #111827

Content
- Typography: Body (16px)
- Color: #111827
- Line-height: 1.6
- If abstract is very long (>1000 words), can add "Show more" toggle
- Margin-bottom: 24px

---

7.4 Full-Text Section (FullTextSection)

Case 1: Full Text Is Legally Available
- Heading: "Full Text" (H2)
- Content: Display permitted full text
- Typography: Body (16px), line-height 1.6
- Margin-bottom: 24px

Case 2: Full Text Is NOT Available
- Status Box (FullTextStatus)
  - Background: #FEF3C7
  - Border-left: 4px solid #F59E0B
  - Padding: 12px 16px
  - Message: "FULL TEXT NOT AVAILABLE"
  - Subtext: "This article's full text is not publicly accessible through available sources."
  - Margin-bottom: 16px

Case 3: Abstract Is Available (when full text not available)
- Heading: "Abstract"
- Content: Display abstract
- Typography: Body (16px), line-height 1.6
- Color: #111827

---

7.5 Actions Section (ArticleActions)

Container
- Margin-top: 24px
- Padding-top: 24px
- Border-top: 1px solid #E5E7EB
- Spacing between buttons: 12px

Primary Action: OPEN ORIGINAL ARTICLE
- Style: solid background #1E40AF, white text
- Size: 14px, weight 600
- Padding: 12px 24px
- On click: open source URL in new tab
- Icon: external link icon (optional)

Secondary Action: EXPORT PDF
- Style: outline border #1E40AF, text #1E40AF
- Size: 14px, weight 600
- Padding: 12px 24px
- On click: navigate to PDF export flow or trigger PDF download

Mobile Layout
- Stack buttons vertically
- Full width: 100%

Desktop Layout
- Inline buttons with gap 12px

---

7.6 Reader Page Responsive Behavior

Mobile (320-767px)
- Single column
- Padding: 16px
- Font sizes reduced (body 14px, headings 18px)
- Full-width content

Tablet (768-1023px)
- Single column, max-width 800px, centered
- Padding: 24px
- Standard font sizes

Desktop (1024px+)
- Single column, max-width 900px, centered
- Padding: 32px
- Full typography scale

---

8. PDF EXPORT INTERACTION

8.1 PDF Modal / Overlay (PDFExportModal)

Trigger
- Button: "EXPORT PDF" on article detail page
- Opens modal or navigates to /export/pdf

Modal Content

Title
- "Export as Research Report"
- Typography: H2

Description
- "Select articles to include in your research report PDF."
- Typography: Body, color #6B7280

Article Selection

Layout
- If viewing single article: auto-include current article
- If viewing results: show checkbox list of all 10 articles

Each item:
- Checkbox (accessible)
- Article title (left-aligned)
- Publication year (right-aligned, muted color)
- Hover state: background #F9FAFB
- Margin-bottom: 8px

Select All (optional)
- Checkbox + text: "Select All"
- Margin-bottom: 16px
- Visible only if showing multiple articles

Export Button
- Text: "EXPORT PDF"
- Style: solid #1E40AF, white text
- Disabled if no articles selected
- Disabled text: "Select at least one article"

Cancel Button (mobile/tablet)
- Text: "CANCEL"
- Style: outline
- On click: close modal

8.2 Loading State (PDFLoading)
- Display: "Generating report..." with spinner
- Button disabled

8.3 Success State (PDFSuccess)
- Display: "Report ready for download"
- Auto-download triggered (browser download)
- Close modal after 2 seconds or on user click

8.4 Error State (PDFError)
- Message: "Could not generate report. Please try again."
- Button: "RETRY"
- On click: retry export

---

9. COMPONENT LIBRARY

Reusable components for AI-03 to implement:

9.1 AppHeader
- Props: title (string)
- Displays application name or page title
- Height: 64px (desktop), 56px (mobile)
- Sticky or fixed (navigation preference)

9.2 SearchBox
- Props: placeholder (string), value (string), onChange (function), onSubmit (function)
- Input field with validation
- Supports Enter key submission

9.3 SearchButton
- Props: text (string), loading (boolean), onClick (function)
- Primary CTA button
- Shows loading state

9.4 SearchExamples
- Props: examples (array of strings)
- Displays example searches
- Optional click handler

9.5 SearchLoading
- Displays loading indicator with text "Mencari jurnal..."
- Centered layout

9.6 SearchError
- Props: message (string), onRetry (function)
- Displays error message and retry button

9.7 SearchEmpty
- Props: message (string), onBack (function)
- Displays empty state

9.8 SearchPartialFailure
- Props: message (string)
- Warning banner with specific styling

9.9 ResultsHeader
- Props: count (number), query (string)
- Displays "10 Latest Relevant Journals" title

9.10 ArticleCard
- Props: article (object), onRead (function), onOpenSource (function)
- Displays article metadata and actions

9.11 ArticleMetadata
- Props: article (object)
- Reusable metadata display for cards and detail pages

9.12 ArticleActions
- Props: article (object), onRead (function), onOpenSource (function)
- Displays action buttons

9.13 ArticleReader
- Props: article (object), onExportPDF (function)
- Full article detail page

9.14 FullTextStatus
- Props: isAvailable (boolean), message (string)
- Displays full-text availability status

9.15 PdfExport
- Props: articles (array), onExport (function), onCancel (function)
- PDF export modal/interface

---

10. LOADING & ERROR STATES

10.1 Search Loading State
- Spinner icon (animated)
- Text: "Mencari jurnal..."
- Display for entire duration of search

10.2 Search Empty State
- Icon: (optional magnifying glass with no results)
- Heading: "Tidak ditemukan jurnal yang sesuai."
- Subtext: "Coba kata kunci lain atau gunakan istilah yang berbeda."
- Button: "← SEARCH AGAIN" or "BACK"

10.3 Search Error State
- Icon: (optional warning or error icon)
- Heading: "Pencarian sedang mengalami gangguan."
- Subtext: "Silakan coba lagi nanti. Jika masalah berlanjut, hubungi dukungan."
- Button: "← TRY AGAIN"
- Do not show technical error details

10.4 Partial Source Failure State
- Banner at top of results
- Icon: (warning icon)
- Message: "Sebagian sumber jurnal sedang tidak tersedia. Hasil yang tersedia tetap ditampilkan."
- Background: #FEF3C7, border-left #F59E0B
- Dismissible or persistent (persistent recommended)
- Results displayed below

10.5 PDF Generation Loading
- Modal/overlay with spinner
- Text: "Generating your research report..."
- Duration: typically < 5 seconds

10.6 PDF Generation Error
- Modal with error message
- Button: "RETRY"
- Suggest user try again or contact support

---

11. RESPONSIVE BEHAVIOR

11.1 Mobile (320-767px)

Search Page
- Full-screen single column
- Header: "JOURNAL AI RESEARCH" (24px)
- Search prompt: "Cari jurnal apa?" (20px)
- Search box: 100% width
- Search button: 100% width
- Examples: stacked list

Results Page
- Single column cards
- No sidebar
- Back button at top
- Card width: 100%

Article Detail Page
- Single column
- Full-width content
- Metadata stacked vertically
- Abstract full-width
- Buttons stacked vertically

11.2 Tablet (768-1023px)

Search Page
- Centered max-width 600px
- Padding: 20px

Results Page
- Single column, max-width 700px, centered
- Cards: 100% width

Article Detail Page
- Single column, max-width 800px, centered
- Larger padding

11.3 Desktop (1024px+)

Search Page
- Centered max-width 700px
- Padding: 32px

Results Page
- Optional 2-column grid (each column max-width 350px, gap 24px)
- Or single column max-width 800px

Article Detail Page
- Single column, max-width 900px, centered
- Optimal reading line length
- Sidebar (optional) for related links

---

12. ACCESSIBILITY REQUIREMENTS

12.1 Semantic HTML

- Use <header>, <main>, <section>, <article> appropriately
- Use <h1>, <h2>, <h3> for proper heading hierarchy
- Use <button> for buttons, not <div> styled as button
- Use <a> for links, with href attribute

12.2 Labels and Inputs

- All form inputs must have associated <label>
- Search input: aria-label="Search for scientific journals"
- Use htmlFor attribute linking label to input id

12.3 Focus Management

- Visible focus indicator on all interactive elements
- Focus ring: 3px solid #1E40AF, offset 2px
- Tab order must be logical and intuitive
- No focus traps

12.4 Screen Readers

- Use aria-live="polite" for dynamic content updates (search results, loading)
- Use aria-label or aria-describedby for icon-only buttons
- Announce loading and error states
- Example: aria-live="polite" on results container

12.5 Keyboard Navigation

- Tab: move focus to next interactive element
- Shift+Tab: move focus to previous interactive element
- Enter: activate button or submit form
- Escape: close modals or overlay

12.6 Color Contrast

- All text must meet WCAG AA standards (4.5:1 for normal text, 3:1 for large text)
- Primary text #111827 on white: 12.6:1 (exceeds requirement)
- Button text #FFFFFF on #1E40AF: 7.5:1 (exceeds requirement)
- Test all color combinations

12.7 Touch Targets

- Minimum touch target size: 44x44px
- Buttons: 44px minimum height
- Links: 44px minimum clickable area
- Spacing between targets: minimum 8px

12.8 Text Sizing

- Do not disable user text zoom (no user-scalable=no)
- Responsive font sizes that scale with viewport
- Minimum 16px base font size

12.9 Alternative Text

- Images: provide meaningful alt text or aria-label
- Icons: provide aria-label if not supplemented by text
- Example: "External link" for link icon

12.10 Error Messages

- Clear, actionable error messages
- Associate error with form field using aria-describedby
- Example: <span id="search-error">Please enter a search term</span>

---

13. INTERACTION SPECIFICATION

13.1 Search Input Interactions

Placeholder Text
- Placeholder: "Contoh: jurnal terbaru tentang operasi katarak"
- Color: #9CA3AF
- Disappears on focus

Focus State
- Border: 1px solid #1E40AF
- Outline: 3px solid #1E40AF (offset 2px)
- Background: #FFFFFF
- Box-shadow: optional glow

Typing
- No restrictions or special handling
- Character counter optional

13.2 Search Button Interactions

Normal State
- Background: #1E40AF
- Text: "SEARCH" (#FFFFFF)
- Cursor: pointer

Hover State
- Background: #1E3A8A (darker)
- Text: #FFFFFF

Active State
- Background: #1E3A8A
- Scale: 0.98 (slight press effect, optional)

Focus State
- Outline: 3px solid #1E40AF (offset 2px)

Disabled State
- Background: #D1D5DB
- Text: #9CA3AF
- Cursor: not-allowed
- Opacity: 0.6

Loading State
- Disable button
- Change text to "SEARCHING..."
- Show spinner icon left of text
- Cursor: wait

13.3 Article Card Interactions

Card Hover (Desktop)
- Slight shadow increase (from small to medium)
- Optional scale: 1.02
- Border remains subtle

Card Focus (Keyboard)
- Outline: 3px solid #1E40AF
- Tab through card elements (title, buttons)

Button Hover
- "BACA JURNAL": background darker (#1E3A8A)
- "OPEN ORIGINAL": background light (#F3F4F6), border darker

13.4 External Link Interactions

- Clicking "OPEN ORIGINAL ARTICLE" opens URL in new tab
- Add target="_blank" and rel="noopener noreferrer"
- Icon: external link icon (optional, after text or before)

13.5 Modal Interactions

Open
- Display overlay/modal
- Focus trap: focus moves to modal
- Background: can use semi-transparent overlay

Close
- Click "CANCEL" button
- Click outside modal (optional)
- Press Escape key
- Focus returns to trigger button

---

14. DESIGN TOKEN REFERENCE TABLE

For AI-03 implementation:

Colors
```
Primary:        #1E40AF
Primary-Dark:   #1E3A8A
Primary-Light:  #3B82F6
Secondary:      #0EA5E9
Success:        #10B981
Warning:        #F59E0B
Error:          #EF4444

Background:     #FFFFFF
Surface:        #F9FAFB
Text-Primary:   #111827
Text-Secondary: #6B7280
Text-Muted:     #9CA3AF
Border:         #E5E7EB
Focus:          #1E40AF
```

Typography
```
H1:          32px desktop, 24px mobile | 700 weight | 1.2 line-height
H2:          24px desktop, 20px mobile | 600 weight | 1.35 line-height
Article:     20px desktop, 18px mobile | 600 weight | 1.4 line-height
Body:        16px | 400 weight | 1.6 line-height
Metadata:    14px | 400 weight | 1.5 line-height
Caption:     12px | 500 weight | 1.5 line-height
```

Spacing
```
2xs:  4px
xs:   8px
sm:   12px
md:   16px
lg:   24px
xl:   32px
2xl:  40px
3xl:  48px
4xl:  64px
```

Radius
```
Input/Button:  6px
Card:          8px
Modal:         12px
```

Shadows
```
Subtle:   0 1px 2px 0 rgba(0, 0, 0, 0.05)
Small:    0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)
Medium:   0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)
Large:    0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)
```

Breakpoints
```
Mobile:   320px - 767px
Tablet:   768px - 1023px
Desktop:  1024px+
```

---

15. CONTENT LANGUAGE REFERENCE

Approved terms and phrases:

Search
- "SEARCH" (button)
- "Mencari jurnal..." (loading state)
- "Cari jurnal apa?" (prompt)

Results
- "10 Latest Relevant Journals" (header)
- "Tidak ditemukan jurnal yang sesuai." (no results)
- "Coba kata kunci lain." (no results suggestion)

Article
- "BACA JURNAL" (read article)
- "OPEN ORIGINAL ARTICLE" (open source)
- "FULL TEXT AVAILABLE" (full text status)
- "FULL TEXT NOT AVAILABLE" (restricted text)

Errors
- "Pencarian sedang mengalami gangguan." (search error)
- "Silakan coba lagi." (retry prompt)
- "Sebagian sumber jurnal sedang tidak tersedia." (partial failure)
- "Hasil yang tersedia tetap ditampilkan." (partial results)

PDF
- "EXPORT PDF" (button)
- "Export as Research Report" (modal title)
- "Select articles to include in your research report PDF." (description)

---

16. NAVIGATION FLOW

16.1 Main Flow

Entry: Search Page
  ↓ (on search submission)
Exit: Results Page
  ↓ (on article click or "BACA JURNAL")
Destination: Article Detail Page
  ↓ (optional actions)
  - "EXPORT PDF": show PDF modal
  - "OPEN ORIGINAL": open new tab
  - Back: return to Results

16.2 Back Navigation

- Results page: show "← Results" or back button
- Article detail: show "← Results" or back button
- Mobile: include back button in header
- Desktop: can use browser back or explicit button

---

17. DECISIONS MADE

DECISION-001
- Search interface remains a single search box with no advanced filters in V1.

DECISION-002
- Results are presented as a vertical list of cards on all screen sizes, with optional 2-column on desktop.

DECISION-003
- Article detail page prioritizes readability with a single-column centered layout.

DECISION-004
- Full-text availability is indicated with clear status labels; no paywall bypass mechanisms.

DECISION-005
- PDF export is positioned as a research report export, not an article PDF copy.

DECISION-006
- Responsive breakpoints are: mobile (320-767), tablet (768-1023), desktop (1024+).

DECISION-007
- Color scheme uses professional blue (#1E40AF) as primary to convey scientific credibility.

DECISION-008
- Typography hierarchy emphasizes article titles and readability over metadata.

DECISION-009
- Accessibility is implemented to WCAG AA standards minimum.

DECISION-010
- All interactive elements support keyboard navigation and screen readers.

---

18. KNOWN LIMITATIONS

- PDF export is a generated research report; it does not include full-text copies of paywalled articles.
- The design does not include advanced filtering, faceted search, or saved searches in V1.
- No user authentication or personalization in V1; all searches are anonymous.
- External source latency is not within UX control; loading state manages user expectations.
- The design assumes next.js will handle routing; SPA or server-side patterns can be adjusted by AI-03 if needed.

---

19. OPEN DECISIONS

OPEN DECISION-001
- Exact animation/transition durations (e.g., modal open, button hover) are not specified and can be adjusted by AI-03 based on performance and platform conventions.
- Why it matters: animations affect perceived responsiveness and browser performance.
- Suggested owner: AI-03 Frontend
- Required decision before: implementation completion

OPEN DECISION-002
- Whether to use a dedicated PDF library (e.g., PDFKit, jsPDF) or server-side generation is not finalized; backend (AI-04) may influence this decision.
- Why it matters: PDF quality and feature parity depend on library choice and architecture.
- Suggested owner: AI-04 Backend, AI-03 Frontend coordination
- Required decision before: PDF export implementation

---

20. IMPLEMENTATION GUIDE FOR AI-03

20.1 Frontend Stack

- Framework: Next.js with TypeScript
- Styling: CSS Modules or Tailwind CSS (tokens mapped to design system)
- State management: React hooks or Context API (simple application)
- HTTP client: fetch or axios
- Testing: Jest + React Testing Library

20.2 Project Structure Recommendation

```
src/
  ├── components/
  │   ├── AppHeader.tsx
  │   ├── SearchBox.tsx
  │   ├── SearchButton.tsx
  │   ├── ArticleCard.tsx
  │   ├── ArticleReader.tsx
  │   ├── FullTextStatus.tsx
  │   ├── PdfExport.tsx
  │   └── ... (other components)
  ├── pages/
  │   ├── index.tsx (search page)
  │   ├── results.tsx (results page)
  │   ├── articles/[id].tsx (article detail)
  │   └── ... (other pages)
  ├── styles/
  │   ├── globals.css
  │   ├── tokens.css (design tokens)
  │   └── ... (component styles)
  ├── types/
  │   ├── article.ts
  │   ├── search.ts
  │   └── ... (type definitions)
  ├── lib/
  │   ├── api.ts (API calls)
  │   ├── utils.ts (utilities)
  │   └── ... (helpers)
  └── tests/
      └── ... (test files)
```

20.3 Implementation Priorities

1. Page structure and layout (HTML/CSS)
2. Component structure and reusability
3. Responsive behavior
4. Basic interactivity (button clicks, form submission)
5. API integration (fetch data from backend)
6. Loading/error/empty states
7. Accessibility improvements
8. Testing
9. Performance optimization

20.4 API Integration

- Expect backend to provide:
  - POST /api/v1/search (returns articles)
  - GET /api/v1/articles/:id (returns article details)
  - POST /api/v1/export/pdf (returns PDF)

- Handle errors gracefully with user-friendly messages
- Show loading states for all API calls
- Implement request timeout handling

20.5 Browser Support

- Modern browsers (Chrome, Firefox, Safari, Edge)
- Mobile browsers (iOS Safari, Chrome Mobile)
- Minimum: ES2020 JavaScript support

---

21. INSTRUCTIONS FOR AI-03

AI-03 is responsible for implementing the approved UX/UI specification.

AI-03 must:

1. Read "SHARED.md", "HANDOFF_AI01.md", and this "HANDOFF_AI02.md" carefully.
2. Implement the design system tokens exactly as specified.
3. Build reusable components following the component library.
4. Implement responsive behavior per the specified breakpoints.
5. Implement all required states (loading, error, empty, partial failure).
6. Implement accessibility features (focus states, keyboard navigation, screen reader support).
7. Implement interactions as specified (Enter key, button behaviors, navigation flow).
8. Use Next.js + TypeScript.
9. Create comprehensive component tests.
10. Do not modify the backend API contract.
11. Do not invent new UX features beyond the specification.
12. Do not implement authentication, social features, or other out-of-scope functionality.

AI-03 must NOT:
- Redesign the user flow
- Change search behavior or ranking logic
- Modify the full-text legal policy
- Add features like login, profiles, or citation management
- Implement backend functionality

Expected deliverables:
- Working Next.js frontend application
- All pages fully implemented and responsive
- All components tested
- Accessibility compliance verified
- HANDOFF_AI03.md with implementation details

---

22. HANDOFF CHECKLIST FOR AI-03

Before marking AI-02 complete, verify:

- [x] Search page design complete
- [x] Search input behavior specified
- [x] Search button behavior specified
- [x] Results page design complete
- [x] Article card component specified
- [x] Article detail page design complete
- [x] Full-text availability states specified
- [x] PDF export UX specified
- [x] Loading states specified for all interactions
- [x] Empty states specified
- [x] Error states specified
- [x] Partial provider failure state specified
- [x] Mobile layout specified (320-767px)
- [x] Tablet layout specified (768-1023px)
- [x] Desktop layout specified (1024px+)
- [x] Accessibility requirements specified
- [x] Design tokens defined (colors, typography, spacing)
- [x] Components library defined
- [x] Interaction rules specified
- [x] Navigation flow specified
- [x] Content language standardized
- [x] AI-03 implementation guide provided

---

23. NEXT STAGE

Next AI stage:
- AI-03 — FRONTEND

Expected handoff:
- AI-03 produces a working Next.js frontend application with all pages, components, and features implemented.
- AI-03 also produces HANDOFF_AI03.md with implementation details and any technical decisions.

This UX/UI specification is the definitive design baseline for V1 Journal AI Research. AI-03 must implement this specification without redefining the product or introducing out-of-scope features.
