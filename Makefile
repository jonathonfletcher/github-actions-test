.DEFAULT_GOAL := deploy

.PHONY: deploy unittest
deploy unittest:
	${MAKE} -C example $@

