# monitor

The monitor component is ran as a centralized NLX process. It monitors the 
network and updates the directory - every NLX network has one directory 
and one monitor. Currently the monitor simply 
executes health checks on inways, and updates their status in 
the directory database.

# Tests.

Add `postgres` to your `etc/hosts` file to point at 127.0.0.1
or to your docker hosts (non linux).

Start postgres docker from `testing/integration`

    docker-compose up postgres

Run migrations

    docker-compose up directory-db

Run tests

    $ go test -tags integration

Destroy postgres container when you are done.

TODO add this in a script and run this is CI/CD.
