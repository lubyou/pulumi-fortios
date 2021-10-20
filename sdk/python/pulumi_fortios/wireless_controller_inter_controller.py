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

__all__ = ['WirelessControllerInterControllerArgs', 'WirelessControllerInterController']

@pulumi.input_type
class WirelessControllerInterControllerArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fast_failover_max: Optional[pulumi.Input[int]] = None,
                 fast_failover_wait: Optional[pulumi.Input[int]] = None,
                 inter_controller_key: Optional[pulumi.Input[str]] = None,
                 inter_controller_mode: Optional[pulumi.Input[str]] = None,
                 inter_controller_peers: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]]] = None,
                 inter_controller_pri: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerInterController resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[int] fast_failover_max: Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        :param pulumi.Input[int] fast_failover_wait: Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        :param pulumi.Input[str] inter_controller_key: Secret key for inter-controller communications.
        :param pulumi.Input[str] inter_controller_mode: Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]] inter_controller_peers: Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        :param pulumi.Input[str] inter_controller_pri: Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fast_failover_max is not None:
            pulumi.set(__self__, "fast_failover_max", fast_failover_max)
        if fast_failover_wait is not None:
            pulumi.set(__self__, "fast_failover_wait", fast_failover_wait)
        if inter_controller_key is not None:
            pulumi.set(__self__, "inter_controller_key", inter_controller_key)
        if inter_controller_mode is not None:
            pulumi.set(__self__, "inter_controller_mode", inter_controller_mode)
        if inter_controller_peers is not None:
            pulumi.set(__self__, "inter_controller_peers", inter_controller_peers)
        if inter_controller_pri is not None:
            pulumi.set(__self__, "inter_controller_pri", inter_controller_pri)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="fastFailoverMax")
    def fast_failover_max(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        """
        return pulumi.get(self, "fast_failover_max")

    @fast_failover_max.setter
    def fast_failover_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fast_failover_max", value)

    @property
    @pulumi.getter(name="fastFailoverWait")
    def fast_failover_wait(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        """
        return pulumi.get(self, "fast_failover_wait")

    @fast_failover_wait.setter
    def fast_failover_wait(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fast_failover_wait", value)

    @property
    @pulumi.getter(name="interControllerKey")
    def inter_controller_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret key for inter-controller communications.
        """
        return pulumi.get(self, "inter_controller_key")

    @inter_controller_key.setter
    def inter_controller_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inter_controller_key", value)

    @property
    @pulumi.getter(name="interControllerMode")
    def inter_controller_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        """
        return pulumi.get(self, "inter_controller_mode")

    @inter_controller_mode.setter
    def inter_controller_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inter_controller_mode", value)

    @property
    @pulumi.getter(name="interControllerPeers")
    def inter_controller_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]]]:
        """
        Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        """
        return pulumi.get(self, "inter_controller_peers")

    @inter_controller_peers.setter
    def inter_controller_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]]]):
        pulumi.set(self, "inter_controller_peers", value)

    @property
    @pulumi.getter(name="interControllerPri")
    def inter_controller_pri(self) -> Optional[pulumi.Input[str]]:
        """
        Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        """
        return pulumi.get(self, "inter_controller_pri")

    @inter_controller_pri.setter
    def inter_controller_pri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inter_controller_pri", value)

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
class _WirelessControllerInterControllerState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fast_failover_max: Optional[pulumi.Input[int]] = None,
                 fast_failover_wait: Optional[pulumi.Input[int]] = None,
                 inter_controller_key: Optional[pulumi.Input[str]] = None,
                 inter_controller_mode: Optional[pulumi.Input[str]] = None,
                 inter_controller_peers: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]]] = None,
                 inter_controller_pri: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerInterController resources.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[int] fast_failover_max: Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        :param pulumi.Input[int] fast_failover_wait: Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        :param pulumi.Input[str] inter_controller_key: Secret key for inter-controller communications.
        :param pulumi.Input[str] inter_controller_mode: Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]] inter_controller_peers: Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        :param pulumi.Input[str] inter_controller_pri: Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fast_failover_max is not None:
            pulumi.set(__self__, "fast_failover_max", fast_failover_max)
        if fast_failover_wait is not None:
            pulumi.set(__self__, "fast_failover_wait", fast_failover_wait)
        if inter_controller_key is not None:
            pulumi.set(__self__, "inter_controller_key", inter_controller_key)
        if inter_controller_mode is not None:
            pulumi.set(__self__, "inter_controller_mode", inter_controller_mode)
        if inter_controller_peers is not None:
            pulumi.set(__self__, "inter_controller_peers", inter_controller_peers)
        if inter_controller_pri is not None:
            pulumi.set(__self__, "inter_controller_pri", inter_controller_pri)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    @pulumi.getter(name="fastFailoverMax")
    def fast_failover_max(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        """
        return pulumi.get(self, "fast_failover_max")

    @fast_failover_max.setter
    def fast_failover_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fast_failover_max", value)

    @property
    @pulumi.getter(name="fastFailoverWait")
    def fast_failover_wait(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        """
        return pulumi.get(self, "fast_failover_wait")

    @fast_failover_wait.setter
    def fast_failover_wait(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fast_failover_wait", value)

    @property
    @pulumi.getter(name="interControllerKey")
    def inter_controller_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret key for inter-controller communications.
        """
        return pulumi.get(self, "inter_controller_key")

    @inter_controller_key.setter
    def inter_controller_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inter_controller_key", value)

    @property
    @pulumi.getter(name="interControllerMode")
    def inter_controller_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        """
        return pulumi.get(self, "inter_controller_mode")

    @inter_controller_mode.setter
    def inter_controller_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inter_controller_mode", value)

    @property
    @pulumi.getter(name="interControllerPeers")
    def inter_controller_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]]]:
        """
        Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        """
        return pulumi.get(self, "inter_controller_peers")

    @inter_controller_peers.setter
    def inter_controller_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerInterControllerInterControllerPeerArgs']]]]):
        pulumi.set(self, "inter_controller_peers", value)

    @property
    @pulumi.getter(name="interControllerPri")
    def inter_controller_pri(self) -> Optional[pulumi.Input[str]]:
        """
        Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        """
        return pulumi.get(self, "inter_controller_pri")

    @inter_controller_pri.setter
    def inter_controller_pri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inter_controller_pri", value)

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


