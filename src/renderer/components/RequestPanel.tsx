import { useState } from 'react';
import { useRequestStore } from '../stores/requestStore';
import type { HttpMethod } from '../../shared/types';
import KeyValueEditor from './KeyValueEditor';

const METHODS: HttpMethod[] = ['GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'HEAD', 'OPTIONS'];

type Tab = 'params' | 'headers' | 'body';

export default function RequestPanel() {
  const { method, url, headers, params, body, loading, setMethod, setUrl, setHeaders, setParams, setBody, sendRequest } =
    useRequestStore();
  const [activeTab, setActiveTab] = useState<Tab>('params');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    sendRequest();
  };

  return (
    <div className="p-4 px-5 border-b border-surface0">
      <form className="flex gap-2" onSubmit={handleSubmit}>
        <select
          value={method}
          onChange={(e) => setMethod(e.target.value as HttpMethod)}
          className="bg-surface0 text-accent border border-surface1 rounded-md px-3 py-2 text-sm font-semibold cursor-pointer outline-none focus:border-accent"
        >
          {METHODS.map((m) => (
            <option key={m} value={m}>{m}</option>
          ))}
        </select>
        <input
          type="text"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="https://api.example.com/endpoint"
          className="flex-1 bg-surface0 text-text border border-surface1 rounded-md px-3 py-2 text-sm outline-none focus:border-accent"
        />
        <button
          type="submit"
          disabled={loading}
          className="bg-accent text-white border-none rounded-md px-6 py-2 text-sm font-semibold cursor-pointer hover:bg-accent-hover disabled:opacity-60 disabled:cursor-not-allowed"
        >
          {loading ? 'Sending...' : 'Send'}
        </button>
      </form>

      <div className="flex mt-3 border-b border-surface0">
        {(['params', 'headers', 'body'] as Tab[]).map((tab) => (
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
        {activeTab === 'params' && (
          <KeyValueEditor pairs={params} onChange={setParams} />
        )}
        {activeTab === 'headers' && (
          <KeyValueEditor pairs={headers} onChange={setHeaders} />
        )}
        {activeTab === 'body' && (
          <textarea
            className="w-full min-h-[120px] bg-surface0 text-text border border-surface1 rounded-md p-2.5 font-mono text-[13px] resize-y outline-none focus:border-accent"
            value={body}
            onChange={(e) => setBody(e.target.value)}
            placeholder='{"key": "value"}'
          />
        )}
      </div>
    </div>
  );
}
