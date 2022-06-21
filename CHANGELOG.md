## v0.4.0 [2022-06-21]
_What's new?_
* Update `RemoveFromStringSlice` to accepts multiple values to remove. Add additional string slice function unit tests.  ([#24](https://github.com/turbot/go-kit/issues/24))
* Add `LintName` function to helpers.  ([#15](https://github.com/turbot/go-kit/issues/15))
* Add `EscapePropertyName` and `UnescapePropertyName` functions to escape dots in property names. ([#21](https://github.com/turbot/go-kit/issues/21))

_Bug fixes_
* Fix `GetNestedFieldValueFromInterface` returning nil, if a column name contains dot. (Update `GetFieldValueFromInterface` and `GetNestedFieldValueFromInterface` to use `UnescapePropertyName` to support names with dots.) ([#23](https://github.com/turbot/go-kit/issues/23))

* ## v0.3.0 [2021-10-08]
_What's new?_
* Add `CombineErrors` and `CombineErrorsWithPrefix`. ([#12](https://github.com/turbot/go-kit/issues/12)) 
* Add `IsNil`. ([#18](https://github.com/turbot/go-kit/issues/18))
* Add `TabifyStringSlice`. ([#13](https://github.com/turbot/go-kit/issues/13)) 
* Add `StringSliceDistinct` and `StringSliceHasDuplicates`. ([#14](https://github.com/turbot/go-kit/issues/14)) 
* Update `TruncateString` to take newlines into account. ([#16](https://github.com/turbot/go-kit/issues/16)) 

## v0.2.1 [2021-05-14]
* Fix length bug in `TruncateString`.


## v0.2.0 [2021-05-13]
_What's new?_
* Add `ListFiles`. ([#9](https://github.com/turbot/go-kit/issues/9)) 
* Add `Tabify`. ([#9](https://github.com/turbot/go-kit/issues/10)) 
* Add `TruncateString`. ([#9](https://github.com/turbot/go-kit/issues/7)) 

_Bug fixes_
* Handle null pointers in `ToBool`
 
## v0.1.3 [2021-03-17]

_Bug fixes_
* Handle null pointers in `ToBool`

## v0.1.2 [2021-03-15]

_What's new?_
* Add `ToBoolPtr` to convert interface to a bool pointer. If the value is nil, or there is a conversion error, return nil. ([#3](https://github.com/turbot/go-kit/issues/3))
  
## v0.1.1 [2021-02-11]

_What's new?_
* Add type conversion function `CastString`, which casts an interface to a string. This is the same as `SafeString` except it also returns a success flag.