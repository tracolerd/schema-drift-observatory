<script setup lang="ts">
import { ref } from 'vue'
import sampleReport from './assets/sample-report.json'
const report = ref(sampleReport)
</script>
<template>
  <div class="dashboard">
    <header>
      <h1>Schema Drift Observatory</h1>
      <p>API Compatibility Report</p>
    </header>
    <main v-if="report">
      <section class="summary" :class="{ 'is-compatible': report.compatible, 'is-incompatible': !report.compatible }">
        <h2>Status: {{ report.compatible ? 'Compatible' : 'Incompatible' }}</h2>
        <div class="counts">
          <span class="badge breaking">Breaking: {{ report.counts.breaking }}</span>
          <span class="badge potentially">Potentially Breaking: {{ report.counts.potentially_breaking }}</span>
          <span class="badge non-breaking">Non-Breaking: {{ report.counts.non_breaking }}</span>
        </div>
      </section>
      <section class="findings">
        <h2>Findings ({{ report.findings.length }})</h2>
        <div v-for="(finding, index) in report.findings" :key="index" class="card" :class="finding.category">
          <h3><span class="method-badge">{{ finding.method }}</span> {{ finding.path }}</h3>
          <p><strong>Rule:</strong> {{ finding.rule_id }} | <strong>Category:</strong> {{ finding.category.toUpperCase() }}</p>
          <p><strong>Explanation:</strong> {{ finding.explanation }}</p>
          <p><strong>Impact:</strong> {{ finding.impact }}</p>
          <p><strong>Remediation:</strong> {{ finding.remediation }}</p>
          <code class="location">{{ finding.location }}</code>
        </div>
      </section>
    </main>
  </div>
</template>
<style scoped>
.dashboard { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; max-width: 900px; margin: 0 auto; padding: 20px; color: #2c3e50; }
header { border-bottom: 2px solid #eaeaea; margin-bottom: 20px; padding-bottom: 10px; }
h1 { margin: 0; color: #1a252f; }
.summary { padding: 20px; border-radius: 8px; margin-bottom: 30px; box-shadow: 0 2px 4px rgba(0,0,0,0.05); }
.is-compatible { background-color: #f0fdf4; border: 1px solid #bbf7d0; }
.is-incompatible { background-color: #fef2f2; border: 1px solid #fecaca; }
.counts { display: flex; gap: 10px; margin-top: 15px; }
.badge { padding: 6px 12px; border-radius: 4px; font-weight: bold; font-size: 0.85em; }
.badge.breaking { background: #ef4444; color: white; }
.badge.potentially { background: #f97316; color: white; }
.badge.non-breaking { background: #22c55e; color: white; }
.card { border: 1px solid #e5e7eb; padding: 20px; border-radius: 8px; margin-bottom: 15px; background: #fff; box-shadow: 0 1px 3px rgba(0,0,0,0.05); }
.card.breaking { border-left: 5px solid #ef4444; }
.card.potentially_breaking { border-left: 5px solid #f97316; }
.card.non_breaking { border-left: 5px solid #22c55e; }
h3 { margin-top: 0; display: flex; align-items: center; gap: 10px; font-size: 1.1em;}
.method-badge { background: #e2e8f0; color: #475569; padding: 3px 8px; border-radius: 4px; font-size: 0.8em; }
.location { display: block; margin-top: 10px; padding: 8px; background: #f8fafc; border-radius: 4px; font-size: 0.85em; color: #64748b; }
</style>
