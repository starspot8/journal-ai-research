import { SearchResponse, Article, FullTextResponse } from './types';
import { getMockSearchResponse, MOCK_ARTICLES } from './mock';

const API_BASE = process.env.NEXT_PUBLIC_API_URL || '';

export async function searchArticles(query: string, limit: number = 10): Promise<SearchResponse> {
  const isMock = process.env.NEXT_PUBLIC_ENABLE_MOCK === 'true';
  if (isMock) {
    return getMockSearchResponse(query);
  }

  try {
    const res = await fetch(`${API_BASE}/api/v1/search`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ query, limit }),
    });

    if (!res.ok) {
      const errData = await res.json().catch(() => ({}));
      throw new Error(errData.error || `Search failed with status: ${res.status}`);
    }

    return await res.json();
  } catch (err: any) {
    // If backend isn't reached and mock isn't explicitly disabled, inform user clearly
    throw new Error(err.message || 'Gagal menghubungi server pencarian jurnal.');
  }
}

export async function getArticle(id: string): Promise<Article> {
  if (process.env.NEXT_PUBLIC_ENABLE_MOCK === 'true') {
    const found = MOCK_ARTICLES.find(a => a.id === id);
    if (found) return found;
    throw new Error('Artikel tidak ditemukan dalam mock database.');
  }

  const res = await fetch(`${API_BASE}/api/v1/articles/${encodeURIComponent(id)}`);
  if (!res.ok) {
    throw new Error(`Failed to load article: ${res.status}`);
  }
  return await res.json();
}

export async function getArticleFullText(id: string): Promise<FullTextResponse> {
  const res = await fetch(`${API_BASE}/api/v1/articles/${encodeURIComponent(id)}/fulltext`);
  if (!res.ok) {
    const errData = await res.json().catch(() => ({}));
    throw new Error(errData.error || 'Full text tidak tersedia secara legal.');
  }
  return await res.json();
}

export async function exportPdfReport(query: string, articleIds: string[]): Promise<Blob> {
  const res = await fetch(`${API_BASE}/api/v1/export/pdf`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query, article_ids: articleIds }),
  });

  if (!res.ok) {
    const errData = await res.json().catch(() => ({}));
    throw new Error(errData.error || 'Gagal membuat laporan PDF.');
  }

  return await res.blob();
}
