-- 001_initial_schema.sql
-- Schema for Journal AI Research (PostgreSQL)

CREATE TABLE IF NOT EXISTS articles (
    id VARCHAR(64) PRIMARY KEY,
    title TEXT NOT NULL,
    abstract TEXT,
    authors JSONB NOT NULL DEFAULT '[]'::jsonb,
    journal VARCHAR(255),
    publication_date VARCHAR(64),
    publication_year INTEGER,
    doi VARCHAR(255),
    pmid VARCHAR(64),
    pmcid VARCHAR(64),
    url TEXT,
    pdf_url TEXT,
    source VARCHAR(64) NOT NULL,
    is_open_access BOOLEAN NOT NULL DEFAULT false,
    full_text TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes for fast identifier lookups and sorting
CREATE INDEX IF NOT EXISTS articles_doi_idx ON articles (doi) WHERE doi IS NOT NULL;
CREATE INDEX IF NOT EXISTS articles_pmid_idx ON articles (pmid) WHERE pmid IS NOT NULL;
CREATE INDEX IF NOT EXISTS articles_pmcid_idx ON articles (pmcid) WHERE pmcid IS NOT NULL;
CREATE INDEX IF NOT EXISTS articles_publication_year_idx ON articles (publication_year DESC NULLS LAST);
CREATE INDEX IF NOT EXISTS articles_created_at_idx ON articles (created_at DESC);

-- Searches history table
CREATE TABLE IF NOT EXISTS searches (
    id VARCHAR(64) PRIMARY KEY,
    query TEXT NOT NULL,
    total_results INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS searches_created_at_idx ON searches (created_at DESC);

-- Association table between searches and articles ranked
CREATE TABLE IF NOT EXISTS search_results (
    id VARCHAR(64) PRIMARY KEY,
    search_id VARCHAR(64) NOT NULL REFERENCES searches(id) ON DELETE CASCADE,
    article_id VARCHAR(64) NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
    rank_position INTEGER NOT NULL,
    score NUMERIC(5, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT search_results_unique_search_article UNIQUE (search_id, article_id)
);

CREATE INDEX IF NOT EXISTS search_results_search_rank_idx ON search_results (search_id, rank_position ASC);
CREATE INDEX IF NOT EXISTS search_results_article_idx ON search_results (article_id);
