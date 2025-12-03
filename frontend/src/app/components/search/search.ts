import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { CveService } from '../../services/cve.service';
import { CveResponse, SearchParams } from '../../models/cve.model';
import { ResultsComponent } from '../results/results';

@Component({
  selector: 'app-search',
  standalone: true,
  imports: [CommonModule, FormsModule, ResultsComponent],
  templateUrl: './search.html',
  styleUrl: './search.css'
})
export class SearchComponent {
  params: SearchParams = {
    limit: 10
  };
  results: CveResponse | null = null;
  loading = false;
  error = '';

  severities = [
    { label: 'BAJA', value: 'LOW' },
    { label: 'MEDIA', value: 'MEDIUM' },
    { label: 'ALTA', value: 'HIGH' },
    { label: 'CRÍTICA', value: 'CRITICAL' }
  ];
  showFilters = false;

  constructor(private cveService: CveService) { }

  toggleFilters() {
    this.showFilters = !this.showFilters;
  }

  getActiveFilters(): { key: string, value: any, label: string }[] {
    const active: { key: string, value: any, label: string }[] = [];
    if (this.params.severity) {
      const severityLabel = this.severities.find(s => s.value === this.params.severity)?.label || this.params.severity;
      active.push({ key: 'severity', value: this.params.severity, label: `Severidad: ${severityLabel}` });
    }
    if (this.params.startDate) active.push({ key: 'startDate', value: this.params.startDate, label: `Inicio: ${this.params.startDate}` });
    if (this.params.endDate) active.push({ key: 'endDate', value: this.params.endDate, label: `Fin: ${this.params.endDate}` });
    if (this.params.cpe) active.push({ key: 'cpe', value: this.params.cpe, label: `CPE: ${this.params.cpe}` });
    if (this.params.cwe) active.push({ key: 'cwe', value: this.params.cwe, label: `CWE: ${this.params.cwe}` });
    if (this.params.limit && this.params.limit !== 10) active.push({ key: 'limit', value: this.params.limit, label: `Límite: ${this.params.limit}` });
    return active;
  }

  removeFilter(key: string) {
    switch (key) {
      case 'severity': this.params.severity = undefined; break;
      case 'startDate': this.params.startDate = undefined; break;
      case 'endDate': this.params.endDate = undefined; break;
      case 'cpe': this.params.cpe = ''; break;
      case 'cwe': this.params.cwe = ''; break;
      case 'limit': this.params.limit = 10; break;
    }
    this.search();
  }

  clearAllFilters() {
    this.params = { limit: 10, keyword: this.params.keyword };
    this.search();
  }

  search() {
    this.loading = true;
    this.error = '';
    this.results = null;

    // Clean empty params
    const searchParams: SearchParams = {};
    if (this.params.keyword) searchParams.keyword = this.params.keyword;
    if (this.params.severity) searchParams.severity = this.params.severity;
    if (this.params.startDate) searchParams.startDate = this.params.startDate;
    if (this.params.endDate) searchParams.endDate = this.params.endDate;
    if (this.params.cpe) searchParams.cpe = this.params.cpe;
    if (this.params.cwe) searchParams.cwe = this.params.cwe;
    if (this.params.limit) searchParams.limit = this.params.limit;

    this.cveService.searchCves(searchParams).subscribe({
      next: (data) => {
        this.results = data;
        this.loading = false;
      },
      error: (err) => {
        this.error = 'Error al obtener resultados. Por favor, inténtelo de nuevo.';
        console.error(err);
        this.loading = false;
      }
    });
  }
}
