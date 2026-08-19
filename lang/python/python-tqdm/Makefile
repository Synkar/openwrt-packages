#
# This is free software, licensed under the GNU General Public License v2.
#

include $(TOPDIR)/rules.mk

PKG_NAME:=python-tqdm
PKG_VERSION:=4.70.0
PKG_RELEASE:=1

PYPI_NAME:=tqdm
PKG_HASH:=55b0b0dbd97462d06ebee91e4dac24ed4d4702be82b24f07e6c1d27e08cea220

PKG_MAINTAINER:=Casper da Costa-Luis <casper.dcl@physics.org>
PKG_LICENSE:=MIT MPL-2.0
PKG_LICENSE_FILES:=LICENCE

PKG_BUILD_DEPENDS:=python-setuptools-scm/host

include ../pypi.mk
include $(INCLUDE_DIR)/package.mk
include ../python3-package.mk

define Package/python3-tqdm
  SUBMENU:=Python
  SECTION:=lang
  CATEGORY:=Languages
  TITLE:=Fast, extensible progress bar
  URL:=https://tqdm.github.io/
  DEPENDS:=+python3-light
endef

define Package/python3-tqdm/description
  A fast, extensible progress bar for Python CLI applications.
endef

$(eval $(call Py3Package,python3-tqdm))
$(eval $(call BuildPackage,python3-tqdm))
$(eval $(call BuildPackage,python3-tqdm-src))
