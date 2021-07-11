select * from items;
select * from people;
select * from address;
select * from users;
select * from lu_user_types;

CREATE TABLE IF NOT EXISTS address(
  id SERIAL primary key,
  line_1 VARCHAR(100) NOT NULL,
  line_2 VARCHAR(100),
  city VARCHAR(100) NOT NULL,
  state VARCHAR(100) NOT NULL,
  postcode VARCHAR(16),
  country VARCHAR(100) NOT NULL
);

drop table if exists people;

CREATE TABLE IF NOT EXISTS people(
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  middle_name VARCHAR(100),
  last_name VARCHAR(100) NOT NULL,
  d_o_b DATE,
  phone_number text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  address_id INT,
  cognito_id text,
  constraint fk_address
  	foreign key(address_id)
  		references address(id)
);

create table if not exists lu_user_types (
	id int unique not null,
	name text
);

create table if not exists users (
	id int not null,
	email text,
	password text,
	user_type_id int not null,
	constraint pk_user primary key (id),
	constraint fk_user_person foreign key (id) references people(id),
	constraint fk_type foreign key (user_type_id) references lu_user_types(id)
);

insert into lu_user_types (id, name) values (1000, 'Student');
insert into lu_user_types (id, name) values (2000, 'Parent');
insert into lu_user_types (id, name) values (3000, 'Tutor');
insert into lu_user_types (id, name) values (4000, 'Franchisee');
insert into lu_user_types (id, name) values (5000, 'Admin');

