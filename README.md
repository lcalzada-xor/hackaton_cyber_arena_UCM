# Hackaton Cyber Arena UCM

Este proyecto contiene el código fuente para el desafío Hackaton Cyber Arena UCM.

## Estructura del Proyecto

- **backend/**: Contiene el servicio backend en Go.
  - API REST con soporte para búsqueda avanzada y ordenación.
  - Integración con ExploitDB para enriquecimiento de vulnerabilidades.
- **frontend/**: Contiene la aplicación frontend en Angular.

## Primeros Pasos

### Backend

Navega al directorio `backend`:

```bash
cd backend
go run cmd/server/main.go
```

El servidor se iniciará en el puerto 8080.

### Frontend

Navega al directorio `frontend`:

```bash
cd frontend
npm install
ng serve
```

La aplicación estará disponible en `http://localhost:4200/`.
