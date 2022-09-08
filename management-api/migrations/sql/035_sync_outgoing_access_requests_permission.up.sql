BEGIN transaction;

INSERT INTO nlx_management.permissions (code) VALUES ('permissions.outgoing_access_requests.sync');
INSERT INTO nlx_management.permissions_roles (role_code, permission_code) VALUES ('admin', 'permissions.outgoing_access_requests.sync');

COMMIT;
