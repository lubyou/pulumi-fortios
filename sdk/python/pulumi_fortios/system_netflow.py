# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemNetflowArgs', 'SystemNetflow']

@pulumi.input_type
class SystemNetflowArgs:
    def __init__(__self__, *,
                 active_flow_timeout: Optional[pulumi.Input[int]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 inactive_flow_timeout: Optional[pulumi.Input[int]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 template_tx_counter: Optional[pulumi.Input[int]] = None,
                 template_tx_timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemNetflow resource.
        :param pulumi.Input[int] active_flow_timeout: Timeout to report active flows (1 - 60 min, default = 30).
        :param pulumi.Input[str] collector_ip: Collector IP.
        :param pulumi.Input[int] collector_port: NetFlow collector port number.
        :param pulumi.Input[int] inactive_flow_timeout: Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        :param pulumi.Input[str] source_ip: Source IP address for communication with the NetFlow agent.
        :param pulumi.Input[int] template_tx_counter: Counter of flowset records before resending a template flowset record.
        :param pulumi.Input[int] template_tx_timeout: Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if active_flow_timeout is not None:
            pulumi.set(__self__, "active_flow_timeout", active_flow_timeout)
        if collector_ip is not None:
            pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if inactive_flow_timeout is not None:
            pulumi.set(__self__, "inactive_flow_timeout", inactive_flow_timeout)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if template_tx_counter is not None:
            pulumi.set(__self__, "template_tx_counter", template_tx_counter)
        if template_tx_timeout is not None:
            pulumi.set(__self__, "template_tx_timeout", template_tx_timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="activeFlowTimeout")
    def active_flow_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout to report active flows (1 - 60 min, default = 30).
        """
        return pulumi.get(self, "active_flow_timeout")

    @active_flow_timeout.setter
    def active_flow_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "active_flow_timeout", value)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Collector IP.
        """
        return pulumi.get(self, "collector_ip")

    @collector_ip.setter
    def collector_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collector_ip", value)

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> Optional[pulumi.Input[int]]:
        """
        NetFlow collector port number.
        """
        return pulumi.get(self, "collector_port")

    @collector_port.setter
    def collector_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "collector_port", value)

    @property
    @pulumi.getter(name="inactiveFlowTimeout")
    def inactive_flow_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        """
        return pulumi.get(self, "inactive_flow_timeout")

    @inactive_flow_timeout.setter
    def inactive_flow_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "inactive_flow_timeout", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communication with the NetFlow agent.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="templateTxCounter")
    def template_tx_counter(self) -> Optional[pulumi.Input[int]]:
        """
        Counter of flowset records before resending a template flowset record.
        """
        return pulumi.get(self, "template_tx_counter")

    @template_tx_counter.setter
    def template_tx_counter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_tx_counter", value)

    @property
    @pulumi.getter(name="templateTxTimeout")
    def template_tx_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        """
        return pulumi.get(self, "template_tx_timeout")

    @template_tx_timeout.setter
    def template_tx_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_tx_timeout", value)

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
class _SystemNetflowState:
    def __init__(__self__, *,
                 active_flow_timeout: Optional[pulumi.Input[int]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 inactive_flow_timeout: Optional[pulumi.Input[int]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 template_tx_counter: Optional[pulumi.Input[int]] = None,
                 template_tx_timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemNetflow resources.
        :param pulumi.Input[int] active_flow_timeout: Timeout to report active flows (1 - 60 min, default = 30).
        :param pulumi.Input[str] collector_ip: Collector IP.
        :param pulumi.Input[int] collector_port: NetFlow collector port number.
        :param pulumi.Input[int] inactive_flow_timeout: Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        :param pulumi.Input[str] source_ip: Source IP address for communication with the NetFlow agent.
        :param pulumi.Input[int] template_tx_counter: Counter of flowset records before resending a template flowset record.
        :param pulumi.Input[int] template_tx_timeout: Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if active_flow_timeout is not None:
            pulumi.set(__self__, "active_flow_timeout", active_flow_timeout)
        if collector_ip is not None:
            pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if inactive_flow_timeout is not None:
            pulumi.set(__self__, "inactive_flow_timeout", inactive_flow_timeout)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if template_tx_counter is not None:
            pulumi.set(__self__, "template_tx_counter", template_tx_counter)
        if template_tx_timeout is not None:
            pulumi.set(__self__, "template_tx_timeout", template_tx_timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="activeFlowTimeout")
    def active_flow_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout to report active flows (1 - 60 min, default = 30).
        """
        return pulumi.get(self, "active_flow_timeout")

    @active_flow_timeout.setter
    def active_flow_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "active_flow_timeout", value)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Collector IP.
        """
        return pulumi.get(self, "collector_ip")

    @collector_ip.setter
    def collector_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collector_ip", value)

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> Optional[pulumi.Input[int]]:
        """
        NetFlow collector port number.
        """
        return pulumi.get(self, "collector_port")

    @collector_port.setter
    def collector_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "collector_port", value)

    @property
    @pulumi.getter(name="inactiveFlowTimeout")
    def inactive_flow_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        """
        return pulumi.get(self, "inactive_flow_timeout")

    @inactive_flow_timeout.setter
    def inactive_flow_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "inactive_flow_timeout", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communication with the NetFlow agent.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="templateTxCounter")
    def template_tx_counter(self) -> Optional[pulumi.Input[int]]:
        """
        Counter of flowset records before resending a template flowset record.
        """
        return pulumi.get(self, "template_tx_counter")

    @template_tx_counter.setter
    def template_tx_counter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_tx_counter", value)

    @property
    @pulumi.getter(name="templateTxTimeout")
    def template_tx_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        """
        return pulumi.get(self, "template_tx_timeout")

    @template_tx_timeout.setter
    def template_tx_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_tx_timeout", value)

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


class SystemNetflow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_flow_timeout: Optional[pulumi.Input[int]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 inactive_flow_timeout: Optional[pulumi.Input[int]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 template_tx_counter: Optional[pulumi.Input[int]] = None,
                 template_tx_timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure NetFlow.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemNetflow("trname",
            active_flow_timeout=30,
            collector_ip="0.0.0.0",
            collector_port=2055,
            inactive_flow_timeout=15,
            source_ip="0.0.0.0",
            template_tx_counter=20,
            template_tx_timeout=30)
        ```

        ## Import

        System Netflow can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemNetflow:SystemNetflow labelname SystemNetflow
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] active_flow_timeout: Timeout to report active flows (1 - 60 min, default = 30).
        :param pulumi.Input[str] collector_ip: Collector IP.
        :param pulumi.Input[int] collector_port: NetFlow collector port number.
        :param pulumi.Input[int] inactive_flow_timeout: Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        :param pulumi.Input[str] source_ip: Source IP address for communication with the NetFlow agent.
        :param pulumi.Input[int] template_tx_counter: Counter of flowset records before resending a template flowset record.
        :param pulumi.Input[int] template_tx_timeout: Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemNetflowArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure NetFlow.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemNetflow("trname",
            active_flow_timeout=30,
            collector_ip="0.0.0.0",
            collector_port=2055,
            inactive_flow_timeout=15,
            source_ip="0.0.0.0",
            template_tx_counter=20,
            template_tx_timeout=30)
        ```

        ## Import

        System Netflow can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemNetflow:SystemNetflow labelname SystemNetflow
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemNetflowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemNetflowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_flow_timeout: Optional[pulumi.Input[int]] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 inactive_flow_timeout: Optional[pulumi.Input[int]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 template_tx_counter: Optional[pulumi.Input[int]] = None,
                 template_tx_timeout: Optional[pulumi.Input[int]] = None,
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
            __props__ = SystemNetflowArgs.__new__(SystemNetflowArgs)

            __props__.__dict__["active_flow_timeout"] = active_flow_timeout
            __props__.__dict__["collector_ip"] = collector_ip
            __props__.__dict__["collector_port"] = collector_port
            __props__.__dict__["inactive_flow_timeout"] = inactive_flow_timeout
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["template_tx_counter"] = template_tx_counter
            __props__.__dict__["template_tx_timeout"] = template_tx_timeout
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemNetflow, __self__).__init__(
            'fortios:index/systemNetflow:SystemNetflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_flow_timeout: Optional[pulumi.Input[int]] = None,
            collector_ip: Optional[pulumi.Input[str]] = None,
            collector_port: Optional[pulumi.Input[int]] = None,
            inactive_flow_timeout: Optional[pulumi.Input[int]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            template_tx_counter: Optional[pulumi.Input[int]] = None,
            template_tx_timeout: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemNetflow':
        """
        Get an existing SystemNetflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] active_flow_timeout: Timeout to report active flows (1 - 60 min, default = 30).
        :param pulumi.Input[str] collector_ip: Collector IP.
        :param pulumi.Input[int] collector_port: NetFlow collector port number.
        :param pulumi.Input[int] inactive_flow_timeout: Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        :param pulumi.Input[str] source_ip: Source IP address for communication with the NetFlow agent.
        :param pulumi.Input[int] template_tx_counter: Counter of flowset records before resending a template flowset record.
        :param pulumi.Input[int] template_tx_timeout: Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemNetflowState.__new__(_SystemNetflowState)

        __props__.__dict__["active_flow_timeout"] = active_flow_timeout
        __props__.__dict__["collector_ip"] = collector_ip
        __props__.__dict__["collector_port"] = collector_port
        __props__.__dict__["inactive_flow_timeout"] = inactive_flow_timeout
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["template_tx_counter"] = template_tx_counter
        __props__.__dict__["template_tx_timeout"] = template_tx_timeout
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemNetflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeFlowTimeout")
    def active_flow_timeout(self) -> pulumi.Output[int]:
        """
        Timeout to report active flows (1 - 60 min, default = 30).
        """
        return pulumi.get(self, "active_flow_timeout")

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> pulumi.Output[str]:
        """
        Collector IP.
        """
        return pulumi.get(self, "collector_ip")

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> pulumi.Output[int]:
        """
        NetFlow collector port number.
        """
        return pulumi.get(self, "collector_port")

    @property
    @pulumi.getter(name="inactiveFlowTimeout")
    def inactive_flow_timeout(self) -> pulumi.Output[int]:
        """
        Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
        """
        return pulumi.get(self, "inactive_flow_timeout")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        Source IP address for communication with the NetFlow agent.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="templateTxCounter")
    def template_tx_counter(self) -> pulumi.Output[int]:
        """
        Counter of flowset records before resending a template flowset record.
        """
        return pulumi.get(self, "template_tx_counter")

    @property
    @pulumi.getter(name="templateTxTimeout")
    def template_tx_timeout(self) -> pulumi.Output[int]:
        """
        Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
        """
        return pulumi.get(self, "template_tx_timeout")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

