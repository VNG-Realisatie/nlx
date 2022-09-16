-- Copyright © VNG Realisatie 2022
-- Licensed under the EUPL

BEGIN transaction;

DROP TABLE nlx_management.permissions_roles;
DROP TABLE nlx_management.users_roles;
DROP TABLE nlx_management.permissions;
DROP TABLE nlx_management.roles;

CREATE TABLE nlx_management.permissions (
    code VARCHAR(250) PRIMARY KEY
);

CREATE TABLE nlx_management.roles (
  code VARCHAR(250) PRIMARY KEY
);

CREATE TABLE nlx_management.permissions_roles (
      role_code VARCHAR(250) NOT NULL,
      permission_code VARCHAR(250) NOT NULL,
      PRIMARY KEY (role_code, permission_code),
      created_at timestamp with time zone NOT NULL DEFAULT now(),
      updated_at timestamp with time zone NOT NULL DEFAULT now(),

      CONSTRAINT fk_role
          FOREIGN KEY (role_code)
              REFERENCES nlx_management.roles(code)
              ON DELETE RESTRICT,
      CONSTRAINT fk_permission_code
          FOREIGN KEY (permission_code)
              REFERENCES nlx_management.permissions(code)
              ON DELETE RESTRICT
);

CREATE TABLE nlx_management.users_roles (
    user_id SERIAL NOT NULL,
    role_code varchar(250) NOT NULL,
    PRIMARY KEY (user_id, role_code),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES nlx_management.users(id)
            ON DELETE RESTRICT,
    CONSTRAINT fk_role
        FOREIGN KEY (role_code)
            REFERENCES nlx_management.roles(code)
            ON DELETE RESTRICT
);

INSERT INTO nlx_management.permissions (code) VALUES
    ('permissions.incoming_access_request.approve'),
    ('permissions.incoming_access_request.reject'),
    ('permissions.incoming_access_requests.read'),
    ('permissions.outgoing_access_request.create'),
    ('permissions.outgoing_access_request.update'),
    ('permissions.access_grants.read'),
    ('permissions.access_grant.revoke'),
    ('permissions.audit_logs.read'),
    ('permissions.finance_report.read'),
    ('permissions.inway.read'),
    ('permissions.inway.update'),
    ('permissions.inway.delete'),
    ('permissions.inways.read'),
    ('permissions.outgoing_order.create'),
    ('permissions.outgoing_order.update'),
    ('permissions.outgoing_order.revoke'),
    ('permissions.outgoing_orders.read'),
    ('permissions.incoming_orders.read'),
    ('permissions.incoming_orders.synchronize'),
    ('permissions.outways.read'),
    ('permissions.service.create'),
    ('permissions.service.read'),
    ('permissions.service.update'),
    ('permissions.service.delete'),
    ('permissions.services.read'),
    ('permissions.services_statistics.read'),
    ('permissions.organization_settings.read'),
    ('permissions.organization_settings.update'),
    ('permissions.terms_of_service.accept'),
    ('permissions.terms_of_service_status.read'),
    ('permissions.transaction_logs.read');

INSERT INTO nlx_management.roles (code) VALUES ('admin');

INSERT INTO nlx_management.permissions_roles (role_code, permission_code) SELECT 'admin', code FROM nlx_management.permissions;

INSERT INTO nlx_management.users_roles (user_id, role_code) (SELECT id, 'admin' FROM nlx_management.users);

COMMIT;
