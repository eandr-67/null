# Go nullable generic types without pointers

The null package implements the generic nullable type null.Null\[any].

The null.Null type is based on the generic type sql.Null, but has a private implementation.
It is API and configuration oriented, and provides serialization/deserialization of JSON and YAML (the YAML implementation uses the gopkg.in/yaml.v3 package).

The equality check is implemented not by a method, but by the Equal function, which supports a subset of the null.Null[comparable] types.
