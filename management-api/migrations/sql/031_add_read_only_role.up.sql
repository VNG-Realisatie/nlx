BEGIN transaction;

INSERT INTO nlx_management.roles (code) VALUES ('readonly');

INSERT INTO nlx_management.permissions_roles
    (role_code, permission_code)
VALUES
    ('readonly', 'permissions.incoming_access_requests.read'),
    ('readonly', 'permissions.access_grants.read'),
    ('readonly', 'permissions.audit_logs.read'),
    ('readonly', 'permissions.finance_report.read'),
    ('readonly', 'permissions.inway.read'),
    ('readonly', 'permissions.inways.read'),
    ('readonly', 'permissions.outgoing_orders.read'),
    ('readonly', 'permissions.incoming_orders.read'),
    ('readonly', 'permissions.outways.read'),
    ('readonly', 'permissions.service.read'),
    ('readonly', 'permissions.services.read'),
    ('readonly', 'permissions.services_statistics.read'),
    ('readonly', 'permissions.organization_settings.read'),
    ('readonly', 'permissions.terms_of_service_status.read'),
    ('readonly', 'permissions.transaction_logs.read');

COMMIT;
