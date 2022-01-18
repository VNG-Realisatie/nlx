begin transaction;

alter table nlx_management.access_requests_outgoing drop constraint access_requests_outgoing_lock_id_key;
alter table nlx_management.access_requests_outgoing add column synchronize_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP;

commit;
