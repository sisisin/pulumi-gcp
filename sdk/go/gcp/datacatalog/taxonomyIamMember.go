// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} * {{project}}/{{region}}/{{taxonomy}} * {{region}}/{{taxonomy}} * {{taxonomy}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog taxonomy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/taxonomyIamMember:TaxonomyIamMember editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/taxonomyIamMember:TaxonomyIamMember editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/taxonomyIamMember:TaxonomyIamMember editor projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TaxonomyIamMember struct {
	pulumi.CustomResourceState

	Condition TaxonomyIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringOutput `pulumi:"taxonomy"`
}

// NewTaxonomyIamMember registers a new resource with the given unique name, arguments, and options.
func NewTaxonomyIamMember(ctx *pulumi.Context,
	name string, args *TaxonomyIamMemberArgs, opts ...pulumi.ResourceOption) (*TaxonomyIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Taxonomy == nil {
		return nil, errors.New("invalid value for required argument 'Taxonomy'")
	}
	var resource TaxonomyIamMember
	err := ctx.RegisterResource("gcp:datacatalog/taxonomyIamMember:TaxonomyIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxonomyIamMember gets an existing TaxonomyIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxonomyIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxonomyIamMemberState, opts ...pulumi.ResourceOption) (*TaxonomyIamMember, error) {
	var resource TaxonomyIamMember
	err := ctx.ReadResource("gcp:datacatalog/taxonomyIamMember:TaxonomyIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxonomyIamMember resources.
type taxonomyIamMemberState struct {
	Condition *TaxonomyIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy *string `pulumi:"taxonomy"`
}

type TaxonomyIamMemberState struct {
	Condition TaxonomyIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringPtrInput
}

func (TaxonomyIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamMemberState)(nil)).Elem()
}

type taxonomyIamMemberArgs struct {
	Condition *TaxonomyIamMemberCondition `pulumi:"condition"`
	Member    string                      `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy string `pulumi:"taxonomy"`
}

// The set of arguments for constructing a TaxonomyIamMember resource.
type TaxonomyIamMemberArgs struct {
	Condition TaxonomyIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `datacatalog.TaxonomyIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringInput
}

func (TaxonomyIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamMemberArgs)(nil)).Elem()
}

type TaxonomyIamMemberInput interface {
	pulumi.Input

	ToTaxonomyIamMemberOutput() TaxonomyIamMemberOutput
	ToTaxonomyIamMemberOutputWithContext(ctx context.Context) TaxonomyIamMemberOutput
}

func (TaxonomyIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyIamMember)(nil)).Elem()
}

func (i TaxonomyIamMember) ToTaxonomyIamMemberOutput() TaxonomyIamMemberOutput {
	return i.ToTaxonomyIamMemberOutputWithContext(context.Background())
}

func (i TaxonomyIamMember) ToTaxonomyIamMemberOutputWithContext(ctx context.Context) TaxonomyIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamMemberOutput)
}

type TaxonomyIamMemberOutput struct {
	*pulumi.OutputState
}

func (TaxonomyIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyIamMemberOutput)(nil)).Elem()
}

func (o TaxonomyIamMemberOutput) ToTaxonomyIamMemberOutput() TaxonomyIamMemberOutput {
	return o
}

func (o TaxonomyIamMemberOutput) ToTaxonomyIamMemberOutputWithContext(ctx context.Context) TaxonomyIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TaxonomyIamMemberOutput{})
}