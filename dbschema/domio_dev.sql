--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: domains; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE domains (
    name character varying NOT NULL,
    owner character varying NOT NULL,
    price_per_month integer NOT NULL,
    zone_id character varying(26),
    is_rented boolean DEFAULT false NOT NULL,
    rented_by character varying,
    ns1 character varying,
    ns2 character varying,
    ns3 character varying,
    ns4 character varying,
    CONSTRAINT domain_name_min_length CHECK ((length((name)::text) > 3))
);
ALTER TABLE ONLY domains ALTER COLUMN rented_by SET STORAGE PLAIN;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE users (
    email character varying(30) NOT NULL,
    password text,
    id character(18),
    role character varying(10) NOT NULL
);


--
-- Name: domains domains_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY domains
    ADD CONSTRAINT domains_pkey PRIMARY KEY (name);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (email);


--
-- Name: fki_owner; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_owner ON domains USING btree (owner);


--
-- Name: fki_rented_by; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_rented_by ON domains USING btree (rented_by);


--
-- Name: domains owner; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY domains
    ADD CONSTRAINT owner FOREIGN KEY (owner) REFERENCES users(email);


--
-- Name: domains rented_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY domains
    ADD CONSTRAINT rented_by FOREIGN KEY (rented_by) REFERENCES users(email);


--
-- Name: public; Type: ACL; Schema: -; Owner: -
--

GRANT ALL ON SCHEMA public TO sergeibasharov;


--
-- PostgreSQL database dump complete
--

