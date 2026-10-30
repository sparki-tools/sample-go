# sample-go

A known-good sample for the **Go** starter on [Sparki](https://sparki.tools):
**go build ./... → go test ./... on golang:1.22**. Zero dependencies, go 1.22 in go.mod, and standard library tests.

## Try it

1. In Sparki, open **Projects → New Project** and add `https://github.com/sparki-tools/sample-go` on branch `main`.
2. On the project, create a pipeline from the **Go** starter.
3. Trigger a build. It should go green.

To see a red build, use branch `fail` instead. It has the same code with one assertion
that cannot hold.

## Why it exists

Third-party repositories change under us: a lockfile is dropped, a Node or Go version moves,
a lint hook is added. This repo is owned by Sparki, has no dependencies, and stays inside
the sandbox's limits: non-root user, read-only root filesystem, writable `/workspace` and
`/tmp`, and a 10-minute step timeout. A first build on it shows the platform working, not
someone else's repo.
