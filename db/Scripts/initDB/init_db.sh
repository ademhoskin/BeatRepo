#!/bin/bash

CONT_NAME="PostgresCont"
DB_USER="postgres"
DB_NAME="postgres"
SQL_SCRIPT_PATH="create_table.sql"
PGPASSWORD="postgres"

export PGPASSWORD
docker exec -i $CONT_NAME psql -U $DB_USER -d $DB_NAME -a -f - < $SQL_SCRIPT_PATH
unset PGPASSWORD
