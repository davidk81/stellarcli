## stellar satellite list-passes

Lists available passes of a satellite.

### Synopsis

Lists available passes of a satellite.

```
stellar satellite list-passes [Satellite ID] [flags]
```

### Options

```
  -h, --help                  help for list-passes
      --min-elevation float   The minimum elevation of passes. Passes are listed having max elevation greater than the minimum elevation. (default 10)
  -o, --output string         Output format. One of: csv|wide|json (default "wide")
  -v, --verbose               Output more information in JSON format. (default false)
```

### SEE ALSO

* [stellar satellite](stellar_satellite.md)	 - Commands for working with satellites

