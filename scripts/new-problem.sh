#!/bin/sh
# Scaffold a new LeetCode problem directory without a solution implementation.
# Usage: ./scripts/new-problem.sh <number> <slug> <package> <FuncName>
# Example: ./scripts/new-problem.sh 0002 add-two-numbers addtwonumbers AddTwoNumbers
set -eu

if [ "$#" -ne 4 ]; then
  echo "usage: $0 <number> <slug> <package> <FuncName>" >&2
  echo "example: $0 0002 add-two-numbers addtwonumbers AddTwoNumbers" >&2
  exit 1
fi

NUM="$1"
SLUG="$2"
PKG="$3"
FUNC="$4"
DIR="problems/${NUM}-${SLUG}"

if [ -e "$DIR" ]; then
  echo "error: $DIR already exists" >&2
  exit 1
fi

mkdir -p "$DIR"

cat > "$DIR/solution.go" <<EOF
// Package ${PKG} contains a stub for LeetCode problem ${NUM}. ${FUNC}.
//
// Link: https://leetcode.com/problems/${SLUG}/
//
// Replace this comment with a short summary of the approach once implemented.
package ${PKG}

// ${FUNC} implements the solution.
// TODO: fix the signature to match the LeetCode statement, then implement.
func ${FUNC}() {
	panic("not implemented: ${FUNC}")
}
EOF

cat > "$DIR/solution_test.go" <<EOF
package ${PKG}

import "testing"

func Test${FUNC}(t *testing.T) {
	// TODO: replace with real table-driven cases.
	t.Fatal("solution not implemented")
}
EOF

cat > "$DIR/README.md" <<EOF
${NUM}. ${FUNC}

Link: https://leetcode.com/problems/${SLUG}/

## Notes

Describe the approach here after implementing \`solution.go\`.

## Checklist

- [ ] Fix the function signature in \`solution.go\` to match LeetCode
- [ ] Implement \`${FUNC}\`
- [ ] Add table-driven cases in \`solution_test.go\` (replace \`t.Fatal\` stub)
- [ ] Run \`make test PKG=./${DIR}/...\`
EOF

echo "created $DIR"
