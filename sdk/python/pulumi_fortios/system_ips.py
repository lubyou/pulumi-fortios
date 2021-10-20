# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemIpsArgs', 'SystemIps']

@pulumi.input_type
class SystemIpsArgs:
    def __init__(__self__, *,
                 override_signature_hold_by_id: Optional[pulumi.Input[str]] = None,
                 signature_hold_time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemIps resource.
        :param pulumi.Input[str] override_signature_hold_by_id: Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] signature_hold_time: Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if override_signature_hold_by_id is not None:
            pulumi.set(__self__, "override_signature_hold_by_id", override_signature_hold_by_id)
        if signature_hold_time is not None:
            pulumi.set(__self__, "signature_hold_time", signature_hold_time)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="overrideSignatureHoldById")
    def override_signature_hold_by_id(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override_signature_hold_by_id")

    @override_signature_hold_by_id.setter
    def override_signature_hold_by_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_signature_hold_by_id", value)

    @property
    @pulumi.getter(name="signatureHoldTime")
    def signature_hold_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
        """
        return pulumi.get(self, "signature_hold_time")

    @signature_hold_time.setter
    def signature_hold_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signature_hold_time", value)

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
class _SystemIpsState:
    def __init__(__self__, *,
                 override_signature_hold_by_id: Optional[pulumi.Input[str]] = None,
                 signature_hold_time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemIps resources.
        :param pulumi.Input[str] override_signature_hold_by_id: Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] signature_hold_time: Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if override_signature_hold_by_id is not None:
            pulumi.set(__self__, "override_signature_hold_by_id", override_signature_hold_by_id)
        if signature_hold_time is not None:
            pulumi.set(__self__, "signature_hold_time", signature_hold_time)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="overrideSignatureHoldById")
    def override_signature_hold_by_id(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override_signature_hold_by_id")

    @override_signature_hold_by_id.setter
    def override_signature_hold_by_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_signature_hold_by_id", value)

    @property
    @pulumi.getter(name="signatureHoldTime")
    def signature_hold_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
        """
        return pulumi.get(self, "signature_hold_time")

    @signature_hold_time.setter
    def signature_hold_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signature_hold_time", value)

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


class SystemIps(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 override_signature_hold_by_id: Optional[pulumi.Input[str]] = None,
                 signature_hold_time: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPS system settings. Applies to FortiOS Version `>= 6.4.2`.

        ## Import

        System Ips can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemIps:SystemIps labelname SystemIps
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] override_signature_hold_by_id: Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] signature_hold_time: Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemIpsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPS system settings. Applies to FortiOS Version `>= 6.4.2`.

        ## Import

        System Ips can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemIps:SystemIps labelname SystemIps
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemIpsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemIpsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 override_signature_hold_by_id: Optional[pulumi.Input[str]] = None,
                 signature_hold_time: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemIpsArgs.__new__(SystemIpsArgs)

            __props__.__dict__["override_signature_hold_by_id"] = override_signature_hold_by_id
            __props__.__dict__["signature_hold_time"] = signature_hold_time
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemIps, __self__).__init__(
            'fortios:index/systemIps:SystemIps',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            override_signature_hold_by_id: Optional[pulumi.Input[str]] = None,
            signature_hold_time: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemIps':
        """
        Get an existing SystemIps resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] override_signature_hold_by_id: Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] signature_hold_time: Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemIpsState.__new__(_SystemIpsState)

        __props__.__dict__["override_signature_hold_by_id"] = override_signature_hold_by_id
        __props__.__dict__["signature_hold_time"] = signature_hold_time
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemIps(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="overrideSignatureHoldById")
    def override_signature_hold_by_id(self) -> pulumi.Output[str]:
        """
        Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override_signature_hold_by_id")

    @property
    @pulumi.getter(name="signatureHoldTime")
    def signature_hold_time(self) -> pulumi.Output[str]:
        """
        Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
        """
        return pulumi.get(self, "signature_hold_time")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
