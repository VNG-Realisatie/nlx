# db

Database for centralized components.

## Development

A pgmodeler model (`model/nlx.dbm`) is used to design the database, changes to the scheme are added in migration files (`./migrations/`).

When changing the model, make sure to run diff.sh. It verifies that the model and migration files are in sync, and otherwise shows differences.
