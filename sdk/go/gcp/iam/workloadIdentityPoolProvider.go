// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// WorkloadIdentityPoolProvider can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default {{project}}/{{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default {{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
// ```
type WorkloadIdentityPoolProvider struct {
	pulumi.CustomResourceState

	// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	// * `assertion`: JSON representing the authentication credential issued by the provider.
	// * `google`: The Google attributes mapped from the assertion in the `attributeMappings`.
	// * `attribute`: The custom attributes mapped from the assertion in the `attributeMappings`.
	//   The maximum length of the attribute condition expression is 4096 characters. If
	//   unspecified, all valid authentication credential are accepted.
	//   The following example shows how to only allow credentials with a mapped `google.groups`
	//   value of `admins`:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeCondition pulumi.StringPtrOutput `pulumi:"attributeCondition"`
	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as `subject` and `segment`.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// * `google.subject`: The principal IAM is authenticating. You can reference this value
	//   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
	//   Cannot exceed 127 characters.
	// * `google.groups`: Groups the external identity belongs to. You can grant groups
	//   access to resources using an IAM `principalSet` binding; access applies to all
	//   members of the group.
	//   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
	//   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
	//   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
	//   is 100 characters, and the key may only contain the characters [a-z0-9_].
	//   You can reference these attributes in IAM policies to define fine-grained access for a
	//   workload to Google Cloud resources. For example:
	// * `google.subject`:
	//   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
	// * `google.groups`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
	// * `attribute.{custom_attribute}`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
	//   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	//   function that maps an identity provider credential to the normalized attribute specified
	//   by the corresponding map key.
	//   You can use the `assertion` keyword in the expression to access a JSON representation of
	//   the authentication credential issued by the provider.
	//   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	//   the total size of all mapped attributes must not exceed 8KB.
	//   For AWS providers, the following rules apply:
	// - If no attribute mapping is defined, the following default mapping applies:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// - If any custom attribute mappings are defined, they must include a mapping to the
	//   `google.subject` attribute.
	//   For OIDC providers, the following rules apply:
	// - Custom attribute mappings must be defined, and must include a mapping to the
	//   `google.subject` attribute. For example, the following maps the `sub` claim of the
	//   incoming credential to the `subject` attribute on a Google token.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeMapping pulumi.StringMapOutput `pulumi:"attributeMapping"`
	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	Aws WorkloadIdentityPoolProviderAwsPtrOutput `pulumi:"aws"`
	// A description for the provider. Cannot exceed 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The resource name of the provider as
	// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.
	Name pulumi.StringOutput `pulumi:"name"`
	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	Oidc WorkloadIdentityPoolProviderOidcPtrOutput `pulumi:"oidc"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The state of the provider. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The provider is active, and may be used to
	// validate authentication credentials. * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
	// deleted after approximately 30 days. You can restore a soft-deleted provider using UndeleteWorkloadIdentityPoolProvider.
	// You cannot reuse the ID of a soft-deleted provider until it is permanently deleted.
	State pulumi.StringOutput `pulumi:"state"`
	// The ID used for the pool, which is the final component of the pool resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId pulumi.StringOutput `pulumi:"workloadIdentityPoolId"`
	// The ID for the provider, which becomes the final component of the resource name. This
	// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolProviderId pulumi.StringOutput `pulumi:"workloadIdentityPoolProviderId"`
}

// NewWorkloadIdentityPoolProvider registers a new resource with the given unique name, arguments, and options.
func NewWorkloadIdentityPoolProvider(ctx *pulumi.Context,
	name string, args *WorkloadIdentityPoolProviderArgs, opts ...pulumi.ResourceOption) (*WorkloadIdentityPoolProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.WorkloadIdentityPoolId == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadIdentityPoolId'")
	}
	if args.WorkloadIdentityPoolProviderId == nil {
		return nil, errors.New("invalid value for required argument 'WorkloadIdentityPoolProviderId'")
	}
	var resource WorkloadIdentityPoolProvider
	err := ctx.RegisterResource("gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadIdentityPoolProvider gets an existing WorkloadIdentityPoolProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadIdentityPoolProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadIdentityPoolProviderState, opts ...pulumi.ResourceOption) (*WorkloadIdentityPoolProvider, error) {
	var resource WorkloadIdentityPoolProvider
	err := ctx.ReadResource("gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadIdentityPoolProvider resources.
type workloadIdentityPoolProviderState struct {
	// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	// * `assertion`: JSON representing the authentication credential issued by the provider.
	// * `google`: The Google attributes mapped from the assertion in the `attributeMappings`.
	// * `attribute`: The custom attributes mapped from the assertion in the `attributeMappings`.
	//   The maximum length of the attribute condition expression is 4096 characters. If
	//   unspecified, all valid authentication credential are accepted.
	//   The following example shows how to only allow credentials with a mapped `google.groups`
	//   value of `admins`:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeCondition *string `pulumi:"attributeCondition"`
	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as `subject` and `segment`.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// * `google.subject`: The principal IAM is authenticating. You can reference this value
	//   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
	//   Cannot exceed 127 characters.
	// * `google.groups`: Groups the external identity belongs to. You can grant groups
	//   access to resources using an IAM `principalSet` binding; access applies to all
	//   members of the group.
	//   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
	//   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
	//   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
	//   is 100 characters, and the key may only contain the characters [a-z0-9_].
	//   You can reference these attributes in IAM policies to define fine-grained access for a
	//   workload to Google Cloud resources. For example:
	// * `google.subject`:
	//   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
	// * `google.groups`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
	// * `attribute.{custom_attribute}`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
	//   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	//   function that maps an identity provider credential to the normalized attribute specified
	//   by the corresponding map key.
	//   You can use the `assertion` keyword in the expression to access a JSON representation of
	//   the authentication credential issued by the provider.
	//   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	//   the total size of all mapped attributes must not exceed 8KB.
	//   For AWS providers, the following rules apply:
	// - If no attribute mapping is defined, the following default mapping applies:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// - If any custom attribute mappings are defined, they must include a mapping to the
	//   `google.subject` attribute.
	//   For OIDC providers, the following rules apply:
	// - Custom attribute mappings must be defined, and must include a mapping to the
	//   `google.subject` attribute. For example, the following maps the `sub` claim of the
	//   incoming credential to the `subject` attribute on a Google token.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeMapping map[string]string `pulumi:"attributeMapping"`
	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	Aws *WorkloadIdentityPoolProviderAws `pulumi:"aws"`
	// A description for the provider. Cannot exceed 256 characters.
	Description *string `pulumi:"description"`
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	Disabled *bool `pulumi:"disabled"`
	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of the provider as
	// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.
	Name *string `pulumi:"name"`
	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	Oidc *WorkloadIdentityPoolProviderOidc `pulumi:"oidc"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The state of the provider. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The provider is active, and may be used to
	// validate authentication credentials. * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
	// deleted after approximately 30 days. You can restore a soft-deleted provider using UndeleteWorkloadIdentityPoolProvider.
	// You cannot reuse the ID of a soft-deleted provider until it is permanently deleted.
	State *string `pulumi:"state"`
	// The ID used for the pool, which is the final component of the pool resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId *string `pulumi:"workloadIdentityPoolId"`
	// The ID for the provider, which becomes the final component of the resource name. This
	// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolProviderId *string `pulumi:"workloadIdentityPoolProviderId"`
}

