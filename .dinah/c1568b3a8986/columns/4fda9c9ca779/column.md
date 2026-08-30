---
title: Implement
slug: implement
kind: work
---
Where specs become code in the DinahVisor codebase, a Wails desktop application importing dinah as a Go library. An implementing agent works in a worktree under C:\dinah-scratch on a branch named after the card, self-tests, pushes the branch, opens its pull request, then moves the card on. The trunk does not see this work until Merge.

### What arrives, and what the clarity gate asks

Read the card body, the spec attachment, the checklist and the latest comments before writing code. What the gate demands depends on how the card got here.

**A card that came through Spec arrives with a reviewed contract.** If that contract is ambiguous, missing or self-contradictory, do not start: push back one column with a comment naming the ambiguity. If the ambiguity is a ruling only the operator can give and no work can proceed without it, block in place with `dinah block <card> --kind operator_decision --reason "<the call>"` instead.

**A card routed straight from Triage arrives with no spec, by design.** That is the fast path, and its contract is meant to be obvious from the card alone. So the gate here is different: ask whether you can state, in one sentence, exactly what is changing and how, with no judgement calls left. If you can, build it. If you cannot, the routing call was wrong rather than the card being unclear, so move it to Spec with a comment naming what a spec would have to settle. Re-routing a misclassified card to the head of its proper path is the one sanctioned exception to the one-column-back rule; never push a spec-less card back where it came from, because nothing there will change.

Don't reinterpret silently and don't expand scope: a dependency the spec missed becomes a new card, filed into Intake with the same workstreams and with the relationship recorded in prose on both cards, and a problem bigger than this card's scope is filed separately rather than fixed in the same diff.

### A question for the operator that the work can survive

Not every ruling has to stop the build. When a question is genuinely the operator's but you can finish the card despite it, meaning you made a defensible call and want it ratified or redirected, file it as an open question with `owner: operator`, put your call and its tradeoffs in the note, and carry on. It gets answered at Operator Code Review, two stages ahead; name it in your handoff comment so the reviewer and then the operator see it coming. A question you cannot build around blocks in place, per the gate above.

## Delivering an implementation means delivering a test

This is not negotiable and it is not somebody else's stage. A change that ships without a test covering the behaviour it changed is an incomplete implementation. Writing the test is part of the work you were asked to do, not an addition to it.

If a test already covers the behaviour you changed, extending it satisfies this and is the preferred answer.

**Integrate, do not accumulate.** When an existing test covers the area you touched, extend that test. Create a new test file only when the behaviour genuinely has no home in the suite. A regression suite that grows by accretion becomes slow, then redundant, then something people skip, and a suite people skip protects nothing. Name in WHAT SHIPPED which existing test you extended, or say why a new one was necessary.

**Arm the test.** A test that passes proves nothing on its own, because it also passes when the thing it guards is absent. Break what you fixed, watch the test go red, restore from a byte-identical backup, and confirm green again. Report that you did it. This is how a test earns the right to be trusted later by somebody who was not here.

## Prose ships under the prose standard

Any prose this card produces (README and docs changes, user-facing copy, error text with sentences in it, and the WHAT SHIPPED note) is written to the workbench attachment **prose-standard.md**. Agent Code Review holds it against that standard's tell list, so write to it the first time. When the card rewrites existing prose, the standard's hard constraint applies before style does: meaning cannot change, and the negative clause you are tempted to cut is often the requirement.

## Go source is written to the Go style standard

The workbench attachment **go-style-standard.md** governs every Go file this card touches, and reading it belongs before the code rather than after a review returns it.

Two of its rules produce work at this stage. Run `gofmt -l .` before you hand off and fix whatever it names, because the gofmt job on the pull request runs the same command and any file it lists fails the workflow. And search for prior art before writing a new function, which on a tree this size is `grep -rn '^func ' cmd internal`, then say in WHAT SHIPPED which existing helper you reused or that you searched and none existed. A WHAT SHIPPED silent on reuse is an incomplete handoff.

A new module dependency is named in WHAT SHIPPED together with the standard-library route you rejected and what made it insufficient. The dinah dependency is different in kind: the go.mod pin names the commit this repository builds against, and a change needed on the dinah side gets filed on the Andoneer board rather than made here.

## A diff that outgrows the card goes back rather than forward

A card that reached this column straight from Triage was judged mechanical, with a bounded blast radius. When the work turns out to touch schema, auth, a published contract, or more than one behavioural surface, that judgement was wrong: move the card to Spec with a comment naming what a spec would have to settle, rather than finishing it on the short path.

## The testing boundary, and what belongs to Test

Implement runs unit-level verification only. Build, run the tests covering what this card changed, and nothing more: no repo-wide sweeps, no coverage runs, no benchmarks. The full regression matrix belongs to Test, and this column can cycle with a reviewer more than once on a single card, so every second spent here is paid again on every loop.

The arming proof above is the one exception, because it requires building the broken state deliberately.

