# Journal AI Research — Project Status

## Project

**Journal AI Research**  
AI-Powered Scientific Journal Search & Research Platform

---

## Current Status

**Active Implementation Complete (Core Full-Stack V1)**

---

## Development Pipeline

AI-01 → AI-02 → AI-03 → AI-04 → AI-05 → AI-06 → AI-07 → AI-08

| Stage | Description | Status | Verification Evidence |
| :--- | :--- | :--- | :--- |
| **AI-01** | Product + Architecture | COMPLETED | Specifications, contracts, and boundaries established |
| **AI-02** | UX/UI Design System | COMPLETED | Design tokens, color palette, responsive wireframes defined |
| **AI-03** | Frontend Implementation | COMPLETED | Next.js 14 App Router, components, typecheck & build passing |
| **AI-04** | Backend + Scientific Search | COMPLETED | Go HTTP service, PubMed/EuropePMC/Crossref clients, ranking engine, unit tests passing |
| **AI-05** | Database + PostgreSQL | COMPLETED | PostgreSQL schema & migrations (`articles`, `searches`, `search_results`), isolated container on port 5435 |
| **AI-06** | Full-Stack Integration | COMPLETED | End-to-end verified (`POST /api/v1/search`, article detail, PDF research report generation) |
| **AI-07** | Testing & Security | IN PROGRESS | Middleware security headers, rate limiting, parameterized queries, unit test coverage active |
| **AI-08** | Deployment | READY | Docker Compose with isolated network & volume ready for staging |

---

## Core Concept

«KETIK → CARI → 10 JURNAL → BACA → PDF»

---

## Main Sources

- PubMed (NCBI E-utilities)
- Europe PMC (REST API)
- Crossref (REST API)

---

## Technology

- **Frontend**: Next.js 14, React 18, TypeScript, Vanilla CSS
- **Backend**: Go 1.27
- **Database**: PostgreSQL 16 (isolated container on port 5435)
- **Containerization**: Docker Compose