class WirelessControllerInterController(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fast_failover_max: Optional[pulumi.Input[int]] = None,
                 fast_failover_wait: Optional[pulumi.Input[int]] = None,
                 inter_controller_key: Optional[pulumi.Input[str]] = None,
                 inter_controller_mode: Optional[pulumi.Input[str]] = None,
                 inter_controller_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerInterControllerInterControllerPeerArgs']]]]] = None,
                 inter_controller_pri: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure inter wireless controller operation.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WirelessControllerInterController("trname",
            fast_failover_max=10,
            fast_failover_wait=10,
            inter_controller_key="ENC XXXX",
            inter_controller_mode="disable",
            inter_controller_pri="primary")
        ```

        ## Import

        WirelessController InterController can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerInterController:WirelessControllerInterController labelname WirelessControllerInterController
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[int] fast_failover_max: Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        :param pulumi.Input[int] fast_failover_wait: Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        :param pulumi.Input[str] inter_controller_key: Secret key for inter-controller communications.
        :param pulumi.Input[str] inter_controller_mode: Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerInterControllerInterControllerPeerArgs']]]] inter_controller_peers: Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        :param pulumi.Input[str] inter_controller_pri: Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerInterControllerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure inter wireless controller operation.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WirelessControllerInterController("trname",
            fast_failover_max=10,
            fast_failover_wait=10,
            inter_controller_key="ENC XXXX",
            inter_controller_mode="disable",
            inter_controller_pri="primary")
        ```

        ## Import

        WirelessController InterController can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerInterController:WirelessControllerInterController labelname WirelessControllerInterController
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelessControllerInterControllerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerInterControllerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fast_failover_max: Optional[pulumi.Input[int]] = None,
                 fast_failover_wait: Optional[pulumi.Input[int]] = None,
                 inter_controller_key: Optional[pulumi.Input[str]] = None,
                 inter_controller_mode: Optional[pulumi.Input[str]] = None,
                 inter_controller_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerInterControllerInterControllerPeerArgs']]]]] = None,
                 inter_controller_pri: Optional[pulumi.Input[str]] = None,
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
            __props__ = WirelessControllerInterControllerArgs.__new__(WirelessControllerInterControllerArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["fast_failover_max"] = fast_failover_max
            __props__.__dict__["fast_failover_wait"] = fast_failover_wait
            __props__.__dict__["inter_controller_key"] = inter_controller_key
            __props__.__dict__["inter_controller_mode"] = inter_controller_mode
            __props__.__dict__["inter_controller_peers"] = inter_controller_peers
            __props__.__dict__["inter_controller_pri"] = inter_controller_pri
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerInterController, __self__).__init__(
            'fortios:index/wirelessControllerInterController:WirelessControllerInterController',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fast_failover_max: Optional[pulumi.Input[int]] = None,
            fast_failover_wait: Optional[pulumi.Input[int]] = None,
            inter_controller_key: Optional[pulumi.Input[str]] = None,
            inter_controller_mode: Optional[pulumi.Input[str]] = None,
            inter_controller_peers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerInterControllerInterControllerPeerArgs']]]]] = None,
            inter_controller_pri: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerInterController':
        """
        Get an existing WirelessControllerInterController resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[int] fast_failover_max: Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        :param pulumi.Input[int] fast_failover_wait: Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        :param pulumi.Input[str] inter_controller_key: Secret key for inter-controller communications.
        :param pulumi.Input[str] inter_controller_mode: Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerInterControllerInterControllerPeerArgs']]]] inter_controller_peers: Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        :param pulumi.Input[str] inter_controller_pri: Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerInterControllerState.__new__(_WirelessControllerInterControllerState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fast_failover_max"] = fast_failover_max
        __props__.__dict__["fast_failover_wait"] = fast_failover_wait
        __props__.__dict__["inter_controller_key"] = inter_controller_key
        __props__.__dict__["inter_controller_mode"] = inter_controller_mode
        __props__.__dict__["inter_controller_peers"] = inter_controller_peers
        __props__.__dict__["inter_controller_pri"] = inter_controller_pri
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerInterController(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="fastFailoverMax")
    def fast_failover_max(self) -> pulumi.Output[int]:
        """
        Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
        """
        return pulumi.get(self, "fast_failover_max")

    @property
    @pulumi.getter(name="fastFailoverWait")
    def fast_failover_wait(self) -> pulumi.Output[int]:
        """
        Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
        """
        return pulumi.get(self, "fast_failover_wait")

    @property
    @pulumi.getter(name="interControllerKey")
    def inter_controller_key(self) -> pulumi.Output[Optional[str]]:
        """
        Secret key for inter-controller communications.
        """
        return pulumi.get(self, "inter_controller_key")

    @property
    @pulumi.getter(name="interControllerMode")
    def inter_controller_mode(self) -> pulumi.Output[str]:
        """
        Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
        """
        return pulumi.get(self, "inter_controller_mode")

    @property
    @pulumi.getter(name="interControllerPeers")
    def inter_controller_peers(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerInterControllerInterControllerPeer']]]:
        """
        Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
        """
        return pulumi.get(self, "inter_controller_peers")

    @property
    @pulumi.getter(name="interControllerPri")
    def inter_controller_pri(self) -> pulumi.Output[str]:
        """
        Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
        """
        return pulumi.get(self, "inter_controller_pri")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
