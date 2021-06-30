/*
 * Copyright © VNG Realisatie 2021
 * Licensed under the EUPL
 */

begin transaction;

alter table nlx_management.users add column password varchar(60) null default null;

commit;
