import RequestPanel from './components/RequestPanel';
import ResponsePanel from './components/ResponsePanel';

export default function App() {
  return (
    <div className="flex flex-col h-screen bg-base text-text text-sm">
      <header className="px-5 py-3 bg-mantle border-b border-surface0">
        <h1 className="text-lg font-semibold text-accent">Callisto</h1>
      </header>
      <main className="flex flex-col flex-1 overflow-hidden">
        <RequestPanel />
        <ResponsePanel />
      </main>
    </div>
  );
}
