# Entities for data management in GO

**This is early work**

The idea is simple enough -- design some entities -- structs in Go -- with common
names, common slots, and a common pointer architecture to replicate the kinds of data
needs of the Semantic Web.  Provide methods for loading and unloading the structures
to triples -- this allows for storage and query in triple stores nd the use of the SPARQL
query language for reporting.

The Go data structures are simple, extensible, and fast. Should be useful for many data
problems, and the construction of CRM systems, academic profiling systems such as VIVO,
and many other common data applications.

