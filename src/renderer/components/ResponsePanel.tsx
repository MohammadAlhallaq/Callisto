import { useState } from 'react';
import { useRequestStore } from '../stores/requestStore';

type Tab = 'body' | 'headers';

function formatSize(bytes: number): string {
  if (bytes < 1024) return `${bytes} B`;
  return `${(bytes / 1024).toFixed(1)} KB`;
}

function tryFormatJson(text: string): string {
  try {
    return JSON.stringify(JSON.parse(text), null, 2);
  } catch {
    return text;
  }
}

export default function ResponsePanel() {
  const { response, error, loading } = useRequestStore();
  const [activeTab, setActiveTab] = useState<Tab>('body');

  if (loading) {
    return <div className="flex-1 p-5 flex items-center justify-center text-overlay0">Sending request...</div>;
  }

  if (error) {
    return <div className="flex-1 p-5 text-red">Error: {error}</div>;
  }

  if (!response) {
    return <div className="flex-1 p-5 flex items-center justify-center text-overlay0">Send a request to see the response</div>;
  }

  const statusClass =
    response.status < 300
      ? 'bg-green/15 text-green'
      : response.status < 400
        ? 'bg-yellow/15 text-yellow'
        : 'bg-red/15 text-red';

  return (
    <div className="flex-1 p-4 px-5 overflow-auto">
      <div className="flex items-center gap-4 mb-3">
        <span className={`font-semibold px-2.5 py-1 rounded text-[13px] ${statusClass}`}>
          {response.status} {response.statusText}
        </span>
        <span className="text-overlay0 text-[13px]">{response.duration}ms</span>
        <span className="text-overlay0 text-[13px]">{formatSize(response.size)}</span>
      </div>

      <div className="flex border-b border-surface0">
        {(['body', 'headers'] as Tab[]).map((tab) => (
          <button
            key={tab}
            className={`bg-transparent border-none px-4 py-2 text-[13px] cursor-pointer border-b-2 ${
              activeTab === tab
                ? 'text-text border-b-accent'
                : 'text-overlay0 border-b-transparent'
            }`}
            onClick={() => setActiveTab(tab)}
          >
            {tab.charAt(0).toUpperCase() + tab.slice(1)}
          </button>
        ))}
      </div>

      <div className="mt-3">
        {activeTab === 'body' && (
          <pre className="bg-surface0 rounded-md p-3 font-mono text-[13px] overflow-auto whitespace-pre-wrap break-words">
            {tryFormatJson(response.body)}
          </pre>
        )}
        {activeTab === 'headers' && (
          <div className="flex flex-col gap-1">
            {Object.entries(response.headers).map(([key, value]) => (
              <div key={key} className="flex gap-3 py-1 border-b border-surface0">
                <span className="text-accent font-medium min-w-[200px]">{key}</span>
                <span className="text-text break-all">{value}</span>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
