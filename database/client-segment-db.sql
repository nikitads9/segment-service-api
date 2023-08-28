/*DROP   TABLE  IF EXISTS schemas.clients;
DROP   TABLE  IF EXISTS schemas.segments;
DROP   TABLE  IF EXISTS schemas.clients_segments_junction;
DROP   SCHEMA IF EXISTS schemas;
CREATE SCHEMA schemas;*/
create table clients
(
    id bigserial primary key,
    username text not null
);

create table segments
(
    id bigserial primary key,
    title text not null
);

create table clients_segments_junction
(
    client_id int not null,
    segment_id int not null,
    constraint fk_client
        foreign key(client_id)
            references clients(id)
            on delete cascade
  			on update cascade,
    constraint fk_segments
        foreign key(segment_id)
            references segments(id)
            on delete cascade
  			on update cascade
);

