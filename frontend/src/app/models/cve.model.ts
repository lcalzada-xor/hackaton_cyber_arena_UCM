export interface CveResponse {
    resultsPerPage: number;
    startIndex: number;
    totalResults: number;
    format: string;
    version: string;
    timestamp: string;
    vulnerabilities: Vulnerability[];
}

export interface Vulnerability {
    cve: Cve;
    exploits?: Exploit[];
}

export interface Exploit {
    id: string;
    name: string;
    type: string;
    url: string;
    description: string;
    date: string;
    author: string;
}

export interface Cve {
    id: string;
    sourceIdentifier: string;
    published: string;
    lastModified: string;
    vulnStatus: string;
    descriptions: Description[];
    metrics: Metrics;
    weaknesses?: Weakness[];
    references: Reference[];
}

export interface Description {
    lang: string;
    value: string;
}

export interface Metrics {
    cvssMetricV31?: CvssMetricV3[];
    cvssMetricV30?: CvssMetricV3[];
    cvssMetricV2?: CvssMetricV2[];
}

export interface CvssMetricV3 {
    source: string;
    type: string;
    cvssData: CvssDataV3;
    exploitabilityScore: number;
    impactScore: number;
}

export interface CvssDataV3 {
    version: string;
    vectorString: string;
    attackVector: string;
    attackComplexity: string;
    privilegesRequired: string;
    userInteraction: string;
    scope: string;
    confidentialityImpact: string;
    integrityImpact: string;
    availabilityImpact: string;
    baseScore: number;
    baseSeverity: string;
}

export interface CvssMetricV2 {
    source: string;
    type: string;
    cvssData: CvssDataV2;
    baseSeverity: string;
    exploitabilityScore: number;
    impactScore: number;
    acInsufInfo: boolean;
    obtainAllPrivilege: boolean;
    obtainUserPrivilege: boolean;
    obtainOtherPrivilege: boolean;
    userInteractionRequired: boolean;
}

export interface CvssDataV2 {
    version: string;
    vectorString: string;
    accessVector: string;
    accessComplexity: string;
    authentication: string;
    confidentialityImpact: string;
    integrityImpact: string;
    availabilityImpact: string;
    baseScore: number;
}

export interface Weakness {
    source: string;
    type: string;
    description: Description[];
}

export interface Reference {
    url: string;
    source: string;
    tags?: string[];
}

export interface SearchParams {
    keyword?: string;
    severity?: string;
    startDate?: string;
    endDate?: string;
    cpe?: string;
    cwe?: string;
    limit?: number;
}
