create table IF NOT EXISTS productos
(
    id          serial                         not null
        constraint productos_pkey
            primary key,
    nombre      varchar(100)                   not null,
    descripcion varchar(500),
    activo      boolean          default true,
    created     timestamp        default now() not null,
    update      timestamp,
    precio      double precision default 0.00
);

create table IF NOT EXISTS categorias
(
    id          serial                  not null
        constraint categorias_pkey
            primary key,
    nombre      varchar(100)            not null,
    descripcion varchar(500),
    activo      boolean   default true,
    created     timestamp default now() not null,
    update      timestamp
);

create table IF NOT EXISTS usuarios
(
    id              serial                     not null
        constraint usuarios_pkey
            primary key,
    nombre          varchar(100)               not null,
    apellido        varchar(500),
    email           varchar(100),
    password        varchar(256)               not null,
    fechaNacimiento timestamp,
    created         timestamp    default now() not null,
    update          timestamp
);
