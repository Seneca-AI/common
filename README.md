# common

This repository holds definitions, configurations and types that are used across all Seneca repositories.

## Python Users Please Read
It seems that the protobuf compiler does not add "from . import..." in the generated python code, even though the import in to .proto file specifies that the import is in the same directory.  Not sure why this happens, not interested in learning enough python to fix it properly.  The solution (which you can see in the GitHub Actions) is to just do a find and replace: `s/^import \(.*\) as/from . import \1 as/g`.  This will probably cause someone a big headache one day, so note this.

## Notes

* Do not edit any files in proto_out/
* Do not remove any fields from protos, simple mark them as deprecated.