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
    id          serial                         not null
        constraint categorias_pkey
            primary key,
    nombre      varchar(100)                   not null,
    descripcion varchar(500),
    activo      boolean          default true,
    created     timestamp        default now() not null,
    update      timestamp
);
