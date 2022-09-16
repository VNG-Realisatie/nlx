# Copyright © VNG Realisatie 2022
# Licensed under the EUPL

.PHONY: fix-copyright-headers
fix-copyright-headers:
	docker run -it -v ${PWD}:/src ghcr.io/google/addlicense -ignore "**/node_modules/**" -ignore "management-ui/src/api/**/*" -ignore "helm/**/*.yaml" -f license-header.tmpl .

.PHONY: fix-newlines
fix-newlines:
	./scripts/eol-at-eof-linter.sh -f
