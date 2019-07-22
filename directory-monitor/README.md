# monitor

The monitor component is ran as a centralized NLX process. It monitors the 
network and updates the directory - every NLX network has one directory 
and one monitor. Currently the monitor simply 
executes health checks on inways, and updates their status in 
the directory database.
