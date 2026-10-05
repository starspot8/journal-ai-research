import React, { useState } from 'react';

interface SearchBoxProps {
  onSearch: (query: string) => void;
  isLoading: boolean;
  initialQuery?: string;
}

export const SearchBox: React.FC<SearchBoxProps> = ({ onSearch, isLoading, initialQuery = '' }) => {
  const [query, setQuery] = useState(initialQuery);
  const [errorMsg, setErrorMsg] = useState<string | null>(null);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const cleanQuery = query.trim();
    if (!cleanQuery) {
      setErrorMsg('Masukkan topik jurnal terlebih dahulu.');
      return;
    }
    setErrorMsg(null);
    onSearch(cleanQuery);
  };

  return (
    <div className="search-wrapper">
      <form onSubmit={handleSubmit} role="search" aria-label="Pencarian Jurnal Ilmiah">
        <div className="search-box-card">
          <svg style={{ marginLeft: 14, color: 'var(--text-muted)' }} width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
            <circle cx="11" cy="11" r="8" />
            <line x1="21" y1="21" x2="16.65" y2="16.65" />
          </svg>
          <input
            id="search-input"
            type="text"
            className="search-input-field"
            placeholder="Ketik topik ilmiah atau pertanyaan riset (Bahasa Indonesia / English)..."
            value={query}
            disabled={isLoading}
            onChange={(e) => {
              setQuery(e.target.value);
              if (errorMsg) setErrorMsg(null);
            }}
          />
          <button
            id="search-btn"
            type="submit"
            className="search-submit-btn"
            disabled={isLoading}
          >
            {isLoading ? (
              <>
                <div style={{ width: 16, height: 16, border: '2px solid rgba(255,255,255,0.4)', borderTopColor: '#fff', borderRadius: '50%', animation: 'spin 0.6s linear infinite' }} />
                <span>Mencari...</span>
              </>
            ) : (
              <>
                <span>Cari</span>
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
                  <line x1="5" y1="12" x2="19" y2="12" />
                  <polyline points="12 5 19 12 12 19" />
                </svg>
              </>
            )}
          </button>
        </div>

        {errorMsg && (
          <div className="error-banner" role="alert" style={{ textAlign: 'center' }}>
            {errorMsg}
          </div>
        )}
      </form>
    </div>
  );
};
