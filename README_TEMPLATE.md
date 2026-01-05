# {{Algorithm Name}}

## 0) Interview framing (Google/FAANG)
**What they test:** {{core pattern}}, invariants, edge cases, and complexity.  
**What you should say out loud:**
- "I’m maintaining {{state/invariant}} while scanning {{input}} once.”
- "The key invariant is {{invariant}}.”
- "This avoids O(n²) by {{reason}}; complexity is {{O(...)}}.”

---

## 1) Overview
{{1–3 sentence description of the pattern/algorithm and what it enables.}}

## 2) Problem(s) it solves
- {{Problem A}}
- {{Problem B}}
- {{Problem C}}

## 3) Core idea & invariants
### Invariant(s)
- **Invariant #1:** {{...}}
- **Invariant #2:** {{...}}

### Why it works (high level)
{{Short correctness argument / intuition.}}

## 4) Typical formulations (interview prompts)
- {{Prompt 1}}
- {{Prompt 2}}
- {{Prompt 3}}

## 5) Complexity
- **Time:** {{...}}
- **Space:** {{...}}
- Notes: {{when it degrades / when constants matter}}

## 6) Go implementation notes (idiomatic)
- Prefer **small, explicit helpers** over clever abstractions.
- Use stdlib where it makes sense (`sort`, `container/heap`, etc.).
- Be careful with:
    - `nil` vs empty slices (tests may check both)
    - rune vs byte when dealing with strings (state your assumption)
    - mutation vs returning new slices (document it)

## 7) Common pitfalls (what interviewers look for)
1. {{Pitfall}}
2. {{Pitfall}}
3. {{Pitfall}}

## 8) Practice katas (remember the pattern)
- {{Problem}} (LeetCode #{{id}})
- {{Problem}} (LeetCode #{{id}})
- {{Problem}} (LeetCode #{{id}})

## 9) Variations you should recognize
- {{Variation A}}
- {{Variation B}}
- {{Variation C}}

## 10) Further reading
- {{Link 1}}
- {{Link 2}}
- {{Link 3}}

---

## Contract (your Go API)
Implement these functions/types in this package:

```go
// (Fill in the exact signatures used by the tests in this folder.)
