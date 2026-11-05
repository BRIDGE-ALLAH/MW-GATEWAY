# Moonwitness Gateway — Brand Guide

> **Connect Intelligence. Preserve Accountability.**

## Product identity

Moonwitness Gateway is the connectivity and configured policy-enforcement layer between AI applications, model providers, agent peers, MCP tool servers, and external services.

It belongs to the wider Moonwitness XV vision. It does not replace ARSY as the ecosystem's application and orchestration hub, does not act as a canonical knowledge store, and does not independently verify truth or grant human approval.

## Narrative

AI systems are becoming networks: models call tools, agents collaborate, and services exchange information across organizational boundaries. Those connections need clear routes, explicit policies, and operational visibility.

Moonwitness Gateway provides that boundary. It helps teams connect intelligence while keeping the path inspectable and the responsibility visible.

**The Moon Watches. The Record Remembers. The Evidence Remains Inspectable.**

The gateway can route, authenticate, authorize, enforce configured rules, and emit telemetry. It cannot turn model output into verified fact or replace accountable human judgment.

## Moonwitness XV principles

- Preserve the original.
- Trace each claim.
- Respect each perspective.
- Investigate openly.
- Create freely.
- Keep humans responsible.

Applied to gateway engineering: retain provenance, make configuration reviewable, expose useful operational signals, and never overstate what a policy or model can guarantee.

## Visual system

The design direction is **evidence-first moonpunk with zine-inspired contrast**: deep ink surfaces, paper-like light surfaces, a crescent-and-orbit identity, sharp typographic hierarchy, and clear operational labels.

| Token | Dark theme | Light theme | Use |
|---|---|---|
| Background | `#07060b` | `#f7f4ec` | Page canvas |
| Surface | `#13101d` | `#ede8dc` | UI panels |
| Surface 2 | `#1b172a` | `#e3ddd0` | Nested panels |\n| Text | `#f7f4ec` | `#07060b` | Primary reading text |\n| Muted text | `#9e96b0` | `#4a4557` | Secondary reading text |
| Signal yellow | `#ffd400` | `#7a6300` | Primary highlight |
| Signal pink | `#ff007a` | `#d4005a` | Secondary signal |
| Signal cyan | `#00d4ff` | `#0099b3` | Focus and connectivity signal |
| Moon purple | `#9d4edd` | `#7a3bb3` | Product identity and supporting accent |

Use the colors as a system, not decoration. Status and meaning must never depend on color alone. Maintain readable contrast, visible focus, keyboard navigation, and descriptive labels.

## Logo and assets

- `ui/src/assets/agw-light.svg` — light-theme sidebar wordmark.
- `ui/src/assets/agw-dark.svg` — dark-theme sidebar wordmark.
- `ui/src/assets/agw-icon.svg` and `ui/public/favicon.svg` — crescent/orbit mark.
- `img/banner-light.svg` and `img/banner-dark.svg` — README banner variants.
- `img/architecture.svg` — illustrative gateway request-flow diagram.

The `agw-*` filenames are retained because the UI imports those paths. Their artwork is Moonwitness Gateway branded; preserve the paths unless updating the imports in the same reviewed change.

## Voice and writing

Prefer precise, calm, accountable language. Describe capabilities as configurable and deployment-dependent when appropriate. Distinguish shipped behavior from roadmap concepts.

Use **Moonwitness Gateway** as the public product name. Keep upstream technical identifiers such as the Go module, Rust crates, binary names, APIs, and resource identifiers stable unless a separately reviewed compatibility migration is approved.

## Upstream provenance

This repository is a Moonwitness-branded fork of Agentgateway. Preserve Apache-2.0 licensing, copyright notices, and upstream attribution. Do not imply official affiliation or endorsement by upstream maintainers.


## Logo usage standards

### Approved variants

- **Primary wordmark:** use the light or dark wordmark that provides the strongest contrast with its immediate surface.
- **Symbol:** use the standalone crescent/orbit mark only where the product name is already visible nearby or the available space is too small for a wordmark.
- **README banners:** use the matching light/dark banner variant; do not recolor or crop the logo independently of its background.
- **Architecture artwork:** treat the architecture diagram as explanatory product documentation, not as a substitute for the logo.

### Clear space and minimum size

