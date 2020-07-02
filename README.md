# mergo

Prior to mergo v0.3.9 merging of fields that had a different types worked.

Ex.
```
config := `foo: 1`
default := `foo: "1"`
```

In v0.3.9, this returns and error and fails to merge the fields.  The merge options
WithTypeCheck default to false, so this behavior is unexpected.

This repo demostrates the issue.  In go.mod, switch between v0.3.9 and earlier versions
like v0.3.8.
