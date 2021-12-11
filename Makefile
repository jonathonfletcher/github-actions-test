.DEFAULT_GOAL := deploy

.PHONY: deploy unittest
deploy:
	${MAKE} -C example $@

unittest:
	${MAKE} -C example $@
