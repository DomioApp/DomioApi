--Table: public.users

DROP TABLE public.users;

CREATE TABLE public.users (
  email       varchar(30) NOT NULL,
  "password"  text,
  /* Keys */
  CONSTRAINT users_pkey
    PRIMARY KEY (email)
) WITH (
    OIDS = FALSE
  );

ALTER TABLE public.users
  OWNER TO postgres;
