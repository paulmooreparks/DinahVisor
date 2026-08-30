---
title: Done
slug: done
kind: done
---
The terminal column. A card here has completed its work, its commit is on `origin/main`, and it was accepted there without incident.

### Discipline

- **No work happens here.** Cards in Done are read-only by convention; the card's journal is the audit trail.
- **No re-opening.** If a Done card needs revisiting (regression, follow-on, scope expansion), file a NEW card and record the relationship in the new card's frontmatter links, pointing at the Done card. Don't drag the Done card back into flow. The tool will refuse a forward move out of this column anyway, and a backward one throws away the record of what shipped.
- **Comments after Done are allowed but rare.** Cross-references and follow-on links are fine; substantive new content belongs on the new card, not here.

### What being Done means

The card's spec is satisfied, its acceptance criteria are verified, its commit is on `origin/main`, and the change has been exercised without incident. Merge landed it and Acceptance watched it.

Done therefore means as far as this workbench currently ships. A card is Done when it has passed every stage the route has, whatever those are at the time.
