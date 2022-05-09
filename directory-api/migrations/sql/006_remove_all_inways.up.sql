begin transaction;

delete from directory.availabilities;
delete from directory.inways;

commit;
