---
title: Intake
slug: intake
kind: intake
---
The workbench-wide entrance, standing outside every station. Cards land here raw: a title and a short framing description, often filed from a chat in two seconds. Bodies are optional, specs are explicitly NOT written here, and Intake is the only place on this workbench where unclassified work is allowed to exist.

Intake is an external bucket: a card here sits unheld, with no capacity limit and no cycle-time clock. A card sitting here is captured intent rather than work in flight. Intake age measures capture-to-triage latency, and the flow clock starts at Triage.

### The only exit is Triage

Every card leaves this column into Triage and nowhere else. Triage is where a card's route through the workbench is chosen, and a card moved from here to any other column has skipped that choice and taken the default route by accident, which is the longest one on the workbench. A card that got out another way goes back here rather than having its route guessed downstream.

### Who pulls

Any owner, agent or operator. A pull from Intake commits you to the triage glance and to nothing else: the pull is `dinah pull triage`, it claims the head of this queue and lands the card at Triage, and the commitment is to classify the card, choose its route, and move it on. It is not a commitment to the work itself, and the decision to spend a working station's budget on a card is made later, when somebody claims it at Spec or at Implement.

Nothing is worked in this column. Classification happens in Triage, not here.

### When to leave a card here

Captured but not yet ripe: a feature concept whose user value is not clear, a bug report waiting on reproduction, a design question waiting for a workstream's direction. These are legitimate, and the column listing and `dinah query` surface them when the time comes.
