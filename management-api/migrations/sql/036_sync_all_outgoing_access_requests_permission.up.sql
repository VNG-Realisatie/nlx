-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

INSERT INTO nlx_management.permissions (code) VALUES ('permissions.outgoing_access_requests.sync_all');
INSERT INTO nlx_management.permissions_roles (role_code, permission_code)
VALUES
    ('admin', 'permissions.outgoing_access_requests.sync_all'),
    ('readonly', 'permissions.outgoing_access_requests.sync_all');

COMMIT;
