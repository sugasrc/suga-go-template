# <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" height="24"> Go on Suga <img src="https://github.com/sugasrc.png" height="24">

Create and deploy your Go apps to your Kubernetes cluster with Suga in minutes.

## Getting Started

### 1. Create your repository

Click **Use this template** to create your own repository.

### 2. Develop locally

```bash
go run ./cmd/api
```

Server runs on http://localhost:8080

### 3. Push your changes

This triggers GitHub Actions to build and push your image to GitHub Container Registry.

### 4. Update Suga

In the Suga dashboard, update your service's **Image URI** to:

```
ghcr.io/YOUR_USERNAME/YOUR_REPO:latest
```

> **Note:** Ensure your GHCR package is public.
