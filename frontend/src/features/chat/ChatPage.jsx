import { useEffect, useState } from 'react';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import { askAssistant } from '../../api/assistant.js';

export default function ChatPage() {
  const [status, setStatus] = useState('Checking API…');
  const [question, setQuestion] = useState('');
  const [messages, setMessages] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    const controller = new AbortController();
    fetch('/api/health', { signal: controller.signal })
      .then((response) => { if (!response.ok) throw new Error(); return response.json(); })
      .then((data) => setStatus(data.status === 'ok' ? 'API connected' : 'API unavailable'))
      .catch((err) => { if (err.name !== 'AbortError') setStatus('API unavailable'); });
    return () => controller.abort();
  }, []);

  async function submit(event) {
    event.preventDefault();
    const value = question.trim();
    if (!value || loading) return;
    setMessages((previous) => [...previous, { role: 'You', content: value }]);
    setQuestion(''); setError(''); setLoading(true);
    try {
      const answer = await askAssistant(value);
      setMessages((previous) => [...previous, { role: 'Assistant', content: answer }]);
    } catch (err) { setError(err.message); setQuestion(value); }
    finally { setLoading(false); }
  }

  return <main>
    <header><span className="mark">✳</span><strong>Policy Assistant</strong><span className="status">{status}</span></header>
    <section>
      <p className="eyebrow">YOUR INSURANCE LEARNING COMPANION</p>
      <h1>Insurance, made clearer.</h1>
      <p>Ask a general question about insurance. Policy documents and citations are coming in the next phase.</p>
      <div className="chat" aria-live="polite" aria-busy={loading}>
        {messages.length === 0 && <p className="empty">Try asking: “What is an insurance deductible?”</p>}
        {messages.map((message, index) => <article className="message" key={index}><strong>{message.role}</strong>{message.role === 'Assistant' ? <div className="markdown"><ReactMarkdown remarkPlugins={[remarkGfm]}>{message.content}</ReactMarkdown></div> : <p>{message.content}</p>}</article>)}
        {loading && <p className="empty">Thinking…</p>}
      </div>
      <form onSubmit={submit}>
        <label htmlFor="question">Your question</label>
        <textarea id="question" value={question} onChange={(event) => setQuestion(event.target.value)} placeholder="Ask about insurance…" maxLength={4000} rows={3} disabled={loading} />
        <div className="form-footer"><span>Each question is sent independently.</span><button disabled={loading || !question.trim()}>{loading ? 'Sending…' : 'Ask assistant'}</button></div>
      </form>
      {error && <p className="error" role="alert">{error}</p>}
      <small>Use fictional data. The assistant cannot access your policy and may make mistakes.</small>
    </section>
  </main>;
}
