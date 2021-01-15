BEGIN transaction; 

DROP TABLE nlx_management.permissions_roles;
DROP TABLE nlx_management.users_roles;
DROP TABLE nlx_management.permissions;
DROP TABLE nlx_management.users;
DROP TABLE nlx_management.roles;

COMMIT;
