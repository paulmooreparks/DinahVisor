---
title: Operator Code Review
slug: operator-code-review
kind: work
operator_owned: true
---
The operator's code-side work station, the twin of Operator Design Review on the build half of the route. Every card stops here after Agent Code Review, and what happens next depends on what the card brought.

The diff review happens on the card's pull request. The arriving comment carries a `[PR #<n>](<url>)` Markdown link, so the operator navigates straight to the diff in GitHub's interface or a VS Code pull-request tool, reads it there, and records approval or change-requests on the PR itself. The workbench mirrors the verdict: approval moves the card forward, change-requests move it back to Implement with the requests summarised in a comment. A code card arriving without a PR link is an incomplete handoff; Agent Code Review was told to catch it, and the fallback is `gh pr view <branch>` against the card's recorded branch.

The Approve button is the point of the pipeline identity. Implement creates every PR as `dinah-gh` precisely so the operator is not its author, because GitHub refuses to let an author approve their own pull request. A PR arriving here authored by the operator's own account has taken that button away, and it is a defect rather than a cosmetic difference: report it and have the card's PR recreated under the pipeline identity.

A card carrying pending open questions stamped `owner: operator` waits for the operator to answer them; Implement and Agent Code Review file their questions for this station rather than blocking, whenever the work can be finished despite the question. A card carrying an artifact the operator must accept before it ships (a changed page, a published surface, customer-facing copy) waits for that acceptance. A card carrying neither is the common case, and the operator reads the PR and moves it on.

### The work

Claim while reviewing; work in flight belongs on the workbench, and that covers operators too. Answer each pending owner operator question (an unstamped one reads the same) by editing the item to resolved with the ruling and its reasoning in the note; downstream stages act on the note, so it carries the ruling itself rather than a bare yes or no. A card may not leave this column with such a question still pending.

On acceptance of an artifact, record the criterion as a real acceptance_criterion checklist item naming what was approved, then run `dinah check`. On rejection, move the card back to Implement with what is wrong stated concretely enough to act on. Otherwise move the card forward to Test.

### What the senders owe this station

Implement and Agent Code Review say in their handoff comments what awaits here: the `[PR #<n>](<url>)` link, the questions listed, the one-line acceptance criterion an acceptance would create, or the plain statement that nothing does. A card arriving with no such comment costs the operator the read the sender should have done.

### Questions raised past this station

Test, Merge and Acceptance sit downstream of every operator station, so a question for the operator raised there does not travel: it blocks the card in place with kind operator_decision, and the operator answers and unblocks it there. Age here is decision latency and measures the operator; keep it short by batching.
