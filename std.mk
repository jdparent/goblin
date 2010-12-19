# std.mk - included by most Makefiles in subdirs

include ../config.mk

OFILES		?= ${TARG}.$O
MANFILE		?= ${TARG}.1

all: ${TARG}
	@echo built ${TARG}

install: install-default post-install

install-default: ${TARG}
	@mkdir -p ${DESTDIR}${PREFIX}/bin
	@cp -f ${TARG} ${DESTDIR}${PREFIX}/bin/
	@chmod 755 ${DESTDIR}${PREFIX}/bin/${TARG}
	@mkdir -p ${DESTDIR}${MANPREFIX}/man1
	@cp -f ${MANFILE} ${DESTDIR}${MANPREFIX}/man1/
	@chmod 444 ${DESTDIR}${MANPREFIX}/man1/${MANFILE}

uninstall: pre-uninstall
	rm -f ${DESTDIR}${PREFIX}/bin/${TARG}
	rm -f ${DESTDIR}${MANPREFIX}/man1/${MANFILE}

%.$O:%.go
	$(GC) -o $@ $^

clean:
	rm -f ${OFILES} ${TARG}

${TARG}: ${OFILES}
	$(LD) -o $@ $^

