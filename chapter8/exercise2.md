# What's the difference between an identifier that starts with a capital letter and one that doesn't

Identifiers in a package that begin with a capital letter are exported from the package,
the rest of the identifiers defined in the package are not. An exported identifier is
accessible with dot notation through the package name in any other package that imports it.
