/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */

begin transaction;

alter table nlx_management.users drop column password;

commit;
