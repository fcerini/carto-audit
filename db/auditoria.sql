--
-- PostgreSQL database dump
--

\restrict BHLXPbYfjmp6jgKLdKgADt2d30rYPdwUT3yDyOlvGdrdUxWNZmUBHOBlTox1Lpp

-- Dumped from database version 15.12 (Debian 15.12-0+deb12u2)
-- Dumped by pg_dump version 18.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: auditoria; Type: SCHEMA; Schema: -; Owner: natienza
--

CREATE SCHEMA auditoria;


ALTER SCHEMA auditoria OWNER TO natienza;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cordoba202512; Type: TABLE; Schema: auditoria; Owner: fcerini
--

CREATE TABLE auditoria.cordoba202512 (
    id bigint,
    provincia text,
    partido text,
    localidad text,
    calle text,
    fromleft numeric(11,0),
    toleft numeric(11,0),
    fromright numeric(11,0),
    toright numeric(11,0),
    alias text,
    wkt text,
    calleinter text,
    geom public.geometry(LineString,4326)
);


ALTER TABLE auditoria.cordoba202512 OWNER TO fcerini;

--
-- Name: procesar; Type: TABLE; Schema: auditoria; Owner: fcerini
--

CREATE TABLE auditoria.procesar (
    idpk integer NOT NULL,
    suceso text,
    provincia text,
    partido text,
    localidad text,
    calle text,
    altura text,
    callesuperior text,
    calleinferior text,
    latitud text,
    longitud text,
    resultado text
);


ALTER TABLE auditoria.procesar OWNER TO fcerini;

--
-- Name: sucesos_cordoba; Type: TABLE; Schema: auditoria; Owner: fcerini
--

CREATE TABLE auditoria.sucesos_cordoba (
    fecha text,
    suceso text,
    agencia text,
    tipo text,
    subtipo text,
    relato text,
    operador text,
    despachador text,
    provincia text,
    partido text,
    localidad text,
    calle text,
    altura text,
    comentario_calle text,
    latitud text,
    longitud text,
    id2 integer,
    geom public.geometry(Point,4326),
    geo text,
    relevante text,
    chequeado text,
    callesuperior text,
    calleinferior text,
    calle2 text,
    callesuperior2 text,
    calleinferior2 text,
    altura2 integer
);


ALTER TABLE auditoria.sucesos_cordoba OWNER TO fcerini;

--
-- Name: procesar procesar_pkey; Type: CONSTRAINT; Schema: auditoria; Owner: fcerini
--

ALTER TABLE ONLY auditoria.procesar
    ADD CONSTRAINT procesar_pkey PRIMARY KEY (idpk);


--
-- Name: idx_cordoba202512_geom; Type: INDEX; Schema: auditoria; Owner: fcerini
--

CREATE INDEX idx_cordoba202512_geom ON auditoria.cordoba202512 USING gist (geom);


--
-- Name: SCHEMA auditoria; Type: ACL; Schema: -; Owner: natienza
--

GRANT ALL ON SCHEMA auditoria TO fcerini;


--
-- PostgreSQL database dump complete
--

\unrestrict BHLXPbYfjmp6jgKLdKgADt2d30rYPdwUT3yDyOlvGdrdUxWNZmUBHOBlTox1Lpp

