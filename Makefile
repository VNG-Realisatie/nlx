.PHONY: fix-copyright-headers
fix-copyright-headers:
	docker run -it -v ${PWD}:/src ghcr.io/google/addlicense -ignore "**/node_modules/**" -ignore "management-ui/src/api/**/*" -f license-header.tmpl .

.PHONY: fix-newlines
fix-newlines:
	./scripts/eol-at-eof-linter.sh -f
