import React from 'react';

interface SearchExamplesProps {
  onSelect: (query: string) => void;
  disabled?: boolean;
}

const EXAMPLES = [
  'glaucoma screening artificial intelligence',
  'terapi katarak modern dan phacoemulsification',
  'diabetic retinopathy deep learning',
  'mrna vaccine cancer immunotherapy',
  'stem cell corneal regeneration'
];

export const SearchExamples: React.FC<SearchExamplesProps> = ({ onSelect, disabled }) => {
  return (
    <div className="examples-container">
      <span className="examples-label">Contoh topik:</span>
      {EXAMPLES.map((ex) => (
        <button
          key={ex}
          type="button"
          className="example-chip"
          disabled={disabled}
          onClick={() => onSelect(ex)}
        >
          {ex}
        </button>
      ))}
    </div>
  );
};
