# Table: baleen_rewrite_rule

A rewrite rule allow users to continue to access moved content seamlessly.

Notes:

- List queries require a `namespace`.

## Examples

### List rewrite rule of a namespace

```sql
select
  source,
  destination
from
  baleen_rewrite_rule
where
  namespace='kfuAlneru9fjrG==';
```

### List rewrite rules that rewrite /document URLs

```sql
select
  source,
  destination
from
  baleen_rewrite_rule
where
  namespace='kfuAlneru9fjrG=='
  and source LIKE '/document%';
```
