begin transaction;

alter table directory.inways add column management_api_proxy_address character varying(255);

commit;
