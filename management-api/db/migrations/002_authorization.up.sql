BEGIN transaction; 

CREATE TABLE nlx_management.permissions (
  permission_id SERIAL PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  code VARCHAR(250) NOT NULL UNIQUE,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
);

CREATE TABLE nlx_management.users (
  user_id SERIAL PRIMARY KEY,
  email VARCHAR(250) NOT NULL UNIQUE,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
);

CREATE TABLE nlx_management.roles (
  role_id SERIAL PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  code VARCHAR(250) NOT NULL UNIQUE,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL
);

CREATE TABLE nlx_management.permissions_roles (
  role_id SERIAL NOT NULL,
  permission_id SERIAL NOT NULL,
  PRIMARY KEY (role_id, permission_id),
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,

  CONSTRAINT fk_role
    FOREIGN KEY (role_id)
    REFERENCES nlx_management.roles(role_id)
    ON DELETE RESTRICT,
  CONSTRAINT fk_permission
    FOREIGN KEY (permission_id)
    REFERENCES nlx_management.permissions(permission_id)
    ON DELETE RESTRICT
);

CREATE TABLE nlx_management.users_roles (
  user_id SERIAL NOT NULL,
  role_id SERIAL NOT NULL,
  PRIMARY KEY (user_id, role_id),
  created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES nlx_management.users(user_id)
    ON DELETE RESTRICT,
  CONSTRAINT fk_role
    FOREIGN KEY (role_id)
    REFERENCES nlx_management.roles(role_id)
    ON DELETE RESTRICT
);

INSERT INTO nlx_management.roles (name, code, created_at, updated_at) VALUES ('Administrator', 'admin', NOW(), NOW());

COMMIT;
