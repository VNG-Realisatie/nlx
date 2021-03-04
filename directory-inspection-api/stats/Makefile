.PHONY: all
all: generate

protos = $(wildcard *.proto)
results = $(protos:.proto=.pb.go) $(protos:.proto=.pb.gw.go) $(protos:.proto=.swagger.json) $(protos:.proto=.swagger.go)

%.pb.go %.pb.gw.go: %.proto
	buf generate

%.swagger.go: %.swagger.json
	echo 'package $*           '              >  $@
	echo 'const ('                            >> $@
	echo 'SwaggerJSONDirectoryInspection = `' >> $@
	cat $<                                    >> $@
	echo '`)'                                 >> $@

.PHONY: generate
generate: $(results)

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f *.pb.go
	rm -f *.pb.gw.go
	rm -f *.swagger.json
	rm -f *.swagger.go
