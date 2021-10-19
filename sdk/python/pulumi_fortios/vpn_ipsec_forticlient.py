# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpnIpsecForticlientArgs', 'VpnIpsecForticlient']

@pulumi.input_type
class VpnIpsecForticlientArgs:
    def __init__(__self__, *,
                 phase2name: pulumi.Input[str],
                 usergroupname: pulumi.Input[str],
                 realm: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpnIpsecForticlient resource.
        :param pulumi.Input[str] phase2name: Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
        :param pulumi.Input[str] usergroupname: User group name for FortiClient users.
        :param pulumi.Input[str] realm: FortiClient realm name.
        :param pulumi.Input[str] status: Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "phase2name", phase2name)
        pulumi.set(__self__, "usergroupname", usergroupname)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def phase2name(self) -> pulumi.Input[str]:
        """
        Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
        """
        return pulumi.get(self, "phase2name")

    @phase2name.setter
    def phase2name(self, value: pulumi.Input[str]):
        pulumi.set(self, "phase2name", value)

    @property
    @pulumi.getter
    def usergroupname(self) -> pulumi.Input[str]:
        """
        User group name for FortiClient users.
        """
        return pulumi.get(self, "usergroupname")

    @usergroupname.setter
    def usergroupname(self, value: pulumi.Input[str]):
        pulumi.set(self, "usergroupname", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient realm name.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _VpnIpsecForticlientState:
    def __init__(__self__, *,
                 phase2name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usergroupname: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpnIpsecForticlient resources.
        :param pulumi.Input[str] phase2name: Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
        :param pulumi.Input[str] realm: FortiClient realm name.
        :param pulumi.Input[str] status: Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usergroupname: User group name for FortiClient users.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if phase2name is not None:
            pulumi.set(__self__, "phase2name", phase2name)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if usergroupname is not None:
            pulumi.set(__self__, "usergroupname", usergroupname)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def phase2name(self) -> Optional[pulumi.Input[str]]:
        """
        Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
        """
        return pulumi.get(self, "phase2name")

    @phase2name.setter
    def phase2name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phase2name", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient realm name.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def usergroupname(self) -> Optional[pulumi.Input[str]]:
        """
        User group name for FortiClient users.
        """
        return pulumi.get(self, "usergroupname")

    @usergroupname.setter
    def usergroupname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usergroupname", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class VpnIpsecForticlient(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 phase2name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usergroupname: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiClient policy realm.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        # fortios_vpnipsec_phase1interface.trname2:
        trname4 = fortios.VpnIpsecPhase1Interface("trname4",
            acct_verify="disable",
            add_gw_route="disable",
            add_route="enable",
            assign_ip="enable",
            assign_ip_from="range",
            authmethod="psk",
            authusrgrp="Guest-group",
            auto_discovery_forwarder="disable",
            auto_discovery_psk="disable",
            auto_discovery_receiver="disable",
            auto_discovery_sender="disable",
            auto_negotiate="enable",
            cert_id_validation="enable",
            childless_ike="disable",
            client_auto_negotiate="disable",
            client_keep_alive="disable",
            comments="VPN: Dialup_IPsec (Created by VPN wizard)",
            default_gw="0.0.0.0",
            default_gw_priority=0,
            dhgrp="14 5",
            digital_signature_auth="disable",
            distance=15,
            dns_mode="auto",
            dpd="on-idle",
            dpd_retrycount=3,
            dpd_retryinterval="60",
            eap="disable",
            eap_identity="use-id-payload",
            encap_local_gw4="0.0.0.0",
            encap_local_gw6="::",
            encap_remote_gw4="0.0.0.0",
            encap_remote_gw6="::",
            encapsulation="none",
            encapsulation_address="ike",
            enforce_unique_id="disable",
            exchange_interface_ip="disable",
            exchange_ip_addr4="0.0.0.0",
            exchange_ip_addr6="::",
            forticlient_enforcement="disable",
            fragmentation="enable",
            fragmentation_mtu=1200,
            group_authentication="disable",
            ha_sync_esp_seqno="enable",
            idle_timeout="disable",
            idle_timeoutinterval=15,
            ike_version="1",
            include_local_lan="disable",
            interface="port4",
            ip_version="4",
            ipv4_dns_server1="0.0.0.0",
            ipv4_dns_server2="0.0.0.0",
            ipv4_dns_server3="0.0.0.0",
            ipv4_end_ip="10.10.10.10",
            ipv4_netmask="255.255.255.192",
            ipv4_split_include="FIREWALL_AUTH_PORTAL_ADDRESS",
            ipv4_start_ip="10.10.10.1",
            ipv4_wins_server1="0.0.0.0",
            ipv4_wins_server2="0.0.0.0",
            ipv6_dns_server1="::",
            ipv6_dns_server2="::",
            ipv6_dns_server3="::",
            ipv6_end_ip="::",
            ipv6_prefix=128,
            ipv6_start_ip="::",
            keepalive=10,
            keylife=86400,
            local_gw="0.0.0.0",
            local_gw6="::",
            localid_type="auto",
            mesh_selector_type="disable",
            mode="aggressive",
            mode_cfg="enable",
            monitor_hold_down_delay=0,
            monitor_hold_down_time="00:00",
            monitor_hold_down_type="immediate",
            monitor_hold_down_weekday="sunday",
            nattraversal="enable",
            negotiate_timeout=30,
            net_device="enable",
            passive_mode="disable",
            peertype="any",
            psksecret="NCIEW32930293203932",
            ppk="disable",
            priority=0,
            proposal="aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1",
            reauth="disable",
            rekey="enable",
            remote_gw="0.0.0.0",
            remote_gw6="::",
            rsa_signature_format="pkcs1",
            save_password="enable",
            send_cert_chain="enable",
            signature_hash_alg="sha2-512 sha2-384 sha2-256 sha1",
            suite_b="disable",
            tunnel_search="selectors",
            type="dynamic",
            unity_support="enable",
            wizard_type="dialup-forticlient",
            xauthtype="auto")
        # fortios_vpnipsec_phase2interface.trname1:
        trname3 = fortios.VpnIpsecPhase2Interface("trname3",
            add_route="phase1",
            auto_discovery_forwarder="phase1",
            auto_discovery_sender="phase1",
            auto_negotiate="disable",
            dhcp_ipsec="disable",
            dhgrp="14 5",
            dst_addr_type="subnet",
            dst_end_ip="0.0.0.0",
            dst_end_ip6="::",
            dst_port=0,
            dst_start_ip="0.0.0.0",
            dst_start_ip6="::",
            dst_subnet="0.0.0.0 0.0.0.0",
            dst_subnet6="::/0",
            encapsulation="tunnel-mode",
            keepalive="disable",
            keylife_type="seconds",
            keylifekbs=5120,
            keylifeseconds=43200,
            l2tp="disable",
            pfs="enable",
            phase1name=trname4.name,
            proposal="aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305",
            protocol=0,
            replay="enable",
            route_overlap="use-new",
            single_source="disable",
            src_addr_type="subnet",
            src_end_ip="0.0.0.0",
            src_end_ip6="::",
            src_port=0,
            src_start_ip="0.0.0.0",
            src_start_ip6="::",
            src_subnet="0.0.0.0 0.0.0.0",
            src_subnet6="::/0")
        trname = fortios.VpnIpsecForticlient("trname",
            phase2name=trname3.name,
            realm="1",
            status="enable",
            usergroupname="Guest-group")
        ```

        ## Import

        VpnIpsec Forticlient can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/vpnIpsecForticlient:VpnIpsecForticlient labelname {{realm}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] phase2name: Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
        :param pulumi.Input[str] realm: FortiClient realm name.
        :param pulumi.Input[str] status: Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usergroupname: User group name for FortiClient users.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpnIpsecForticlientArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiClient policy realm.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        # fortios_vpnipsec_phase1interface.trname2:
        trname4 = fortios.VpnIpsecPhase1Interface("trname4",
            acct_verify="disable",
            add_gw_route="disable",
            add_route="enable",
            assign_ip="enable",
            assign_ip_from="range",
            authmethod="psk",
            authusrgrp="Guest-group",
            auto_discovery_forwarder="disable",
            auto_discovery_psk="disable",
            auto_discovery_receiver="disable",
            auto_discovery_sender="disable",
            auto_negotiate="enable",
            cert_id_validation="enable",
            childless_ike="disable",
            client_auto_negotiate="disable",
            client_keep_alive="disable",
            comments="VPN: Dialup_IPsec (Created by VPN wizard)",
            default_gw="0.0.0.0",
            default_gw_priority=0,
            dhgrp="14 5",
            digital_signature_auth="disable",
            distance=15,
            dns_mode="auto",
            dpd="on-idle",
            dpd_retrycount=3,
            dpd_retryinterval="60",
            eap="disable",
            eap_identity="use-id-payload",
            encap_local_gw4="0.0.0.0",
            encap_local_gw6="::",
            encap_remote_gw4="0.0.0.0",
            encap_remote_gw6="::",
            encapsulation="none",
            encapsulation_address="ike",
            enforce_unique_id="disable",
            exchange_interface_ip="disable",
            exchange_ip_addr4="0.0.0.0",
            exchange_ip_addr6="::",
            forticlient_enforcement="disable",
            fragmentation="enable",
            fragmentation_mtu=1200,
            group_authentication="disable",
            ha_sync_esp_seqno="enable",
            idle_timeout="disable",
            idle_timeoutinterval=15,
            ike_version="1",
            include_local_lan="disable",
            interface="port4",
            ip_version="4",
            ipv4_dns_server1="0.0.0.0",
            ipv4_dns_server2="0.0.0.0",
            ipv4_dns_server3="0.0.0.0",
            ipv4_end_ip="10.10.10.10",
            ipv4_netmask="255.255.255.192",
            ipv4_split_include="FIREWALL_AUTH_PORTAL_ADDRESS",
            ipv4_start_ip="10.10.10.1",
            ipv4_wins_server1="0.0.0.0",
            ipv4_wins_server2="0.0.0.0",
            ipv6_dns_server1="::",
            ipv6_dns_server2="::",
            ipv6_dns_server3="::",
            ipv6_end_ip="::",
            ipv6_prefix=128,
            ipv6_start_ip="::",
            keepalive=10,
            keylife=86400,
            local_gw="0.0.0.0",
            local_gw6="::",
            localid_type="auto",
            mesh_selector_type="disable",
            mode="aggressive",
            mode_cfg="enable",
            monitor_hold_down_delay=0,
            monitor_hold_down_time="00:00",
            monitor_hold_down_type="immediate",
            monitor_hold_down_weekday="sunday",
            nattraversal="enable",
            negotiate_timeout=30,
            net_device="enable",
            passive_mode="disable",
            peertype="any",
            psksecret="NCIEW32930293203932",
            ppk="disable",
            priority=0,
            proposal="aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1",
            reauth="disable",
            rekey="enable",
            remote_gw="0.0.0.0",
            remote_gw6="::",
            rsa_signature_format="pkcs1",
            save_password="enable",
            send_cert_chain="enable",
            signature_hash_alg="sha2-512 sha2-384 sha2-256 sha1",
            suite_b="disable",
            tunnel_search="selectors",
            type="dynamic",
            unity_support="enable",
            wizard_type="dialup-forticlient",
            xauthtype="auto")
        # fortios_vpnipsec_phase2interface.trname1:
        trname3 = fortios.VpnIpsecPhase2Interface("trname3",
            add_route="phase1",
            auto_discovery_forwarder="phase1",
            auto_discovery_sender="phase1",
            auto_negotiate="disable",
            dhcp_ipsec="disable",
            dhgrp="14 5",
            dst_addr_type="subnet",
            dst_end_ip="0.0.0.0",
            dst_end_ip6="::",
            dst_port=0,
            dst_start_ip="0.0.0.0",
            dst_start_ip6="::",
            dst_subnet="0.0.0.0 0.0.0.0",
            dst_subnet6="::/0",
            encapsulation="tunnel-mode",
            keepalive="disable",
            keylife_type="seconds",
            keylifekbs=5120,
            keylifeseconds=43200,
            l2tp="disable",
            pfs="enable",
            phase1name=trname4.name,
            proposal="aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305",
            protocol=0,
            replay="enable",
            route_overlap="use-new",
            single_source="disable",
            src_addr_type="subnet",
            src_end_ip="0.0.0.0",
            src_end_ip6="::",
            src_port=0,
            src_start_ip="0.0.0.0",
            src_start_ip6="::",
            src_subnet="0.0.0.0 0.0.0.0",
            src_subnet6="::/0")
        trname = fortios.VpnIpsecForticlient("trname",
            phase2name=trname3.name,
            realm="1",
            status="enable",
            usergroupname="Guest-group")
        ```

        ## Import

        VpnIpsec Forticlient can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/vpnIpsecForticlient:VpnIpsecForticlient labelname {{realm}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VpnIpsecForticlientArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnIpsecForticlientArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 phase2name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usergroupname: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpnIpsecForticlientArgs.__new__(VpnIpsecForticlientArgs)

            if phase2name is None and not opts.urn:
                raise TypeError("Missing required property 'phase2name'")
            __props__.__dict__["phase2name"] = phase2name
            __props__.__dict__["realm"] = realm
            __props__.__dict__["status"] = status
            if usergroupname is None and not opts.urn:
                raise TypeError("Missing required property 'usergroupname'")
            __props__.__dict__["usergroupname"] = usergroupname
            __props__.__dict__["vdomparam"] = vdomparam
        super(VpnIpsecForticlient, __self__).__init__(
            'fortios:index/vpnIpsecForticlient:VpnIpsecForticlient',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            phase2name: Optional[pulumi.Input[str]] = None,
            realm: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            usergroupname: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'VpnIpsecForticlient':
        """
        Get an existing VpnIpsecForticlient resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] phase2name: Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
        :param pulumi.Input[str] realm: FortiClient realm name.
        :param pulumi.Input[str] status: Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usergroupname: User group name for FortiClient users.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpnIpsecForticlientState.__new__(_VpnIpsecForticlientState)

        __props__.__dict__["phase2name"] = phase2name
        __props__.__dict__["realm"] = realm
        __props__.__dict__["status"] = status
        __props__.__dict__["usergroupname"] = usergroupname
        __props__.__dict__["vdomparam"] = vdomparam
        return VpnIpsecForticlient(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def phase2name(self) -> pulumi.Output[str]:
        """
        Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
        """
        return pulumi.get(self, "phase2name")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[str]:
        """
        FortiClient realm name.
        """
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def usergroupname(self) -> pulumi.Output[str]:
        """
        User group name for FortiClient users.
        """
        return pulumi.get(self, "usergroupname")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

