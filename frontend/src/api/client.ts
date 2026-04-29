const API_BASE_URL = 'http://localhost:8080';

export type HealthResponse = {
  status: string;
};

export async function fetchHealth(): Promise<HealthResponse> {
  const res = await fetch(`${API_BASE_URL}/health`, { method: 'GET' });

  if (!res.ok) {
    const text = await res.text().catch(() => '');
    throw new Error(`Health check failed: ${res.status} ${text}`.trim());
  }

  return (await res.json()) as HealthResponse;
}

