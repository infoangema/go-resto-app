create table Pedido
(
    id          int auto_increment
        primary key,
    id_producto int not null,
    id_user     int not null,
    id_mesa     int not null
);

create table Rol
(
    id          int auto_increment
        primary key,
    nombre      varchar(15) not null,
    descripcion varchar(40) null
);

create table User
(
    id       int         null,
    nombre   varchar(40) null,
    apellido varchar(40) null,
    id_rol   int         null,
    password varchar(30) null,
    constraint role___fk
        foreign key (id_rol) references Rol (id)
)
comment 'Usuario del sistema';

create table categorias
(
    id          int          not null
        primary key,
    descripcion varchar(120) not null,
    estado      bit          null,
    imagen      varchar(255) null,
    nombre      varchar(120) not null
);

create table hibernate_sequence
(
    next_val bigint null
)
    engine = MyISAM;

create table productos
(
    id           int                  not null
        primary key,
    nombre       varchar(100)         not null,
    sabor        varchar(100)         not null,
    graduacion   varchar(100)         not null,
    descripcion  varchar(250)         null,
    precio       double               not null,
    id_categoria int                  not null,
    imagen       varchar(255)         null,
    status       tinyint(1) default 1 null,
    constraint fk_productos_categorias
        foreign key (id_categoria) references categorias (id)
)
    charset = utf8mb4;

create index fk_productos_categorias_idx
    on productos (id_categoria);
