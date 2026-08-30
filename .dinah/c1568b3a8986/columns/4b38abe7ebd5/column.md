---
title: Agent Code Review
slug: agent-code-review
kind: work
---
A fresh agent reads the implementer's diff against the spec: clean, forward to Operator Code Review; findings, push back to Implement. The reviewer comes to the card with no memory of writing it, which is the whole value of the seat.

Checks: correctness against the spec and its edge cases; security (secrets handling, authorization on every entry point that has one); style (codebase idioms); self-test level appropriate per the Implement gate; WHAT I CUT honesty.

## Begin at the code

Reading the diff is where this review starts and where most of it happens. Implement already smoked the change and Test owns the authoritative execution pass, so a clean read is a complete review. Do not build, run a suite, or exercise a fixture as a matter of course. A routine build on a diff that read clean buys nothing Test will not buy properly an hour later, and it is how this column loses the cheapness that lets it run on every card.

Expand into execution when, and only when, reading has surfaced a specific problem you must understand before you can write the response. That is the trigger: a concrete question the diff raises and the source will not answer, not a wish for more confidence. Once you are there, run what the question needs; the scope follows the question rather than a fixed permission.

When you expand, record it in the findings comment: what reading could not settle, what you ran, and what it showed. That is the difference between a justified expansion and drifting back into being a second Test.

When you suspect a behavioral defect and choose not to chase it, either push back to Implement with the concrete suspicion, or pass and name the check in your findings comment as a directed verification target for Test ("Test must specifically exercise X"). Named targets are contract that Test picks up.

## The operator's station is next, and your comment briefs it

Operator Code Review sits directly after this column, so a clean card's forward move lands it in front of the operator, who reads the diff on the card's pull request. Verify the `[PR #<n>](<url>)` link from WHAT SHIPPED points at the card's branch and carry the same Markdown link forward in your handoff comment; a code card whose branch has no pull request is a [major], because the operator's review surface is missing.

The comment also tells the operator what else awaits: any pending owner operator questions listed (Implement files these when the work could be finished despite the question, and they ride forward rather than blocking), the one-line acceptance criterion an artifact's approval would create, or the plain statement that nothing awaits them so one pass sends it on. A pending operator question never rides past that station.

## A change without a test is a blocker, not a nit

The Implement column requires that delivering an implementation means delivering a test for the behaviour that changed, and this column is where that is enforced. Push back when the diff changes behaviour and nothing in it exercises that behaviour. An extended existing test satisfies the rule and is the preferred shape; a new file covering ground an existing test already owns does not, because a suite that grows by accretion becomes slow, then redundant, then skipped.

Check also that the implementer armed the test: broke the fix, watched it go red, restored, watched it pass. A test nobody has ever seen fail is not yet evidence that it guards anything. Two failure shapes are worth knowing because both pass a careless read. An assertion that matches a substring the old and the new behaviour both produce is not a test of the change. And an expected value computed by calling the thing under test moves with it, so the assertion holds in both worlds; the expectation has to be pinned independently.

## Check the diff against the house idiom corpus

The workbench attachment **convention-counterexamples.md** carries the wrong-versus-right idiom pairs this work has already paid for, and consulting it is part of the review rather than optional enrichment. It is large, so do not read it end to end: match its entries against the surfaces this diff actually touches, and read those.

Report the result in your findings comment: which entries you checked, and which areas you skipped because the diff does not touch them. A review that names no entries has not done this step, and "none applied" is a claim that needs the areas listed to be credible.

When you catch a defect whose class is not yet in the corpus, and the same class has now been caught twice, say so in your findings comment and file a card to append it as a new wrong/right pair. The attachment is a snapshot of a hosted original, so refreshing it here is deliberate work rather than an edit in passing. That promotion practice is what keeps the corpus worth reading, and it is the only mechanism this workbench has for turning a lesson into something the next reviewer inherits.

## Prose in the diff is reviewed against the prose standard

The workbench attachment **prose-standard.md** governs any prose this diff touches: README and docs changes, user-facing copy, comments on the card, and the WHAT SHIPPED note itself. Hold that prose against the standard's tell list and tag violations as findings. Its hard constraint cuts both ways here. A diff that rewrites existing prose gets checked for meaning drift word by word: a dropped qualifier, a deleted negative clause, or a weakened commitment is a [major] even when the style improved, because the antithesis a rewrite is tempted to cut is often the specification. And a finding you write asking for a prose fix must itself say what to remove without removing what the sentence promises. Name the tell when you tag one; "reads machine-written" is not actionable.

