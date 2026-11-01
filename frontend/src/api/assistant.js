export async function askAssistant(question) {
  const response = await fetch('/api/chat', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ question }),
  });
  const data = await response.json();
  if (!response.ok) throw new Error(data.error || 'Unable to send your question.');
  return data;
}
