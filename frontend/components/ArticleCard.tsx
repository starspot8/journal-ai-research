import React from 'react';
import { Article } from '../lib/types';

interface ArticleCardProps {
  article: Article;
  rank: number;
  isSelected: boolean;
  onToggleSelect: (id: string) => void;
  onOpenReader: (article: Article) => void;
}

export const ArticleCard: React.FC<ArticleCardProps> = ({
  article,
  rank,
  isSelected,
  onToggleSelect,
  onOpenReader,
}) => {
  const authorNames = article.authors.map(a => a.name).join(', ') || 'Penulis tidak terindeks';

  return (
    <article className="article-card" aria-labelledby={`title-${article.id}`}>
      <div className="article-card-header">
        <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
          <input
            type="checkbox"
            id={`select-${article.id}`}
            checked={isSelected}
            onChange={() => onToggleSelect(article.id)}
            aria-label={`Pilih artikel ${article.title} untuk laporan PDF`}
            style={{ width: 18, height: 18, cursor: 'pointer', accentColor: 'var(--primary)' }}
          />
          <div className="article-rank-badge">#{rank}</div>
          <h3
            id={`title-${article.id}`}
            className="article-title-link"
            onClick={() => onOpenReader(article)}
          >
            {article.title}
          </h3>
        </div>
      </div>

      <div className="article-meta-row">
        <span className="article-journal">{article.journal || 'Journal'}</span>
        {article.publication_year ? <span>({article.publication_year})</span> : null}
        <span style={{ color: 'var(--text-muted)' }}>•</span>
        <span style={{ maxWidth: 360, whiteSpace: 'nowrap', overflow: 'hidden', textOverflow: 'ellipsis' }}>
          {authorNames}
        </span>
        <span className="badge-source">{article.source}</span>
        {article.is_open_access && <span className="badge-oa">Open Access</span>}
      </div>

      <p className="article-abstract-snippet">
        {article.abstract || 'Abstrak belum tersedia untuk artikel ini.'}
      </p>

      <div className="article-actions-row">
        <div style={{ display: 'flex', gap: 10, alignItems: 'center' }}>
          <button
            type="button"
            className="btn-read"
            onClick={() => onOpenReader(article)}
          >
            Baca Abstrak & Metadata
          </button>
          {article.is_open_access && article.pdf_url && (
            <a
              href={article.pdf_url}
              target="_blank"
              rel="noopener noreferrer"
              className="btn-read"
              style={{ background: 'var(--success-bg)', color: 'var(--success)', borderColor: 'rgba(16,185,129,0.3)' }}
            >
              PDF Open Access
            </a>
          )}
        </div>

        {article.url && (
          <a
            href={article.url}
            target="_blank"
            rel="noopener noreferrer"
            className="btn-external"
          >
            <span>Tautan Asli</span>
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6" />
              <polyline points="15 3 21 3 21 9" />
              <line x1="10" y1="14" x2="21" y2="3" />
            </svg>
          </a>
        )}
      </div>
    </article>
  );
};
