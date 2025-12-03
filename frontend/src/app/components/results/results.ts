import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CveResponse, Vulnerability } from '../../models/cve.model';

@Component({
  selector: 'app-results',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './results.html',
  styleUrl: './results.css'
})
export class ResultsComponent {
  @Input() data: CveResponse | null = null;

  get vulnerabilities(): Vulnerability[] {
    return this.data?.vulnerabilities || [];
  }

  getSeverityColor(v: Vulnerability): string {
    const metrics = v.cve.metrics;
    let score = 0;
    let severity = 'UNKNOWN';

    if (metrics.cvssMetricV31?.length) {
      score = metrics.cvssMetricV31[0].cvssData.baseScore;
      severity = metrics.cvssMetricV31[0].cvssData.baseSeverity;
    } else if (metrics.cvssMetricV30?.length) {
      score = metrics.cvssMetricV30[0].cvssData.baseScore;
      severity = metrics.cvssMetricV30[0].cvssData.baseSeverity;
    } else if (metrics.cvssMetricV2?.length) {
      score = metrics.cvssMetricV2[0].cvssData.baseScore;
      severity = metrics.cvssMetricV2[0].baseSeverity;
    }

    switch (severity.toUpperCase()) {
      case 'CRITICAL': return 'border-red-500 bg-red-500/10 text-red-400 shadow-red-500/20';
      case 'HIGH': return 'border-orange-500 bg-orange-500/10 text-orange-400 shadow-orange-500/20';
      case 'MEDIUM': return 'border-yellow-500 bg-yellow-500/10 text-yellow-400 shadow-yellow-500/20';
      case 'LOW': return 'border-green-500 bg-green-500/10 text-green-400 shadow-green-500/20';
      default: return 'border-gray-500 bg-gray-500/10 text-gray-400';
    }
  }

  getScore(v: Vulnerability): number {
    const metrics = v.cve.metrics;
    if (metrics.cvssMetricV31?.length) return metrics.cvssMetricV31[0].cvssData.baseScore;
    if (metrics.cvssMetricV30?.length) return metrics.cvssMetricV30[0].cvssData.baseScore;
    if (metrics.cvssMetricV2?.length) return metrics.cvssMetricV2[0].cvssData.baseScore;
    return 0;
  }
}
