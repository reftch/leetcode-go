# leetcode-go

Go skeleton for LeetCode practice. Problems are organized by number with
stubbed solutions only — no implementations.

## Layout

```text
internal/ds/            # shared ListNode, TreeNode helpers + tests
problems/0001-two-sum/  # example stub: solution.go, solution_test.go, README.md
scripts/new-problem.sh  # scaffolding script (no solution logic)
```

Each problem directory contains:

- `solution.go` — stubbed function that panics with `not implemented`
- `solution_test.go` — table-driven test that fails until implemented
- `README.md` — problem link + checklist

## Usage

```sh
go test ./...
go vet ./...
make test PKG=./problems/0001-two-sum/...

# Scaffold a new problem (stub only):
make new-problem NUM=0002 SLUG=add-two-numbers PKG=addtwonumbers FUNC=AddTwoNumbers
```

## Conventions

- Directory: `problems/<zero-padded-number>-<slug>/`
- Package: lowercase, no leading digits (e.g. `twosum`)
- Tests are table-driven and fail (`t.Fatal` / stub panic) until implemented
- Shared types live in `internal/ds`; do not redefine `ListNode`/`TreeNode` per problem
