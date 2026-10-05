import React from 'react';
import Link from 'next/link';

export const AppHeader: React.FC = () => {
  return (
    <header className="app-header">
      <div className="container header-content">
        <Link href="/" className="brand">
          <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20" />
            <path d="M6.5 2H20v20H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2z" />
            <circle cx="12" cy="9" r="2" />
            <path d="M12 13v3" />
          </svg>
          <span>Journal AI Research</span>
          <span className="brand-badge">V1 Science</span>
        </Link>

        <nav className="nav-links">
          <span className="nav-item">PubMed • Europe PMC • Crossref</span>
          <Link href="/about" className="nav-item" style={{ marginLeft: 12 }}>
            Pedoman Legal
          </Link>
        </nav>
      </div>
    </header>
  );
};
