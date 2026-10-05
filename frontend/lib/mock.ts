import { SearchResponse, Article } from './types';

export const MOCK_ARTICLES: Article[] = [
  {
    id: 'mock-1',
    title: 'Deep Learning Approaches for Multi-Modal Glaucoma Screening in Clinical Settings',
    abstract: 'Automated screening of retinal fundus photographs and optical coherence tomography using convolutional neural networks demonstrates high sensitivity (96.4%) and specificity (94.2%) in multi-center validation.',
    authors: [
      { name: 'Siti Rahmawati, MD' },
      { name: 'Budi Santoso, PhD' },
      { name: 'Dr. John Davies' }
    ],
    journal: 'Journal of AI in Clinical Ophthalmology',
    publication_date: '2026-03',
    publication_year: 2026,
    doi: '10.1016/j.jclinoph.2026.03.001',
    pmid: '39988112',
    pmcid: 'PMC1122334',
    url: 'https://doi.org/10.1016/j.jclinoph.2026.03.001',
    pdf_url: 'https://europepmc.org/articles/PMC1122334/pdf',
    source: 'europepmc',
    is_open_access: true,
    full_text: 'Open Access Research Article: Clinical evaluation of deep convolutional networks across 12,000 patient retinal scans.',
    relevance_score: 0.94,
    recency_score: 1.0,
    completeness_score: 0.95,
    total_score: 0.959,
  },
  {
    id: 'mock-2',
    title: 'Advancements in Femtosecond Laser-Assisted Cataract Surgery: Outcomes and Safety',
    abstract: 'Evaluation of visual rehabilitation, corneal endothelial loss, and refractive accuracy following laser cataract procedures versus manual phacoemulsification.',
    authors: [
      { name: 'Elena Rostova, MD' },
      { name: 'Michael Chang, PhD' }
    ],
    journal: 'British Journal of Ophthalmology',
    publication_date: '2025-11',
    publication_year: 2025,
    doi: '10.1136/bjo-2025-321100',
    pmid: '39441122',
    url: 'https://doi.org/10.1136/bjo-2025-321100',
    source: 'pubmed',
    is_open_access: false,
    relevance_score: 0.88,
    recency_score: 0.95,
    completeness_score: 0.85,
    total_score: 0.898,
  }
];

export function getMockSearchResponse(query: string): SearchResponse {
  return {
    query,
    results: MOCK_ARTICLES,
    total: MOCK_ARTICLES.length,
    search_id: 'mock-demo-session',
    sources: { pubmed: 1, europepmc: 1 },
    warnings: ['[MOCK MODE ACTIVE] Displaying synthetic validation data.'],
    timestamp: new Date().toISOString(),
  };
}
