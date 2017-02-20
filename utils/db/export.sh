#!/usr/bin/env bash

mv ../../dbschema/domio_dev.sql ../../dbschema/domio_dev.previous.sql
pg_dump -U postgres --schema-only --no-owner domio_dev > ../../dbschema/domio_dev.sql