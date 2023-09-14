---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_pcaconnectorad_connector Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::PCAConnectorAD::Connector
---

# awscc_pcaconnectorad_connector (Data Source)

Data Source schema for AWS::PCAConnectorAD::Connector



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `certificate_authority_arn` (String)
- `connector_arn` (String)
- `directory_id` (String)
- `tags` (Map of String)
- `vpc_information` (Attributes) (see [below for nested schema](#nestedatt--vpc_information))

<a id="nestedatt--vpc_information"></a>
### Nested Schema for `vpc_information`

Read-Only:

- `security_group_ids` (List of String)