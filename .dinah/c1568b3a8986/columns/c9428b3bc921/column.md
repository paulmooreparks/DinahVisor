---
title: Test
slug: test
kind: work
---
Validate the implementation against the spec and the acceptance criteria by running it. This is the workbench's one authoritative execution stage and the last gate before Merge puts the code on `origin/main`, so the repo-wide sweep runs here and only here, when it is warranted.

### Put the worktree on the branch, then bring the trunk into it

The card's code is on the branch named after the card, recorded in its body or a comment, and it is not in the trunk, so a worktree left on `main` verifies nothing at all. Work in a worktree under C:\dinah-scratch, never inside the repository. Before you run anything:

```
git fetch origin
git checkout --detach origin/<branch>
git merge origin/main
```

Detach rather than `git checkout -B <branch>`. Git refuses to check out a branch that another worktree already holds, and the implementer's worktree usually still exists and still holds this one. Checking out the remote-tracking commit detached gets you the same tree with no collision, and it works whether or not that worktree is still around.

The merge is not housekeeping. Nobody tests the trunk with this card in it before Merge lands it, so the merged state is what you are being asked about, and testing the branch alone would pass a card that breaks on contact with whatever shipped while it sat in review. If the merge conflicts, push back to Implement with the conflicting files named. If it changed anything, push it with `git push origin HEAD:refs/heads/<branch>` so Merge and every later reader see the same tree you tested. If git reports "Already up to date", nothing changed and there is nothing to push; say so rather than pushing an empty update.

### Choose the scope deliberately, then name what you ran

Targeted is the default. Identify the touched packages from WHAT SHIPPED plus the diff, then run `go test` on those packages and their direct dependents. Reserve the full `go test ./... -count=1` sweep for cards touching stored formats, shared infrastructure, cross-package renames, the dinah dependency pin, or anything under cmd/. The report says which one you ran and why.

### The card's new tests have to run inside the sweep you executed

Regression integration is a check here, not a formality. Confirm the new tests actually ran rather than assuming they did. Where an existing test already covered the area, the new coverage belongs inside it; a parallel file testing the same ground twice is duplication every future card pays for, and you report it rather than letting it settle. A suite growing faster than the code it guards is a defect in its own right.

### Criteria already marked verified are the thing to be careful about

By the time a card reaches you the implementer has marked its criteria and two reviewers have accepted those marks by reading. That is exactly the condition where an execution stage earns its keep by being suspicious instead of confirmatory. Re-verify by running, and re-mark anything whose evidence does not survive an actual run. A criterion marked verified on a thin note is a finding.

Watch for the two test shapes that pass without proving anything. An assertion that matches a substring the old behavior also produced tells you nothing about the change. An expected value computed by calling the thing under test moves with the defect and can never fail. Either one is a failed criterion, not a passing one.

### Verification kinds, as the card demands

Package tests, stored-format runs (a workbench written before the change plus one written after) where the card touches the format, and direct exercise of the affected pages or controls where the change touches what the operator sees. Pick up any directed verification targets the Agent Code Review findings comment named; those are contract.

User-facing output gets realistic data shapes rather than the happy path alone: empty, single row, boundary at any cap, overflow past the cap so that any "N more" affordance fires, and very-long text. A criterion that mentions a cap, pagination, or "top N" is only verified once the overflow condition actually ran; otherwise mark it pending operator verification.

### Every open question you file names its owner

A pending open question on a live card reads as the operator's whenever it carries no owner. You are past both operator stations here, so a question for the operator does not travel with the card: it blocks the card in place with `dinah block <card> --kind operator_decision --reason "<the call>"`, and the operator answers and unblocks.

Before you file a question that does ride the card, place it in one of three classes and stamp it:

1. **The operator's.** It commits them to something durable or awkward to reverse, or it is a scheduling call only they can make. That is the block above, not a travelling item.
2. **A later stage's.** Merge or Acceptance decides it as a matter of course. File it with `owner: holder`, say which stage and what they need in order to decide. A criterion that genuinely cannot close until the change is on the trunk belongs here rather than being marked failed: leave it pending, say it closes at Acceptance, and say why.
3. **Another card's.** Resolve it here with a note stating the follow-up at title quality, so whoever files it does not have to re-read your report. A question waiting on unscheduled work is a backlog item, not a pending decision.

Report the counts: how many you filed, and whether anything blocked for the operator. Run `dinah check` after editing any item.

### Report and route

Mark each acceptance criterion verified (with what was checked) or failed (expected versus actual). Post a `## Test report` comment carrying status, scope, an output sample, failures, and the verified/pending criteria. A pass moves the card to Merge. A failure pushes back to Implement. A problem outside this card's scope becomes a separate card, filed into Intake.

The work here is mechanical: build, run, diff, check criteria. The judgment about what to verify already came from Spec and Agent Code Review, so a card that seems to need fresh design judgement at this stage has a defect upstream of you; say so in a comment rather than redesigning it here.

### What this stage cannot cover

Every card is tested on its own branch with the trunk merged in as of the moment you merged it. A failure that only appears once this card sits alongside siblings that landed after that merge is invisible here by construction. Acceptance exercises the integrated trunk for exactly that reason.
