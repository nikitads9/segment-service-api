DROP   TABLE  IF EXISTS users;
DROP   TABLE  IF EXISTS segments;
DROP   TABLE  IF EXISTS users_segments_junction;

create table users
(
    id bigserial primary key,
    username text not null
    unique (username)
);

create table segments
(
    id bigserial primary key,
    slug text not null,
    unique(slug)
);

create table users_segments_junction
(
    user_id int not null,
    segment_id int not null,
    added_at timestamp,
    time_of_expire timestamp,
    state bool,
    constraint fk_users
        foreign key(user_id)
            references users(id)
            on delete cascade
  			on update cascade,
    constraint fk_segments
        foreign key(segment_id)
            references segments(id)
            on delete cascade
  			on update cascade,
    unique(user_id, segment_id)
);

