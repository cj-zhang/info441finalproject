create table if not exists users (
    id int not null auto_increment primary key,
    email varchar(254) not null unique,
    username varchar(255) not null unique,
    pass_hash char(150) not null,
    first_name varchar(64) not null,
    last_name varchar(128) not null,
    photo_url varchar(2083) not null
);

create table if not exists tournaments (
    id int not null auto_increment primary key,
    website varchar(2083) not null,
    tounament_location varchar(255) not null,
    tournament_organizer_id int not null,
    photo_url varchar(2083) not null,
    database_address varchar(2083) not null
);

create table if not exists registration (
    u_id int not null,
    tournament_id int not null
);

create table if not exists games (
    id int not null auto_increment primary key,
    player_one int not null,
    player_two int not null,
    date_time datetime not null, 
    bracket_id int not null,
    tournament_organizer_id int not null,
    in_progress boolean not null,
    completed boolean not null,
    result varchar(10) not null
)

create table if not exists brackets {
    id int not null auto_increment primary key,
    tournament_organizer_id int not null,
    tournament_location varchar(255) not null,
    in_progress boolean not null,
    completed boolean not null
}

create table if not exists single_tournament_users (
    id int not null auto_increment primary key,
    email varchar(254) not null unique,
    username varchar(255) not null unique,
    pass_hash char(150) not null,
    first_name varchar(64) not null,
    last_name varchar(128) not null,
    photo_url varchar(2083) not null,
    player int not null,
    tournament_organizer int not null
);