import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CveResponse, SearchParams } from '../models/cve.model';

import { environment } from '../../environments/environment';

@Injectable({
    providedIn: 'root'
})
export class CveService {
    private apiUrl = `${environment.apiUrl}/search`;

    constructor(private http: HttpClient) { }

    searchCves(params: SearchParams): Observable<CveResponse> {
        let httpParams = new HttpParams();

        if (params.keyword) httpParams = httpParams.set('keyword', params.keyword);
        if (params.severity) httpParams = httpParams.set('severity', params.severity);
        if (params.startDate) httpParams = httpParams.set('startDate', params.startDate);
        if (params.endDate) httpParams = httpParams.set('endDate', params.endDate);
        if (params.cpe) httpParams = httpParams.set('cpe', params.cpe);
        if (params.cwe) httpParams = httpParams.set('cwe', params.cwe);
        if (params.limit) httpParams = httpParams.set('limit', params.limit);

        return this.http.get<CveResponse>(this.apiUrl, { params: httpParams });
    }
}
