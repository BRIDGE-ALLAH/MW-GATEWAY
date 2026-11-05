# MOONWITNESS GATEWAY — PRODUCT & TECHNICAL BLUEPRINT

> **Connect Intelligence. Preserve Accountability.**
>
> **The Moon Watches. The Record Remembers. The Evidence Remains Inspectable.**

| Field | Decision |
|---|---|
| Product | Moonwitness Gateway |
| Repository | `BRIDGE-ALLAH/MW-GATEWAY` |
| Product class | AI connectivity gateway and policy-enforcement proxy |
| Parent vision | Moonwitness XV |
| Implementation baseline | Fork of `agentgateway/agentgateway` |
| License | Apache-2.0; preserve upstream license and attribution |
| Status | Product and brand baseline; not a production-readiness certificate |

This blueprint defines the product identity, responsibility boundaries, design language, and compatibility policy for Moonwitness Gateway. It does not claim that every capability described in the wider Moonwitness XV vision is implemented in this repository.

---

## 1. Product definition

**Moonwitness Gateway is the connectivity and policy-enforcement layer between AI applications, model providers, agent peers, tool servers, and external services.**

It provides a configurable proxy surface for LLM, MCP, A2A, and related HTTP traffic. Its inherited implementation supplies routing, security controls, guardrail integrations, and observability capabilities according to the selected runtime and configuration.

The gateway's purpose is to make AI connectivity:

1. **Connected** — models, agents, and tools communicate through explicit interfaces.
2. **Governed** — authentication, authorization, routing, and policy boundaries can be configured and enforced.
3. **Observable** — operations produce signals that operators can inspect and investigate.
4. **Accountable** — people remain responsible for configuration, access, deployment, and consequential decisions.

### Product promise

> Connect intelligence without hiding the path, removing the boundary, or transferring human responsibility to a model.

### What the product is not

Moonwitness Gateway is not:
- A truth-verification engine or an authority that certifies model responses as facts.
- A replacement for ARSY, the central Hub in Moonwitness XV.
- A canonical identity provider, business database, or ecosystem-wide source of truth.
- A human approval authority or a system that grants agents unrestricted autonomy.
- An implementation of the entire Moonwitness XV knowledge, witness, research, or creative ecosystem.

A model response is a model response. A trace is an operational record. Neither is automatically a verified claim.

---

## 2. Relationship to Moonwitness XV

