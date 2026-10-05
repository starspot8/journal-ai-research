import React from 'react';

interface ResultsHeaderProps {
  query: string;
  total: number;
  sources: Record<string, number>;
}

export const ResultsHeader: React.FC<ResultsHeaderProps> = ({ query, total, sources }) => {
  const sourceDetails = Object.entries(sources)
    .filter(([_, count]) => count > 0)
    .map(([name, count]) => `${name}: ${count}`)
    .join(' • ');

  return (
    <div className="results-header-bar">
      <div>
        <h2 className="results-title">10 Latest Relevant Journals</h2>
        <div className="results-subtitle">
          Pencarian untuk: <strong style={{ color: 'var(--primary)' }}>&ldquo;{query}&rdquo;</strong>
          {sourceDetails && <span> ({sourceDetails})</span>}
        </div>
      </div>
      <div className="results-count-badge">
        {total} Artikel Ditemukan
      </div>
    </div>
  );
};