type WorkloadIdentityPoolProviderState struct {
	// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	// * `assertion`: JSON representing the authentication credential issued by the provider.
	// * `google`: The Google attributes mapped from the assertion in the `attributeMappings`.
	// * `attribute`: The custom attributes mapped from the assertion in the `attributeMappings`.
	//   The maximum length of the attribute condition expression is 4096 characters. If
	//   unspecified, all valid authentication credential are accepted.
	//   The following example shows how to only allow credentials with a mapped `google.groups`
	//   value of `admins`:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeCondition pulumi.StringPtrInput
	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as `subject` and `segment`.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// * `google.subject`: The principal IAM is authenticating. You can reference this value
	//   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
	//   Cannot exceed 127 characters.
	// * `google.groups`: Groups the external identity belongs to. You can grant groups
	//   access to resources using an IAM `principalSet` binding; access applies to all
	//   members of the group.
	//   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
	//   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
	//   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
	//   is 100 characters, and the key may only contain the characters [a-z0-9_].
	//   You can reference these attributes in IAM policies to define fine-grained access for a
	//   workload to Google Cloud resources. For example:
	// * `google.subject`:
	//   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
	// * `google.groups`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
	// * `attribute.{custom_attribute}`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
	//   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	//   function that maps an identity provider credential to the normalized attribute specified
	//   by the corresponding map key.
	//   You can use the `assertion` keyword in the expression to access a JSON representation of
	//   the authentication credential issued by the provider.
	//   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	//   the total size of all mapped attributes must not exceed 8KB.
	//   For AWS providers, the following rules apply:
	// - If no attribute mapping is defined, the following default mapping applies:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// - If any custom attribute mappings are defined, they must include a mapping to the
	//   `google.subject` attribute.
	//   For OIDC providers, the following rules apply:
	// - Custom attribute mappings must be defined, and must include a mapping to the
	//   `google.subject` attribute. For example, the following maps the `sub` claim of the
	//   incoming credential to the `subject` attribute on a Google token.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeMapping pulumi.StringMapInput
	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	Aws WorkloadIdentityPoolProviderAwsPtrInput
	// A description for the provider. Cannot exceed 256 characters.
	Description pulumi.StringPtrInput
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	Disabled pulumi.BoolPtrInput
	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrInput
	// The resource name of the provider as
	// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.
	Name pulumi.StringPtrInput
	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	Oidc WorkloadIdentityPoolProviderOidcPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The state of the provider. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The provider is active, and may be used to
	// validate authentication credentials. * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
	// deleted after approximately 30 days. You can restore a soft-deleted provider using UndeleteWorkloadIdentityPoolProvider.
	// You cannot reuse the ID of a soft-deleted provider until it is permanently deleted.
	State pulumi.StringPtrInput
	// The ID used for the pool, which is the final component of the pool resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId pulumi.StringPtrInput
	// The ID for the provider, which becomes the final component of the resource name. This
	// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolProviderId pulumi.StringPtrInput
}

func (WorkloadIdentityPoolProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadIdentityPoolProviderState)(nil)).Elem()
}

type workloadIdentityPoolProviderArgs struct {
	// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	// * `assertion`: JSON representing the authentication credential issued by the provider.
	// * `google`: The Google attributes mapped from the assertion in the `attributeMappings`.
	// * `attribute`: The custom attributes mapped from the assertion in the `attributeMappings`.
	//   The maximum length of the attribute condition expression is 4096 characters. If
	//   unspecified, all valid authentication credential are accepted.
	//   The following example shows how to only allow credentials with a mapped `google.groups`
	//   value of `admins`:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeCondition *string `pulumi:"attributeCondition"`
	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as `subject` and `segment`.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// * `google.subject`: The principal IAM is authenticating. You can reference this value
	//   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
	//   Cannot exceed 127 characters.
	// * `google.groups`: Groups the external identity belongs to. You can grant groups
	//   access to resources using an IAM `principalSet` binding; access applies to all
	//   members of the group.
	//   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
	//   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
	//   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
	//   is 100 characters, and the key may only contain the characters [a-z0-9_].
	//   You can reference these attributes in IAM policies to define fine-grained access for a
	//   workload to Google Cloud resources. For example:
	// * `google.subject`:
	//   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
	// * `google.groups`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
	// * `attribute.{custom_attribute}`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
	//   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	//   function that maps an identity provider credential to the normalized attribute specified
	//   by the corresponding map key.
	//   You can use the `assertion` keyword in the expression to access a JSON representation of
	//   the authentication credential issued by the provider.
	//   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	//   the total size of all mapped attributes must not exceed 8KB.
	//   For AWS providers, the following rules apply:
	// - If no attribute mapping is defined, the following default mapping applies:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// - If any custom attribute mappings are defined, they must include a mapping to the
	//   `google.subject` attribute.
	//   For OIDC providers, the following rules apply:
	// - Custom attribute mappings must be defined, and must include a mapping to the
	//   `google.subject` attribute. For example, the following maps the `sub` claim of the
	//   incoming credential to the `subject` attribute on a Google token.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeMapping map[string]string `pulumi:"attributeMapping"`
	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	Aws *WorkloadIdentityPoolProviderAws `pulumi:"aws"`
	// A description for the provider. Cannot exceed 256 characters.
	Description *string `pulumi:"description"`
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	Disabled *bool `pulumi:"disabled"`
	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName *string `pulumi:"displayName"`
	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	Oidc *WorkloadIdentityPoolProviderOidc `pulumi:"oidc"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The ID used for the pool, which is the final component of the pool resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId string `pulumi:"workloadIdentityPoolId"`
	// The ID for the provider, which becomes the final component of the resource name. This
	// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolProviderId string `pulumi:"workloadIdentityPoolProviderId"`
}

// The set of arguments for constructing a WorkloadIdentityPoolProvider resource.
type WorkloadIdentityPoolProviderArgs struct {
	// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
	// plain text, to restrict what otherwise valid authentication credentials issued by the
	// provider should not be accepted.
	// The expression must output a boolean representing whether to allow the federation.
	// The following keywords may be referenced in the expressions:
	// * `assertion`: JSON representing the authentication credential issued by the provider.
	// * `google`: The Google attributes mapped from the assertion in the `attributeMappings`.
	// * `attribute`: The custom attributes mapped from the assertion in the `attributeMappings`.
	//   The maximum length of the attribute condition expression is 4096 characters. If
	//   unspecified, all valid authentication credential are accepted.
	//   The following example shows how to only allow credentials with a mapped `google.groups`
	//   value of `admins`:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeCondition pulumi.StringPtrInput
	// Maps attributes from authentication credentials issued by an external identity provider
	// to Google Cloud attributes, such as `subject` and `segment`.
	// Each key must be a string specifying the Google Cloud IAM attribute to map to.
	// The following keys are supported:
	// * `google.subject`: The principal IAM is authenticating. You can reference this value
	//   in IAM bindings. This is also the subject that appears in Cloud Logging logs.
	//   Cannot exceed 127 characters.
	// * `google.groups`: Groups the external identity belongs to. You can grant groups
	//   access to resources using an IAM `principalSet` binding; access applies to all
	//   members of the group.
	//   You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
	//   where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
	//   define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
	//   is 100 characters, and the key may only contain the characters [a-z0-9_].
	//   You can reference these attributes in IAM policies to define fine-grained access for a
	//   workload to Google Cloud resources. For example:
	// * `google.subject`:
	//   `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
	// * `google.groups`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
	// * `attribute.{custom_attribute}`:
	//   `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
	//   Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
	//   function that maps an identity provider credential to the normalized attribute specified
	//   by the corresponding map key.
	//   You can use the `assertion` keyword in the expression to access a JSON representation of
	//   the authentication credential issued by the provider.
	//   The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
	//   the total size of all mapped attributes must not exceed 8KB.
	//   For AWS providers, the following rules apply:
	// - If no attribute mapping is defined, the following default mapping applies:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// - If any custom attribute mappings are defined, they must include a mapping to the
	//   `google.subject` attribute.
	//   For OIDC providers, the following rules apply:
	// - Custom attribute mappings must be defined, and must include a mapping to the
	//   `google.subject` attribute. For example, the following maps the `sub` claim of the
	//   incoming credential to the `subject` attribute on a Google token.
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	AttributeMapping pulumi.StringMapInput
	// An Amazon Web Services identity provider. Not compatible with the property oidc.
	// Structure is documented below.
	Aws WorkloadIdentityPoolProviderAwsPtrInput
	// A description for the provider. Cannot exceed 256 characters.
	Description pulumi.StringPtrInput
	// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
	// However, existing tokens still grant access.
	Disabled pulumi.BoolPtrInput
	// A display name for the provider. Cannot exceed 32 characters.
	DisplayName pulumi.StringPtrInput
	// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
	// Structure is documented below.
	Oidc WorkloadIdentityPoolProviderOidcPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The ID used for the pool, which is the final component of the pool resource name. This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolId pulumi.StringInput
	// The ID for the provider, which becomes the final component of the resource name. This
	// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// `gcp-` is reserved for use by Google, and may not be specified.
	WorkloadIdentityPoolProviderId pulumi.StringInput
}

func (WorkloadIdentityPoolProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadIdentityPoolProviderArgs)(nil)).Elem()
}

type WorkloadIdentityPoolProviderInput interface {
	pulumi.Input

	ToWorkloadIdentityPoolProviderOutput() WorkloadIdentityPoolProviderOutput
	ToWorkloadIdentityPoolProviderOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOutput
}

func (WorkloadIdentityPoolProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPoolProvider)(nil)).Elem()
}

func (i WorkloadIdentityPoolProvider) ToWorkloadIdentityPoolProviderOutput() WorkloadIdentityPoolProviderOutput {
	return i.ToWorkloadIdentityPoolProviderOutputWithContext(context.Background())
}

func (i WorkloadIdentityPoolProvider) ToWorkloadIdentityPoolProviderOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadIdentityPoolProviderOutput)
}

type WorkloadIdentityPoolProviderOutput struct {
	*pulumi.OutputState
}

func (WorkloadIdentityPoolProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadIdentityPoolProviderOutput)(nil)).Elem()
}

func (o WorkloadIdentityPoolProviderOutput) ToWorkloadIdentityPoolProviderOutput() WorkloadIdentityPoolProviderOutput {
	return o
}

func (o WorkloadIdentityPoolProviderOutput) ToWorkloadIdentityPoolProviderOutputWithContext(ctx context.Context) WorkloadIdentityPoolProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkloadIdentityPoolProviderOutput{})
}