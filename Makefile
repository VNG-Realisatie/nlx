# Copyright © VNG Realisatie 2022
# Licensed under the EUPL

.PHONY: fix-copyright-headers
fix-copyright-headers:
	docker run -it -v ${PWD}:/src ghcr.io/google/addlicense -ignore "**/node_modules/**" -ignore "management-ui/src/api/**/*" -ignore "helm/**/*.yaml" -f license-header.tmpl .

.PHONY: fix-newlines
fix-newlines:
	./scripts/eol-at-eof-linter.sh -f

.PHONY: fix-proto-formatting
fix-proto-formatting:
	docker run --rm \
		-v ${PWD}/management-api/api:/src/management-api/api \
		-v ${PWD}/outway/api:/src/outway/api \
		-v ${PWD}/txlog-api/api:/src/txlog-api/api \
		-v ${PWD}/management-api/api:/src/management-api/api \
 		--workdir /src bufbuild/buf:1.8.0 format -w
