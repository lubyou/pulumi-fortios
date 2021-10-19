# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemLicenseVDOMArgs', 'SystemLicenseVDOM']

@pulumi.input_type
class SystemLicenseVDOMArgs:
    def __init__(__self__, *,
                 license: pulumi.Input[str]):
        """
        The set of arguments for constructing a SystemLicenseVDOM resource.
        :param pulumi.Input[str] license: Registration code.
        """
        pulumi.set(__self__, "license", license)

    @property
    @pulumi.getter
    def license(self) -> pulumi.Input[str]:
        """
        Registration code.
        """
        return pulumi.get(self, "license")

    @license.setter
    def license(self, value: pulumi.Input[str]):
        pulumi.set(self, "license", value)


@pulumi.input_type
class _SystemLicenseVDOMState:
    def __init__(__self__, *,
                 license: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemLicenseVDOM resources.
        :param pulumi.Input[str] license: Registration code.
        """
        if license is not None:
            pulumi.set(__self__, "license", license)

    @property
    @pulumi.getter
    def license(self) -> Optional[pulumi.Input[str]]:
        """
        Registration code.
        """
        return pulumi.get(self, "license")

    @license.setter
    def license(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license", value)


class SystemLicenseVDOM(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 license: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to add a VDOM license for FortiOS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test2 = fortios.SystemLicenseVDOM("test2", license="license")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] license: Registration code.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemLicenseVDOMArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to add a VDOM license for FortiOS.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test2 = fortios.SystemLicenseVDOM("test2", license="license")
        ```

        :param str resource_name: The name of the resource.
        :param SystemLicenseVDOMArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemLicenseVDOMArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 license: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemLicenseVDOMArgs.__new__(SystemLicenseVDOMArgs)

            if license is None and not opts.urn:
                raise TypeError("Missing required property 'license'")
            __props__.__dict__["license"] = license
        super(SystemLicenseVDOM, __self__).__init__(
            'fortios:index/systemLicenseVDOM:SystemLicenseVDOM',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            license: Optional[pulumi.Input[str]] = None) -> 'SystemLicenseVDOM':
        """
        Get an existing SystemLicenseVDOM resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] license: Registration code.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemLicenseVDOMState.__new__(_SystemLicenseVDOMState)

        __props__.__dict__["license"] = license
        return SystemLicenseVDOM(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def license(self) -> pulumi.Output[str]:
        """
        Registration code.
        """
        return pulumi.get(self, "license")

