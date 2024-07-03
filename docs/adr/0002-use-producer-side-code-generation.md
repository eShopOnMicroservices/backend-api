# 2. Use producer side code generation

Date: 2024-07-03

## Status

Accepted

## Context

This repository is effectively a monorepo with various versioning units.
The only consumers of these versioning units are internal backend services.

## Decision

Use producer side code generation.

## Rationale

1. **Consistency**: This repository will be a single source of
   truth for generated code.

2. **Breaking change detection in CI**: Simply won't be possible to detect them in
   generated code in case of consumer side generation.

## Implications

1. **Source code dependency resolution**: Since we have generated code in this
   repository, we have to provide various approaches for source code dependency
   resolution.
