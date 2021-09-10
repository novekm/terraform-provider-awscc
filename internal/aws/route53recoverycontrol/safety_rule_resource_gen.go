// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoverycontrol

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
	registry.AddResourceTypeFactory("awscc_route53recoverycontrol_safety_rule", safetyRuleResourceType)
}

// safetyRuleResourceType returns the Terraform awscc_route53recoverycontrol_safety_rule resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53RecoveryControl::SafetyRule resource type.
func safetyRuleResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"assertion_rule": {
			// Property: AssertionRule
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met. Otherwise, the change to the routing control is not accepted.",
			//   "properties": {
			//     "AssertedControls": {
			//       "description": "The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed. For example, you might include three routing controls, one for each of three AWS Regions.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "WaitPeriodMs": {
			//       "description": "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "AssertedControls",
			//     "WaitPeriodMs"
			//   ],
			//   "type": "object"
			// }
			Description: "An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met. Otherwise, the change to the routing control is not accepted.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"asserted_controls": {
						// Property: AssertedControls
						Description: "The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed. For example, you might include three routing controls, one for each of three AWS Regions.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
					},
					"wait_period_ms": {
						// Property: WaitPeriodMs
						Description: "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
						Type:        types.NumberType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"control_panel_arn": {
			// Property: ControlPanelArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the control panel.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the control panel.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // ControlPanelArn is a force-new property.
			},
		},
		"gating_rule": {
			// Property: GatingRule
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A gating rule verifies that a set of gating controls evaluates as true, based on a rule configuration that you specify. If the gating rule evaluates to true, Amazon Route 53 Application Recovery Controller allows a set of routing control state changes to run and complete against the set of target controls.",
			//   "properties": {
			//     "GatingControls": {
			//       "description": "The gating controls for the gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "TargetControls": {
			//       "description": "Routing controls that can only be set or unset if the specified RuleConfig evaluates to true for the specified GatingControls. For example, say you have three gating controls, one for each of three AWS Regions. Now you specify AtLeast 2 as your RuleConfig. With these settings, you can only change (set or unset) the routing controls that you have specified as TargetControls if that rule evaluates to true. \nIn other words, your ability to change the routing controls that you have specified as TargetControls is gated by the rule that you set for the routing controls in GatingControls.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "WaitPeriodMs": {
			//       "description": "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "WaitPeriodMs",
			//     "TargetControls",
			//     "GatingControls"
			//   ],
			//   "type": "object"
			// }
			Description: "A gating rule verifies that a set of gating controls evaluates as true, based on a rule configuration that you specify. If the gating rule evaluates to true, Amazon Route 53 Application Recovery Controller allows a set of routing control state changes to run and complete against the set of target controls.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"gating_controls": {
						// Property: GatingControls
						Description: "The gating controls for the gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
					},
					"target_controls": {
						// Property: TargetControls
						Description: "Routing controls that can only be set or unset if the specified RuleConfig evaluates to true for the specified GatingControls. For example, say you have three gating controls, one for each of three AWS Regions. Now you specify AtLeast 2 as your RuleConfig. With these settings, you can only change (set or unset) the routing controls that you have specified as TargetControls if that rule evaluates to true. \nIn other words, your ability to change the routing controls that you have specified as TargetControls is gated by the rule that you set for the routing controls in GatingControls.",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
					},
					"wait_period_ms": {
						// Property: WaitPeriodMs
						Description: "An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail. This helps prevent \"flapping\" of state. The wait period is 5000 ms by default, but you can choose a custom value.",
						Type:        types.NumberType,
						Required:    true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the safety rule.",
			//   "type": "string"
			// }
			Description: "The name for the safety rule.",
			Type:        types.StringType,
			Optional:    true,
		},
		"rule_config": {
			// Property: RuleConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The rule configuration for an assertion rule or gating rule. This is the criteria that you set for specific assertion controls (routing controls) or gating controls. This configuration specifies how many controls must be enabled after a transaction completes.",
			//   "properties": {
			//     "Inverted": {
			//       "description": "Logical negation of the rule. If the rule would usually evaluate true, it's evaluated as false, and vice versa.",
			//       "type": "boolean"
			//     },
			//     "Threshold": {
			//       "description": "The value of N, when you specify an ATLEAST rule type. That is, Threshold is the number of controls that must be set when you specify an ATLEAST type.",
			//       "type": "integer"
			//     },
			//     "Type": {
			//       "description": "A rule can be one of the following: ATLEAST, AND, or OR.",
			//       "enum": [
			//         "AND",
			//         "OR",
			//         "ATLEAST"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type",
			//     "Threshold",
			//     "Inverted"
			//   ],
			//   "type": "object"
			// }
			Description: "The rule configuration for an assertion rule or gating rule. This is the criteria that you set for specific assertion controls (routing controls) or gating controls. This configuration specifies how many controls must be enabled after a transaction completes.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"inverted": {
						// Property: Inverted
						Description: "Logical negation of the rule. If the rule would usually evaluate true, it's evaluated as false, and vice versa.",
						Type:        types.BoolType,
						Required:    true,
					},
					"threshold": {
						// Property: Threshold
						Description: "The value of N, when you specify an ATLEAST rule type. That is, Threshold is the number of controls that must be set when you specify an ATLEAST type.",
						Type:        types.NumberType,
						Required:    true,
					},
					"type": {
						// Property: Type
						Description: "A rule can be one of the following: ATLEAST, AND, or OR.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"AND",
								"OR",
								"ATLEAST",
							}),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // RuleConfig is a force-new property.
			},
		},
		"safety_rule_arn": {
			// Property: SafetyRuleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the safety rule.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the safety rule.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			//   "enum": [
			//     "PENDING",
			//     "DEPLOYED",
			//     "PENDING_DELETION"
			//   ],
			//   "type": "string"
			// }
			Description: "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS Route53 Recovery Control basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryControl::SafetyRule").WithTerraformTypeName("awscc_route53recoverycontrol_safety_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asserted_controls": "AssertedControls",
		"assertion_rule":    "AssertionRule",
		"control_panel_arn": "ControlPanelArn",
		"gating_controls":   "GatingControls",
		"gating_rule":       "GatingRule",
		"inverted":          "Inverted",
		"name":              "Name",
		"rule_config":       "RuleConfig",
		"safety_rule_arn":   "SafetyRuleArn",
		"status":            "Status",
		"target_controls":   "TargetControls",
		"threshold":         "Threshold",
		"type":              "Type",
		"wait_period_ms":    "WaitPeriodMs",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	opts = opts.WithRequiredAttributesValidators(validate.OneOfRequired(
		validate.Required(
			"assertion_rule",
			"name",
			"control_panel_arn",
			"rule_config",
		),
		validate.Required(
			"gating_rule",
			"name",
			"control_panel_arn",
			"rule_config",
		),
	),
	)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53recoverycontrol_safety_rule", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
