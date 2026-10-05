'use client';

import React, { useState } from 'react';
import { SearchBox } from '../components/SearchBox';
import { SearchExamples } from '../components/SearchExamples';
import { ResultsHeader } from '../components/ResultsHeader';
import { ArticleCard } from '../components/ArticleCard';
import { ArticleReaderModal } from '../components/ArticleReaderModal';
import { PdfExportBar } from '../components/PdfExportBar';
import { searchArticles, exportPdfReport } from '../lib/api';
import { Article, SearchResponse } from '../lib/types';

export default function HomePage() {
  const [query, setQuery] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [searchResponse, setSearchResponse] = useState<SearchResponse | null>(null);
  const [selectedArticleIds, setSelectedArticleIds] = useState<string[]>([]);
  const [activeArticle, setActiveArticle] = useState<Article | null>(null);
  const [errorMsg, setErrorMsg] = useState<string | null>(null);
  const [isExporting, setIsExporting] = useState(false);

  const handleSearch = async (searchQuery: string) => {
    setQuery(searchQuery);
    setIsLoading(true);
    setErrorMsg(null);
    setSearchResponse(null);

    try {
      const data = await searchArticles(searchQuery, 10);
      setSearchResponse(data);
      if (data.results.length === 0) {
        setErrorMsg('Tidak ditemukan jurnal yang sesuai. Coba kata kunci lain.');
      } else {
        // Select all by default for convenience
        setSelectedArticleIds(data.results.map((a) => a.id));
      }
    } catch (err: any) {
      setErrorMsg(err.message || 'Pencarian jurnal sedang mengalami gangguan. Silakan coba lagi.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleToggleSelect = (id: string) => {
    setSelectedArticleIds((prev) =>
      prev.includes(id) ? prev.filter((item) => item !== id) : [...prev, id]
    );
  };

  const handleToggleSelectAll = () => {
    if (!searchResponse) return;
    if (selectedArticleIds.length === searchResponse.results.length) {
      setSelectedArticleIds([]);
    } else {
      setSelectedArticleIds(searchResponse.results.map((a) => a.id));
    }
  };

  const handleExportPdf = async () => {
    if (!searchResponse || selectedArticleIds.length === 0) return;
    setIsExporting(true);
    try {
      const blob = await exportPdfReport(query, selectedArticleIds);
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', `journal-ai-research-${Date.now()}.pdf`);
      document.body.appendChild(link);
      link.click();
      link.remove();
      window.URL.revokeObjectURL(url);
    } catch (err: any) {
      alert(err.message || 'Gagal mengunduh laporan PDF.');
    } finally {
      setIsExporting(false);
    }
  };

  return (
    <div className="container">
      {/* Hero Section */}
      <section className="hero-section">
        <div className="hero-tagline">AI-Powered Scientific Literature</div>
        <h1 className="hero-title">
          Ketik Riset, Dapatkan 10 Jurnal Terkini
        </h1>
        <p className="hero-subtitle">
          Pencarian lintas database ilmiah PubMed, Europe PMC, dan Crossref secara paralel. 
          Transparan, bebas bias komersial, dan patuh legalitas Open Access.
        </p>

        <SearchBox
          onSearch={handleSearch}
          isLoading={isLoading}
          initialQuery={query}
        />

        <SearchExamples
          onSelect={(q) => handleSearch(q)}
          disabled={isLoading}
        />
      </section>

      {/* Warnings Banner if any */}
      {searchResponse?.warnings && searchResponse.warnings.length > 0 && (
        <div style={{ maxWidth: 760, margin: '0 auto 16px', background: '#fffbeb', border: '1px solid #fef3c7', borderRadius: 8, padding: '10px 16px', fontSize: '0.85rem', color: '#b45309' }}>
          <strong>Catatan Provider:</strong> {searchResponse.warnings.join('; ')}
        </div>
      )}

      {/* Loading state */}
      {isLoading && (
        <div className="state-box">
          <div className="spinner" />
          <h3 style={{ fontSize: '1.1rem', marginBottom: 6 }}>Menghubungi Sumber Ilmiah...</h3>
          <p style={{ color: 'var(--text-secondary)', fontSize: '0.9rem' }}>
            Mengagregasi PubMed, Europe PMC, dan Crossref • Normalisasi • Deduplikasi • Ranking Relevansi
          </p>
        </div>
      )}

      {/* Error / Empty state */}
      {errorMsg && !isLoading && (
        <div className="state-box">
          <svg style={{ margin: '0 auto 12px', color: 'var(--text-muted)' }} width="44" height="44" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
            <circle cx="12" cy="12" r="10" />
            <line x1="8" y1="12" x2="16" y2="12" />
          </svg>
          <h3 style={{ fontSize: '1.1rem', marginBottom: 6, color: 'var(--text-main)' }}>Informasi Pencarian</h3>
          <p style={{ color: 'var(--text-secondary)', fontSize: '0.92rem' }}>{errorMsg}</p>
        </div>
      )}

      {/* Results Section */}
      {searchResponse && searchResponse.results.length > 0 && !isLoading && (
        <section className="results-section">
          <ResultsHeader
            query={searchResponse.query}
            total={searchResponse.total}
            sources={searchResponse.sources}
          />

          <div className="articles-grid">
            {searchResponse.results.map((article, index) => (
              <ArticleCard
                key={article.id}
                article={article}
                rank={index + 1}
                isSelected={selectedArticleIds.includes(article.id)}
                onToggleSelect={handleToggleSelect}
                onOpenReader={(a) => setActiveArticle(a)}
              />
            ))}
          </div>

          <PdfExportBar
            selectedCount={selectedArticleIds.length}
            totalCount={searchResponse.results.length}
            isAllSelected={
              selectedArticleIds.length === searchResponse.results.length &&
              searchResponse.results.length > 0
            }
            onToggleSelectAll={handleToggleSelectAll}
            onExport={handleExportPdf}
            isExporting={isExporting}
          />
        </section>
      )}

      {/* Article Reader Modal */}
      <ArticleReaderModal
        article={activeArticle}
        onClose={() => setActiveArticle(null)}
      />

      {/* Bottom Disclaimer */}
      <footer style={{ borderTop: '1px solid var(--border-light)', padding: '32px 0 48px', marginTop: 40, textAlign: 'center', fontSize: '0.8rem', color: 'var(--text-muted)' }}>
        <p>
          Journal AI Research adalah platform penelusuran pustaka ilmiah berbasis AI. 
          Metadata dan abstrak diambil dari PubMed, Europe PMC, dan Crossref.
        </p>
        <p style={{ marginTop: 4 }}>
          Platform ini tidak menyediakan diagnosis medis, rekomendasi pengobatan, maupun keputusan klinis individual.
        </p>
      </footer>
    </div>
  );
}
