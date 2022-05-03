# Table: baleen_custom_static_rule

A static rule allow to classify traffic with some conditions.

Notes:

- List queries require a `namespace`.

## Examples

### List static rule of a namespace

```sql
select
  id,
  category,
  description
from
  baleen_namespace_custom_static_rule
where
  namespace='kfuAlneru9fjrG==';
```

### List static rules of all namespaces

```sql
select
  n.name,
  csr.id,
  csr.category,
  csr.description
from
  baleen_custom_static_rule csr
join
  baleen_namespace n on csr.namespace = n.id;
```
