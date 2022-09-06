BEGIN transaction;

INSERT INTO nlx_management.permissions (code) VALUES ('permissions.outgoing_access_request.send');
INSERT INTO nlx_management.permissions_roles (role_code, permission_code) VALUES ('admin', 'permissions.outgoing_access_request.send');

COMMIT;
