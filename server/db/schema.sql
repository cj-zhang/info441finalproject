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
    website varchar(2083),
    tounament_location varchar(255) not null,
    tournament_organizer_id int not null,
    photo_url varchar(2083) not null
);

create table if not exists players (
    u_id int not null,
    tournament_id int not null,
    FOREIGN KEY (u_id) REFERENCES users(id),
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id)
);

create table if not exists tournament_organizers (
    u_id int not null,
    tournament_id int not null,
    FOREIGN KEY (u_id) REFERENCES users(id),
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id)
);

create table if not exists brackets (
    id int not null auto_increment primary key,
    tournament_id int not null,
    tournament_organizer_id int not null,
    bracket_location varchar(255) not null,
    in_progress boolean not null,
    completed boolean not null,
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id),
    FOREIGN KEY (tournament_organizer_id) REFERENCES tournament_organizers(u_id)
);

create table if not exists games (
    id int not null auto_increment primary key,
    tournament_id int not null,
    player_one int not null,
    player_two int not null,
    date_time datetime not null, 
    bracket_id int not null,
    tournament_organizer_id int not null,
    in_progress boolean not null,
    completed boolean not null,
    result varchar(10) not null,
    FOREIGN KEY (bracket_id) REFERENCES brackets(id),
    FOREIGN KEY (tournament_organizer_id) REFERENCES tournament_organizers(u_id),
    FOREIGN KEY (player_one) REFERENCES players(u_id),
    FOREIGN KEY (player_two) REFERENCES players(u_id)
);



-- create table if not exists single_tournament_users (
--     id int not null auto_increment primary key,
--     email varchar(254) not null unique,
--     username varchar(255) not null unique,
--     pass_hash char(150) not null,
--     first_name varchar(64) not null,
--     last_name varchar(128) not null,
--     photo_url varchar(2083) not null,
--     player int not null,
--     tournament_organizer int not null
-- );