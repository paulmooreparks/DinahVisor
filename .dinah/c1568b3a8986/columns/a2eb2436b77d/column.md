---
title: Triage
slug: triage
kind: work
---
The routing station. **Every card entering this workbench stops here first, and nothing leaves Intake by any other door.** Triage is the only place a card's path through the workbench is chosen, and every stage after it inherits that choice without revisiting it. A card that skips this column does not get "no route", it gets the default one, which is the longest and most expensive on the workbench.

Think of it the way a hospital does. Most arrivals need a glance and a direction, a few need one look and an immediate answer, and nothing is admitted without somebody deciding where it goes. The glance is cheap precisely because it is a glance.

That cheapness is the discipline. A card arriving with severity and priority already stamped is confirmed rather than re-derived: read it, agree or correct, choose the route, move it on. Spending deep reasoning on a classification that is already obvious is the waste this column exists to avoid.

### The route

Three decisions, and they travel together.

**Does this card need a spec?** Triage asks the question and never answers it, because writing the spec is Spec's work. Work needing a contract goes to Spec, which is the normal path. A mechanical fix whose contract is already obvious from the card alone goes straight to Implement, meaning nothing touching schema, auth, a published contract, or more than one behavioural surface. A card that turns out not to be ripe goes back to Intake.

**Will the operator have to approve something?** Call it here, as a prediction rather than a finding. A card that will produce a UX sketch (the page or interaction the operator will see, described or mocked in text), an external interface, a change to how the workbench itself works, or published copy is going to stop at Operator Design Review before anything commits to it. Say so in a comment on the card so the route is visible from the start. You are not the last word: Spec knows for certain once it has produced an artifact or not, and the operator stations are ordinary columns on the route, so a predicted stop that proves unnecessary costs the operator one glance.

**The workstream.** Read the set with `dinah workstream` and attach the card with `dinah join`, to more than one where the work genuinely spans them. A card belonging to no workstream is invisible to every workstream-scoped read from here on, so leaving it unattached quietly removes it from the views work actually gets planned from. When nothing fits, say so in a comment rather than inventing a workstream to fill the field.

### Severity and priority

Set both with `dinah card set`, and set priority in context rather than in isolation. Severity asks how bad the thing is on its own terms, and the card alone usually answers it. Priority asks when the card should be worked, which the card alone cannot answer, because that answer is a claim about this card relative to everything already in flight.

So read the workbench before you stamp. `dinah query "priority:now"` names every other card already at the top level, so the incumbents are one command away rather than a column-by-column scan; the language filters without ranking, so which of two matters more is still your judgement. Ask after the columns holding work in flight rather than Intake, which is a backlog where a stale top-priority stamp costs nothing.

**The top priority is a cap of one.** If two cards are both `now`, neither of them is, and the field has stopped carrying signal for everybody who reads it. When a card genuinely displaces the incumbent, demote the incumbent in the same sitting (`dinah card set <card> priority next`) and name both in a comment. When it does not displace the incumbent, stamp `next` and leave the incumbent alone. A second `now` stamped to avoid making that call is how a priority field decays into decoration. When the cap is already blown, one triage pass cannot restore it: do not add to the pile, say in a comment that the cap is over so the count stays visible, and leave the repair to an operator pass across the whole set.

A card arriving with severity and priority already set is confirmed rather than re-derived, and the workbench read above applies to that confirmation exactly as it does to a fresh stamp.

Duplicate-check while you are here, against work already in flight: search with `dinah query`, and when you find a twin, record the relationship in the card's frontmatter as a `duplicates` link, then archive the loser with `dinah archive` or merge its content into the live card.

### Who works it

A triage agent works this column when one is dispatched, and the operator works it directly when nobody is. The mandate is the routing decision above and nothing else, so the card leaves as soon as the three calls are made: release the claim, then `dinah move` the card to its destination, because a move changes the card's column and nothing else about it.

Cards arrive by a pull from Intake (`dinah pull triage`), by an agent or by the operator, and that pull commits the puller to the triage glance rather than to the work. It is the only way in, and it is the only way out of Intake. A card moved from Intake to any other column has skipped its route, and the right correction is to send it back there rather than to guess the route downstream.

Batch the sitting where you can: pull several cards from Intake, classify the set, route them all. A pull claims what it lands, so the classification is visible work from its first act, and that rule covers operators too.
