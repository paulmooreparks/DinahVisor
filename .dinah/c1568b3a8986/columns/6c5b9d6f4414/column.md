---
title: Merge
slug: merge
kind: work
---
Merge the card's branch into the trunk. This is the integration, not a confirmation of one: until you run it, the card's code is on its branch and the trunk has never seen it. Merge is a mechanical land-and-record gate rather than a review or sign-off gate. Agent Code Review, the operator stations and Test are the per-card review, and Acceptance is the operator's last look.

### The merge

Read the branch name from the card's body or comments. The branch is named after the card, so card dinahvisor-7 lives on branch dinahvisor-7.

**The card gets merged through its pull request.** Work in a worktree under C:\dinah-scratch, never in the operator's checkout:

```
git fetch origin
gh pr view <branch> --json number,url,state
gh pr merge <branch> --squash --delete-branch --subject "<card reference>: <one-line summary>" --body "<mirrors WHAT SHIPPED>"
git fetch --prune origin
git log origin/main -1 --format="%H %s"
```

One squashed commit per card, so the trunk reads as one line per card and the hash you record is the card's whole footprint. `--delete-branch` removes the remote branch as part of the merge, and GitHub closes the PR, which is the record the pull-request discipline exists to create. Verify the deletion with `git branch -r --list "*<card reference>*"` after the prune, which must come back empty, and confirm the squash commit is the trunk tip. Leave the branch name recorded on the card: it is the record of where the work happened, and it stays true after the branch is gone.

**The trunk accepts no direct push.** The repository's ruleset requires a pull request with four status checks green (`test` on Linux, Windows and macOS, plus `gofmt`), blocks force pushes, and restricts deletion of the default branch, with an empty bypass list, so the pull request is the only door and the remote refuses everything else. A branch with no PR is therefore a process miss rather than a blocker: Implement was told to open one. Create the missing PR yourself with `gh pr create --head <branch> ...`, giving it the title and body a PR should have carried, wait for the required checks to finish, merge it by the path above, and report the miss in your handoff comment so it stays visible.

Two outcomes of the PR path are expected rather than problems. `gh pr view` reporting the PR already MERGED before you touched it is a re-entry after an earlier attempt landed the merge but not the bookkeeping: recover the hash with `git log --oneline --grep=<card reference> origin/main`, record it, and say so in your comment. And `gh pr merge` refusing because a required check is still running is a wait, not a failure: poll `gh pr checks <branch>` and merge when they settle.

### What to check before you merge

Spot-check the message (card-reference prefix, mirrors WHAT SHIPPED) and the diff footprint: no secrets, no accidental large files, nothing outside the card's stated scope. `git diff origin/main...origin/<branch> --stat` is the cheap read.

**A card routed straight from Triage to Implement skipped Spec, and every card here has passed Test**, so the reading you owe is narrow. Check that the diff is what the card said it would be. You are not judging correctness, which would make this a review stage with a loop.

**A merge conflict is not yours to resolve.** `gh pr merge` refuses a conflicting PR exactly as raw git would. Write `dinah block <card> --kind other --reason "merge conflict merging <branch> into main: <files>"` and stop. The operator resolves on the branch and unblocks. The implementer is told to keep the trunk merged into the branch, so a conflict here means something changed under them.

**Judge working-tree drift by whether it belongs to this card.** An uncommitted or untracked change that is part of this card's work is a real gate failure: the card is not landed, so push back. A stray file with nothing to do with this card is somebody else's mess, and it is not this card's problem to be held hostage by. Report it in a comment, name the file, and carry on; never commit it as part of this card, and never revert it, because you do not know who is mid-work on it.

One file deserves particular care. This repository carries the live workbench under `.dinah/`, git-tracked and mutated by every dinah call anyone makes. A diff touching those files is almost never the card's work, and merging one lands somebody's in-flight board state on the trunk. Treat it as drift, name it, and leave it alone.

### Exit

Move the card to Acceptance, and post a comment carrying:

```
COMMIT
<hash> (origin/main): <one-line summary>
[PR #<n>](<url>)
```

The hash is the squash commit and the PR link is the one the merge closed. Say which path ran (PR merge, or created-then-merged), name the branch that was deleted, and note anything expected-but-odd you saw.

A question for the operator raised at this stage blocks the card in place; both operator stations are behind you.

A missing commit gets pushed to the card's BRANCH before claiming, never with `--force`. Secrets or a wrong landing: don't rewrite history from Merge; push back to Implement with the cleanup described.
