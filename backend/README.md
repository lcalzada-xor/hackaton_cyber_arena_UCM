# Herramienta de Búsqueda de CVE

Una potente herramienta de línea de comandos (CLI) para buscar Vulnerabilidades y Exposiciones Comunes (CVEs) utilizando la API del NVD.

## Características

-   **Búsqueda por Palabra Clave**: Encuentra CVEs por palabra clave (ej. "log4j").
-   **Filtrado Avanzado**: Filtra por severidad CVSS, fecha de publicación, fecha de modificación, CPE, CWE y más.
-   **Salida Flexible**: Salida legible para humanos con resaltado de colores o JSON para automatización.
-   **Ordenación**: Ordena los resultados por fecha de publicación, fecha de modificación o puntuación CVSS.

## Instalación

Asegúrate de tener Go 1.24+ instalado.

```bash
go install github.com/lcalzada-xor/hackaton_cyber_arena_UCM/cmd/cve-search@latest
```

O compila desde el código fuente:

```bash
git clone https://github.com/lcalzada-xor/hackaton_cyber_arena_UCM.git
cd hackaton_cyber_arena_UCM
go build -o cve-search ./cmd/cve-search
```

## Uso

```bash
./cve-search [flags]
```

### Ejemplos

**Buscar vulnerabilidades de "log4j":**
```bash
./cve-search -k log4j
```

**Encontrar vulnerabilidades críticas publicadas en 2023:**
```bash
./cve-search -s CRITICAL -d 2023-01-01 -e 2023-12-31
```

**Buscar por CPE (ej. Apache Log4j):**
```bash
./cve-search -c cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*
```

**Salida como JSON:**
```bash
./cve-search -k wordpress -o json > resultados.json
```

### Flags (Opciones)

| Flag | Corto | Descripción |
|------|-------|-------------|
| `--keyword` | `-k` | Palabra clave a buscar |
| `--severity` | `-s` | Severidad CVSS v3 (LOW, MEDIUM, HIGH, CRITICAL) |
| `--start-date` | `-d` | Fecha de Inicio (AAAA-MM-DD) |
| `--end-date` | `-e` | Fecha de Fin (AAAA-MM-DD) |
| `--cpe` | `-c` | Nombre CPE |
| `--cwe` | `-w` | ID CWE |
| `--limit` | `-l` | Número de resultados a devolver (por defecto 10) |
| `--output` | `-o` | Formato de salida (human, json) |
| `--apikey` | `-a` | Clave API del NVD (opcional) |

## Desarrollo

### Ejecutar Tests

```bash
go test ./...
```

## API del Servidor

El servidor expone una API REST para realizar búsquedas.

### Endpoints

#### `GET /api/search`

Realiza una búsqueda de CVEs.

**Parámetros:**

| Parámetro | Descripción |
|-----------|-------------|
| `keyword` | Palabra clave a buscar |
| `severity` | Severidad CVSS v3 (LOW, MEDIUM, HIGH, CRITICAL) |
| `startDate` | Fecha de Inicio (AAAA-MM-DD) |
| `endDate` | Fecha de Fin (AAAA-MM-DD) |
| `cpe` | Nombre CPE |
| `cwe` | ID CWE |
| `cvssV2Severity` | Severidad CVSS v2 |
| `modStartDate` | Fecha de Inicio de Última Modificación |
| `modEndDate` | Fecha de Fin de Última Modificación |
| `source` | Identificador de Fuente |
| `limit` | Número de resultados (por defecto 10) |
| `sort` | Campo para ordenar (`published`, `modified`, `score`) |
| `direction` | Dirección del ordenamiento (`asc`, `desc`) |

**Ejemplo de uso:**

```bash
curl "http://localhost:8081/api/search?keyword=log4j&sort=published&direction=desc"
```