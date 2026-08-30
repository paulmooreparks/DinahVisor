---
title: Acceptance
slug: acceptance
kind: work
operator_owned: true
---
Cards whose commit is on the trunk: delivered, but not yet proven in a running build.

A card sits here as the operator's awareness that the change is out and living in whatever build carries it. Passing Acceptance is the operator's acknowledgement that nothing untoward happened as a result of the card shipping. It is an acknowledgement rather than an investigation; nobody assembles an evidence file here, and a quiet build is itself the evidence.

Claim the card while judging it. The rule against working an unclaimed card covers operators too.

A card arriving here with a pending owner operator open question was misrouted: questions are answered at the operator review stations, or at the block where they arose, and never presented for acceptance. If you are the operator, answer the question before judging the card; anyone else blocks the card with the question as the reason rather than working around it.

### When disruption appears

Evidence of disruption in the corresponding build while the card sits here sends it back: to Implement when the contract was right and the change was wrong, or to Spec when the contract itself was wrong. Say what was observed in a comment. The card that shipped the change owns the consequence, and pushing that card back rather than filing a fresh one is what keeps the link between a change and its effect, which is the reason this stage exists. A problem that turns out to predate the merge is a separate card, as it would be anywhere else.

Any acceptance criterion the card still carries, including one minted at an operator station earlier in its life, is closed honestly here or left pending with the reason recorded.

### Age here

Age in this column is decision latency and it measures the operator rather than any supplier. Keep it short by batching: acknowledge several cards in one sitting rather than one at a time.

### Leaving

Move to Done when the card has sat in a running build without incident long enough to be acknowledged. Done is the terminal column, so Done here means merged and accepted.

A card leaves this column by a `dinah move` to Done. Release the claim first, because a move refuses to carry a card its holder is still holding into a column where nobody works it.
