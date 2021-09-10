---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudwatch_metric_stream Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudWatch::MetricStream
---

# awscc_cloudwatch_metric_stream (Data Source)

Data Source schema for AWS::CloudWatch::MetricStream



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) Amazon Resource Name of the metric stream.
- **creation_date** (String) The date of creation of the metric stream.
- **exclude_filters** (Attributes List) Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null. (see [below for nested schema](#nestedatt--exclude_filters))
- **firehose_arn** (String) The ARN of the Kinesis Firehose where to stream the data.
- **include_filters** (Attributes List) Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null. (see [below for nested schema](#nestedatt--include_filters))
- **last_update_date** (String) The date of the last update of the metric stream.
- **name** (String) Name of the metric stream.
- **output_format** (String) The output format of the data streamed to the Kinesis Firehose.
- **role_arn** (String) The ARN of the role that provides access to the Kinesis Firehose.
- **state** (String) Displays the state of the Metric Stream.
- **tags** (Attributes List) A set of tags to assign to the delivery stream. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--exclude_filters"></a>
### Nested Schema for `exclude_filters`

Read-Only:

- **namespace** (String) Only metrics with Namespace matching this value will be streamed.


<a id="nestedatt--include_filters"></a>
### Nested Schema for `include_filters`

Read-Only:

- **namespace** (String) Only metrics with Namespace matching this value will be streamed.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String) A unique identifier for the tag.
- **value** (String) An optional string, which you can use to describe or define the tag.

