import { StatusBar } from 'expo-status-bar';
import { useEffect, useState } from 'react';
import { ActivityIndicator, StyleSheet, Text, View } from 'react-native';

import { fetchHealth, type HealthResponse } from './src/api/client';

export default function App() {
  const [loading, setLoading] = useState(true);
  const [health, setHealth] = useState<HealthResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;

    async function run() {
      try {
        setLoading(true);
        setError(null);
        const result = await fetchHealth();
        if (cancelled) return;
        setHealth(result);
      } catch (e) {
        if (cancelled) return;
        setError(e instanceof Error ? e.message : String(e));
      } finally {
        if (cancelled) return;
        setLoading(false);
      }
    }

    run();

    return () => {
      cancelled = true;
    };
  }, []);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>routeradar-yyc</Text>
      <Text>Backend health:</Text>

      {loading ? (
        <ActivityIndicator style={{ marginTop: 12 }} />
      ) : error ? (
        <Text style={styles.error}>{error}</Text>
      ) : (
        <Text style={styles.ok}>{health?.status ?? 'unknown'}</Text>
      )}
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
  title: {
    fontSize: 20,
    fontWeight: '600',
    marginBottom: 12,
  },
  ok: {
    marginTop: 8,
    fontWeight: '600',
    color: '#0A7A0A',
  },
  error: {
    marginTop: 8,
    color: '#b00020',
    textAlign: 'center',
    paddingHorizontal: 24,
  },
});
