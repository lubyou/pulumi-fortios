# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebfilterFortiguardArgs', 'WebfilterFortiguard']

@pulumi.input_type
class WebfilterFortiguardArgs:
    def __init__(__self__, *,
                 cache_mem_percent: Optional[pulumi.Input[int]] = None,
                 cache_mode: Optional[pulumi.Input[str]] = None,
                 cache_prefix_match: Optional[pulumi.Input[str]] = None,
                 close_ports: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_https: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_port: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_http: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https_flow: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_warning: Optional[pulumi.Input[int]] = None,
                 request_packet_size_limit: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_auth_https: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebfilterFortiguard resource.
        :param pulumi.Input[int] cache_mem_percent: Maximum percentage of available memory allocated to caching (1 - 15%).
        :param pulumi.Input[str] cache_mode: Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        :param pulumi.Input[str] cache_prefix_match: Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] close_ports: Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ovrd_auth_https: Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] ovrd_auth_port: Port to use for FortiGuard Web Filter override authentication.
        :param pulumi.Input[int] ovrd_auth_port_http: Port to use for FortiGuard Web Filter HTTP override authentication
        :param pulumi.Input[int] ovrd_auth_port_https: Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        :param pulumi.Input[int] ovrd_auth_port_https_flow: Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        :param pulumi.Input[int] ovrd_auth_port_warning: Port to use for FortiGuard Web Filter Warning override authentication.
        :param pulumi.Input[int] request_packet_size_limit: Limit size of URL request packets sent to FortiGuard server (0 for default).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] warn_auth_https: Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        """
        if cache_mem_percent is not None:
            pulumi.set(__self__, "cache_mem_percent", cache_mem_percent)
        if cache_mode is not None:
            pulumi.set(__self__, "cache_mode", cache_mode)
        if cache_prefix_match is not None:
            pulumi.set(__self__, "cache_prefix_match", cache_prefix_match)
        if close_ports is not None:
            pulumi.set(__self__, "close_ports", close_ports)
        if ovrd_auth_https is not None:
            pulumi.set(__self__, "ovrd_auth_https", ovrd_auth_https)
        if ovrd_auth_port is not None:
            pulumi.set(__self__, "ovrd_auth_port", ovrd_auth_port)
        if ovrd_auth_port_http is not None:
            pulumi.set(__self__, "ovrd_auth_port_http", ovrd_auth_port_http)
        if ovrd_auth_port_https is not None:
            pulumi.set(__self__, "ovrd_auth_port_https", ovrd_auth_port_https)
        if ovrd_auth_port_https_flow is not None:
            pulumi.set(__self__, "ovrd_auth_port_https_flow", ovrd_auth_port_https_flow)
        if ovrd_auth_port_warning is not None:
            pulumi.set(__self__, "ovrd_auth_port_warning", ovrd_auth_port_warning)
        if request_packet_size_limit is not None:
            pulumi.set(__self__, "request_packet_size_limit", request_packet_size_limit)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if warn_auth_https is not None:
            pulumi.set(__self__, "warn_auth_https", warn_auth_https)

    @property
    @pulumi.getter(name="cacheMemPercent")
    def cache_mem_percent(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum percentage of available memory allocated to caching (1 - 15%).
        """
        return pulumi.get(self, "cache_mem_percent")

    @cache_mem_percent.setter
    def cache_mem_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cache_mem_percent", value)

    @property
    @pulumi.getter(name="cacheMode")
    def cache_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        """
        return pulumi.get(self, "cache_mode")

    @cache_mode.setter
    def cache_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_mode", value)

    @property
    @pulumi.getter(name="cachePrefixMatch")
    def cache_prefix_match(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_prefix_match")

    @cache_prefix_match.setter
    def cache_prefix_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_prefix_match", value)

    @property
    @pulumi.getter(name="closePorts")
    def close_ports(self) -> Optional[pulumi.Input[str]]:
        """
        Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "close_ports")

    @close_ports.setter
    def close_ports(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "close_ports", value)

    @property
    @pulumi.getter(name="ovrdAuthHttps")
    def ovrd_auth_https(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ovrd_auth_https")

    @ovrd_auth_https.setter
    def ovrd_auth_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ovrd_auth_https", value)

    @property
    @pulumi.getter(name="ovrdAuthPort")
    def ovrd_auth_port(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter override authentication.
        """
        return pulumi.get(self, "ovrd_auth_port")

    @ovrd_auth_port.setter
    def ovrd_auth_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port", value)

    @property
    @pulumi.getter(name="ovrdAuthPortHttp")
    def ovrd_auth_port_http(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter HTTP override authentication
        """
        return pulumi.get(self, "ovrd_auth_port_http")

    @ovrd_auth_port_http.setter
    def ovrd_auth_port_http(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_http", value)

    @property
    @pulumi.getter(name="ovrdAuthPortHttps")
    def ovrd_auth_port_https(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        """
        return pulumi.get(self, "ovrd_auth_port_https")

    @ovrd_auth_port_https.setter
    def ovrd_auth_port_https(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_https", value)

    @property
    @pulumi.getter(name="ovrdAuthPortHttpsFlow")
    def ovrd_auth_port_https_flow(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        """
        return pulumi.get(self, "ovrd_auth_port_https_flow")

    @ovrd_auth_port_https_flow.setter
    def ovrd_auth_port_https_flow(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_https_flow", value)

    @property
    @pulumi.getter(name="ovrdAuthPortWarning")
    def ovrd_auth_port_warning(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter Warning override authentication.
        """
        return pulumi.get(self, "ovrd_auth_port_warning")

    @ovrd_auth_port_warning.setter
    def ovrd_auth_port_warning(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_warning", value)

    @property
    @pulumi.getter(name="requestPacketSizeLimit")
    def request_packet_size_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Limit size of URL request packets sent to FortiGuard server (0 for default).
        """
        return pulumi.get(self, "request_packet_size_limit")

    @request_packet_size_limit.setter
    def request_packet_size_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_packet_size_limit", value)

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

    @property
    @pulumi.getter(name="warnAuthHttps")
    def warn_auth_https(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "warn_auth_https")

    @warn_auth_https.setter
    def warn_auth_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "warn_auth_https", value)


@pulumi.input_type
class _WebfilterFortiguardState:
    def __init__(__self__, *,
                 cache_mem_percent: Optional[pulumi.Input[int]] = None,
                 cache_mode: Optional[pulumi.Input[str]] = None,
                 cache_prefix_match: Optional[pulumi.Input[str]] = None,
                 close_ports: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_https: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_port: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_http: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https_flow: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_warning: Optional[pulumi.Input[int]] = None,
                 request_packet_size_limit: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_auth_https: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebfilterFortiguard resources.
        :param pulumi.Input[int] cache_mem_percent: Maximum percentage of available memory allocated to caching (1 - 15%).
        :param pulumi.Input[str] cache_mode: Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        :param pulumi.Input[str] cache_prefix_match: Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] close_ports: Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ovrd_auth_https: Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] ovrd_auth_port: Port to use for FortiGuard Web Filter override authentication.
        :param pulumi.Input[int] ovrd_auth_port_http: Port to use for FortiGuard Web Filter HTTP override authentication
        :param pulumi.Input[int] ovrd_auth_port_https: Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        :param pulumi.Input[int] ovrd_auth_port_https_flow: Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        :param pulumi.Input[int] ovrd_auth_port_warning: Port to use for FortiGuard Web Filter Warning override authentication.
        :param pulumi.Input[int] request_packet_size_limit: Limit size of URL request packets sent to FortiGuard server (0 for default).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] warn_auth_https: Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        """
        if cache_mem_percent is not None:
            pulumi.set(__self__, "cache_mem_percent", cache_mem_percent)
        if cache_mode is not None:
            pulumi.set(__self__, "cache_mode", cache_mode)
        if cache_prefix_match is not None:
            pulumi.set(__self__, "cache_prefix_match", cache_prefix_match)
        if close_ports is not None:
            pulumi.set(__self__, "close_ports", close_ports)
        if ovrd_auth_https is not None:
            pulumi.set(__self__, "ovrd_auth_https", ovrd_auth_https)
        if ovrd_auth_port is not None:
            pulumi.set(__self__, "ovrd_auth_port", ovrd_auth_port)
        if ovrd_auth_port_http is not None:
            pulumi.set(__self__, "ovrd_auth_port_http", ovrd_auth_port_http)
        if ovrd_auth_port_https is not None:
            pulumi.set(__self__, "ovrd_auth_port_https", ovrd_auth_port_https)
        if ovrd_auth_port_https_flow is not None:
            pulumi.set(__self__, "ovrd_auth_port_https_flow", ovrd_auth_port_https_flow)
        if ovrd_auth_port_warning is not None:
            pulumi.set(__self__, "ovrd_auth_port_warning", ovrd_auth_port_warning)
        if request_packet_size_limit is not None:
            pulumi.set(__self__, "request_packet_size_limit", request_packet_size_limit)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if warn_auth_https is not None:
            pulumi.set(__self__, "warn_auth_https", warn_auth_https)

    @property
    @pulumi.getter(name="cacheMemPercent")
    def cache_mem_percent(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum percentage of available memory allocated to caching (1 - 15%).
        """
        return pulumi.get(self, "cache_mem_percent")

    @cache_mem_percent.setter
    def cache_mem_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cache_mem_percent", value)

    @property
    @pulumi.getter(name="cacheMode")
    def cache_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        """
        return pulumi.get(self, "cache_mode")

    @cache_mode.setter
    def cache_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_mode", value)

    @property
    @pulumi.getter(name="cachePrefixMatch")
    def cache_prefix_match(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_prefix_match")

    @cache_prefix_match.setter
    def cache_prefix_match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_prefix_match", value)

    @property
    @pulumi.getter(name="closePorts")
    def close_ports(self) -> Optional[pulumi.Input[str]]:
        """
        Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "close_ports")

    @close_ports.setter
    def close_ports(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "close_ports", value)

    @property
    @pulumi.getter(name="ovrdAuthHttps")
    def ovrd_auth_https(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ovrd_auth_https")

    @ovrd_auth_https.setter
    def ovrd_auth_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ovrd_auth_https", value)

    @property
    @pulumi.getter(name="ovrdAuthPort")
    def ovrd_auth_port(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter override authentication.
        """
        return pulumi.get(self, "ovrd_auth_port")

    @ovrd_auth_port.setter
    def ovrd_auth_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port", value)

    @property
    @pulumi.getter(name="ovrdAuthPortHttp")
    def ovrd_auth_port_http(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter HTTP override authentication
        """
        return pulumi.get(self, "ovrd_auth_port_http")

    @ovrd_auth_port_http.setter
    def ovrd_auth_port_http(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_http", value)

    @property
    @pulumi.getter(name="ovrdAuthPortHttps")
    def ovrd_auth_port_https(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        """
        return pulumi.get(self, "ovrd_auth_port_https")

    @ovrd_auth_port_https.setter
    def ovrd_auth_port_https(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_https", value)

    @property
    @pulumi.getter(name="ovrdAuthPortHttpsFlow")
    def ovrd_auth_port_https_flow(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        """
        return pulumi.get(self, "ovrd_auth_port_https_flow")

    @ovrd_auth_port_https_flow.setter
    def ovrd_auth_port_https_flow(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_https_flow", value)

    @property
    @pulumi.getter(name="ovrdAuthPortWarning")
    def ovrd_auth_port_warning(self) -> Optional[pulumi.Input[int]]:
        """
        Port to use for FortiGuard Web Filter Warning override authentication.
        """
        return pulumi.get(self, "ovrd_auth_port_warning")

    @ovrd_auth_port_warning.setter
    def ovrd_auth_port_warning(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ovrd_auth_port_warning", value)

    @property
    @pulumi.getter(name="requestPacketSizeLimit")
    def request_packet_size_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Limit size of URL request packets sent to FortiGuard server (0 for default).
        """
        return pulumi.get(self, "request_packet_size_limit")

    @request_packet_size_limit.setter
    def request_packet_size_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_packet_size_limit", value)

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

    @property
    @pulumi.getter(name="warnAuthHttps")
    def warn_auth_https(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "warn_auth_https")

    @warn_auth_https.setter
    def warn_auth_https(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "warn_auth_https", value)


class WebfilterFortiguard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_mem_percent: Optional[pulumi.Input[int]] = None,
                 cache_mode: Optional[pulumi.Input[str]] = None,
                 cache_prefix_match: Optional[pulumi.Input[str]] = None,
                 close_ports: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_https: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_port: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_http: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https_flow: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_warning: Optional[pulumi.Input[int]] = None,
                 request_packet_size_limit: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_auth_https: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiGuard Web Filter service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WebfilterFortiguard("trname",
            cache_mem_percent=2,
            cache_mode="ttl",
            cache_prefix_match="enable",
            close_ports="disable",
            ovrd_auth_https="enable",
            ovrd_auth_port=8008,
            ovrd_auth_port_http=8008,
            ovrd_auth_port_https=8010,
            ovrd_auth_port_https_flow=8015,
            ovrd_auth_port_warning=8020,
            warn_auth_https="enable")
        ```

        ## Import

        Webfilter Fortiguard can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webfilterFortiguard:WebfilterFortiguard labelname WebfilterFortiguard
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webfilterFortiguard:WebfilterFortiguard labelname WebfilterFortiguard
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cache_mem_percent: Maximum percentage of available memory allocated to caching (1 - 15%).
        :param pulumi.Input[str] cache_mode: Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        :param pulumi.Input[str] cache_prefix_match: Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] close_ports: Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ovrd_auth_https: Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] ovrd_auth_port: Port to use for FortiGuard Web Filter override authentication.
        :param pulumi.Input[int] ovrd_auth_port_http: Port to use for FortiGuard Web Filter HTTP override authentication
        :param pulumi.Input[int] ovrd_auth_port_https: Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        :param pulumi.Input[int] ovrd_auth_port_https_flow: Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        :param pulumi.Input[int] ovrd_auth_port_warning: Port to use for FortiGuard Web Filter Warning override authentication.
        :param pulumi.Input[int] request_packet_size_limit: Limit size of URL request packets sent to FortiGuard server (0 for default).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] warn_auth_https: Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WebfilterFortiguardArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiGuard Web Filter service.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WebfilterFortiguard("trname",
            cache_mem_percent=2,
            cache_mode="ttl",
            cache_prefix_match="enable",
            close_ports="disable",
            ovrd_auth_https="enable",
            ovrd_auth_port=8008,
            ovrd_auth_port_http=8008,
            ovrd_auth_port_https=8010,
            ovrd_auth_port_https_flow=8015,
            ovrd_auth_port_warning=8020,
            warn_auth_https="enable")
        ```

        ## Import

        Webfilter Fortiguard can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webfilterFortiguard:WebfilterFortiguard labelname WebfilterFortiguard
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webfilterFortiguard:WebfilterFortiguard labelname WebfilterFortiguard
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WebfilterFortiguardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebfilterFortiguardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_mem_percent: Optional[pulumi.Input[int]] = None,
                 cache_mode: Optional[pulumi.Input[str]] = None,
                 cache_prefix_match: Optional[pulumi.Input[str]] = None,
                 close_ports: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_https: Optional[pulumi.Input[str]] = None,
                 ovrd_auth_port: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_http: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_https_flow: Optional[pulumi.Input[int]] = None,
                 ovrd_auth_port_warning: Optional[pulumi.Input[int]] = None,
                 request_packet_size_limit: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 warn_auth_https: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebfilterFortiguardArgs.__new__(WebfilterFortiguardArgs)

            __props__.__dict__["cache_mem_percent"] = cache_mem_percent
            __props__.__dict__["cache_mode"] = cache_mode
            __props__.__dict__["cache_prefix_match"] = cache_prefix_match
            __props__.__dict__["close_ports"] = close_ports
            __props__.__dict__["ovrd_auth_https"] = ovrd_auth_https
            __props__.__dict__["ovrd_auth_port"] = ovrd_auth_port
            __props__.__dict__["ovrd_auth_port_http"] = ovrd_auth_port_http
            __props__.__dict__["ovrd_auth_port_https"] = ovrd_auth_port_https
            __props__.__dict__["ovrd_auth_port_https_flow"] = ovrd_auth_port_https_flow
            __props__.__dict__["ovrd_auth_port_warning"] = ovrd_auth_port_warning
            __props__.__dict__["request_packet_size_limit"] = request_packet_size_limit
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["warn_auth_https"] = warn_auth_https
        super(WebfilterFortiguard, __self__).__init__(
            'fortios:index/webfilterFortiguard:WebfilterFortiguard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cache_mem_percent: Optional[pulumi.Input[int]] = None,
            cache_mode: Optional[pulumi.Input[str]] = None,
            cache_prefix_match: Optional[pulumi.Input[str]] = None,
            close_ports: Optional[pulumi.Input[str]] = None,
            ovrd_auth_https: Optional[pulumi.Input[str]] = None,
            ovrd_auth_port: Optional[pulumi.Input[int]] = None,
            ovrd_auth_port_http: Optional[pulumi.Input[int]] = None,
            ovrd_auth_port_https: Optional[pulumi.Input[int]] = None,
            ovrd_auth_port_https_flow: Optional[pulumi.Input[int]] = None,
            ovrd_auth_port_warning: Optional[pulumi.Input[int]] = None,
            request_packet_size_limit: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            warn_auth_https: Optional[pulumi.Input[str]] = None) -> 'WebfilterFortiguard':
        """
        Get an existing WebfilterFortiguard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cache_mem_percent: Maximum percentage of available memory allocated to caching (1 - 15%).
        :param pulumi.Input[str] cache_mode: Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        :param pulumi.Input[str] cache_prefix_match: Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] close_ports: Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ovrd_auth_https: Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] ovrd_auth_port: Port to use for FortiGuard Web Filter override authentication.
        :param pulumi.Input[int] ovrd_auth_port_http: Port to use for FortiGuard Web Filter HTTP override authentication
        :param pulumi.Input[int] ovrd_auth_port_https: Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        :param pulumi.Input[int] ovrd_auth_port_https_flow: Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        :param pulumi.Input[int] ovrd_auth_port_warning: Port to use for FortiGuard Web Filter Warning override authentication.
        :param pulumi.Input[int] request_packet_size_limit: Limit size of URL request packets sent to FortiGuard server (0 for default).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] warn_auth_https: Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebfilterFortiguardState.__new__(_WebfilterFortiguardState)

        __props__.__dict__["cache_mem_percent"] = cache_mem_percent
        __props__.__dict__["cache_mode"] = cache_mode
        __props__.__dict__["cache_prefix_match"] = cache_prefix_match
        __props__.__dict__["close_ports"] = close_ports
        __props__.__dict__["ovrd_auth_https"] = ovrd_auth_https
        __props__.__dict__["ovrd_auth_port"] = ovrd_auth_port
        __props__.__dict__["ovrd_auth_port_http"] = ovrd_auth_port_http
        __props__.__dict__["ovrd_auth_port_https"] = ovrd_auth_port_https
        __props__.__dict__["ovrd_auth_port_https_flow"] = ovrd_auth_port_https_flow
        __props__.__dict__["ovrd_auth_port_warning"] = ovrd_auth_port_warning
        __props__.__dict__["request_packet_size_limit"] = request_packet_size_limit
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["warn_auth_https"] = warn_auth_https
        return WebfilterFortiguard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cacheMemPercent")
    def cache_mem_percent(self) -> pulumi.Output[int]:
        """
        Maximum percentage of available memory allocated to caching (1 - 15%).
        """
        return pulumi.get(self, "cache_mem_percent")

    @property
    @pulumi.getter(name="cacheMode")
    def cache_mode(self) -> pulumi.Output[str]:
        """
        Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
        """
        return pulumi.get(self, "cache_mode")

    @property
    @pulumi.getter(name="cachePrefixMatch")
    def cache_prefix_match(self) -> pulumi.Output[str]:
        """
        Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "cache_prefix_match")

    @property
    @pulumi.getter(name="closePorts")
    def close_ports(self) -> pulumi.Output[str]:
        """
        Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "close_ports")

    @property
    @pulumi.getter(name="ovrdAuthHttps")
    def ovrd_auth_https(self) -> pulumi.Output[str]:
        """
        Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ovrd_auth_https")

    @property
    @pulumi.getter(name="ovrdAuthPort")
    def ovrd_auth_port(self) -> pulumi.Output[int]:
        """
        Port to use for FortiGuard Web Filter override authentication.
        """
        return pulumi.get(self, "ovrd_auth_port")

    @property
    @pulumi.getter(name="ovrdAuthPortHttp")
    def ovrd_auth_port_http(self) -> pulumi.Output[int]:
        """
        Port to use for FortiGuard Web Filter HTTP override authentication
        """
        return pulumi.get(self, "ovrd_auth_port_http")

    @property
    @pulumi.getter(name="ovrdAuthPortHttps")
    def ovrd_auth_port_https(self) -> pulumi.Output[int]:
        """
        Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
        """
        return pulumi.get(self, "ovrd_auth_port_https")

    @property
    @pulumi.getter(name="ovrdAuthPortHttpsFlow")
    def ovrd_auth_port_https_flow(self) -> pulumi.Output[int]:
        """
        Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
        """
        return pulumi.get(self, "ovrd_auth_port_https_flow")

    @property
    @pulumi.getter(name="ovrdAuthPortWarning")
    def ovrd_auth_port_warning(self) -> pulumi.Output[int]:
        """
        Port to use for FortiGuard Web Filter Warning override authentication.
        """
        return pulumi.get(self, "ovrd_auth_port_warning")

    @property
    @pulumi.getter(name="requestPacketSizeLimit")
    def request_packet_size_limit(self) -> pulumi.Output[int]:
        """
        Limit size of URL request packets sent to FortiGuard server (0 for default).
        """
        return pulumi.get(self, "request_packet_size_limit")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="warnAuthHttps")
    def warn_auth_https(self) -> pulumi.Output[str]:
        """
        Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "warn_auth_https")

