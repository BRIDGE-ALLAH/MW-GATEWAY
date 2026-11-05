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

| Token | Value | Use |
|---|---|---|
| Ink | `#07060b` | Dark canvas |
| Night surface | `#13101d` | Dark UI surfaces |
| Paper | `#f7f4ec` | Light canvas and high-contrast text |
| Signal yellow | `#ffd400` | Primary highlight |
| Signal pink | `#ff007a` | Secondary signal |
| Signal cyan | `#00d4ff` | Focus and connectivity signal |
| Moon purple | `#9d4edd` | Product identity and supporting accent |

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
