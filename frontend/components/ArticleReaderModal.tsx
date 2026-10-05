import React from 'react';
import { Article } from '../lib/types';

interface ArticleReaderModalProps {
  article: Article | null;
  onClose: () => void;
}

export const ArticleReaderModal: React.FC<ArticleReaderModalProps> = ({ article, onClose }) => {
  if (!article) return null;

  return (
    <div className="modal-overlay" role="dialog" aria-modal="true" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <div>
            <div style={{ display: 'flex', gap: 8, alignItems: 'center', marginBottom: 6 }}>
              <span className="badge-source">{article.source}</span>
              {article.is_open_access ? (
                <span className="badge-oa">Open Access</span>
              ) : (
                <span style={{ fontSize: '0.75rem', color: 'var(--text-muted)' }}>Publisher Restricted</span>
              )}
            </div>
            <h2 className="modal-title">{article.title}</h2>
          </div>
          <button
            type="button"
            className="modal-close-btn"
            onClick={onClose}
            aria-label="Tutup jendela"
          >
            &times;
          </button>
        </div>

        <div className="modal-body">
          <div className="article-meta-row" style={{ marginBottom: 16 }}>
            <strong>Jurnal:</strong> {article.journal || '-'} ({article.publication_date || article.publication_year || '-'})
          </div>

          <div className="article-meta-row" style={{ marginBottom: 16 }}>
            <strong>Penulis:</strong>{' '}
            {article.authors.length > 0
              ? article.authors.map((a) => a.name).join(', ')
              : 'Informasi penulis tidak tersedia'}
          </div>

          {/* Identifier block */}
          <div style={{ display: 'flex', flexWrap: 'wrap', gap: 14, background: 'var(--bg-subtle)', padding: 12, borderRadius: 'var(--radius-sm)', fontSize: '0.85rem', marginBottom: 18 }}>
            {article.doi && <div><strong>DOI:</strong> <a href={`https://doi.org/${article.doi}`} target="_blank" rel="noopener noreferrer">{article.doi}</a></div>}
            {article.pmid && <div><strong>PMID:</strong> <a href={`https://pubmed.ncbi.nlm.nih.gov/${article.pmid}`} target="_blank" rel="noopener noreferrer">{article.pmid}</a></div>}
            {article.pmcid && <div><strong>PMCID:</strong> {article.pmcid}</div>}
          </div>

          <h3 className="reader-section-title">Abstrak</h3>
          <p className="reader-abstract">
            {article.abstract || 'Abstrak tidak dicantumkan oleh penyedia metadata.'}
          </p>

          <h3 className="reader-section-title">Status Legal Teks Lengkap (Full-Text)</h3>
          {article.is_open_access ? (
            <div style={{ background: 'var(--success-bg)', border: '1px solid rgba(16,185,129,0.3)', padding: 14, borderRadius: 'var(--radius-sm)', fontSize: '0.9rem', color: '#065f46' }}>
              ✓ Artikel ini berlisensi <strong>Open Access</strong>. Anda dapat mengakses teks lengkap melalui dokumen resmi penerbit.
              {article.pdf_url && (
                <div style={{ marginTop: 8 }}>
                  <a href={article.pdf_url} target="_blank" rel="noopener noreferrer" style={{ color: '#047857', fontWeight: 600, textDecoration: 'underline' }}>
                    Unduh Dokumen PDF Resmi
                  </a>
                </div>
              )}
            </div>
          ) : (
            <div style={{ background: 'var(--bg-subtle)', border: '1px solid var(--border-light)', padding: 14, borderRadius: 'var(--radius-sm)', fontSize: '0.9rem', color: 'var(--text-secondary)' }}>
              <strong>FULL TEXT NOT AVAILABLE VIA REPOSITORY</strong>
              <p style={{ marginTop: 6 }}>
                Sesuai kebijakan legalitas Journal AI Research, sistem tidak membobol paywall atau menduplikasi karya berhak cipta. Silakan baca langsung pada situs resmi penerbit.
              </p>
            </div>
          )}
        </div>

        <div className="modal-footer">
          <button type="button" className="btn-read" onClick={onClose}>
            Tutup
          </button>
          {article.url && (
            <a
              href={article.url}
              target="_blank"
              rel="noopener noreferrer"
              className="search-submit-btn"
              style={{ padding: '8px 18px', fontSize: '0.88rem' }}
            >
              <span>Buka Tautan Asli Penerbit</span>
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6" />
                <polyline points="15 3 21 3 21 9" />
                <line x1="10" y1="14" x2="21" y2="3" />
              </svg>
            </a>
          )}
        </div>
      </div>
    </div>
  );
};
