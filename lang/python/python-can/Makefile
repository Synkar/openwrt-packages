#
# Copyright (C) 2026 Synkar
#
# This is free software, licensed under the GNU General Public License v2.
# See /LICENSE for more information.
#

include $(TOPDIR)/rules.mk

PKG_NAME:=python-can
PKG_VERSION:=4.6.1
PKG_RELEASE:=1

PYPI_NAME:=python-can
PYPI_SOURCE_NAME:=python_can
PKG_HASH:=290fea135d04b8504ebff33889cc6d301e2181a54099116609f940825ffe5005

PKG_MAINTAINER:=Synkar <dev@synkar.com>
PKG_LICENSE:=LGPL-3.0-only
PKG_LICENSE_FILES:=LICENSE

HOST_BUILD_DEPENDS:= \
	python-typing-extensions/host \
	python-packaging/host \
	python-setuptools-scm/host

PKG_BUILD_DEPENDS:=python-setuptools-scm/host

include ../pypi.mk
include $(INCLUDE_DIR)/package.mk
include ../python3-package.mk
include ../python3-host-build.mk

define Package/python3-can
  SUBMENU:=Python
  SECTION:=lang
  CATEGORY:=Languages
  TITLE:=CAN bus support for Python (SocketCAN)
  URL:=https://github.com/hardbyte/python-can
  DEPENDS:= \
	+python3-light \
	+python3-ctypes \
	+python3-logging \
	+python3-packaging \
	+python3-typing-extensions
endef

define Package/python3-can/description
  The CAN (Controller Area Network) support package for Python, providing
  common API for different CAN interfaces, including SocketCAN on Linux.
endef

$(eval $(call Py3Package,python3-can))
$(eval $(call BuildPackage,python3-can))
$(eval $(call BuildPackage,python3-can-src))
