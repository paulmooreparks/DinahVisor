---
format: 1
profile: dinah-core/0.7
title: DinahVisor development
slug: dinahvisor
operator: paul
columns:
  - 5ea2db0272fc   # Intake
  - a2eb2436b77d   # Triage
  - ca3badf49985   # Spec
  - 0d86ad99cdbc   # Agent Design Review
  - 5729d4578008   # Operator Design Review
  - 4fda9c9ca779   # Implement
  - 4b38abe7ebd5   # Agent Code Review
  - ee29487fad76   # Operator Code Review
  - c9428b3bc921   # Test
  - 6c5b9d6f4414   # Merge
  - b69abf918c42   # Acceptance
  - aa6cd1c6ae5f   # Done
levels:
  severity: [trivial, minor, major, critical]
  priority: [later, soon, next, now]
groups:
  DESIGN: [ca3badf49985, 0d86ad99cdbc, 5729d4578008]
  BUILD: [4fda9c9ca779, 4b38abe7ebd5, ee29487fad76]
  VERIFY: [c9428b3bc921, 6c5b9d6f4414, b69abf918c42]
---
This workbench is where the DinahVisor repository is developed. DinahVisor
is a GUI desktop tool for Dinah: a Wails shell around the pages and the
command log, styled by the Tela Design Language, importing dinah as a Go
library and running the HTTP head's handler in process. The repository
lives at C:\Users\paul\source\repos\DinahVisor on trunk branch main,
published at github.com/paulmooreparks/DinahVisor. Development of the
dinah repository itself stays on the hosted board in Andoneer; the work
that moved here is the GUI.

## Where this bench came from

This bench was instantiated from the dinah repository's own workbench
(`.dinah/149f228d48c3` there) through `dinah init --from`, so its id is
fresh. That source bench is a hand extraction of the hosted board, and
its body records the fifteen findings of that extraction. A definition
travels whole, so this bench inherited every one of those losses, and
the source body, in the dinah repository, is the place to read what the
format could not carry.

This bench is live where that one is a definition: cards get filed
here, claims get taken here, journals accumulate. This body therefore
spends the standing-context slot on working rules rather than on
extraction archaeology. The hosted board remains the arbiter for dinah's
own development, and this bench is the arbiter for DinahVisor.

## What the translation changed

The source bench's column bodies speak the hosted product's surface:
claim verbs named for MCP tools, gates, lanes, zones, tiers, loop
limits, workbench documents. None of that is served here. Each body was
translated onto this tool's own surface, and the choices are recorded
once here so the columns stay short.

1. Verbs are CLI spellings (`dinah claim`, `dinah move`, `dinah release`,
   `dinah block`, `dinah card`, `dinah comment`), because the CLI
   spelling is the stable one while heads vary. A move carries no note
   on this surface, so wherever the source bodies said move-note, the
   words travel as a `dinah comment` on the card.

2. Checklist items are hand-written files under the card's
   `checklist/<id>/item.md`; docs/design/format.md in the dinah
   repository is the authority on their frontmatter, and `dinah check`
   validates the structure, so run it after writing one. The owner
   value is the plain word: `owner: operator` says the operator answers
   the item, `owner: holder` says whoever next works the card does.
   The three kinds survive: acceptance_criterion, open_question and
   decision. A decision records a call already made, and two tests
   gate filing one.
   The call must not commit the operator to something durable and
   awkward to reverse, meaning the shape of their repository's history,
   a published interface, a stored data format, a price, or anything a
   customer will see. The stated justification must not rest on a claim
   about the operator's world that has not been verified, meaning what
   their trunk actually looks like, how they run another project, or
   what they would prefer. A call that fails either test is an open
   question stamped owner operator.

3. Gates have no slot, so a card that must wait on another card carries
   that in prose on both cards.

4. Links are frontmatter declarations that nothing enforces, and they
   are same-workbench only, so a dependency on a dinah card travels as
   prose plus the go.mod pin rather than as a link.

5. The hosted branch_name field has no counterpart, so the working
   branch is named in the card's body or a comment at the moment work
   starts on it. The branch is named after the card's reference, so
   card dinahvisor-7 works on branch dinahvisor-7.

6. Zones, capacity limits, lanes, expected tiers, loop limits and
   subagent directives are declared nowhere on this bench; the
   discipline they carried lives as prose in the column bodies. Pull
   order is the queue's entry order, earliest first with ties broken
   by ascending creation ordinal, so severity and priority inform
   people and never the queue.

7. Pull requests on this repository run under the operator's own
   identity. The pipeline bot, login dinah-gh, holds read-only access
   to paulmooreparks/DinahVisor: it can pull and cannot push, probed
   2026-08-29. It therefore cannot create branches or PRs here, and
   the source bench's rule that a PR authored by the operator's
   account is a defect does not translate. Granting the bot write
   access would restore that rule, and the grant is the operator's
   call.

8. Four standards hang off this bench as workbench attachments:
   convention-counterexamples.md, prose-standard.md,
   go-style-standard.md and working-safely.md, snapshotted 2026-08-29
   from the hosted originals. The originals remain the living
   versions, and refreshing a snapshot here is deliberate work.

9. The source bench carried a queue column in front of each working
   station, Design Queue before Spec and Build Queue before Implement,
   where cards waited for a station to have room. Neither column exists
   here. A card waits at the station it is bound for and whoever works
   that station claims it there, which is the same pull discipline
   with one fewer place to stand. Twelve columns run the flow, and the
   demand signal the queues carried is read off the waiting cards at
   each station instead.

## How this bench runs

Implementation work goes in worktrees under C:\dinah-scratch, never
inside this repository, because a worktree of this repository carries a
git-tracked copy of this bench and a mutation landing there would
diverge from the live one silently. Name this bench by absolute path in
every tool call, because the server serves no default and discovery
from inside a worktree finds the copy rather than the original.

The columns list above is the route. There are no lanes and no
adjacency rule, so every move the list allows is legal to the tool, and
the discipline the hosted board enforced through lanes and typed
push-back directives lives entirely in the column bodies. An agent that
skips reading them skips the flow.

The cross-repo contract is dinah's HTTP handler, importable and
mountable, with its route roster a published interface; the dinah side
carries it on dinah-152. The go.mod pin names the commit this repository
builds against, and a change needed on the dinah side gets filed on the
Andoneer board rather than here.

Every prose surface on this bench answers the style constraints the
operator holds everywhere: no em-dashes, no machine-written sentence
shapes, complete sentences in prose, headings that do not borrow their
paragraph's subject, and one register held within each list. The
prose-standard attachment carries the tell list.