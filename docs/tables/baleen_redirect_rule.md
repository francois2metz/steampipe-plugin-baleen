# Table: baleen_redirect_rule

A redirect rule allow users to access moved content while maintaining the natural referencing of the pages.

Notes:

- List queries require a `namespace`.

## Examples

### List redirect rule of a namespace

```sql
select
  source,
  destination
from
  baleen_redirect_rule
where
  namespace='kfuAlneru9fjrG==';
```
