---
organization: francois2metz
category: ["saas"]
brand_color: "#0528f2"
display_name: "Baleen"
short_name: "baleen"
description: "Steampipe plugin for querying Baleen."
og_description: "Query Baleen with SQL! Open source CLI. No DB required."
icon_url: "/images/plugins/francois2metz/baleen.svg"
og_image: "/images/plugins/francois2metz/baleen-social-graphic.png"
---

# Baleen + Steampipe

[Baleen](https://baleen.cloud/) is a content delivery network and DDoS mitigation company.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  name,
  url
from
  baleen_namespace
```

```
+--------------------+-------------+----------------------+
| id                 | name        | url                  |
+--------------------+-------------+----------------------+
| HQSd02Tjhba3ue==   | Test        | https://example.net/ |
| c1x6H2wuyJArcwM==  | Plop        | https://example.com  |
+--------------------+-------------+----------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/baleen/tables)**

## Get started

### Install

Download and install the latest Baleen plugin:

```bash
steampipe plugin install francois2metz/baleen
```

### Configuration

Installing the latest baleen plugin will create a config file (`~/.steampipe/config/baleen.spc`) with a single connection named `baleen`:

```hcl
connection "baleen" {
    plugin = "francois2metz/baleen"

    # Personal access token
    # Ask the support to get it: https://support.baleen.cloud/hc/fr/articles/360017482439-G%C3%A9n%C3%A9ral-Utiliser-les-APIs
    token = "xxxxx-xxx-xxxx-xxxx-xxxx"
}
```

You can also use environment variables:

- `BALEEN_TOKEN` for the API token (ex: xxxxx-xxx-xxxx-xxxx-xxxx)

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-baleen
