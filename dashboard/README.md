<p align="center">
  <img src="../assets/schema-drift-banner.svg" alt="Schema Drift Observatory animated banner" width="100%" />
</p>

<h1 align="center">Schema Drift Observatory — Dashboard Prototype</h1>

<p align="center">
  <strong>A Vue 3 interface prototype for exploring API contract-diff report examples.</strong>
</p>

<p align="center">
  <a href="../README.md">Main project README</a> ·
  <a href="../README.md#known-limitations-and-audit-notes">Known limitations</a>
</p>

## Current scope

This folder contains a Vue 3 + TypeScript + Vite dashboard prototype. It is a visual companion to the Schema Drift Observatory CLI, not a live-connected frontend for the comparison engine. The currently bundled sample/example data should not be interpreted as the result of a fresh CLI run.

## Stack

- Vue 3
- TypeScript
- Vite

## Run locally

From this directory:

```bash
npm install
npm run dev
```

Create a static build:

```bash
npm run build
npm run preview
```

The generated assets are for the dashboard prototype only. The root README explains how the Go CLI is built and run.

## Current limitations

- The dashboard does not currently invoke the Go comparison engine.
- Example/sample report figures are not live analysis results.
- Type/enum comparison and other coverage gaps are described in the root project's audit notes.

## Creator, Developer & Maintainer

**Created, designed, developed, and maintained by [Nurul Faiyaz](https://github.com/tracolerd).**

This dashboard prototype is independently built and maintained by Nurul Faiyaz.

<p align="center">
  <a href="https://www.facebook.com/faiyaz.888"><img src="https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white" alt="Facebook" /></a>
  <a href="https://www.instagram.com/tracolerd"><img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white" alt="Instagram" /></a>
  <a href="https://www.linkedin.com/in/faiyaz-xyz"><img src="https://img.shields.io/badge/LinkedIn-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white" alt="LinkedIn" /></a>
  <a href="https://x.com/faiyaz_explores"><img src="https://img.shields.io/badge/X-000000?style=for-the-badge&logo=x&logoColor=white" alt="X" /></a>
  <a href="https://github.com/tracolerd"><img src="https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white" alt="GitHub" /></a>
  <a href="mailto:scrollfaiyaz@gmail.com"><img src="https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white" alt="Email" /></a>
  <a href="https://nurul-faiyaz.netlify.app/"><img src="https://img.shields.io/badge/Portfolio-Visit%20Website-0F766E?style=for-the-badge&logo=googlechrome&logoColor=white" alt="Personal website" /></a>
</p>
