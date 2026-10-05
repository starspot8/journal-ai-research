'use client';

import React, { useEffect, useState } from 'react';
import Link from 'next/link';
import { useParams } from 'next/navigation';
import { getArticle } from '../../../lib/api';
import { Article } from '../../../lib/types';

export default function ArticleDetailPage() {
  const params = useParams();
  const id = params?.id as string;

  const [article, setArticle] = useState<Article | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    setIsLoading(true);
    getArticle(id)
      .then((data) => {
        setArticle(data);
        setError(null);
      })
      .catch((err) => {
        setError(err.message || 'Gagal memuat detail artikel.');
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, [id]);

  if (isLoading) {
    return (
      <div className="container" style={{ padding: '60px 0' }}>
        <div className="state-box">
          <div className="spinner" />
          <p>Memuat data artikel ilmiah...</p>
        </div>
      </div>
    );
  }

  if (error || !article) {
    return (
      <div className="container" style={{ padding: '60px 0' }}>
        <div className="state-box">
          <h3>Artikel Tidak Ditemukan</h3>
          <p style={{ margin: '12px 0 20px', color: 'var(--text-secondary)' }}>
            {error || 'Data artikel tidak tersedia.'}
          </p>
          <Link href="/" className="btn-read">
            &larr; Kembali ke Pencarian
          </Link>
        </div>
      </div>
    );
  }

  return (
    <div className="container" style={{ padding: '40px 0 80px', maxWidth: 860 }}>
      <div style={{ marginBottom: 20 }}>
        <Link href="/" style={{ fontSize: '0.9rem', color: 'var(--primary)', display: 'inline-flex', alignItems: 'center', gap: 6 }}>
          <span>&larr;</span> Kembali ke Hasil Pencarian
        </Link>
      </div>

      <article className="article-card" style={{ padding: 32 }}>
        <div style={{ display: 'flex', gap: 8, alignItems: 'center', marginBottom: 12 }}>
          <span className="badge-source">{article.source}</span>
          {article.is_open_access ? (
            <span className="badge-oa">Open Access</span>
          ) : (
            <span style={{ fontSize: '0.75rem', color: 'var(--text-muted)' }}>Publisher Restricted</span>
          )}
        </div>

        <h1 style={{ fontSize: '1.75rem', fontWeight: 800, lineHeight: 1.35, marginBottom: 16 }}>
          {article.title}
        </h1>

        <div className="article-meta-row" style={{ fontSize: '0.95rem', marginBottom: 20 }}>
          <span className="article-journal">{article.journal || 'Journal'}</span>
          <span>({article.publication_date || article.publication_year || '-'})</span>
        </div>

        <div style={{ marginBottom: 24, fontSize: '0.92rem', color: 'var(--text-secondary)' }}>
          <strong>Penulis: </strong>
          {article.authors.length > 0
            ? article.authors.map((a) => a.name).join(', ')
            : 'Informasi penulis tidak tersedia'}
        </div>

        {/* Identifiers */}
        <div style={{ display: 'flex', flexWrap: 'wrap', gap: 16, background: 'var(--bg-subtle)', padding: 16, borderRadius: 'var(--radius-sm)', marginBottom: 24, fontSize: '0.88rem' }}>
          {article.doi && <div><strong>DOI:</strong> <a href={`https://doi.org/${article.doi}`} target="_blank" rel="noopener noreferrer">{article.doi}</a></div>}
          {article.pmid && <div><strong>PMID:</strong> <a href={`https://pubmed.ncbi.nlm.nih.gov/${article.pmid}`} target="_blank" rel="noopener noreferrer">{article.pmid}</a></div>}
          {article.pmcid && <div><strong>PMCID:</strong> {article.pmcid}</div>}
        </div>

        <h2 style={{ fontSize: '1.1rem', fontWeight: 700, marginBottom: 10 }}>Abstrak</h2>
        <p style={{ fontSize: '1rem', lineHeight: 1.7, color: '#334155', background: 'var(--bg-subtle)', padding: 20, borderRadius: 'var(--radius-sm)', borderLeft: '4px solid var(--primary)', marginBottom: 28 }}>
          {article.abstract || 'Abstrak tidak tersedia.'}
        </p>

        <h2 style={{ fontSize: '1.1rem', fontWeight: 700, marginBottom: 10 }}>Akses Dokumen Resmi</h2>
        {article.is_open_access ? (
          <div style={{ background: 'var(--success-bg)', border: '1px solid rgba(16,185,129,0.3)', padding: 18, borderRadius: 'var(--radius-sm)', color: '#065f46', marginBottom: 24 }}>
            <p>✓ Artikel ini berlisensi <strong>Open Access</strong>. Anda dapat mengunduh naskah lengkap secara langsung.</p>
            {article.pdf_url && (
              <div style={{ marginTop: 10 }}>
                <a href={article.pdf_url} target="_blank" rel="noopener noreferrer" className="btn-read" style={{ background: '#059669', color: 'white', borderColor: 'transparent', display: 'inline-block' }}>
                  Unduh Dokumen PDF Resmi
                </a>
              </div>
            )}
          </div>
        ) : (
          <div style={{ background: 'var(--bg-subtle)', border: '1px solid var(--border-light)', padding: 18, borderRadius: 'var(--radius-sm)', color: 'var(--text-secondary)', marginBottom: 24 }}>
            <strong>FULL TEXT NOT AVAILABLE VIA REPOSITORY</strong>
            <p style={{ marginTop: 6, fontSize: '0.9rem' }}>
              Sesuai kebijakan legalitas Journal AI Research, sistem tidak membobol paywall atau menduplikasi karya berhak cipta. Silakan buka tautan penerbit resmi di bawah ini.
            </p>
          </div>
        )}

        {article.url && (
          <div>
            <a
              href={article.url}
              target="_blank"
              rel="noopener noreferrer"
              className="search-submit-btn"
              style={{ display: 'inline-flex' }}
            >
              <span>Buka Tautan Asli Penerbit</span>
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6" />
                <polyline points="15 3 21 3 21 9" />
                <line x1="10" y1="14" x2="21" y2="3" />
              </svg>
            </a>
          </div>
        )}
      </article>
    </div>
  );
}
