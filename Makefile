.DEFAULT_GOAL := all

OUTPUTDIR ?= .

.PHONY: all
all:
	@echo $(OUTPUTDIR)
	@mkdir -p ${OUTPUTDIR}
	${MAKE} -C example OUTPUTDIR=$(OUTPUTDIR)

