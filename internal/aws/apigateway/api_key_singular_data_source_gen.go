// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_apigateway_api_key", apiKeyDataSourceType)
}

// apiKeyDataSourceType returns the Terraform awscc_apigateway_api_key data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ApiGateway::ApiKey resource type.
func apiKeyDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"api_key_id": {
			// Property: APIKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs ",
			//   "type": "string"
			// }
			Description: "A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs ",
			Type:        types.StringType,
			Computed:    true,
		},
		"customer_id": {
			// Property: CustomerId
			// CloudFormation resource type schema:
			// {
			//   "description": "An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.",
			//   "type": "string"
			// }
			Description: "An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the purpose of the API key.",
			//   "type": "string"
			// }
			Description: "A description of the purpose of the API key.",
			Type:        types.StringType,
			Computed:    true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the API key can be used by clients.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the API key can be used by clients.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"generate_distinct_id": {
			// Property: GenerateDistinctId
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.",
			//   "type": "string"
			// }
			Description: "A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"stage_keys": {
			// Property: StageKeys
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of stages to associate with this API key.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "RestApiId": {
			//         "description": "The ID of a RestApi resource that includes the stage with which you want to associate the API key.",
			//         "type": "string"
			//       },
			//       "StageName": {
			//         "description": "The name of the stage with which to associate the API key. The stage must be included in the RestApi resource that you specified in the RestApiId property. ",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of stages to associate with this API key.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"rest_api_id": {
						// Property: RestApiId
						Description: "The ID of a RestApi resource that includes the stage with which you want to associate the API key.",
						Type:        types.StringType,
						Computed:    true,
					},
					"stage_name": {
						// Property: StageName
						Description: "The name of the stage with which to associate the API key. The stage must be included in the RestApi resource that you specified in the RestApiId property. ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of arbitrary tags (key-value pairs) to associate with the API key.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of arbitrary tags (key-value pairs) to associate with the API key.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"value": {
			// Property: Value
			// CloudFormation resource type schema:
			// {
			//   "description": "The value of the API key. Must be at least 20 characters long.",
			//   "type": "string"
			// }
			Description: "The value of the API key. Must be at least 20 characters long.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ApiGateway::ApiKey",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::ApiKey").WithTerraformTypeName("awscc_apigateway_api_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_key_id":           "APIKeyId",
		"customer_id":          "CustomerId",
		"description":          "Description",
		"enabled":              "Enabled",
		"generate_distinct_id": "GenerateDistinctId",
		"key":                  "Key",
		"name":                 "Name",
		"rest_api_id":          "RestApiId",
		"stage_keys":           "StageKeys",
		"stage_name":           "StageName",
		"tags":                 "Tags",
		"value":                "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_apigateway_api_key", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}