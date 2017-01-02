move domio_dev.sql domio_dev.previous.sql
pg_dump -U sergeibasharov --schema-only --no-owner domio_dev > domio_dev.sql