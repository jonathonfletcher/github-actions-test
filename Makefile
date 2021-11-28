.DEFAULT_GOAL := deploy

.PHONY: deploy
deploy:
	${MAKE} -C example

