---
logo:
    alt: 'errpath logo: golang gopher determinedly walking through a blue and red maze'
    source: logo.jpg
    width: 300
---

![Code Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)


Package errpath provides utilities for creating and managing detailed error paths.
It allows users to construct error messages that include the full path to the error,
which can be particularly useful when traversing complex data structures such as JSON
or YAML files.

Example for an error in an OpenAPI: `components.schemas["Pet"].allOf[0]: invalid schema`

The package defines several error types that can be used to represent different kinds
of errors, such as missing required values, invalid values, and errors occurring at
specific fields, indices, or keys within a data structure. These error types implement
a chaining mechanism that builds a detailed error path.
