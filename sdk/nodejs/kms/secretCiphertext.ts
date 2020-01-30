// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_secret_ciphertext.html.markdown.
 */
export class SecretCiphertext extends pulumi.CustomResource {
    /**
     * Get an existing SecretCiphertext resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretCiphertextState, opts?: pulumi.CustomResourceOptions): SecretCiphertext {
        return new SecretCiphertext(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/secretCiphertext:SecretCiphertext';

    /**
     * Returns true if the given object is an instance of SecretCiphertext.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretCiphertext {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretCiphertext.__pulumiType;
    }

    /**
     * Contains the result of encrypting the provided plaintext, encoded in base64.
     */
    public /*out*/ readonly ciphertext!: pulumi.Output<string>;
    /**
     * The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
     * ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
     */
    public readonly cryptoKey!: pulumi.Output<string>;
    /**
     * The plaintext to be encrypted.
     */
    public readonly plaintext!: pulumi.Output<string>;

    /**
     * Create a SecretCiphertext resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretCiphertextArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretCiphertextArgs | SecretCiphertextState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretCiphertextState | undefined;
            inputs["ciphertext"] = state ? state.ciphertext : undefined;
            inputs["cryptoKey"] = state ? state.cryptoKey : undefined;
            inputs["plaintext"] = state ? state.plaintext : undefined;
        } else {
            const args = argsOrState as SecretCiphertextArgs | undefined;
            if (!args || args.cryptoKey === undefined) {
                throw new Error("Missing required property 'cryptoKey'");
            }
            if (!args || args.plaintext === undefined) {
                throw new Error("Missing required property 'plaintext'");
            }
            inputs["cryptoKey"] = args ? args.cryptoKey : undefined;
            inputs["plaintext"] = args ? args.plaintext : undefined;
            inputs["ciphertext"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretCiphertext.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretCiphertext resources.
 */
export interface SecretCiphertextState {
    /**
     * Contains the result of encrypting the provided plaintext, encoded in base64.
     */
    readonly ciphertext?: pulumi.Input<string>;
    /**
     * The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
     * ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
     */
    readonly cryptoKey?: pulumi.Input<string>;
    /**
     * The plaintext to be encrypted.
     */
    readonly plaintext?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretCiphertext resource.
 */
export interface SecretCiphertextArgs {
    /**
     * The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
     * ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
     */
    readonly cryptoKey: pulumi.Input<string>;
    /**
     * The plaintext to be encrypted.
     */
    readonly plaintext: pulumi.Input<string>;
}