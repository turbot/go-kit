

## v1.0.0 [2025-02-06]
_Breaking changes_
* Remove `MergeMaps`
* Remove `MergeStringMaps`
* Remove `CloneStringMap`

_What's new?_
* Add `helpers.AppendSliceUnique`. ([#86](https://github.com/turbot/go-kit/issues/86))
* Add ByteMapToStringMap
 
 
## v0.9.0 [2023-12-18]
_What's new?_
* Add `ListFilesWithContext` function to handle context cancellations.([#70](https://github.com/turbot/go-kit/issues/70))
* Add `AnySliceToTypedSlice` and `ToTypedSlice` functions.([#78](https://github.com/turbot/go-kit/issues/78))
* Add `TrimBlankLine` function.([#79](https://github.com/turbot/go-kit/issues/79))
* Add StringSliceEqualIgnoreOrder function to handle string slices comparison ignoring order. ([#741](https://github.com/turbot/go-kit/issues/74))
* Add type check functions.

## v0.8.1 [2023-09-14]
_Bug fixes_
* Fixes issue with `RotatingLogWriter` where it panics when starting up.([#67](https://github.com/turbot/go-kit/issues/67))

## v0.8.0 [2023-09-13]
_What's new?_
* Adds `FilterMap`.([#59](https://github.com/turbot/go-kit/issues/59))
* Adds `RotatingLogWriter`.([#63](https://github.com/turbot/go-kit/issues/63))

_Bug fixes_
* Remove `UpgradeRWMutex`, which had a potential race condition.([#60](https://github.com/turbot/go-kit/issues/60))

## v0.7.0 [2023-06-08]
_What's new?_ 
* Adds flag in 'ListFiles' to skip over empty directories. ([#53](https://github.com/turbot/go-kit/issues/53))

## v0.6.0 [2023-06-05]
_What's new?_
* Add `UpgradeRWMutex`. ([#42](https://github.com/turbot/go-kit/issues/42))
* Optimise list files to recursively traverse directories only if there's a match.  ([#49](https://github.com/turbot/go-kit/issues/49))
* Add utilities to work with golang maps. ([#50](https://github.com/turbot/go-kit/issues/50))
* Add MaxResults to ListOptions type to limit number of files listing. ([#45](https://github.com/turbot/go-kit/issues/45))
* 
## v0.5.0 [2022-11-30]
_What's new?_
* Add file watcher. ([#26](https://github.com/turbot/go-kit/issues/26))
* Adds helper functions for strings and hashes. ([#29](https://github.com/turbot/go-kit/issues/29))
* Add DirectoryExists.
* Add GlobRoot.
* Add SplitPath.
* Deprecate `Tildefy` and `FileExists` in `helpers` package and add to `files` package.

## v0.4.0 [2022-06-21]
_What's new?_
* Update `RemoveFromStringSlice` to accepts multiple values to remove. Add additional string slice function unit tests.  ([#24](https://github.com/turbot/go-kit/issues/24))
* Add `LintName` function to helpers.  ([#15](https://github.com/turbot/go-kit/issues/15))
* Add `EscapePropertyName` and `UnescapePropertyName` functions to escape dots in property names. ([#21](https://github.com/turbot/go-kit/issues/21))

_Bug fixes_
* Fix `GetNestedFieldValueFromInterface` returning nil, if a column name contains dot. (Update `GetFieldValueFromInterface` and `GetNestedFieldValueFromInterface` to use `UnescapePropertyName` to support names with dots.) ([#23](https://github.com/turbot/go-kit/issues/23))
* Fix StringSliceDistinct not returning all the distinct items, instead removing the duplicate one from output. ([#20](https://github.com/turbot/go-kit/issues/20))
* 
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