## Go style and reuse

The workbench attachment **go-style-standard.md** is the rule set for Go source in the diff, and this column is where it is enforced. Take the mechanical floor first, since `gofmt -l .` settles that part without judgement, then read for the judged rules.

Weight findings the way that document weights them. A file gofmt would reformat is a [major]. So is a new function duplicating one the codebase already has, or a second solution to a problem the codebase already solves a particular way. The judged rules covering line density, argument breaking, guard clauses, doc comments and error strings are [minor] findings, recorded in the findings comment and not on their own a reason to send the card back. Where a [minor] will outlive the card, the same document says what becomes of it.

Check the reuse claim and not only the code. WHAT SHIPPED either names a helper that was reused or asserts none existed, and that assertion is checkable in one grep. A claim absent altogether is itself a finding.

Cite the rule you are applying. A style finding the document does not state is out of scope here, and the way to raise one is to propose it as an addition to the document.

## Copy references are load-bearing

For every user-facing string that names a destination (a URL, a page, a button, a route, a command, a "see X", a "configured under X"), grep that the destination exists or is created in this diff. A plausible reference to vapor is a [blocker]; the copy gets corrected or the destination gets a follow-on card.

## A card whose deliverable is not repo code

Some cards deliver a document, a UX sketch or a specification rather than a commit, and those land as an attachment on the card or as an edit to this workbench's own prose. Such a card has no branch and no diff, so the git commands below find nothing. **That absence is the expected state for these cards, not a push-back.** Tell the two apart by the deliverable the spec names, never by the absence of a branch alone: a card whose spec promises code and produced none is still a push-back with the absence stated.

Review the artifact itself. Read it with `dinah show`, `dinah attachments` or `dinah instructions` as the deliverable demands, and hold it against the spec, against this workbench's writing conventions, and against the reference rule above, which applies to prose exactly as it applies to copy in code. Check what the implementer named in WHAT SHIPPED against what you read: when they disagree, something moved after the handoff, and you are reviewing a surface the implementer did not hand you. Say so rather than reviewing it silently. The findings format and the tagging are unchanged.

## Where the diff is

Findings go in a `dinah comment` under `## Findings`, tagged [blocker]/[major]/[minor]/[nit], with the counts in the same comment. Read WHAT SHIPPED (for the diff base) and the prior comments first; don't duplicate findings.

The card's branch is named after the card and is recorded in its body or a comment. Fetch first, then read the branch against the trunk:

```
git fetch origin
git diff origin/main...origin/<branch>
git log --oneline origin/main..origin/<branch>
```

Three dots, not two: that is the diff since the merge base, so it stays correct after the implementer merges the trunk into the branch, which they are told to do. The log gives you the commit series, which is how you see what a second or third round added without re-reading what you already passed. `gh pr diff <branch>` reads the same diff through the pull request when that is more convenient.

This code is NOT in the trunk. Pushing back costs the trunk nothing, which is the whole reason this stage sits where it does.

**A recorded branch that no longer exists on `origin` means the card was already merged.** Merge deletes both copies when it lands the squash commit, and the remote deletion drops the remote-tracking ref at once, so the commands above fail with `fatal: ambiguous argument 'origin/main...origin/<branch>': unknown revision or path not in the working tree`. That is the ordinary end state of a merged card, so do not read it as a missing branch and do not push back for it. Read the squash commit on the trunk instead: `git fetch --prune origin`, then `git log --oneline --grep=<card reference> origin/main`, then `git show --stat <hash>`. If the trunk carries nothing under the card's reference either, the work is genuinely nowhere, and that is a push-back with the absence stated.

A card that should have a branch and has none recorded is a push-back on its own: the implementer skipped recording it and nobody downstream can find the work.

## Repeated rounds

This surface declares no loop limit, so nothing counts your rounds for you. Judge instead. A card you have now sent back three times on the same class of finding is not being reviewed, it is being argued with, and the argument is itself a question for the operator: post `## Escalation` with the current findings and the round history, then `dinah block <card> --kind operator_decision --reason "<what needs deciding>"` and do not move the card.

The distinction that matters: a card that came back clean, with no blocker and no major, has finished rather than circled, whatever number of rounds it took to get there. Blocking that card asks the operator a question with no content. Say plainly which of the two you are looking at, and record the reasoning either way so it can be overruled.
