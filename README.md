# Hackaton Cyber Arena UCM

Este proyecto contiene el código fuente para el desafío Hackaton Cyber Arena UCM.

### Tabla de Contenidos
----
1. [[#Descripción del Proyecto]]
	- [[#Estructura del Proyecto]]
	- [[#Tecnología Utilizada]]
2. [[#Primeros Pasos]]
	- [[#Instalación y Ejecución (Docker)]]
	- [[#Instalación y Ejecución (sin Docker)]]
3. [[#Contacto]]

### Descripción del Proyecto
----
Este proyecto consiste en una herramienta capaz de buscar CVEs que permita localizar vulnerabilidades por nombre, ID, palabras clave y otros filtros. Dispone de un frontend intuitivo e incluye información sobre exploits relacionados.

También estamos considerando integrar un modelo de IA que resuma el resultado generado para extraer los puntos más relevantes.

#### Estructura del Proyecto

- **backend/**: Contiene el servicio backend en Go.
  - API REST con soporte para búsqueda avanzada y ordenación.
  - Integración con ExploitDB para enriquecimiento de vulnerabilidades.
- **frontend/**: Contiene la aplicación frontend en Angular.

#### Tecnología Utilizada
- Node.js
- Golang
- Angular

### Primeros Pasos
---
#### Instalación y Ejecución (Docker)

1.  Clona el repositorio:  
```
git clone https://github.com/lcalzada-xor/hackaton_cyber_arena_UCM
```  

2. Cambia de directorio:  
`` cd hackaton_cyber_arena_UCM `` 

3. Ejecuta la siguiente instrucción de docker-compose para montar el entorno: 
```
docker-compose up --build
```  

#### Instalación y Ejecución (sin Docker)

1.  Clona el repositorio:  
```
git clone https://github.com/lcalzada-xor/hackaton_cyber_arena_UCM
```  

2. Cambia de directorio:  
`` cd hackaton_cyber_arena_UCM `` 

3. Navega al directorio `backend` y corre:
```bash
cd backend
go run cmd/server/main.go
```

El servidor se iniciará en el puerto 8080.

4. Navega al directorio `frontend` y corre:
```bash
cd frontend
npm install
ng serve
```

La aplicación estará disponible en `http://localhost:4200/`.

### Contacto
---
Para cualquier detalle, contacta a los propietarios del proyecto en LinkedIn:

[lucas-calzada-del-pozo](https://www.linkedin.com/in/lucas-calzada-del-pozo-562571304/ )
[alberto-meléndez-garcía](https://www.linkedin.com/in/alberto-melendez-garcia-4713a1264/)
[javier-julve-yubero](https://www.linkedin.com/in/javier-julve-yubero-188203384/ "https://www.linkedin.com/in/javier-julve-yubero-188203384/")
[pablo-garcía-viña](https://www.linkedin.com/in/pablo-garc%C3%ADa-vi%C3%B1a/)

