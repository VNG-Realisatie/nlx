-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

INSERT INTO nlx_management.permissions (code) VALUES ('permissions.outgoing_access_request.send');
INSERT INTO nlx_management.permissions_roles (role_code, permission_code) VALUES ('admin', 'permissions.outgoing_access_request.send');


DELETE FROM nlx_management.permissions_roles WHERE permission_code = 'permissions.outgoing_access_request.create';
DELETE FROM nlx_management.permissions WHERE code = 'permissions.outgoing_access_request.create';

DELETE FROM nlx_management.access_requests_outgoing WHERE state = 'created';

COMMIT;
