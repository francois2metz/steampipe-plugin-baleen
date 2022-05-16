# Table: baleen_owasp

List Open Web Application Security Project rules.

Notes:

- List queries require a `namespace`.

## Examples

### List current status of owasp rules

```sql
select
  name, enabled
from
  baleen_owasp
where
  namespace='kfuAlneru9fjrG==';
```

### Get disabled rules

```sql
select
  name, enabled
from
  baleen_owasp
where
  namespace='kfuAlneru9fjrG=='
  and not enabled;
```
