// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_globalaccelerator_listener", listenerResource)
}

// listenerResource returns the Terraform awscc_globalaccelerator_listener resource.
// This Terraform resource corresponds to the CloudFormation AWS::GlobalAccelerator::Listener resource.
func listenerResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AcceleratorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the accelerator.",
		//	  "type": "string"
		//	}
		"accelerator_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the accelerator.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClientAffinity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "NONE",
		//	  "description": "Client affinity lets you direct all requests from a user to the same endpoint.",
		//	  "enum": [
		//	    "NONE",
		//	    "SOURCE_IP"
		//	  ],
		//	  "type": "string"
		//	}
		"client_affinity": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Client affinity lets you direct all requests from a user to the same endpoint.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"NONE",
					"SOURCE_IP",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("NONE"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ListenerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the listener.",
		//	  "type": "string"
		//	}
		"listener_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the listener.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PortRanges
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A port range to support for connections from  clients to your accelerator.",
		//	    "properties": {
		//	      "FromPort": {
		//	        "description": "A network port number",
		//	        "maximum": 65535,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "ToPort": {
		//	        "description": "A network port number",
		//	        "maximum": 65535,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "FromPort",
		//	      "ToPort"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"port_ranges": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FromPort
					"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "A network port number",
						Required:    true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 65535),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: ToPort
					"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "A network port number",
						Required:    true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 65535),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Protocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "TCP",
		//	  "description": "The protocol for the listener.",
		//	  "enum": [
		//	    "TCP",
		//	    "UDP"
		//	  ],
		//	  "type": "string"
		//	}
		"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The protocol for the listener.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"TCP",
					"UDP",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("TCP"),
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::GlobalAccelerator::Listener",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::Listener").WithTerraformTypeName("awscc_globalaccelerator_listener")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accelerator_arn": "AcceleratorArn",
		"client_affinity": "ClientAffinity",
		"from_port":       "FromPort",
		"listener_arn":    "ListenerArn",
		"port_ranges":     "PortRanges",
		"protocol":        "Protocol",
		"to_port":         "ToPort",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
