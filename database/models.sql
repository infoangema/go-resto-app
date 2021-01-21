create table if not exists productos
(
    id          serial                         not null
        constraint productos_pkey
            primary key,
    nombre      varchar(100)                   not null,
    descripcion varchar(500),
    activo      boolean          default true,
    Created_at     timestamp        default now() not null,
    update      timestamp        default now() not null,
    precio      double precision default 0.00,
    categoriaid integer
);

alter table productos
    add constraint producto_categoria_fk1 foreign key (id) references categorias;

create table if not exists categorias
(
    id          serial                  not null
        constraint categorias_pkey
            primary key,
    nombre      varchar(100)            not null,
    descripcion varchar(500),
    activo      boolean   default true,
    Created_at     timestamp default now() not null,
    Updated_at     timestamp default now() not null
);

create table if not exists usuarios
(
    id              serial                  not null
        constraint usuarios_pkey
            primary key,
    nombre          varchar(100)            not null,
    apellido        varchar(500),
    email           varchar(100),
    password        varchar(256)            not null,
    fechanacimiento timestamp,
    Created_at         timestamp default now() not null,
    Updated_at         timestamp default now() not null
);

