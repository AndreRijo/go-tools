# Go Tools

A collection of useful tools for any Go-based project.
As of now, the only tool included is a tool to read configuration files in a folder.

## ConfigLoader
Fully implemented in configLoader.go. Given a folder, reads all text files with a `.cfg` extension (other extensions can be read by changing the `configType` variable).

Each setting should be written in its own line, with the format `name = value`.
Empty lines are ignored.
Comments can be added by prefixing the line with `#` or `//`.

Example of a configuration file:

```
#Connection settings
localServer = true
serverIP = 192.168.40.112:8432

//Test settings
nInstances = 24
updateRate = 0.8
queryList = 1 2 3 4 5
```

Multiple methods to facilitate configuration reading are provided. 
It is possible to read directly the fields with their intended type (e.g., `GetInt64Config`). 
Methods for most basic types are included.

Furthermore, it is possible to provide a default value to be returned when the key is not present. 
E.g., `GetInt64Config("myKey", -1)` will return the value associated to "myKey" if present, otherwise returns -1 (in `int64` form).

GetConfig, GetAndHasConfig and GetOrDefault can be used for values of any type. 
They return the value in string form. 
GetAndHasConfig also returns a boolean representing if the key is present or not.
Finally, GetOrDefault receives an extra argument corresponding to the default value to return if the key isn't present.