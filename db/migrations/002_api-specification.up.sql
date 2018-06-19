ALTER TABLE directory.services
	ADD COLUMN api_specification_type character varying(20);

ALTER TABLE directory.services ADD CONSTRAINT services_check_typespec CHECK (
    (
        (api_specification_type IS NULL) OR (
            (api_specification_type::text = 'OpenAPI2'::text)
            OR
            (api_specification_type::text = 'OpenAPI3'::text)
        )
    )
);
