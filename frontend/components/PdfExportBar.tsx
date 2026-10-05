import React from 'react';

interface PdfExportBarProps {
  selectedCount: number;
  totalCount: number;
  isAllSelected: boolean;
  onToggleSelectAll: () => void;
  onExport: () => void;
  isExporting: boolean;
}

export const PdfExportBar: React.FC<PdfExportBarProps> = ({
  selectedCount,
  totalCount,
  isAllSelected,
  onToggleSelectAll,
  onExport,
  isExporting,
}) => {
  return (
    <div className="pdf-export-bar">
      <div className="pdf-export-info">
        <label style={{ display: 'flex', alignItems: 'center', gap: 8, cursor: 'pointer' }}>
          <input
            type="checkbox"
            checked={isAllSelected}
            onChange={onToggleSelectAll}
            style={{ width: 16, height: 16, accentColor: 'var(--primary)' }}
          />
          <span>{selectedCount} dari {totalCount} artikel dipilih untuk laporan</span>
        </label>
      </div>

      <button
        type="button"
        className="btn-export-pdf"
        disabled={selectedCount === 0 || isExporting}
        onClick={onExport}
      >
        {isExporting ? (
          <>
            <div style={{ width: 14, height: 14, border: '2px solid rgba(255,255,255,0.4)', borderTopColor: '#fff', borderRadius: '50%', animation: 'spin 0.6s linear infinite' }} />
            <span>Membuat PDF...</span>
          </>
        ) : (
          <>
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z" />
              <polyline points="14 2 14 8 20 8" />
              <line x1="16" y1="13" x2="8" y2="13" />
              <line x1="16" y1="17" x2="8" y2="17" />
              <polyline points="10 9 9 9 8 9" />
            </svg>
            <span>Unduh Laporan Riset (PDF)</span>
          </>
        )}
      </button>
    </div>
  );
};
