# AI-05 — HANDOFF
## DATABASE + POSTGRESQL

---

### 1. STATUS
**COMPLETED**

---

### 2. DATABASE
- **Database Engine**: PostgreSQL
- **Version**: 16-alpine (isolated container `journal-ai-postgres`)
- **Port Mapping**: `127.0.0.1:5435 -> 5432` (strictly avoiding port conflicts with other local containers)
- **Driver**: `github.com/lib/pq` (v1.12.3)

---

### 3. SCHEMA & TABLES
- `articles`: Stores normalized scientific metadata from PubMed, Europe PMC, and Crossref.
- `searches`: Audit log of search queries and result counts.
- `search_results`: Many-to-many association mapping searches to ranked articles with rank positions (1-10) and total scores.

---

### 4. ARTICLES TABLE STRUCTURE
- `id`: VARCHAR(64) PRIMARY KEY
- `title`: TEXT NOT NULL
- `abstract`: TEXT
- `authors`: JSONB NOT NULL DEFAULT '[]'::jsonb
- `journal`: VARCHAR(255)
- `publication_date`: VARCHAR(64)
- `publication_year`: INTEGER
- `doi`: VARCHAR(255)
- `pmid`: VARCHAR(64)
- `pmcid`: VARCHAR(64)
- `url`: TEXT
- `pdf_url`: TEXT
- `source`: VARCHAR(64) NOT NULL
- `is_open_access`: BOOLEAN NOT NULL DEFAULT false
- `full_text`: TEXT
- `created_at`: TIMESTAMP WITH TIME ZONE DEFAULT NOW()
- `updated_at`: TIMESTAMP WITH TIME ZONE DEFAULT NOW()

Indexes:
- `articles_doi_idx`: ON articles (doi) WHERE doi IS NOT NULL
- `articles_pmid_idx`: ON articles (pmid) WHERE pmid IS NOT NULL
- `articles_pmcid_idx`: ON articles (pmcid) WHERE pmcid IS NOT NULL
- `articles_publication_year_idx`: ON articles (publication_year DESC NULLS LAST)
- `articles_created_at_idx`: ON articles (created_at DESC)

---

### 5. REPOSITORIES IMPLEMENTED
- `PostgresRepository` in `backend/internal/repository/postgres_repository.go`:
  - `SaveSearch(ctx, query, totalResults)`
  - `SaveArticles(ctx, articles)` (using upsert `ON CONFLICT (id) DO UPDATE`)
  - `SaveSearchResults(ctx, searchID, articles)` (transactional rank position persistence)
  - `GetArticleByID(ctx, id)`
  - `GetArticlesByIDs(ctx, ids)`
  - `GetRecentSearches(ctx, limit)`

---

### 6. CONNECTION POOL CONFIGURATION
- `MaxOpenConns`: 25
- `MaxIdleConns`: 5
- `ConnMaxLifetime`: 15 minutes

---

### 7. VERIFICATION
- Migrations executed successfully in container `journal-ai-postgres`.
- Verification query: `SELECT count(*) FROM searches; SELECT count(*) FROM articles; SELECT count(*) FROM search_results;` returned populated counts following real searches.
