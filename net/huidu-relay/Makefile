include $(TOPDIR)/rules.mk

PKG_NAME:=huidu-relay
PKG_VERSION:=1.0.0
PKG_RELEASE:=1

# Fonte: relay/ em https://github.com/Synkar/synkar_huidu_led_ros (privado),
# commit 0ea7170489d6b5acb910aa56def7581e64b297a3. Vendored em src/ para
# build hermetico (o repo de origem exige autenticacao).

PKG_LICENSE:=MIT
PKG_MAINTAINER:=Synkar <dev@synkar.com>

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_BUILD_FLAGS:=no-mips16

GO_PKG:=huidu_relay
PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)-$(PKG_VERSION)
LOCAL_SRC_DIR:=$(CURDIR)/src

include $(INCLUDE_DIR)/package.mk
include ../../lang/golang/golang-package.mk

define Package/huidu-relay
  SECTION:=synkar
  CATEGORY:=Synkar
  TITLE:=Relay VPN para controladoras Huidu LED
  URL:=https://github.com/Synkar/synkar_huidu_led_ros
  DEPENDS:=$(GO_ARCH_DEPENDS)
endef

define Package/huidu-relay/description
  Relay UDP/TCP que permite a clientes na VPN descobrir e controlar uma
  placa Huidu de LED conectada na LAN local do dispositivo: reescreve o
  IP da placa nas respostas de discovery e faz proxy TCP do canal de
  controle com politica de um PC por vez.
endef

define Build/Prepare
	rm -rf $(PKG_BUILD_DIR)
	mkdir -p $(PKG_BUILD_DIR)
	$(CP) $(LOCAL_SRC_DIR)/. $(PKG_BUILD_DIR)/
endef

define Package/huidu-relay/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(GO_PKG_BUILD_BIN_DIR)/huidu_relay $(1)/usr/bin/huidu-relay

	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/huidu-relay.init $(1)/etc/init.d/huidu-relay

	$(INSTALL_DIR) $(1)/etc/uci-defaults
	( \
		echo '#!/bin/sh'; \
		echo '/etc/init.d/huidu-relay enable'; \
		echo '/etc/init.d/huidu-relay start'; \
		echo 'exit 0'; \
	) > $(1)/etc/uci-defaults/99-huidu-relay-enable.sh
	chmod +x $(1)/etc/uci-defaults/99-huidu-relay-enable.sh
endef

$(eval $(call BuildPackage,huidu-relay))
