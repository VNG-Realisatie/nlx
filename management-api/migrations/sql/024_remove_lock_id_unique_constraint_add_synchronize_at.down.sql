begin transaction;

alter table nlx_management.access_requests_outgoing add unique (lock_id);
alter table nlx_management.access_requests_outgoing drop column synchronize_at;

commit;
