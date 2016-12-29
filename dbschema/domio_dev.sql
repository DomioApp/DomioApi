--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.5
-- Dumped by pg_dump version 9.5.5

SET statement_timeout = 0;
SET lock_timeout = 0;
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
    is_rented boolean DEFAULT false NOT NULL,
    period_end date,
    CONSTRAINT domain_name_min_length CHECK ((length((name)::text) > 3))
);


--
-- Name: payments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE payments (
    months_count integer NOT NULL,
    amount_in_cents integer NOT NULL,
    paid_by character varying NOT NULL,
    id bigint NOT NULL,
    stripe_id character varying
);


--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE payments_id_seq OWNED BY payments.id;


--
-- Name: rentals; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE rentals (
    domain_name character varying NOT NULL,
    period_range daterange,
    id bigint NOT NULL,
    renter_id character varying NOT NULL,
    payment_id bigint
);


--
-- Name: rentals_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE rentals_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: rentals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE rentals_id_seq OWNED BY rentals.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE users (
    email character varying(30) NOT NULL,
    password text,
    id character(18)
);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY payments ALTER COLUMN id SET DEFAULT nextval('payments_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY rentals ALTER COLUMN id SET DEFAULT nextval('rentals_id_seq'::regclass);


--
-- Name: domains_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY domains
    ADD CONSTRAINT domains_pkey PRIMARY KEY (name);


--
-- Name: id; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY rentals
    ADD CONSTRAINT id PRIMARY KEY (id);


--
-- Name: id_serial; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY payments
    ADD CONSTRAINT id_serial PRIMARY KEY (id);


--
-- Name: users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (email);


--
-- Name: fki_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_id ON rentals USING btree (domain_name);


--
-- Name: fki_owner; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_owner ON domains USING btree (owner);


--
-- Name: fki_payer_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_payer_id ON payments USING btree (paid_by);


--
-- Name: fki_renter_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX fki_renter_id ON rentals USING btree (renter_id);


--
-- Name: domain_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY rentals
    ADD CONSTRAINT domain_id FOREIGN KEY (domain_name) REFERENCES domains(name);


--
-- Name: owner; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY domains
    ADD CONSTRAINT owner FOREIGN KEY (owner) REFERENCES users(email);


--
-- Name: payer_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY payments
    ADD CONSTRAINT payer_id FOREIGN KEY (paid_by) REFERENCES users(email);


--
-- Name: rentals_payment_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY rentals
    ADD CONSTRAINT rentals_payment_id_fkey FOREIGN KEY (payment_id) REFERENCES payments(id);


--
-- Name: renter_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY rentals
    ADD CONSTRAINT renter_id FOREIGN KEY (renter_id) REFERENCES users(email);


--
-- Name: public; Type: ACL; Schema: -; Owner: -
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM sergeibasharov;
GRANT ALL ON SCHEMA public TO sergeibasharov;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

