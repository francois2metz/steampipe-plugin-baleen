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
  baleen_custom_static_rule
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

### List enabled static rules of a namespace

```sql
select
  id,
  category,
  description
from
  baleen_custom_static_rule
where
  namespace='kfuAlneru9fjrG=='
  and enabled;
```

### List static custom that block traffic

```sql
select
  id,
  category,
  description
from
  baleen_custom_static_rule
where
  namespace='kfuAlneru9fjrG=='
  and category='block';
```
