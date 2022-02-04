# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FirewallAccessProxyArgs', 'FirewallAccessProxy']

@pulumi.input_type
class FirewallAccessProxyArgs:
    def __init__(__self__, *,
                 api_gateway6s: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]]] = None,
                 api_gateways: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]]] = None,
                 auth_portal: Optional[pulumi.Input[str]] = None,
                 auth_virtual_host: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 decrypted_traffic_mirror: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 empty_cert_action: Optional[pulumi.Input[str]] = None,
                 log_blocked_traffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallAccessProxy resource.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]] api_gateway6s: Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]] api_gateways: Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        :param pulumi.Input[str] auth_portal: Enable/disable authentication portal. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] auth_virtual_host: Virtual host for authentication portal.
        :param pulumi.Input[str] client_cert: Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] decrypted_traffic_mirror: Decrypted traffic mirror.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] empty_cert_action: Action of an empty client certificate. Valid values: `accept`, `block`.
        :param pulumi.Input[str] log_blocked_traffic: Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Server host key name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vip: Virtual IP name.
        """
        if api_gateway6s is not None:
            pulumi.set(__self__, "api_gateway6s", api_gateway6s)
        if api_gateways is not None:
            pulumi.set(__self__, "api_gateways", api_gateways)
        if auth_portal is not None:
            pulumi.set(__self__, "auth_portal", auth_portal)
        if auth_virtual_host is not None:
            pulumi.set(__self__, "auth_virtual_host", auth_virtual_host)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if decrypted_traffic_mirror is not None:
            pulumi.set(__self__, "decrypted_traffic_mirror", decrypted_traffic_mirror)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if empty_cert_action is not None:
            pulumi.set(__self__, "empty_cert_action", empty_cert_action)
        if log_blocked_traffic is not None:
            pulumi.set(__self__, "log_blocked_traffic", log_blocked_traffic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vip is not None:
            pulumi.set(__self__, "vip", vip)

    @property
    @pulumi.getter(name="apiGateway6s")
    def api_gateway6s(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]]]:
        """
        Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        """
        return pulumi.get(self, "api_gateway6s")

    @api_gateway6s.setter
    def api_gateway6s(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]]]):
        pulumi.set(self, "api_gateway6s", value)

    @property
    @pulumi.getter(name="apiGateways")
    def api_gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]]]:
        """
        Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        """
        return pulumi.get(self, "api_gateways")

    @api_gateways.setter
    def api_gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]]]):
        pulumi.set(self, "api_gateways", value)

    @property
    @pulumi.getter(name="authPortal")
    def auth_portal(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable authentication portal. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "auth_portal")

    @auth_portal.setter
    def auth_portal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_portal", value)

    @property
    @pulumi.getter(name="authVirtualHost")
    def auth_virtual_host(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual host for authentication portal.
        """
        return pulumi.get(self, "auth_virtual_host")

    @auth_virtual_host.setter
    def auth_virtual_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_virtual_host", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter(name="decryptedTrafficMirror")
    def decrypted_traffic_mirror(self) -> Optional[pulumi.Input[str]]:
        """
        Decrypted traffic mirror.
        """
        return pulumi.get(self, "decrypted_traffic_mirror")

    @decrypted_traffic_mirror.setter
    def decrypted_traffic_mirror(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decrypted_traffic_mirror", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="emptyCertAction")
    def empty_cert_action(self) -> Optional[pulumi.Input[str]]:
        """
        Action of an empty client certificate. Valid values: `accept`, `block`.
        """
        return pulumi.get(self, "empty_cert_action")

    @empty_cert_action.setter
    def empty_cert_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "empty_cert_action", value)

    @property
    @pulumi.getter(name="logBlockedTraffic")
    def log_blocked_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log_blocked_traffic")

    @log_blocked_traffic.setter
    def log_blocked_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_blocked_traffic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server host key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    @pulumi.getter
    def vip(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual IP name.
        """
        return pulumi.get(self, "vip")

    @vip.setter
    def vip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip", value)


@pulumi.input_type
class _FirewallAccessProxyState:
    def __init__(__self__, *,
                 api_gateway6s: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]]] = None,
                 api_gateways: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]]] = None,
                 auth_portal: Optional[pulumi.Input[str]] = None,
                 auth_virtual_host: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 decrypted_traffic_mirror: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 empty_cert_action: Optional[pulumi.Input[str]] = None,
                 log_blocked_traffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vip: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallAccessProxy resources.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]] api_gateway6s: Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]] api_gateways: Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        :param pulumi.Input[str] auth_portal: Enable/disable authentication portal. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] auth_virtual_host: Virtual host for authentication portal.
        :param pulumi.Input[str] client_cert: Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] decrypted_traffic_mirror: Decrypted traffic mirror.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] empty_cert_action: Action of an empty client certificate. Valid values: `accept`, `block`.
        :param pulumi.Input[str] log_blocked_traffic: Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Server host key name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vip: Virtual IP name.
        """
        if api_gateway6s is not None:
            pulumi.set(__self__, "api_gateway6s", api_gateway6s)
        if api_gateways is not None:
            pulumi.set(__self__, "api_gateways", api_gateways)
        if auth_portal is not None:
            pulumi.set(__self__, "auth_portal", auth_portal)
        if auth_virtual_host is not None:
            pulumi.set(__self__, "auth_virtual_host", auth_virtual_host)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if decrypted_traffic_mirror is not None:
            pulumi.set(__self__, "decrypted_traffic_mirror", decrypted_traffic_mirror)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if empty_cert_action is not None:
            pulumi.set(__self__, "empty_cert_action", empty_cert_action)
        if log_blocked_traffic is not None:
            pulumi.set(__self__, "log_blocked_traffic", log_blocked_traffic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vip is not None:
            pulumi.set(__self__, "vip", vip)

    @property
    @pulumi.getter(name="apiGateway6s")
    def api_gateway6s(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]]]:
        """
        Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        """
        return pulumi.get(self, "api_gateway6s")

    @api_gateway6s.setter
    def api_gateway6s(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGateway6Args']]]]):
        pulumi.set(self, "api_gateway6s", value)

    @property
    @pulumi.getter(name="apiGateways")
    def api_gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]]]:
        """
        Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        """
        return pulumi.get(self, "api_gateways")

    @api_gateways.setter
    def api_gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAccessProxyApiGatewayArgs']]]]):
        pulumi.set(self, "api_gateways", value)

    @property
    @pulumi.getter(name="authPortal")
    def auth_portal(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable authentication portal. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "auth_portal")

    @auth_portal.setter
    def auth_portal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_portal", value)

    @property
    @pulumi.getter(name="authVirtualHost")
    def auth_virtual_host(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual host for authentication portal.
        """
        return pulumi.get(self, "auth_virtual_host")

    @auth_virtual_host.setter
    def auth_virtual_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_virtual_host", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter(name="decryptedTrafficMirror")
    def decrypted_traffic_mirror(self) -> Optional[pulumi.Input[str]]:
        """
        Decrypted traffic mirror.
        """
        return pulumi.get(self, "decrypted_traffic_mirror")

    @decrypted_traffic_mirror.setter
    def decrypted_traffic_mirror(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decrypted_traffic_mirror", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="emptyCertAction")
    def empty_cert_action(self) -> Optional[pulumi.Input[str]]:
        """
        Action of an empty client certificate. Valid values: `accept`, `block`.
        """
        return pulumi.get(self, "empty_cert_action")

    @empty_cert_action.setter
    def empty_cert_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "empty_cert_action", value)

    @property
    @pulumi.getter(name="logBlockedTraffic")
    def log_blocked_traffic(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log_blocked_traffic")

    @log_blocked_traffic.setter
    def log_blocked_traffic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_blocked_traffic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server host key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    @pulumi.getter
    def vip(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual IP name.
        """
        return pulumi.get(self, "vip")

    @vip.setter
    def vip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip", value)


class FirewallAccessProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_gateway6s: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGateway6Args']]]]] = None,
                 api_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGatewayArgs']]]]] = None,
                 auth_portal: Optional[pulumi.Input[str]] = None,
                 auth_virtual_host: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 decrypted_traffic_mirror: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 empty_cert_action: Optional[pulumi.Input[str]] = None,
                 log_blocked_traffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPv4 access proxy. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        Firewall AccessProxy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallAccessProxy:FirewallAccessProxy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGateway6Args']]]] api_gateway6s: Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGatewayArgs']]]] api_gateways: Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        :param pulumi.Input[str] auth_portal: Enable/disable authentication portal. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] auth_virtual_host: Virtual host for authentication portal.
        :param pulumi.Input[str] client_cert: Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] decrypted_traffic_mirror: Decrypted traffic mirror.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] empty_cert_action: Action of an empty client certificate. Valid values: `accept`, `block`.
        :param pulumi.Input[str] log_blocked_traffic: Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Server host key name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vip: Virtual IP name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallAccessProxyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPv4 access proxy. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        Firewall AccessProxy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallAccessProxy:FirewallAccessProxy labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallAccessProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallAccessProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_gateway6s: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGateway6Args']]]]] = None,
                 api_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGatewayArgs']]]]] = None,
                 auth_portal: Optional[pulumi.Input[str]] = None,
                 auth_virtual_host: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 decrypted_traffic_mirror: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 empty_cert_action: Optional[pulumi.Input[str]] = None,
                 log_blocked_traffic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vip: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallAccessProxyArgs.__new__(FirewallAccessProxyArgs)

            __props__.__dict__["api_gateway6s"] = api_gateway6s
            __props__.__dict__["api_gateways"] = api_gateways
            __props__.__dict__["auth_portal"] = auth_portal
            __props__.__dict__["auth_virtual_host"] = auth_virtual_host
            __props__.__dict__["client_cert"] = client_cert
            __props__.__dict__["decrypted_traffic_mirror"] = decrypted_traffic_mirror
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["empty_cert_action"] = empty_cert_action
            __props__.__dict__["log_blocked_traffic"] = log_blocked_traffic
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vip"] = vip
        super(FirewallAccessProxy, __self__).__init__(
            'fortios:index/firewallAccessProxy:FirewallAccessProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_gateway6s: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGateway6Args']]]]] = None,
            api_gateways: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGatewayArgs']]]]] = None,
            auth_portal: Optional[pulumi.Input[str]] = None,
            auth_virtual_host: Optional[pulumi.Input[str]] = None,
            client_cert: Optional[pulumi.Input[str]] = None,
            decrypted_traffic_mirror: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            empty_cert_action: Optional[pulumi.Input[str]] = None,
            log_blocked_traffic: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vip: Optional[pulumi.Input[str]] = None) -> 'FirewallAccessProxy':
        """
        Get an existing FirewallAccessProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGateway6Args']]]] api_gateway6s: Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAccessProxyApiGatewayArgs']]]] api_gateways: Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        :param pulumi.Input[str] auth_portal: Enable/disable authentication portal. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] auth_virtual_host: Virtual host for authentication portal.
        :param pulumi.Input[str] client_cert: Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] decrypted_traffic_mirror: Decrypted traffic mirror.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] empty_cert_action: Action of an empty client certificate. Valid values: `accept`, `block`.
        :param pulumi.Input[str] log_blocked_traffic: Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Server host key name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vip: Virtual IP name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallAccessProxyState.__new__(_FirewallAccessProxyState)

        __props__.__dict__["api_gateway6s"] = api_gateway6s
        __props__.__dict__["api_gateways"] = api_gateways
        __props__.__dict__["auth_portal"] = auth_portal
        __props__.__dict__["auth_virtual_host"] = auth_virtual_host
        __props__.__dict__["client_cert"] = client_cert
        __props__.__dict__["decrypted_traffic_mirror"] = decrypted_traffic_mirror
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["empty_cert_action"] = empty_cert_action
        __props__.__dict__["log_blocked_traffic"] = log_blocked_traffic
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vip"] = vip
        return FirewallAccessProxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiGateway6s")
    def api_gateway6s(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallAccessProxyApiGateway6']]]:
        """
        Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
        """
        return pulumi.get(self, "api_gateway6s")

    @property
    @pulumi.getter(name="apiGateways")
    def api_gateways(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallAccessProxyApiGateway']]]:
        """
        Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
        """
        return pulumi.get(self, "api_gateways")

    @property
    @pulumi.getter(name="authPortal")
    def auth_portal(self) -> pulumi.Output[str]:
        """
        Enable/disable authentication portal. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "auth_portal")

    @property
    @pulumi.getter(name="authVirtualHost")
    def auth_virtual_host(self) -> pulumi.Output[str]:
        """
        Virtual host for authentication portal.
        """
        return pulumi.get(self, "auth_virtual_host")

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> pulumi.Output[str]:
        """
        Enable/disable to request client certificate. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "client_cert")

    @property
    @pulumi.getter(name="decryptedTrafficMirror")
    def decrypted_traffic_mirror(self) -> pulumi.Output[str]:
        """
        Decrypted traffic mirror.
        """
        return pulumi.get(self, "decrypted_traffic_mirror")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="emptyCertAction")
    def empty_cert_action(self) -> pulumi.Output[str]:
        """
        Action of an empty client certificate. Valid values: `accept`, `block`.
        """
        return pulumi.get(self, "empty_cert_action")

    @property
    @pulumi.getter(name="logBlockedTraffic")
    def log_blocked_traffic(self) -> pulumi.Output[str]:
        """
        Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log_blocked_traffic")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Server host key name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vip(self) -> pulumi.Output[str]:
        """
        Virtual IP name.
        """
        return pulumi.get(self, "vip")
