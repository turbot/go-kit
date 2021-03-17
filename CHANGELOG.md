## v0.1.3 [2021-03-17]

_Bug fixes?_
* Handle null pointers in ToBool

## v0.1.2 [2021-03-15]

_What's new?_
* Add ToBoolPtr to convert interface to a bool pointer. If the value is nil, or there is a conversion error, return nil. [#3](https://github.com/turbot/go-kit/issues/3).
  
## v0.1.1 [2021-02-11]

_What's new?_
* Add type conversion function CastString, which casts an interface to a string. This is the same as SafeString except it also returns a success flag.