---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_network_performance_metric_subscription Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_ec2_network_performance_metric_subscription (Resource)

An example resource schema demonstrating some basic constructs and validation rules.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination` (String) The destination is a mandatory element for the metric subscription.
- `metric` (String) The metric type for the metric subscription.
- `source` (String) The source is a mandatory element for the metric subscription.
- `statistic` (String) The statistic type for the metric subscription.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_network_performance_metric_subscription.example <resource ID>
```