#
# Copyright (C) 2026 Synkar
#
# This is free software, licensed under the GNU General Public License v2.
# See /LICENSE for more information.
#

include $(TOPDIR)/rules.mk

PKG_NAME:=gst1-plugins-rockchip
PKG_VERSION:=1.22.6
PKG_RELEASE:=1

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://github.com/Meonardo/gst-rockchip.git
PKG_SOURCE_VERSION:=fa0d862a2a5135f0b5e19cdb19aac2355bad2dd3
PKG_SOURCE_DATE:=2024-05-28
PKG_SOURCE_SUBDIR:=$(PKG_NAME)-$(PKG_VERSION)
PKG_SOURCE:=$(PKG_SOURCE_SUBDIR).tar.gz
PKG_MIRROR_HASH:=80582ac0fbde89a1c4221415ab2eeb9c196200897661404a13e6108414d7ee98

PKG_MAINTAINER:=Synkar <dev@synkar.com>
PKG_LICENSE:=LGPL-2.1-or-later
PKG_LICENSE_FILES:=COPYING

PKG_BUILD_PARALLEL:=1
PKG_INSTALL:=1

include $(INCLUDE_DIR)/package.mk
include $(INCLUDE_DIR)/meson.mk

define Package/gst1-plugins-rockchip
  SECTION:=multimedia
  CATEGORY:=Multimedia
  TITLE:=GStreamer Rockchip MPP plugin (mpph264enc, mpph265enc, mppdec)
  URL:=https://github.com/Meonardo/gst-rockchip
  DEPENDS:=+libgstreamer1 +libgst1video +libgst1pbutils +libgst1allocators +librockchip-mpp +libdrm
endef

define Package/gst1-plugins-rockchip/description
  GStreamer plugin wrapping the Rockchip MPP library (librockchip_mpp).
  Provides hardware accelerated H.264/H.265/VP8 encode (mpph264enc,
  mpph265enc, mppvp8enc) and H.264/H.265/JPEG decode (mppdec, mppjpegdec)
  on the RK3588, backed by the out-of-tree rkvenc driver via /dev/mpp_service.
endef

MESON_ARGS += \
	-Drkximage=disabled \
	-Dkmssrc=disabled \
	-Drockchipmpp=enabled \
	-Drga=disabled

define Build/InstallDev
	true
endef

define Package/gst1-plugins-rockchip/install
	$(INSTALL_DIR) $(1)/usr/lib/gstreamer-1.0
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/gstreamer-1.0/libgstrockchipmpp.so $(1)/usr/lib/gstreamer-1.0/
endef

$(eval $(call BuildPackage,gst1-plugins-rockchip))
