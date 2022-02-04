# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemVdomArgs', 'SystemVdom']

@pulumi.input_type
class SystemVdomArgs:
    def __init__(__self__, *,
                 flag: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[int]] = None,
                 vcluster_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemVdom resource.
        :param pulumi.Input[int] flag: Flag.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[int] temporary: Temporary.
        :param pulumi.Input[int] vcluster_id: Virtual cluster ID (0 - 4294967295).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if flag is not None:
            pulumi.set(__self__, "flag", flag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if short_name is not None:
            pulumi.set(__self__, "short_name", short_name)
        if temporary is not None:
            pulumi.set(__self__, "temporary", temporary)
        if vcluster_id is not None:
            pulumi.set(__self__, "vcluster_id", vcluster_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def flag(self) -> Optional[pulumi.Input[int]]:
        """
        Flag.
        """
        return pulumi.get(self, "flag")

    @flag.setter
    def flag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "flag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM short name.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)

    @property
    @pulumi.getter
    def temporary(self) -> Optional[pulumi.Input[int]]:
        """
        Temporary.
        """
        return pulumi.get(self, "temporary")

    @temporary.setter
    def temporary(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "temporary", value)

    @property
    @pulumi.getter(name="vclusterId")
    def vcluster_id(self) -> Optional[pulumi.Input[int]]:
        """
        Virtual cluster ID (0 - 4294967295).
        """
        return pulumi.get(self, "vcluster_id")

    @vcluster_id.setter
    def vcluster_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vcluster_id", value)

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
class _SystemVdomState:
    def __init__(__self__, *,
                 flag: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[int]] = None,
                 vcluster_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemVdom resources.
        :param pulumi.Input[int] flag: Flag.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[int] temporary: Temporary.
        :param pulumi.Input[int] vcluster_id: Virtual cluster ID (0 - 4294967295).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if flag is not None:
            pulumi.set(__self__, "flag", flag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if short_name is not None:
            pulumi.set(__self__, "short_name", short_name)
        if temporary is not None:
            pulumi.set(__self__, "temporary", temporary)
        if vcluster_id is not None:
            pulumi.set(__self__, "vcluster_id", vcluster_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def flag(self) -> Optional[pulumi.Input[int]]:
        """
        Flag.
        """
        return pulumi.get(self, "flag")

    @flag.setter
    def flag(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "flag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        VDOM short name.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)

    @property
    @pulumi.getter
    def temporary(self) -> Optional[pulumi.Input[int]]:
        """
        Temporary.
        """
        return pulumi.get(self, "temporary")

    @temporary.setter
    def temporary(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "temporary", value)

    @property
    @pulumi.getter(name="vclusterId")
    def vcluster_id(self) -> Optional[pulumi.Input[int]]:
        """
        Virtual cluster ID (0 - 4294967295).
        """
        return pulumi.get(self, "vcluster_id")

    @vcluster_id.setter
    def vcluster_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vcluster_id", value)

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


class SystemVdom(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flag: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[int]] = None,
                 vcluster_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure virtual domain.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemVdom("trname",
            short_name="testvdom",
            temporary=0)
        ```

        ## Import

        System Vdom can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemVdom:SystemVdom labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] flag: Flag.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[int] temporary: Temporary.
        :param pulumi.Input[int] vcluster_id: Virtual cluster ID (0 - 4294967295).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemVdomArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure virtual domain.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemVdom("trname",
            short_name="testvdom",
            temporary=0)
        ```

        ## Import

        System Vdom can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemVdom:SystemVdom labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemVdomArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemVdomArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flag: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 temporary: Optional[pulumi.Input[int]] = None,
                 vcluster_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemVdomArgs.__new__(SystemVdomArgs)

            __props__.__dict__["flag"] = flag
            __props__.__dict__["name"] = name
            __props__.__dict__["short_name"] = short_name
            __props__.__dict__["temporary"] = temporary
            __props__.__dict__["vcluster_id"] = vcluster_id
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemVdom, __self__).__init__(
            'fortios:index/systemVdom:SystemVdom',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            flag: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None,
            temporary: Optional[pulumi.Input[int]] = None,
            vcluster_id: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemVdom':
        """
        Get an existing SystemVdom resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] flag: Flag.
        :param pulumi.Input[str] name: VDOM name.
        :param pulumi.Input[str] short_name: VDOM short name.
        :param pulumi.Input[int] temporary: Temporary.
        :param pulumi.Input[int] vcluster_id: Virtual cluster ID (0 - 4294967295).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemVdomState.__new__(_SystemVdomState)

        __props__.__dict__["flag"] = flag
        __props__.__dict__["name"] = name
        __props__.__dict__["short_name"] = short_name
        __props__.__dict__["temporary"] = temporary
        __props__.__dict__["vcluster_id"] = vcluster_id
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemVdom(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def flag(self) -> pulumi.Output[int]:
        """
        Flag.
        """
        return pulumi.get(self, "flag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        VDOM name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> pulumi.Output[str]:
        """
        VDOM short name.
        """
        return pulumi.get(self, "short_name")

    @property
    @pulumi.getter
    def temporary(self) -> pulumi.Output[int]:
        """
        Temporary.
        """
        return pulumi.get(self, "temporary")

    @property
    @pulumi.getter(name="vclusterId")
    def vcluster_id(self) -> pulumi.Output[int]:
        """
        Virtual cluster ID (0 - 4294967295).
        """
        return pulumi.get(self, "vcluster_id")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

