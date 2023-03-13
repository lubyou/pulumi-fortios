// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("fortios");

/**
 * CA Bundle file content
 */
export declare const cabundlecontent: string | undefined;
Object.defineProperty(exports, "cabundlecontent", {
    get() {
        return __config.get("cabundlecontent");
    },
    enumerable: true,
});

/**
 * CA Bundle file
 */
export declare const cabundlefile: string | undefined;
Object.defineProperty(exports, "cabundlefile", {
    get() {
        return __config.get("cabundlefile") ?? utilities.getEnv("FORTIOS_CA_CABUNDLE");
    },
    enumerable: true,
});

/**
 * CA certtificate(Optional)
 */
export declare const cacert: string | undefined;
Object.defineProperty(exports, "cacert", {
    get() {
        return __config.get("cacert");
    },
    enumerable: true,
});

/**
 * User certificate
 */
export declare const clientcert: string | undefined;
Object.defineProperty(exports, "clientcert", {
    get() {
        return __config.get("clientcert");
    },
    enumerable: true,
});

/**
 * User private key
 */
export declare const clientkey: string | undefined;
Object.defineProperty(exports, "clientkey", {
    get() {
        return __config.get("clientkey");
    },
    enumerable: true,
});

/**
 * CA Bundle file
 */
export declare const fmgCabundlefile: string | undefined;
Object.defineProperty(exports, "fmgCabundlefile", {
    get() {
        return __config.get("fmgCabundlefile") ?? utilities.getEnv("FORTIOS_FMG_CABUNDLE");
    },
    enumerable: true,
});

/**
 * Hostname/IP address of the FortiManager to connect to
 */
export declare const fmgHostname: string | undefined;
Object.defineProperty(exports, "fmgHostname", {
    get() {
        return __config.get("fmgHostname") ?? utilities.getEnv("FORTIOS_FMG_HOSTNAME");
    },
    enumerable: true,
});

export declare const fmgInsecure: boolean | undefined;
Object.defineProperty(exports, "fmgInsecure", {
    get() {
        return __config.getObject<boolean>("fmgInsecure") ?? utilities.getEnvBoolean("FORTIOS_FMG_INSECURE");
    },
    enumerable: true,
});

export declare const fmgPasswd: string | undefined;
Object.defineProperty(exports, "fmgPasswd", {
    get() {
        return __config.get("fmgPasswd") ?? utilities.getEnv("FORTIOS_FMG_PASSWORD");
    },
    enumerable: true,
});

export declare const fmgUsername: string | undefined;
Object.defineProperty(exports, "fmgUsername", {
    get() {
        return __config.get("fmgUsername") ?? utilities.getEnv("FORTIOS_FMG_USERNAME");
    },
    enumerable: true,
});

/**
 * The hostname/IP address of the FortiOS to be connected
 */
export declare const hostname: string | undefined;
Object.defineProperty(exports, "hostname", {
    get() {
        return __config.get("hostname") ?? utilities.getEnv("FORTIOS_ACCESS_HOSTNAME");
    },
    enumerable: true,
});

/**
 * HTTP proxy address
 */
export declare const httpProxy: string | undefined;
Object.defineProperty(exports, "httpProxy", {
    get() {
        return __config.get("httpProxy");
    },
    enumerable: true,
});

export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure") ?? utilities.getEnvBoolean("FORTIOS_INSECURE");
    },
    enumerable: true,
});

/**
 * Enable/disable peer authentication, can be 'enable' or 'disable'
 */
export declare const peerauth: string | undefined;
Object.defineProperty(exports, "peerauth", {
    get() {
        return __config.get("peerauth");
    },
    enumerable: true,
});

export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token") ?? utilities.getEnv("FORTIOS_ACCESS_TOKEN");
    },
    enumerable: true,
});

export declare const vdom: string | undefined;
Object.defineProperty(exports, "vdom", {
    get() {
        return __config.get("vdom");
    },
    enumerable: true,
});

