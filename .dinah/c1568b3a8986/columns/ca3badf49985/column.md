---
title: Spec
slug: spec
kind: work
---
Write the spec: the contract another reader could implement without further questions, against the DinahVisor codebase, a Wails desktop application importing dinah as a Go library. UX-touching cards include a UX sketch: the page or interaction the operator will see, described or mocked in text. Spawn related cards as they surface; inherit workstream membership on each, and file them into Intake unless the card says otherwise.

Field discipline: the card body stays a one-paragraph framing; the contract travels as a spec attachment (a spec.md attached to the card), and the rationale and alternatives the contract grew from travel with it in that file. Acceptance criteria, open questions and decisions go in structured checklist items (kind acceptance_criterion / open_question / decision, hand-written under `checklist/<id>/item.md`), never as prose headings inside the spec attachment. A body that arrives carrying spec content (acceptance criteria, out-of-scope lists, implementation prescriptions, enumerable deliverables) goes back to Triage with a suggested one-paragraph rewrite; a merely thin body gets refined in place.

Write every prose surface to the workbench attachment **prose-standard.md** from the first draft: the spec attachment, checklist text, documents, published copy. Agent Design Review holds prose against that standard's tell list, so a spec written without it buys a review cycle a first draft would have avoided. Its meaning-preservation rules matter here in the authoring direction too: when the contract needs an exclusion, write the "X, not Y" and keep it, because the standard protects load-bearing antitheses rather than banning them.

### Open questions, and who they belong to

Don't fabricate; record an open question instead. Every question you leave pending must carry an owner in its frontmatter, which is a stored field rather than a sentence: file it with `owner: holder` when whoever works the card next answers it in the course of the work, and with `owner: operator` when the ruling is genuinely the operator's. An unstamped question is read everywhere downstream as needing the operator, so leaving the owner off is a decision to put the card in his queue rather than an absence of one. Make that decision here.

The note still carries the reasoning, the recommendation and the tradeoffs, and it is no longer where the addressee lives. Prose cannot be filtered, counted or routed, which is how a card raising five questions for its own next implementer adds five items to the operator's queue.

An owner operator question is answered at Operator Design Review, which the route places two stages ahead of you, directly after Agent Design Review. A question the spec can be finished around simply rides the card forward and gets answered at that station; name it in a comment before you move the card so the reviewer and then the operator see it coming. The hosted board kept a mechanical backstop on such questions at its output queue; this surface has no gates, so the backstop is prose: nothing reaches Done owing the operator a ruling, and the column bodies carry that rule.

A question that makes the spec undeliverable does not travel at all: file it, then `dinah block <card> --kind operator_decision --reason "<the question, posed so it can be answered without opening the card>"` before the card moves anywhere. The block frees the card and halts it here; only the operator lifts it, so nothing follows the block, no move and no release.

### Decisions, and what you are allowed to decide

A decision item records a call that has already been made. Filing one closes the question, and no stage downstream reopens it, so the bar for filing one is that the call was yours to make.

Two tests apply before you file. Ask first whether the call commits the operator to something durable and awkward to reverse, such as the shape of their repository's history, a published interface, a stored data format, a price, or anything a customer will see. Ask second what your stated justification rests on. If it contains a claim about the operator's world that you have not verified, meaning what their trunk actually looks like, how they run another project, or what they would prefer, then you are guessing and the guess is carrying the argument.

A call that fails either test is an open question. File it as one, stamp it owner operator, put your recommendation and its tradeoffs in the note, and let the operator rule at their station. When the operator agrees with you, asking cost one line of confirmation. When a wrong call ships as a resolved decision, it costs an implementation, because nobody downstream questions a decision that already reads as settled.

Verify claims about this repository before you rest a decision on one. Most of them are a `git log` away. Put the count in the note.

### An artifact is a file, and a file belongs on the card's branch

Spec artifacts are UX sketches (page and interaction sketches in text) and, sparingly, `docs/specs/<card-slug>-design-notes.md`. Files the implementer will edit are never created here.

**When you write one, put it on the card's branch and push it before the card moves.** Work happens in a worktree under C:\dinah-scratch, never inside the repository, so what you write lands in a throwaway tree and is lost unless you commit it. Writing into the operator's working directory instead leaves an untracked file for the life of the card, which is how a spec artifact ends up blocking work on a card it does not belong to.

The branch belongs to the card rather than to a stage, and whichever stage first needs one creates it. That is Spec whenever Spec produces a file. The branch takes the card's reference as its name, so card dinahvisor-7 works on branch dinahvisor-7. Record the branch name in the card's body or a comment before committing anything, then detach at the branch if `git ls-remote --heads origin <branch>` shows it and at `origin/main` if it does not. Push with the full `HEAD:refs/heads/<branch>` form. Implement's own branch section already finds a branch that exists, so creating it here costs the implementer nothing.

### Artifacts the operator will have to accept

If this card produced something whose form nobody can test until a person accepts it, the operator will rule on it at Operator Design Review. Your job is to name the cargo, not to route it: say in a comment before you move the card what the artifact SHOWS and state the one-line acceptance criterion the operator's approval would create. A UX sketch is the common case; an external interface, a change to how the workbench itself works, published copy and a customer-facing schema are the others.

Say what the artifact shows, not only where it lives. A path tells the operator a file exists and leaves them to find out what is in it, which is how a ruling gets given on a summary instead of on the thing. Name the sections, name what differs between the options, and name the detail that decides it.

Naming it is not optional politeness. If you produced a UX sketch and say nothing, the operator waves the card through their station blind, and it reaches an implementer with no criterion governing the form it builds.

### Exits

**Forward: Agent Design Review.** The columns list is the route and there are no lanes, so the stage after Spec is Agent Design Review for every card here.

Move forward when the spec stands: the contract is complete, every criterion is testable, and every question is resolved or carries an owner. Post a comment that summarises the contract in two sentences, lists any pending operator questions, and names any artifact awaiting approval along with the branch it is on and the criterion its approval would create, then move the card.

**Back:** push to Triage when the card's shape is wrong, meaning it is really two cards, or a duplicate, or its body carries spec content.

**Halt:** the operator-decision block described above.
