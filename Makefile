# goblin - Plan 9 command line tools proted to Go

include config.mk

SUBDIRS =	basename\
	cal\
	cat\
	echo\
	md5sum\
	mkdir\
	pbd\
	sleep

all:
	@for i in ${SUBDIRS}; do cd $$i; ${MAKE} || exit; cd ..; done;

clean:
	@for i in ${SUBDIRS}; do cd $$i; ${MAKE} clean || exit; cd ..; done
	@rm -f goblin-${VERSION}.tar.gz
	@echo cleaned goblin

install: all
	@for i in ${SUBDIRS}; do cd $$i; ${MAKE} install || exit; cd ..; done
	@echo installed goblin to ${DESTDIR}${PREFIX}

uninstall:
	@for i in ${SUBDIRS}; do cd $$i; ${MAKE} uninstall || exit; cd ..; done
	@echo uninstalled goblin

dist: clean
	@mkdir -p goblin-${VERSION}
	@cp -R Makefile README LICENSE std.mk config.mk ${SUBDIRS} goblin-${VERSION}
	@tar -cf goblin-${VERSION}.tar goblin-${VERSION}
	@gzip -9 goblin-${VERSION}.tar
	@rm -rf goblin-${VERSION}
	@echo created distribution goblin-${VERSION}.tar.gz
