# go-get-data-model
# go-get-data-model_git.bb (recipes)
LICENSE = "CLOSED"
LIC_FILES_CHKSUM = ""
SRC_URI = "git://github.com/rohini970/go-get-data-model;branch=main"
SRCREV = "8c06f477a4c14ea4ff749e8310f1e6e3fa817d69"
GO_IMPORT = "go-get-data-model"

GO_IMPORT = "go-get-data-model"
GO_INSTALL = "${GO_IMPORT}"
GO_WORKDIR = "${GO_IMPORT}"
export GO111MODULE="off"

# -- DEPENDS BEGIN

RDEPENDS:${PN} += " \
bash \
"
RDEPENDS:${PN}-dev += " \
bash \
"
inherit go

do_compile:append() {
    bbnote "Compiling get-data-model"
    cd ${GOPATH}/src/go-get-data-model/
    go build -o go-get-data-model main.go
    bbnote "Compile done"
}

do_install:append() {
    bbnote "Installing get-data-model"
    install -d ${D}${bindir}
    install -m 0755 ${GOPATH}/src/go-get-data-model/go-get-data-model ${D}${bindir}/
    bbnote "Install done"
}

CGO_CFLAGS += "-I${WORKDIR}/recipe-sysroot/usr/include"