Moonwitness Gateway adopts the principles in the [Moonwitness XV Master Vision Narrative](https://github.com/bjo163/moonwitnessxv/blob/main/docs/vision/MASTER_VISION_NARRATIVE_XV.md) and the [Moonwitness XV Master Design](https://github.com/bjo163/moonwitnessxv/blob/main/MASTER_DESIGN.md).

The relevant project principle is:

**Preserve the original. Trace each claim. Respect each perspective. Investigate openly. Create freely. Keep humans responsible.**

For this repository, the principle is applied to infrastructure boundaries:
- Preserve request and configuration context needed for operational investigation.
- Route traffic through explicit, reviewable configuration.
- Apply authorization and policy controls at the point where they are configured.
- Expose observability signals without presenting them as proof of semantic truth.
- Keep approval, rights, consent, and publication decisions with their designated human-governed systems.

### Ecosystem responsibility map

| Component | Ecosystem responsibility | Gateway relationship |
|---|---|---|
| Moonwitness XV | Overall knowledge, witness, research, creation, and accountable-action vision | Parent product context; not a runtime dependency |
| ARSY | Single application Hub and canonical run-state owner in XV | Gateway may serve configured traffic; it does not become a second Hub |
| AI Gateway | Model connectivity and inference-routing capability in the ecosystem | Define one clear routing boundary per deployment; avoid duplicate authorities |
| LUMEN | Accepted/planned agent framework and internal worker boundary | Agent execution must remain inside approved ARSY and human-governance contracts |
| HEADQUARTERS / 100CENTER | Human review, governance, and approval | Keeps consequential decisions with authorized people |
| ORBIT | Planned operational presentation surface | May display authorized gateway/runtime state; owns no backend or identity |

The table describes responsibility boundaries. It does not assert that integrations with these components are already implemented in this repository.

---

## 3. Functional scope

The following is the inherited capability set of the Agentgateway implementation. Exact availability depends on runtime mode, configuration, external providers, and the deployed environment.

### 3.1 LLM connectivity
- Unified OpenAI-compatible API for supported model traffic.
- Multiple model providers and provider-specific configuration.
- Routing, load balancing, and failover.
- Budget and spend controls where configured.
- Prompt enrichment and related LLM traffic policies.

### 3.2 MCP connectivity
- Model Context Protocol client/server connectivity.
- Tool federation and external data/tool access.
- Supported stdio and network transports, including HTTP-based MCP modes.
- OpenAPI integration and supported authentication flows.

### 3.3 A2A connectivity
- Agent-to-agent communication using the A2A protocol.
- Agent capability discovery and task-oriented exchanges.
- Interoperability across compatible agent implementations.

### 3.4 Inference routing
- Routing support for configured self-hosted and Kubernetes inference environments.
- Infrastructure-aware decisions when the relevant signals and integrations are available, such as GPU utilization, queue depth, KV cache, or LoRA adapter information.

### 3.5 Policy and guardrails
- Authentication and authorization mechanisms supported by the implementation.
- RBAC and CEL-based policy evaluation.
- Rate limiting and TLS.
- Configurable content-filtering and guardrail integrations, including external providers and webhooks.

Guardrails enforce configured rules. They do not establish the truth of a response or guarantee that a model cannot produce harmful or incorrect output.

### 3.6 Observability
- Metrics, logs, and distributed tracing.
- OpenTelemetry integrations.
- Operational context for debugging, capacity analysis, and policy investigation.

Telemetry should be protected and retained according to the deployment's privacy, security, and retention requirements.

---

## 4. System boundary

### Gateway owns
- Configured gateway listeners and proxy routing.
- Configured LLM, MCP, A2A, and HTTP connectivity.
- Policy evaluation and enforcement supported by the implementation.
- Gateway configuration and runtime state exposed by its own interfaces.
- Configured gateway metrics, logs, and traces.

### Gateway does not own
- Moonwitness XV's canonical identity or application run state.
- A separate SQL or ecosystem-wide data authority.
- Human approvals, rights decisions, consent, or publication authority.
- Canonical source texts, research conclusions, or witness archives.
- An autonomous agent's authority to expand its permissions.

### Request-flow model

`Client / Agent → Moonwitness Gateway → Configured Model, Tool, or Agent Destination`

Within the gateway boundary, configured authentication, authorization, routing, and guardrail policies are applied. Operational signals may be emitted to configured observability backends.

A request path is not a guarantee that every destination is trustworthy. Operators must validate destination identity, TLS, credentials, policy, and data handling.

---

## 5. Security and governance requirements

1. **Least privilege:** use narrowly scoped credentials and policies.
2. **Explicit authentication:** do not expose administrative or sensitive surfaces without an appropriate authentication design.
3. **Authorization before access:** verify the policy applied to each protected route and resource.
4. **Secret handling:** do not place tokens, passwords, or private keys in source control, logs, screenshots, or public issue reports.
5. **Transport protection:** configure TLS and verify certificate and trust boundaries for production traffic.
6. **Rate and resource controls:** configure limits appropriate to the gateway's workload and threat model.
7. **Guardrail transparency:** document which controls are enabled, their providers, their failure modes, and their scope.
8. **Observability hygiene:** avoid logging secrets or unnecessary personal/sensitive content; set retention and access policies.
9. **Human accountability:** policy ownership, approvals, incident response, and production rollout remain assigned to people.
10. **No invented assurance:** do not describe a route as secure, verified, compliant, or production-ready without corresponding evidence.

---

## 6. Brand constitution

### Product name
**Moonwitness Gateway**

Use this capitalization in product-facing documentation, UI titles, diagrams, release notes, and presentations.

### Tagline
**Connect Intelligence. Preserve Accountability.**

### Supporting line
**Secure Connectivity. Observable Intelligence. Accountable Technology.**

### Brand narrative

The Moon watches. The record remembers. The evidence remains inspectable.

Moonwitness Gateway connects models, agents, tools, and services through an explicit infrastructure boundary. It makes configured paths easier to govern and operational behavior easier to inspect, while preserving the distinction between machine output, evidence, interpretation, and human decision.

### Brand voice
- Precise, calm, technically credible, and human-readable.
- Confident about implemented behavior; explicit about limitations.
- Never imply that AI is a source of divine, legal, moral, or absolute authority.
- Do not use "truth engine", "self-governing intelligence", or "autonomous authority" as product claims.
- Distinguish current functionality from roadmap proposals.

### Name separation
- **Moonwitness Gateway** — this repository's gateway product and connectivity layer.
- **Moonwitness XV** — the wider ecosystem vision and implementation repository.
- **ARSY** — the single Hub in Moonwitness XV; not a synonym for this gateway.
- **AI Gateway** — the ecosystem's model-connectivity capability; deployment topology must avoid duplicate routing authorities.
- **Agentgateway** — the upstream project and inherited technical identifiers. Use this name when describing provenance, retained package names, or upstream documentation.

---

## 7. Visual design system

### Visual grammar

The interface follows the Moonwitness XV **evidence-first moonpunk / zine** direction, adapted for an operational gateway console:

- High-contrast dark surfaces and readable light-mode surfaces.
- Strong, simple hierarchy; uppercase display treatments used sparingly.
- Yellow, pink, cyan, and moon-purple accents with clear semantic roles.
- Restrained geometric/orbital marks and crisp vector artwork.
- Functional density without sacrificing readable labels, focus states, or whitespace.
- No decorative status indicators that imply a system state not confirmed by the runtime.

### Canonical palette

The palette is derived from the Moonwitness XV shared design tokens.

| Token | Dark value | Light value | Role |
|---|---|---|---|
| Ink / background | `#07060b` | `#f7f4ec` | Main canvas |
| Surface | `#13101d` | `#ede8dc` | Panels and sidebar |
| Surface 2 | `#1b172a` | `#e3ddd0` | Secondary surfaces |
| Foreground | `#f7f4ec` | `#07060b` | Main text |
| Muted | `#9e96b0` | `#4a4557` | Secondary text |
| Yellow | `#ffd400` | `#7a6300` | Primary emphasis |
| Pink | `#ff007a` | `#d4005a` | Secondary emphasis |
| Cyan | `#00d4ff` | `#0099b3` | Signal and keyboard focus |
| Moon | `#9d4edd` | `#7a3bb3` | Brand identity |
| Success | `#00e676` | `#00b35c` | Positive operational state |

Status colors must always be paired with text or an accessible icon/label. Color alone must not communicate a security, health, or request state.

### Typography
- UI/body: Geist Sans with Inter and system sans-serif fallbacks.
- Display: Impact, Anton, Arial Black, or a system heavy sans-serif fallback.
- Code/configuration: JetBrains Mono, DejaVu Sans Mono, or a system monospace fallback.
- Preserve readable line length and avoid using display fonts for long explanations, logs, or configuration.

### Logo system
- Primary mark: moon/eclipse motif crossed by an orbit or signal path.
- Wordmark: **MOONWITNESS** with **GATEWAY** as the product descriptor.
- Use the light wordmark on light surfaces and the dark-surface wordmark on dark surfaces.
- Keep the mark legible at favicon scale; do not add fine detail that disappears at small sizes.
- Prefer vector SVG for the logo, favicon, and diagrams.
- Do not reuse upstream Agentgateway marks as Moonwitness identity.

### Dashboard composition
- One navigation shell and one main content area.
- Sidebar identity remains visible but does not overpower operational information.
- Clear page titles, breadcrumbs, section labels, and predictable controls.
- Dense configuration and telemetry tables remain scannable.
- Forms use explicit labels, visible validation feedback, and clear save/apply states.
- Dialogs and destructive operations must identify the affected resource and allow a deliberate confirmation.
- Respect keyboard navigation, visible focus, screen-reader labels, and reduced-motion preferences.

### Operational language
Use explicit words for runtime states such as `LIVE`, `STALE`, `OFFLINE`, `UNKNOWN`, `BLOCKED`, `PENDING`, and `ERROR` only when the underlying runtime supports that state. Never invent counts, success indicators, or freshness badges for visual effect.

---

## 8. Compatibility and naming policy

This repository is derived from Agentgateway. The implementation contains established identifiers that may be referenced by build files, code imports, CLI users, containers, Kubernetes resources, test fixtures, examples, environment variables, APIs, and upstream documentation.

### Branding changes allowed without a compatibility migration
- README and product-facing documentation.
- UI title, visible wordmarks, favicon, and help links.
- User-facing welcome copy and descriptive labels.
- Product banners and first-party architecture illustrations.
- Product metadata where it can be changed without altering runtime contracts.

### Identifiers retained for compatibility
- Rust crate/package identifiers such as `agentgateway` and `agentgateway-app`.
- The current CLI command name `agentgateway`.
- Go module/import paths under `github.com/agentgateway/agentgateway`.
- Existing API paths, environment-variable names, configuration keys, and protocol fields.
- Kubernetes resource names, Helm/chart identifiers, container labels, and examples that are part of existing operational contracts.
- Upstream directory and source provenance.

A future technical rename must be handled as a separately reviewed migration with an explicit inventory, compatibility aliases or a deprecation window, updated generators and tests, and rollback evidence. Global search-and-replace is not an acceptable migration strategy.

---

## 9. Upstream provenance and licensing

Moonwitness Gateway is a branded fork of [agentgateway/agentgateway](https://github.com/agentgateway/agentgateway).

- Preserve the Apache License 2.0 and all upstream copyright and attribution notices.
- Preserve required notices in source files and distributions.
- Do not imply that this fork is the official Agentgateway project or is endorsed by the Linux Foundation or upstream maintainers.
- Keep upstream documentation links clearly identified as upstream references until Moonwitness-hosted documentation is available.
- Product branding does not transfer ownership of upstream code, trademarks, or third-party assets.

---

## 10. Delivery phases

| Phase | Scope | Acceptance evidence |
|---|---|---|
| P0 — Brand baseline | README, this blueprint, title, UI logo, favicon, visible links, design tokens, product diagrams | Source review, rendered UI review, documentation review |
| P1 — Brand consistency | Audit remaining first-party UI strings, CLI help, image references, release metadata, examples, and docs entry points | Search inventory and compatibility review |
| P2 — Moonwitness docs | Publish a Moonwitness documentation entry point while maintaining explicit upstream technical references | Versioned docs, tested links, release mapping |
| P3 — Ecosystem integration | Configure and document explicit contracts with Moonwitness XV components where needed | Integration tests, authorization review, owner approval |
| P4 — Release governance | Release process, security review, upgrade/rollback guidance, and operational evidence | Exact-version build/test, deployment rehearsal, restore and rollback records |

No phase is complete merely because files exist or a local build passes. Acceptance requires evidence appropriate to the change.

---

## 11. Verification checklist for branding changes

- [ ] README describes the product accurately and attributes the upstream project.
- [ ] UI title, favicon, logos, and help links use Moonwitness Gateway identity.
- [ ] UI remains readable in light and dark themes.
- [ ] Accent colors have sufficient contrast and are never the only status signal.
- [ ] User-facing copy does not imply autonomous or absolute authority.
- [ ] Existing API, CLI, package, and deployment contracts remain unchanged.
- [ ] UI build, typecheck, lint, and relevant tests pass.
- [ ] Gateway build and relevant tests pass.
- [ ] Generated assets and screenshots are reviewed in both themes.
- [ ] Production readiness is not claimed without deployment-specific evidence.

**Verification status of this branding patch:** source files are being updated; automated build, UI tests, and deployment validation must still be run in the project environment.

---

## 12. Change ownership and AI-assistance disclosure

This branding and documentation revision was AI-assisted at the explicit request of the repository owner. The repository owner/maintainer remains responsible for reviewing, understanding, testing, and accepting the changes in accordance with the repository's contribution and community policies.

The brand baseline does not override upstream governance, the Apache license, or any project's security and contribution requirements.
