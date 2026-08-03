#
# Copyright (C) 2026 Synkar
#
# This is free software, licensed under the GNU General Public License v2.
# See /LICENSE for more information.
#

include $(TOPDIR)/rules.mk

PKG_NAME:=libnice
PKG_VERSION:=0.1.22
PKG_RELEASE:=1

PKG_SOURCE:=$(PKG_NAME)-$(PKG_VERSION).tar.gz
PKG_SOURCE_URL:=https://libnice.freedesktop.org/releases
PKG_HASH:=a5f724cf09eae50c41a7517141d89da4a61ec9eaca32da4a0073faed5417ad7e

PKG_MAINTAINER:=Synkar <dev@synkar.com>
PKG_LICENSE:=MPL-2.0 OR LGPL-2.1-or-later
PKG_LICENSE_FILES:=COPYING

PKG_BUILD_PARALLEL:=1
PKG_INSTALL:=1

include $(INCLUDE_DIR)/package.mk
include $(INCLUDE_DIR)/meson.mk

define Package/libnice
  SECTION:=libs
  CATEGORY:=Libraries
  TITLE:=GLib ICE (Interactive Connectivity Establishment) library
  URL:=https://libnice.freedesktop.org/
  DEPENDS:=+glib2 +libgnutls +libtasn1
  ABI_VERSION:=10
endef

define Package/libnice/description
 libnice implements the IETF Interactive Connectivity Establishment (ICE)
 standard (RFC 8445) and related NAT traversal techniques. It is the ICE
 backend required by GStreamer's webrtcbin element.
endef

MESON_ARGS += \
	-Dcrypto-library=gnutls \
	-Dgupnp=disabled \
	-Dgstreamer=disabled \
	-Dintrospection=disabled \
	-Dtests=disabled \
	-Dexamples=disabled \
	-Dgtk_doc=disabled

define Build/InstallDev
	$(INSTALL_DIR) $(1)/usr/include/nice
	$(CP) $(PKG_INSTALL_DIR)/usr/include/nice/*.h $(1)/usr/include/nice/
	$(INSTALL_DIR) $(1)/usr/lib
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/libnice.so* $(1)/usr/lib/
	$(INSTALL_DIR) $(1)/usr/lib/pkgconfig
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/pkgconfig/nice.pc $(1)/usr/lib/pkgconfig/
endef

define Package/libnice/install
	$(INSTALL_DIR) $(1)/usr/lib
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/libnice.so.$(ABI_VERSION)* $(1)/usr/lib/
endef

$(eval $(call BuildPackage,libnice))
