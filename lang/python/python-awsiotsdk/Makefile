#
# This is free software, licensed under the GNU General Public License v2.
#

include $(TOPDIR)/rules.mk

PKG_NAME:=awsiotsdk
PKG_VERSION:=1.31.0
PKG_RELEASE:=1

PYPI_NAME:=$(PKG_NAME)
PKG_HASH:=652c7c357f09ea983330a22b5b1b560defc23e738f85d1db7d2ca9517c033fde

PKG_MAINTAINER:=AWS SDK Common Runtime Team
PKG_LICENSE:=MIT

include ../pypi.mk
include $(INCLUDE_DIR)/package.mk
include ../python3-package.mk

define Package/python-awsiotsdk
  CATEGORY:=Languages
  SECTION:=lang
  SUBMENU:=Python
  TITLE:=python-awsiotsdk
  URL:=https://github.com/aws/aws-iot-device-sdk-python-v2
  DEPENDS:=+python3 \
  +python-awscrt
endef

define Package/python-awsiotsdk/description
AWS IoT SDK based on the AWS Common Runtime
endef

$(eval $(call Py3Package,python-awsiotsdk))
$(eval $(call BuildPackage,python-awsiotsdk))
$(eval $(call BuildPackage,python-awsiotsdk-src))

