# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetRouterIsisResult',
    'AwaitableGetRouterIsisResult',
    'get_router_isis',
    'get_router_isis_output',
]

@pulumi.output_type
class GetRouterIsisResult:
    """
    A collection of values returned by GetRouterIsis.
    """
    def __init__(__self__, adjacency_check=None, adjacency_check6=None, adv_passive_only=None, adv_passive_only6=None, auth_keychain_l1=None, auth_keychain_l2=None, auth_mode_l1=None, auth_mode_l2=None, auth_password_l1=None, auth_password_l2=None, auth_sendonly_l1=None, auth_sendonly_l2=None, default_originate=None, default_originate6=None, dynamic_hostname=None, id=None, ignore_lsp_errors=None, is_type=None, isis_interfaces=None, isis_nets=None, lsp_gen_interval_l1=None, lsp_gen_interval_l2=None, lsp_refresh_interval=None, max_lsp_lifetime=None, metric_style=None, overload_bit=None, overload_bit_on_startup=None, overload_bit_suppress=None, redistribute6_l1=None, redistribute6_l1_list=None, redistribute6_l2=None, redistribute6_l2_list=None, redistribute6s=None, redistribute_l1=None, redistribute_l1_list=None, redistribute_l2=None, redistribute_l2_list=None, redistributes=None, spf_interval_exp_l1=None, spf_interval_exp_l2=None, summary_address6s=None, summary_addresses=None, vdomparam=None):
        if adjacency_check and not isinstance(adjacency_check, str):
            raise TypeError("Expected argument 'adjacency_check' to be a str")
        pulumi.set(__self__, "adjacency_check", adjacency_check)
        if adjacency_check6 and not isinstance(adjacency_check6, str):
            raise TypeError("Expected argument 'adjacency_check6' to be a str")
        pulumi.set(__self__, "adjacency_check6", adjacency_check6)
        if adv_passive_only and not isinstance(adv_passive_only, str):
            raise TypeError("Expected argument 'adv_passive_only' to be a str")
        pulumi.set(__self__, "adv_passive_only", adv_passive_only)
        if adv_passive_only6 and not isinstance(adv_passive_only6, str):
            raise TypeError("Expected argument 'adv_passive_only6' to be a str")
        pulumi.set(__self__, "adv_passive_only6", adv_passive_only6)
        if auth_keychain_l1 and not isinstance(auth_keychain_l1, str):
            raise TypeError("Expected argument 'auth_keychain_l1' to be a str")
        pulumi.set(__self__, "auth_keychain_l1", auth_keychain_l1)
        if auth_keychain_l2 and not isinstance(auth_keychain_l2, str):
            raise TypeError("Expected argument 'auth_keychain_l2' to be a str")
        pulumi.set(__self__, "auth_keychain_l2", auth_keychain_l2)
        if auth_mode_l1 and not isinstance(auth_mode_l1, str):
            raise TypeError("Expected argument 'auth_mode_l1' to be a str")
        pulumi.set(__self__, "auth_mode_l1", auth_mode_l1)
        if auth_mode_l2 and not isinstance(auth_mode_l2, str):
            raise TypeError("Expected argument 'auth_mode_l2' to be a str")
        pulumi.set(__self__, "auth_mode_l2", auth_mode_l2)
        if auth_password_l1 and not isinstance(auth_password_l1, str):
            raise TypeError("Expected argument 'auth_password_l1' to be a str")
        pulumi.set(__self__, "auth_password_l1", auth_password_l1)
        if auth_password_l2 and not isinstance(auth_password_l2, str):
            raise TypeError("Expected argument 'auth_password_l2' to be a str")
        pulumi.set(__self__, "auth_password_l2", auth_password_l2)
        if auth_sendonly_l1 and not isinstance(auth_sendonly_l1, str):
            raise TypeError("Expected argument 'auth_sendonly_l1' to be a str")
        pulumi.set(__self__, "auth_sendonly_l1", auth_sendonly_l1)
        if auth_sendonly_l2 and not isinstance(auth_sendonly_l2, str):
            raise TypeError("Expected argument 'auth_sendonly_l2' to be a str")
        pulumi.set(__self__, "auth_sendonly_l2", auth_sendonly_l2)
        if default_originate and not isinstance(default_originate, str):
            raise TypeError("Expected argument 'default_originate' to be a str")
        pulumi.set(__self__, "default_originate", default_originate)
        if default_originate6 and not isinstance(default_originate6, str):
            raise TypeError("Expected argument 'default_originate6' to be a str")
        pulumi.set(__self__, "default_originate6", default_originate6)
        if dynamic_hostname and not isinstance(dynamic_hostname, str):
            raise TypeError("Expected argument 'dynamic_hostname' to be a str")
        pulumi.set(__self__, "dynamic_hostname", dynamic_hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_lsp_errors and not isinstance(ignore_lsp_errors, str):
            raise TypeError("Expected argument 'ignore_lsp_errors' to be a str")
        pulumi.set(__self__, "ignore_lsp_errors", ignore_lsp_errors)
        if is_type and not isinstance(is_type, str):
            raise TypeError("Expected argument 'is_type' to be a str")
        pulumi.set(__self__, "is_type", is_type)
        if isis_interfaces and not isinstance(isis_interfaces, list):
            raise TypeError("Expected argument 'isis_interfaces' to be a list")
        pulumi.set(__self__, "isis_interfaces", isis_interfaces)
        if isis_nets and not isinstance(isis_nets, list):
            raise TypeError("Expected argument 'isis_nets' to be a list")
        pulumi.set(__self__, "isis_nets", isis_nets)
        if lsp_gen_interval_l1 and not isinstance(lsp_gen_interval_l1, int):
            raise TypeError("Expected argument 'lsp_gen_interval_l1' to be a int")
        pulumi.set(__self__, "lsp_gen_interval_l1", lsp_gen_interval_l1)
        if lsp_gen_interval_l2 and not isinstance(lsp_gen_interval_l2, int):
            raise TypeError("Expected argument 'lsp_gen_interval_l2' to be a int")
        pulumi.set(__self__, "lsp_gen_interval_l2", lsp_gen_interval_l2)
        if lsp_refresh_interval and not isinstance(lsp_refresh_interval, int):
            raise TypeError("Expected argument 'lsp_refresh_interval' to be a int")
        pulumi.set(__self__, "lsp_refresh_interval", lsp_refresh_interval)
        if max_lsp_lifetime and not isinstance(max_lsp_lifetime, int):
            raise TypeError("Expected argument 'max_lsp_lifetime' to be a int")
        pulumi.set(__self__, "max_lsp_lifetime", max_lsp_lifetime)
        if metric_style and not isinstance(metric_style, str):
            raise TypeError("Expected argument 'metric_style' to be a str")
        pulumi.set(__self__, "metric_style", metric_style)
        if overload_bit and not isinstance(overload_bit, str):
            raise TypeError("Expected argument 'overload_bit' to be a str")
        pulumi.set(__self__, "overload_bit", overload_bit)
        if overload_bit_on_startup and not isinstance(overload_bit_on_startup, int):
            raise TypeError("Expected argument 'overload_bit_on_startup' to be a int")
        pulumi.set(__self__, "overload_bit_on_startup", overload_bit_on_startup)
        if overload_bit_suppress and not isinstance(overload_bit_suppress, str):
            raise TypeError("Expected argument 'overload_bit_suppress' to be a str")
        pulumi.set(__self__, "overload_bit_suppress", overload_bit_suppress)
        if redistribute6_l1 and not isinstance(redistribute6_l1, str):
            raise TypeError("Expected argument 'redistribute6_l1' to be a str")
        pulumi.set(__self__, "redistribute6_l1", redistribute6_l1)
        if redistribute6_l1_list and not isinstance(redistribute6_l1_list, str):
            raise TypeError("Expected argument 'redistribute6_l1_list' to be a str")
        pulumi.set(__self__, "redistribute6_l1_list", redistribute6_l1_list)
        if redistribute6_l2 and not isinstance(redistribute6_l2, str):
            raise TypeError("Expected argument 'redistribute6_l2' to be a str")
        pulumi.set(__self__, "redistribute6_l2", redistribute6_l2)
        if redistribute6_l2_list and not isinstance(redistribute6_l2_list, str):
            raise TypeError("Expected argument 'redistribute6_l2_list' to be a str")
        pulumi.set(__self__, "redistribute6_l2_list", redistribute6_l2_list)
        if redistribute6s and not isinstance(redistribute6s, list):
            raise TypeError("Expected argument 'redistribute6s' to be a list")
        pulumi.set(__self__, "redistribute6s", redistribute6s)
        if redistribute_l1 and not isinstance(redistribute_l1, str):
            raise TypeError("Expected argument 'redistribute_l1' to be a str")
        pulumi.set(__self__, "redistribute_l1", redistribute_l1)
        if redistribute_l1_list and not isinstance(redistribute_l1_list, str):
            raise TypeError("Expected argument 'redistribute_l1_list' to be a str")
        pulumi.set(__self__, "redistribute_l1_list", redistribute_l1_list)
        if redistribute_l2 and not isinstance(redistribute_l2, str):
            raise TypeError("Expected argument 'redistribute_l2' to be a str")
        pulumi.set(__self__, "redistribute_l2", redistribute_l2)
        if redistribute_l2_list and not isinstance(redistribute_l2_list, str):
            raise TypeError("Expected argument 'redistribute_l2_list' to be a str")
        pulumi.set(__self__, "redistribute_l2_list", redistribute_l2_list)
        if redistributes and not isinstance(redistributes, list):
            raise TypeError("Expected argument 'redistributes' to be a list")
        pulumi.set(__self__, "redistributes", redistributes)
        if spf_interval_exp_l1 and not isinstance(spf_interval_exp_l1, str):
            raise TypeError("Expected argument 'spf_interval_exp_l1' to be a str")
        pulumi.set(__self__, "spf_interval_exp_l1", spf_interval_exp_l1)
        if spf_interval_exp_l2 and not isinstance(spf_interval_exp_l2, str):
            raise TypeError("Expected argument 'spf_interval_exp_l2' to be a str")
        pulumi.set(__self__, "spf_interval_exp_l2", spf_interval_exp_l2)
        if summary_address6s and not isinstance(summary_address6s, list):
            raise TypeError("Expected argument 'summary_address6s' to be a list")
        pulumi.set(__self__, "summary_address6s", summary_address6s)
        if summary_addresses and not isinstance(summary_addresses, list):
            raise TypeError("Expected argument 'summary_addresses' to be a list")
        pulumi.set(__self__, "summary_addresses", summary_addresses)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="adjacencyCheck")
    def adjacency_check(self) -> str:
        """
        Enable/disable adjacency check.
        """
        return pulumi.get(self, "adjacency_check")

    @property
    @pulumi.getter(name="adjacencyCheck6")
    def adjacency_check6(self) -> str:
        """
        Enable/disable IPv6 adjacency check.
        """
        return pulumi.get(self, "adjacency_check6")

    @property
    @pulumi.getter(name="advPassiveOnly")
    def adv_passive_only(self) -> str:
        """
        Enable/disable IS-IS advertisement of passive interfaces only.
        """
        return pulumi.get(self, "adv_passive_only")

    @property
    @pulumi.getter(name="advPassiveOnly6")
    def adv_passive_only6(self) -> str:
        """
        Enable/disable IPv6 IS-IS advertisement of passive interfaces only.
        """
        return pulumi.get(self, "adv_passive_only6")

    @property
    @pulumi.getter(name="authKeychainL1")
    def auth_keychain_l1(self) -> str:
        """
        Authentication key-chain for level 1 PDUs.
        """
        return pulumi.get(self, "auth_keychain_l1")

    @property
    @pulumi.getter(name="authKeychainL2")
    def auth_keychain_l2(self) -> str:
        """
        Authentication key-chain for level 2 PDUs.
        """
        return pulumi.get(self, "auth_keychain_l2")

    @property
    @pulumi.getter(name="authModeL1")
    def auth_mode_l1(self) -> str:
        """
        Level 1 authentication mode.
        """
        return pulumi.get(self, "auth_mode_l1")

    @property
    @pulumi.getter(name="authModeL2")
    def auth_mode_l2(self) -> str:
        """
        Level 2 authentication mode.
        """
        return pulumi.get(self, "auth_mode_l2")

    @property
    @pulumi.getter(name="authPasswordL1")
    def auth_password_l1(self) -> str:
        """
        Authentication password for level 1 PDUs.
        """
        return pulumi.get(self, "auth_password_l1")

    @property
    @pulumi.getter(name="authPasswordL2")
    def auth_password_l2(self) -> str:
        """
        Authentication password for level 2 PDUs.
        """
        return pulumi.get(self, "auth_password_l2")

    @property
    @pulumi.getter(name="authSendonlyL1")
    def auth_sendonly_l1(self) -> str:
        """
        Enable/disable level 1 authentication send-only.
        """
        return pulumi.get(self, "auth_sendonly_l1")

    @property
    @pulumi.getter(name="authSendonlyL2")
    def auth_sendonly_l2(self) -> str:
        """
        Enable/disable level 2 authentication send-only.
        """
        return pulumi.get(self, "auth_sendonly_l2")

    @property
    @pulumi.getter(name="defaultOriginate")
    def default_originate(self) -> str:
        """
        Enable/disable distribution of default route information.
        """
        return pulumi.get(self, "default_originate")

    @property
    @pulumi.getter(name="defaultOriginate6")
    def default_originate6(self) -> str:
        """
        Enable/disable distribution of default IPv6 route information.
        """
        return pulumi.get(self, "default_originate6")

    @property
    @pulumi.getter(name="dynamicHostname")
    def dynamic_hostname(self) -> str:
        """
        Enable/disable dynamic hostname.
        """
        return pulumi.get(self, "dynamic_hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreLspErrors")
    def ignore_lsp_errors(self) -> str:
        """
        Enable/disable ignoring of LSP errors with bad checksums.
        """
        return pulumi.get(self, "ignore_lsp_errors")

    @property
    @pulumi.getter(name="isType")
    def is_type(self) -> str:
        """
        IS type.
        """
        return pulumi.get(self, "is_type")

    @property
    @pulumi.getter(name="isisInterfaces")
    def isis_interfaces(self) -> Sequence['outputs.GetRouterIsisIsisInterfaceResult']:
        """
        IS-IS interface configuration. The structure of `isis_interface` block is documented below.
        """
        return pulumi.get(self, "isis_interfaces")

    @property
    @pulumi.getter(name="isisNets")
    def isis_nets(self) -> Sequence['outputs.GetRouterIsisIsisNetResult']:
        """
        IS-IS net configuration. The structure of `isis_net` block is documented below.
        """
        return pulumi.get(self, "isis_nets")

    @property
    @pulumi.getter(name="lspGenIntervalL1")
    def lsp_gen_interval_l1(self) -> int:
        """
        Minimum interval for level 1 LSP regenerating.
        """
        return pulumi.get(self, "lsp_gen_interval_l1")

    @property
    @pulumi.getter(name="lspGenIntervalL2")
    def lsp_gen_interval_l2(self) -> int:
        """
        Minimum interval for level 2 LSP regenerating.
        """
        return pulumi.get(self, "lsp_gen_interval_l2")

    @property
    @pulumi.getter(name="lspRefreshInterval")
    def lsp_refresh_interval(self) -> int:
        """
        LSP refresh time in seconds.
        """
        return pulumi.get(self, "lsp_refresh_interval")

    @property
    @pulumi.getter(name="maxLspLifetime")
    def max_lsp_lifetime(self) -> int:
        """
        Maximum LSP lifetime in seconds.
        """
        return pulumi.get(self, "max_lsp_lifetime")

    @property
    @pulumi.getter(name="metricStyle")
    def metric_style(self) -> str:
        """
        Use old-style (ISO 10589) or new-style packet formats
        """
        return pulumi.get(self, "metric_style")

    @property
    @pulumi.getter(name="overloadBit")
    def overload_bit(self) -> str:
        """
        Enable/disable signal other routers not to use us in SPF.
        """
        return pulumi.get(self, "overload_bit")

    @property
    @pulumi.getter(name="overloadBitOnStartup")
    def overload_bit_on_startup(self) -> int:
        """
        Overload-bit only temporarily after reboot.
        """
        return pulumi.get(self, "overload_bit_on_startup")

    @property
    @pulumi.getter(name="overloadBitSuppress")
    def overload_bit_suppress(self) -> str:
        """
        Suppress overload-bit for the specific prefixes.
        """
        return pulumi.get(self, "overload_bit_suppress")

    @property
    @pulumi.getter(name="redistribute6L1")
    def redistribute6_l1(self) -> str:
        """
        Enable/disable redistribution of level 1 IPv6 routes into level 2.
        """
        return pulumi.get(self, "redistribute6_l1")

    @property
    @pulumi.getter(name="redistribute6L1List")
    def redistribute6_l1_list(self) -> str:
        """
        Access-list for IPv6 route redistribution from l1 to l2.
        """
        return pulumi.get(self, "redistribute6_l1_list")

    @property
    @pulumi.getter(name="redistribute6L2")
    def redistribute6_l2(self) -> str:
        """
        Enable/disable redistribution of level 2 IPv6 routes into level 1.
        """
        return pulumi.get(self, "redistribute6_l2")

    @property
    @pulumi.getter(name="redistribute6L2List")
    def redistribute6_l2_list(self) -> str:
        """
        Access-list for IPv6 route redistribution from l2 to l1.
        """
        return pulumi.get(self, "redistribute6_l2_list")

    @property
    @pulumi.getter
    def redistribute6s(self) -> Sequence['outputs.GetRouterIsisRedistribute6Result']:
        """
        IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
        """
        return pulumi.get(self, "redistribute6s")

    @property
    @pulumi.getter(name="redistributeL1")
    def redistribute_l1(self) -> str:
        """
        Enable/disable redistribution of level 1 routes into level 2.
        """
        return pulumi.get(self, "redistribute_l1")

    @property
    @pulumi.getter(name="redistributeL1List")
    def redistribute_l1_list(self) -> str:
        """
        Access-list for route redistribution from l1 to l2.
        """
        return pulumi.get(self, "redistribute_l1_list")

    @property
    @pulumi.getter(name="redistributeL2")
    def redistribute_l2(self) -> str:
        """
        Enable/disable redistribution of level 2 routes into level 1.
        """
        return pulumi.get(self, "redistribute_l2")

    @property
    @pulumi.getter(name="redistributeL2List")
    def redistribute_l2_list(self) -> str:
        """
        Access-list for route redistribution from l2 to l1.
        """
        return pulumi.get(self, "redistribute_l2_list")

    @property
    @pulumi.getter
    def redistributes(self) -> Sequence['outputs.GetRouterIsisRedistributeResult']:
        """
        IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
        """
        return pulumi.get(self, "redistributes")

    @property
    @pulumi.getter(name="spfIntervalExpL1")
    def spf_interval_exp_l1(self) -> str:
        """
        Level 1 SPF calculation delay.
        """
        return pulumi.get(self, "spf_interval_exp_l1")

    @property
    @pulumi.getter(name="spfIntervalExpL2")
    def spf_interval_exp_l2(self) -> str:
        """
        Level 2 SPF calculation delay.
        """
        return pulumi.get(self, "spf_interval_exp_l2")

    @property
    @pulumi.getter(name="summaryAddress6s")
    def summary_address6s(self) -> Sequence['outputs.GetRouterIsisSummaryAddress6Result']:
        """
        IS-IS IPv6 summary address. The structure of `summary_address6` block is documented below.
        """
        return pulumi.get(self, "summary_address6s")

    @property
    @pulumi.getter(name="summaryAddresses")
    def summary_addresses(self) -> Sequence['outputs.GetRouterIsisSummaryAddressResult']:
        """
        IS-IS summary addresses. The structure of `summary_address` block is documented below.
        """
        return pulumi.get(self, "summary_addresses")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterIsisResult(GetRouterIsisResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterIsisResult(
            adjacency_check=self.adjacency_check,
            adjacency_check6=self.adjacency_check6,
            adv_passive_only=self.adv_passive_only,
            adv_passive_only6=self.adv_passive_only6,
            auth_keychain_l1=self.auth_keychain_l1,
            auth_keychain_l2=self.auth_keychain_l2,
            auth_mode_l1=self.auth_mode_l1,
            auth_mode_l2=self.auth_mode_l2,
            auth_password_l1=self.auth_password_l1,
            auth_password_l2=self.auth_password_l2,
            auth_sendonly_l1=self.auth_sendonly_l1,
            auth_sendonly_l2=self.auth_sendonly_l2,
            default_originate=self.default_originate,
            default_originate6=self.default_originate6,
            dynamic_hostname=self.dynamic_hostname,
            id=self.id,
            ignore_lsp_errors=self.ignore_lsp_errors,
            is_type=self.is_type,
            isis_interfaces=self.isis_interfaces,
            isis_nets=self.isis_nets,
            lsp_gen_interval_l1=self.lsp_gen_interval_l1,
            lsp_gen_interval_l2=self.lsp_gen_interval_l2,
            lsp_refresh_interval=self.lsp_refresh_interval,
            max_lsp_lifetime=self.max_lsp_lifetime,
            metric_style=self.metric_style,
            overload_bit=self.overload_bit,
            overload_bit_on_startup=self.overload_bit_on_startup,
            overload_bit_suppress=self.overload_bit_suppress,
            redistribute6_l1=self.redistribute6_l1,
            redistribute6_l1_list=self.redistribute6_l1_list,
            redistribute6_l2=self.redistribute6_l2,
            redistribute6_l2_list=self.redistribute6_l2_list,
            redistribute6s=self.redistribute6s,
            redistribute_l1=self.redistribute_l1,
            redistribute_l1_list=self.redistribute_l1_list,
            redistribute_l2=self.redistribute_l2,
            redistribute_l2_list=self.redistribute_l2_list,
            redistributes=self.redistributes,
            spf_interval_exp_l1=self.spf_interval_exp_l1,
            spf_interval_exp_l2=self.spf_interval_exp_l2,
            summary_address6s=self.summary_address6s,
            summary_addresses=self.summary_addresses,
            vdomparam=self.vdomparam)


def get_router_isis(vdomparam: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterIsisResult:
    """
    Use this data source to get information on fortios router isis


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterIsis:GetRouterIsis', __args__, opts=opts, typ=GetRouterIsisResult).value

    return AwaitableGetRouterIsisResult(
        adjacency_check=__ret__.adjacency_check,
        adjacency_check6=__ret__.adjacency_check6,
        adv_passive_only=__ret__.adv_passive_only,
        adv_passive_only6=__ret__.adv_passive_only6,
        auth_keychain_l1=__ret__.auth_keychain_l1,
        auth_keychain_l2=__ret__.auth_keychain_l2,
        auth_mode_l1=__ret__.auth_mode_l1,
        auth_mode_l2=__ret__.auth_mode_l2,
        auth_password_l1=__ret__.auth_password_l1,
        auth_password_l2=__ret__.auth_password_l2,
        auth_sendonly_l1=__ret__.auth_sendonly_l1,
        auth_sendonly_l2=__ret__.auth_sendonly_l2,
        default_originate=__ret__.default_originate,
        default_originate6=__ret__.default_originate6,
        dynamic_hostname=__ret__.dynamic_hostname,
        id=__ret__.id,
        ignore_lsp_errors=__ret__.ignore_lsp_errors,
        is_type=__ret__.is_type,
        isis_interfaces=__ret__.isis_interfaces,
        isis_nets=__ret__.isis_nets,
        lsp_gen_interval_l1=__ret__.lsp_gen_interval_l1,
        lsp_gen_interval_l2=__ret__.lsp_gen_interval_l2,
        lsp_refresh_interval=__ret__.lsp_refresh_interval,
        max_lsp_lifetime=__ret__.max_lsp_lifetime,
        metric_style=__ret__.metric_style,
        overload_bit=__ret__.overload_bit,
        overload_bit_on_startup=__ret__.overload_bit_on_startup,
        overload_bit_suppress=__ret__.overload_bit_suppress,
        redistribute6_l1=__ret__.redistribute6_l1,
        redistribute6_l1_list=__ret__.redistribute6_l1_list,
        redistribute6_l2=__ret__.redistribute6_l2,
        redistribute6_l2_list=__ret__.redistribute6_l2_list,
        redistribute6s=__ret__.redistribute6s,
        redistribute_l1=__ret__.redistribute_l1,
        redistribute_l1_list=__ret__.redistribute_l1_list,
        redistribute_l2=__ret__.redistribute_l2,
        redistribute_l2_list=__ret__.redistribute_l2_list,
        redistributes=__ret__.redistributes,
        spf_interval_exp_l1=__ret__.spf_interval_exp_l1,
        spf_interval_exp_l2=__ret__.spf_interval_exp_l2,
        summary_address6s=__ret__.summary_address6s,
        summary_addresses=__ret__.summary_addresses,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_isis)
def get_router_isis_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterIsisResult]:
    """
    Use this data source to get information on fortios router isis


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
