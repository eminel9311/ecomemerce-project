## Run migrate database development
sql-migrate up -config=dbconfig.yml -env="development"
sql-migrate down -config=dbconfig.yml -env="development"

## Run migrate database production
sql-migrate up -config=dbconfig.yml -env="production"
sql-migrate down -config=dbconfig.yml -env="production"