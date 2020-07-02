# mergo

Prior to mergo v0.3.9 merging of fileds that had a different type worked.

Ex.
config := `foo: 1`
default := `foo: "1"`

In v0.3.9, this returns and error and fails to merge the fields.  The merge options
WithTypeCheck default to false, so this behavior is unexpected.
