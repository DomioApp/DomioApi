move ../../dbschema/domio_dev.sql ../../dbschema/domio_dev.previous.sql
pg_dump -U sergeibasharov --schema-only --no-owner domio_dev > ../../dbschema/domio_dev.sql