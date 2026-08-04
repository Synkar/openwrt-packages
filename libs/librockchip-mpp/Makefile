#
# Copyright (C) 2026 Synkar
#
# This is free software, licensed under the GNU General Public License v2.
# See /LICENSE for more information.
#

include $(TOPDIR)/rules.mk

PKG_NAME:=librockchip-mpp
PKG_VERSION:=0.1.0
PKG_RELEASE:=1

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://github.com/tsukumijima/mpp-rockchip.git
PKG_SOURCE_VERSION:=750e76ec2d9287babfaf08c8bf395ebc5e8778ea
PKG_SOURCE_DATE:=2026-01-21
PKG_SOURCE_SUBDIR:=$(PKG_NAME)-$(PKG_VERSION)
PKG_SOURCE:=$(PKG_SOURCE_SUBDIR).tar.gz
PKG_MIRROR_HASH:=3c0fa17b8cace50570f4998c3a22c501a51053f8df9e385dcabf604e05502982

PKG_MAINTAINER:=Synkar <dev@synkar.com>
PKG_LICENSE:=Apache-2.0
PKG_LICENSE_FILES:=LICENSES/Apache-2.0

PKG_BUILD_PARALLEL:=1
PKG_INSTALL:=1

include $(INCLUDE_DIR)/package.mk
include $(INCLUDE_DIR)/cmake.mk

define Package/librockchip-mpp
  SECTION:=libs
  CATEGORY:=Libraries
  TITLE:=Rockchip Media Process Platform (VEPU580 H.264/H.265 encoder)
  URL:=https://github.com/tsukumijima/mpp-rockchip
  DEPENDS:=@TARGET_rockchip +libpthread
endef

define Package/librockchip-mpp/description
  Rockchip userspace Media Process Platform library (librockchip_mpp.so).
  This is the tsukumijima/mpp-rockchip fork, pinned at commit 750e76ec
  (2026-01-21), intended for mainline kernels backed by the out-of-tree
  rkvenc driver, which exposes the vendor MPP ioctl ABI via /dev/mpp_service.
  Provides H.264/H.265/JPEG encode (VEPU580) and decode APIs for the RK3588.
  Buffers are allocated through the dma-buf heaps allocator (/dev/dma_heap),
  so the kernel must be built with CONFIG_DMABUF_HEAPS (system + cma).
endef

define Build/InstallDev
	$(INSTALL_DIR) $(1)/usr/include/rockchip
	$(CP) $(PKG_INSTALL_DIR)/usr/include/rockchip/*.h $(1)/usr/include/rockchip/
	$(INSTALL_DIR) $(1)/usr/lib
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/librockchip_mpp.so* $(1)/usr/lib/
	$(INSTALL_DIR) $(1)/usr/lib/pkgconfig
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/pkgconfig/rockchip_mpp.pc $(1)/usr/lib/pkgconfig/
endef

define Package/librockchip-mpp/install
	$(INSTALL_DIR) $(1)/usr/lib
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/librockchip_mpp.so.0 $(1)/usr/lib/
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/librockchip_mpp.so.1 $(1)/usr/lib/
	$(CP) $(PKG_INSTALL_DIR)/usr/lib/libmpp_ext.so $(1)/usr/lib/
endef

define Package/librockchip-mpp-tools
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=Rockchip MPP test tools (mpp_info, mpi_enc, mpi_dec)
  DEPENDS:=+librockchip-mpp
endef

define Package/librockchip-mpp-tools/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(PKG_INSTALL_DIR)/usr/bin/mpp_info_test $(1)/usr/bin/
	$(INSTALL_BIN) $(PKG_INSTALL_DIR)/usr/bin/mpi_enc_test $(1)/usr/bin/
	$(INSTALL_BIN) $(PKG_INSTALL_DIR)/usr/bin/mpi_enc_mt_test $(1)/usr/bin/
	$(INSTALL_BIN) $(PKG_INSTALL_DIR)/usr/bin/mpi_dec_test $(1)/usr/bin/
	$(INSTALL_BIN) $(PKG_INSTALL_DIR)/usr/bin/mpi_dec_mt_test $(1)/usr/bin/
endef

$(eval $(call BuildPackage,librockchip-mpp))
$(eval $(call BuildPackage,librockchip-mpp-tools))
