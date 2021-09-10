// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iot_domain_configuration", domainConfigurationResourceType)
}

// domainConfigurationResourceType returns the Terraform awscc_iot_domain_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::DomainConfiguration resource type.
func domainConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"authorizer_config": {
			// Property: AuthorizerConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AllowAuthorizerOverride": {
			//       "type": "boolean"
			//     },
			//     "DefaultAuthorizerName": {
			//       "maxLength": 128,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allow_authorizer_override": {
						// Property: AllowAuthorizerOverride
						Type:     types.BoolType,
						Optional: true,
					},
					"default_authorizer_name": {
						// Property: DefaultAuthorizerName
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
				},
			),
			Optional: true,
		},
		"domain_configuration_name": {
			// Property: DomainConfigurationName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // DomainConfigurationName is a force-new property.
			},
		},
		"domain_configuration_status": {
			// Property: DomainConfigurationStatus
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ENABLED",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ENABLED",
					"DISABLED",
				}),
			},
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 253,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 253),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // DomainName is a force-new property.
			},
		},
		"domain_type": {
			// Property: DomainType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ENDPOINT",
			//     "AWS_MANAGED",
			//     "CUSTOMER_MANAGED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"server_certificate_arns": {
			// Property: ServerCertificateArns
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 2048,
			//     "minLength": 1,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "maxItems": 1,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // ServerCertificateArns is a force-new property.
			},
			// ServerCertificateArns is a write-only property.
		},
		"server_certificates": {
			// Property: ServerCertificates
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ServerCertificateArn": {
			//         "maxLength": 2048,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ServerCertificateStatus": {
			//         "enum": [
			//           "INVALID",
			//           "VALID"
			//         ],
			//         "type": "string"
			//       },
			//       "ServerCertificateStatusDetail": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"server_certificate_arn": {
						// Property: ServerCertificateArn
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
					"server_certificate_status": {
						// Property: ServerCertificateStatus
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"INVALID",
								"VALID",
							}),
						},
					},
					"server_certificate_status_detail": {
						// Property: ServerCertificateStatusDetail
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"service_type": {
			// Property: ServiceType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DATA",
			//     "CREDENTIAL_PROVIDER",
			//     "JOBS"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DATA",
					"CREDENTIAL_PROVIDER",
					"JOBS",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // ServiceType is a force-new property.
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"validation_certificate_arn": {
			// Property: ValidationCertificateArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // ValidationCertificateArn is a force-new property.
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Create and manage a Domain Configuration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::DomainConfiguration").WithTerraformTypeName("awscc_iot_domain_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_authorizer_override":        "AllowAuthorizerOverride",
		"arn":                              "Arn",
		"authorizer_config":                "AuthorizerConfig",
		"default_authorizer_name":          "DefaultAuthorizerName",
		"domain_configuration_name":        "DomainConfigurationName",
		"domain_configuration_status":      "DomainConfigurationStatus",
		"domain_name":                      "DomainName",
		"domain_type":                      "DomainType",
		"key":                              "Key",
		"server_certificate_arn":           "ServerCertificateArn",
		"server_certificate_arns":          "ServerCertificateArns",
		"server_certificate_status":        "ServerCertificateStatus",
		"server_certificate_status_detail": "ServerCertificateStatusDetail",
		"server_certificates":              "ServerCertificates",
		"service_type":                     "ServiceType",
		"tags":                             "Tags",
		"validation_certificate_arn":       "ValidationCertificateArn",
		"value":                            "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ServerCertificateArns",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iot_domain_configuration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}