The gate before handoff:

- `go build ./...` and `go vet ./...` clean. Non-negotiable.
- `go test` on the package or packages the diff touched.
- Stored-format changes: verify against a workbench written before the change and one written after, with data preserved.
- Interface changes (a page, a control, a route, a command): run the affected surface locally, with a comment naming what was exercised.

## The branch is where this work lives

The card's code goes on a branch named after the card, and you never push the trunk. The branch takes the card's reference as its name, so card dinahvisor-7 works on branch dinahvisor-7. Record it in the card's body or a comment before you write code. That name is what every stage after you reads to find your work, and a branch nobody recorded is a branch nobody downstream can find.

Then put the worktree on the work detached, and let the remote decide where you detach:

```
git fetch origin
git ls-remote --heads origin <branch>
```

A line of output means the branch already exists, so detach onto it with `git checkout --detach origin/<branch>`. No output means it does not exist yet, so detach onto the trunk with `git checkout --detach origin/main`. That single question covers the first pass, the re-entry after a push-back, and every awkward state in between, and the recorded name never changes in any of them.

**Detach rather than `git checkout -b` or `git checkout -B`.** Both bind a local branch name, and each ordinary re-entry state breaks one of them. `git checkout -B <branch> origin/<branch>` dies with `fatal: '<branch>' is already used by worktree at '<path>'` whenever an earlier pass's worktree still holds the name, and since those worktrees outlive the stage that made them, that is the normal case rather than the rare one. `git checkout -b <branch> origin/main` dies with `fatal: a branch named '<branch>' already exists` when a previous pass created the name and was interrupted before it pushed. Worst of the three, `git checkout -b <branch> origin/main` SUCCEEDS when the branch exists only on the remote, quietly tracking the trunk instead, and the push at the end is then rejected non-fast-forward with no legal way out, because you may not force-push. Detaching binds no name, so none of that can happen.

Never rebase work you have already pushed and never force-push. When the trunk has moved and your diff needs what moved, bring the trunk in with `git merge origin/main` and resolve there.

Commit as `git add <specific paths>` (never `git add -A`), with the message `<card reference>: <one-line summary>` and a body mirroring WHAT SHIPPED. Before handing off, bring the trunk in one last time and push:

```
git merge origin/main
git push origin HEAD:refs/heads/<branch>
```

Spell that destination out in full. The shorthand `git push origin HEAD:<branch>` updates a branch that already exists, but on the pass that creates one git refuses it with `error: The destination you provided is not a full refname`. A bare `git push -u origin <branch>` is worse than refused: from a detached HEAD it pushes whatever local branch happens to carry that name, reports `Everything up-to-date`, and exits 0 having shipped none of your work.

Never `--force`, never `--no-verify`. Confirm the branch is on the remote with `git ls-remote --heads origin <branch>`, and confirm you did not land anything on the trunk with `git log origin/main -1`.

## Every branch gets a pull request

After the push, make sure the card's pull request exists. It is the operator's window onto the diff, and CI runs its checks against it for free.

**Create the PR as the pipeline identity, never as the operator.** The machine account `dinah-gh` exists because GitHub refuses to let an author approve their own pull request, and a PR created under the operator's own authentication takes his Approve button away. Its token lives at `$HOME/.dinah-gh-token`, outside every repository; never copy the token itself into any file, comment, or output. PR authorship follows the creator, so only the create call needs the identity, and pushes and reads stay on the default authentication:

```
gh pr view <branch> --json number,url 2>NUL   # an existing PR's number and URL; errors when none exists
GH_TOKEN="$(cat "$HOME/.dinah-gh-token")" gh pr create --head <branch> --title "<card reference>: <one-line summary>" --body "<mirrors WHAT SHIPPED>"
```

Create only when the view found nothing; a re-entry after a push-back already has its PR, and the new commits appear on it. Either way, put the URL in your handoff comment as a Markdown link, `[PR #<n>](<url>)`, so the operator and any PR-reading tool navigate straight to it from the card. A pushed branch with no PR is an incomplete handoff, and a PR authored by the operator's own account is a defect of the same class: it silently disables the review the station downstream exists to give.

## Reference verification

Every user-facing string that names a destination (a URL, a page, a button, a route, a command, a settings location, a card reference, a "see X") must point at something that exists in the codebase or in this diff. Otherwise correct the copy, or file the destination as a follow-on card and record its absence in WHAT I CUT.

## Exits

**Forward: Agent Code Review**, which the columns list places directly after this station.

The handoff comment carries `## WHAT SHIPPED` and `## WHAT I CUT` as Markdown headings on their own lines; the cuts are the reviewer's and tester's map of what is NOT verified. The same comment carries the `[PR #<n>](<url>)` link and lists any pending owner operator questions, so the reviewer and then the operator's station see them coming.

**Back:** one column, with a comment. **Halt:** block in place for an operator ruling the work cannot proceed without.
