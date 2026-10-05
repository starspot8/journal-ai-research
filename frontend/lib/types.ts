export interface Author {
  name: string;
  affiliation?: string;
}

export interface Article {
  id: string;
  title: string;
  abstract: string;
  authors: Author[];
  journal: string;
  publication_date: string;
  publication_year: number;
  doi?: string;
  pmid?: string;
  pmcid?: string;
  url: string;
  pdf_url?: string;
  source: string;
  is_open_access: boolean;
  full_text?: string;
  relevance_score?: number;
  recency_score?: number;
  completeness_score?: number;
  total_score?: number;
  created_at?: string;
}

export interface SearchResponse {
  query: string;
  results: Article[];
  total: number;
  search_id: string;
  sources: Record<string, number>;
  warnings?: string[];
  timestamp: string;
}

export interface FullTextResponse {
  id: string;
  title: string;
  is_open_access: boolean;
  full_text?: string;
  pdf_url?: string;
  original_url: string;
}
