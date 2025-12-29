ALTER TABLE auditoria.cordoba202512 OWNER TO fcerini;

GRANT USAGE ON SCHEMA auditoria TO fcerini;
GRANT CREATE ON SCHEMA auditoria TO fcerini;


ALTER TABLE auditoria.cordoba202512 
ADD COLUMN geom geometry(LineString, 4326);

UPDATE auditoria.cordoba202512 SET geom = ST_GeomFromText(wkt, 4326);

CREATE INDEX idx_cordoba202512_geom ON auditoria.cordoba202512 USING GIST (geom);