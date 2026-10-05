import type { Metadata } from 'next';
import './globals.css';
import { AppHeader } from '../components/AppHeader';

export const metadata: Metadata = {
  title: 'Journal AI Research — AI-Powered Scientific Literature Search',
  description: 'Cari dan telaah 10 jurnal ilmiah terkini dari PubMed, Europe PMC, dan Crossref secara cepat, transparan, dan berlandaskan kepatuhan legal.',
  keywords: ['jurnal ilmiah', 'scientific papers', 'pubmed', 'europe pmc', 'crossref', 'research AI'],
  authors: [{ name: 'Journal AI Research Team' }],
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="id">
      <body>
        <AppHeader />
        <main>{children}</main>
      </body>
    </html>
  );
}
