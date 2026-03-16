import type { KeyValuePair } from '../../shared/types';

interface Props {
  pairs: KeyValuePair[];
  onChange: (pairs: KeyValuePair[]) => void;
}

export default function KeyValueEditor({ pairs, onChange }: Props) {
  const update = (index: number, field: keyof KeyValuePair, value: string | boolean) => {
    const updated = pairs.map((p, i) => (i === index ? { ...p, [field]: value } : p));
    onChange(updated);
  };

  const addRow = () => {
    onChange([...pairs, { key: '', value: '', enabled: true }]);
  };

  const removeRow = (index: number) => {
    if (pairs.length <= 1) return;
    onChange(pairs.filter((_, i) => i !== index));
  };

  return (
    <div className="flex flex-col gap-1.5">
      {pairs.map((pair, i) => (
        <div key={i} className="flex items-center gap-2">
          <input
            type="checkbox"
            checked={pair.enabled}
            onChange={(e) => update(i, 'enabled', e.target.checked)}
            className="accent-accent"
          />
          <input
            type="text"
            placeholder="Key"
            value={pair.key}
            onChange={(e) => update(i, 'key', e.target.value)}
            className="flex-1 bg-surface0 text-text border border-surface1 rounded px-2.5 py-1.5 text-[13px] outline-none focus:border-accent"
          />
          <input
            type="text"
            placeholder="Value"
            value={pair.value}
            onChange={(e) => update(i, 'value', e.target.value)}
            className="flex-1 bg-surface0 text-text border border-surface1 rounded px-2.5 py-1.5 text-[13px] outline-none focus:border-accent"
          />
          <button
            className="bg-transparent border-none text-red text-lg cursor-pointer px-1"
            onClick={() => removeRow(i)}
          >
            &times;
          </button>
        </div>
      ))}
      <button
        className="bg-transparent border border-dashed border-surface1 text-overlay0 rounded py-1.5 text-[13px] cursor-pointer mt-1 hover:text-text hover:border-overlay0"
        onClick={addRow}
      >
        + Add
      </button>
    </div>
  );
}
