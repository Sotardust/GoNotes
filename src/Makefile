
.PHONY : env
env:
	@$(CURDIR)/script/buf.sh env

gen: env
	cd proto && buf mod update
	buf generate --template ./buf.gen.yaml .