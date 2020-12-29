# go-app-resto

## Conectar con POSTGRES
### Conectar a postgresql para utiluzar por terminal
 - sudo -u postgres psql

### Cambiar password
- sudo -u postgres psql -c "ALTER USER postgres PASSWORD '9003';"

### Crear usuario
- CREATE USER gerrdev PASSWORD '9003';

### Crear base de datos y asignar usuario
- CREATE DATABASE restodb OWNER gerrdev;

### Crear tabla productos
- CREATE TABLE productos (
  id SERIAL PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  descripcion VARCHAR(50),
  activo BOOLEAN NOT NULL Default true,
  created TIMESTAMP NOT NULL DEFAULT NOW(),
  updated TIMESTAMP
  );
