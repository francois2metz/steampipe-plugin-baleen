# Table: baleen_namespace

A namespace is a proxy with a configured origin and specific rules.

## Examples

### List namespaces

```sql
select
  id,
  name,
  url
from
  baleen_namespace;
```

### Get namespaces with the WAF enabled

```sql
select
  id,
  name
from
  baleen_namespace
where
  waf_enabled;
```

### Get namespaces with the header X-Frame-Options: deny no send

```sql
select
  id,
  name
from
  baleen_namespace
where
  not headers_deny_frame_options;
```
