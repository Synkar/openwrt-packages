#
# This is free software, licensed under the GNU General Public License v2.
#

include $(TOPDIR)/rules.mk

PKG_NAME:=awscrt
PKG_VERSION:=0.36.1
PKG_RELEASE:=1

PYPI_NAME:=$(PKG_NAME)
PKG_HASH:=bd1f86b092b57a9ec1f95138224007946eecdf944df1e08fd99f75b64fe2ad20

PKG_MAINTAINER:=Amazon Web Services, Inc <aws-sdk-common-runtime@amazon.com>
PKG_LICENSE:=MIT

include ../pypi.mk
include $(INCLUDE_DIR)/package.mk
include ../python3-package.mk

define Package/python-awscrt
  CATEGORY:=Languages
  SECTION:=lang
  SUBMENU:=Python
  TITLE:=python-awscrt
  URL:=https://github.com/awslabs/aws-crt-python
  DEPENDS:=+python3
endef

define Package/python-awscrt/description
A common runtime for AWS Python projects
endef

PYTHON3_PKG_BUILD_VARS:=AWS_CRT_BUILD_USE_SYSTEM_LIBCRYPTO=1

$(eval $(call Py3Package,python-awscrt))
$(eval $(call BuildPackage,python-awscrt))
$(eval $(call BuildPackage,python-awscrt-src))

