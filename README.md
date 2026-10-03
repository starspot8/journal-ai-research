Journal AI Research

AI-Powered Scientific Journal Search & Research Platform

Overview

Journal AI Research adalah platform pencarian jurnal ilmiah yang membantu pengguna menemukan, membaca, dan mengakses artikel ilmiah secara lebih sederhana.

Konsep utama:

«KETIK → CARI → 10 JURNAL → BACA → PDF»

Pengguna cukup memasukkan pertanyaan atau kata kunci dalam bahasa Indonesia atau Inggris, kemudian sistem mencari artikel ilmiah yang relevan dari sumber ilmiah seperti PubMed, Europe PMC, dan Crossref.

Core User Flow

USER
  ↓
ONE SEARCH BOX
  ↓
QUERY PROCESSING / AI UNDERSTANDING
  ↓
PubMed + Europe PMC + Crossref
  ↓
NORMALIZE
  ↓
DEDUPLICATE
  ↓
RELEVANCE + RECENCY RANKING
  ↓
10 LATEST RELEVANT JOURNALS
  ↓
ARTICLE DETAIL / READER
  ↓
FULL TEXT WHEN LEGALLY AVAILABLE
  ↓
ORIGINAL ARTICLE
  ↓
PDF RESEARCH REPORT

Main Features

- Natural-language scientific journal search
- Indonesian and English queries
- Search across multiple scientific sources
- PubMed integration
- Europe PMC integration
- Crossref integration
- Article normalization
- Duplicate removal
- Relevance and recency ranking
- Maximum 10 relevant results
- Article detail page
- Abstract display
- Legal full-text access when available
- Original article access
- Research report PDF export
- Responsive web interface
- Accessibility support
- Security and rate limiting
- Production deployment with Docker

Search Result Policy

The system displays:

10 Latest Relevant Journals

The application does not use labels such as:

- Best Journals
- Top Journals
- Number One Journal

Ranking is based on documented technical criteria rather than an editorial judgment of journal quality.

Initial ranking model:

Relevance          60%
Recency            30%
Metadata Completeness 10%

Scientific Sources

The initial external sources are:

PubMed
Europe PMC
Crossref

Scientific databases are treated as the source of truth for article metadata.

AI functionality is an assisting layer and does not replace the underlying scientific data sources.

Legal Full-Text Policy

The application only displays full text when legally accessible through an appropriate source.

If full text is unavailable, the system may display:

- Article metadata
- Abstract when permitted
- Original article link

The system must not:

- Bypass paywalls
- Circumvent access controls
- Scrape restricted content
- Create unauthorized copies of copyrighted articles

PDF Policy

The PDF feature is primarily a research report generator.

A generated report may contain:

- User search query
- Search date
- Selected article information
- Authors
- Journal
- Publication date
- DOI
- PMID / PMCID when available
- Abstract when permitted
- Source information
- Original article link

The system does not create unauthorized copies of restricted journal PDFs.

Technology Stack

Frontend

Next.js
TypeScript
React

Backend

Go

Database

PostgreSQL

Deployment

Docker
Cloud Infrastructure
HTTPS
Reverse Proxy

Project Architecture

journal-ai-research/
│
├── AI-01_PRODUCT_ARCHITECTURE/
├── AI-02_UX_UI/
├── AI-03_FRONTEND/
├── AI-04_BACKEND_SEARCH/
├── AI-05_DATABASE_POSTGRESQL/
├── AI-06_INTEGRATION/
├── AI-07_TESTING_SECURITY/
├── AI-08_DEPLOYMENT/
│
├── docs/
├── README.md
├── .gitignore
└── .env.example

AI Development Workflow

The project is developed through sequential engineering stages.

AI-01
Product + Architecture
        ↓
AI-02
UX/UI
        ↓
AI-03
Frontend
        ↓
AI-04
Backend + Scientific Search
        ↓
AI-05
Database + PostgreSQL
        ↓
AI-06
Full-Stack Integration
        ↓
AI-07
Testing + Security
        ↓
AI-08
Deployment

Each stage must read the previous stage's handoff before making implementation decisions.

Backend API

Initial API structure:

POST /api/v1/search

GET /api/v1/articles/:id

GET /api/v1/articles/:id/fulltext

POST /api/v1/export/pdf

GET /api/v1/health

Security Principles

The project follows these basic security requirements:

- HTTPS in production
- Input validation
- Rate limiting
- Request timeouts
- Secure HTTP headers
- Parameterized SQL queries
- Secret management
- No secrets in frontend code
- No secrets committed to Git
- Safe external links
- SSRF protection
- XSS protection
- Restricted CORS
- Structured logging
- Dependency security checks

Performance Goal

The application should feel responsive under normal operating conditions.

The initial target is for a normal search to generally feel fast, approximately:

< 3 seconds

This is a performance target, not a guarantee, because external scientific APIs and network conditions can affect response time.

Initial Scope

Included

- Scientific article search
- Multi-source search
- Result normalization
- Deduplication
- Ranking
- Article reader
- Legal full-text access
- Original article links
- Research report PDF
- PostgreSQL persistence
- Testing
- Security
- Production deployment

Not Included in V1

- User accounts
- Social profiles
- Social networking
- Native mobile application
- Citation manager
- Systematic review automation
- Meta-analysis
- Diagnosis
- Treatment recommendations
- Medical decision support
- Paid subscription system

Development Principle

The project follows these principles:

1. Scientific data first.
2. AI assists; it does not replace authoritative sources.
3. Legal access only.
4. Security by design.
5. Clear separation between frontend, backend, database, and external providers.
6. Small, testable development stages.
7. No unnecessary feature expansion.
8. Every major stage documents its implementation in a handoff.

Project Status

Current stage:

Documentation / Architecture Preparation

Development stages:

[ ] AI-01 Product + Architecture
[ ] AI-02 UX/UI
[ ] AI-03 Frontend
[ ] AI-04 Backend + Search
[ ] AI-05 Database + PostgreSQL
[ ] AI-06 Integration
[ ] AI-07 Testing + Security
[ ] AI-08 Deployment

Repository

Project name:

Journal AI Research

Repository name:

journal-ai-research

Project description:

AI-Powered Scientific Journal Search & Research Platform

License

License selection will be finalized before public production release.

Disclaimer

Journal AI Research is a scientific literature search and research assistance platform.

Search results, article metadata, abstracts, and other scientific information should be verified against the original scientific source.

The platform is not intended to provide medical diagnosis, treatment recommendations, or individualized medical advice.