- Keep clear space around the mark equal to at least the diameter of the crescent's inner circle.
- Do not place borders, text, badges, or other logos inside that clear-space area.
- Use the standalone symbol at 24 × 24 CSS pixels or larger; below that size, prefer a simplified favicon treatment and test at actual display size.
- Use the wordmark at 160 CSS pixels wide or larger. If the layout cannot support that width, use the symbol with a visible adjacent product name.
- Preserve the source aspect ratio. Never stretch, skew, rotate, outline, blur, or add shadows to the logo.

### Prohibited treatments

- Do not substitute the upstream Agentgateway mark or imply upstream endorsement.
- Do not place the light wordmark on a light surface or the dark wordmark on a dark surface.
- Do not use the symbol alone as an unexplained status indicator.
- Do not add gradients, glow, animation, or extra orbital elements to the logo artwork.
- Do not split, reorder, or rewrite the MOONWITNESS / GATEWAY wordmark.

## Color application

Use the UI's semantic CSS variables rather than scattering raw brand hex values through new components. The canonical base tokens are:

| Token | Dark theme | Light theme | Intended role |
|---|---|---|---|
| Background | `#07060b` | `#f7f4ec` | Page canvas |
| Surface | `#13101d` | `#ede8dc` | Sidebar, cards, panels |
| Surface 2 | `#1b172a` | `#e3ddd0` | Nested and secondary surfaces |
| Text | `#f7f4ec` | `#07060b` | Primary reading text |
| Muted text | `#9e96b0` | `#4a4557` | Secondary information |
| Brand purple | `#9d4edd` | `#7a3bb3` | Brand identity and interactive emphasis |
| Signal yellow | `#ffd400` | `#7a6300` | Primary emphasis |
| Signal pink | `#ff007a` | `#d4005a` | Secondary emphasis |
| Signal cyan | `#00d4ff` | `#0099b3` | Connectivity and focus accent |

Semantic state colors (success, warning, information, and error) are functional UI colors, not logo colors. Keep their labels and icons visible, and verify text contrast in both themes. Never use the bright accent colors as small body text on a similarly bright surface.

## Typography and hierarchy

- Use Geist Sans for interface and body copy, with the documented system fallbacks.
- Reserve heavy display faces for short page titles, campaign graphics, or section accents; never use them for dense configuration, tables, or long-form reading.
- Use the established monospace stack for code, identifiers, logs, and configuration.
- Keep interface labels concise and sentence case. Reserve uppercase styling for short eyebrow labels and the wordmark.
- Avoid tracking-heavy all-caps paragraphs and avoid using color as the only way to establish hierarchy.

## Voice and approved language

### Preferred

- "Route model and tool traffic through a configurable gateway."
- "Apply configured authentication, authorization, and guardrail policies."
- "Inspect operational logs, metrics, and traces."
- "Human owners remain responsible for access and consequential decisions."

### Avoid

- "Guarantees truthful AI."
- "Makes agents fully autonomous."
- "Certifies every response."
- "Unbreakable security" or "zero-risk AI."

Describe capabilities as **supported**, **configurable**, or **deployment-dependent** when the exact behavior depends on runtime configuration or external providers. Use "Moonwitness Gateway" on first mention in a page or document; "the gateway" is acceptable after the product is unambiguous.

## Product claims and provenance

Every public capability statement should be traceable to one of the following:

1. A shipped implementation in this repository.
2. A documented upstream capability, explicitly identified as inherited.
3. A planned integration, explicitly labelled as planned or not yet implemented.

Do not use Moonwitness ecosystem names as though they were dependencies or integrations unless the relevant contract and implementation exist. Preserve Apache-2.0 notices and upstream copyright statements in all redistributions.

## Brand QA checklist

Before merging a brand-facing change, verify:

- [ ] Product name and tagline use the approved spelling and capitalization.
- [ ] Correct light/dark logo variant is used and remains legible at the rendered size.
- [ ] Logo proportions and clear space are preserved.
- [ ] Both UI themes have readable text, visible focus, and non-color status cues.
- [ ] Product claims distinguish shipped, inherited, and planned capabilities.
- [ ] Upstream provenance and license notices remain intact.
- [ ] Links, favicon, page metadata, and documentation references resolve.
- [ ] Screenshots or previews show the actual implementation rather than an illustrative mockup.
- [ ] Build, lint, and relevant UI tests pass for the exact commit.

This checklist is a review gate, not a claim that those checks have already been run.
