# 🚀 My Go Learning Journey

> **From Django Full-Stack Developer to Production-Ready Go Engineer**  
> 8-week intensive learning plan with public documentation

**Start Date:** 6/10/2025
**Target Completion:** 6/12/2025  
**Time Commitment:** ~1-3 hours/days (Depends)

---

## 📋 Overview

This repository tracks my journey learning Go (Golang) as an experienced Django/React developer. The goal is to build 3 production-quality projects while documenting the process on YouTube.

**Background:**

-   3 years full-stack experience (Django + React)
-   EC2/Kubernetes/Helm exposure for deployment
-   Goal: Become proficient in Go for cloud-native development

---

## 🎯 Learning Goals

By the end of 8 weeks, I will have:

-   ✅ Built 3 production-ready Go projects
-   ✅ Deployed applications to Kubernetes
-   ✅ Understanding of Go concurrency patterns
-   ✅ Published 4 YouTube videos documenting the journey
-   ✅ Active GitHub profile with consistent Go commits

---

## 📁 Repository Structure

```
go-learning-journey/
├── README.md                 # This file - main tracker
├── week-01-02-foundations/
│   ├── cli-tool-1/          # First simple CLI tool
│   └── concurrent-program/   # Concurrency practice project
├── week-03-04-rest-api/
│   └── [project-name]/      # Main REST API project
├── week-05-06-k8s-deploy/
│   ├── helm-charts/         # Helm configurations
│   └── ci-cd/               # GitHub Actions workflows
├── week-07-08-specialization/
│   └── [specialized-project]/
├── resources/
│   ├── notes.md             # Learning notes
│   └── useful-links.md      # Resource collection
└── youtube/
    └── video-scripts.md     # Video planning and scripts
```

---

## 📅 8-Week Roadmap

### **PHASE 1: Foundations (Week 1-2)**

#### Week 1: Core Syntax & First Code

-   [ ] **Day 1-3:** Complete [Tour of Go](https://go.dev/tour/)
-   [ ] **Day 1-3:** Build first CLI tool: `_______________`
-   [ ] **Day 4-7:** Study error handling, defer, pointers, structs
-   [ ] **Day 4-7:** Read [Effective Go](https://go.dev/doc/effective_go)
-   [ ] **Day 4-7:** Refactor CLI tool with proper patterns
-   [ ] **Day 4-7:** Write tests for CLI tool
-   [ ] **Milestone:** 1 CLI tool on GitHub with tests

#### Week 2: Concurrency

-   [ ] **Day 8-10:** Study goroutines, channels, select statements
-   [ ] **Day 8-10:** Watch [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
-   [ ] **Day 11-14:** Build concurrent web scraper OR parallel file processor
-   [ ] **Day 11-14:** Benchmark and prove performance improvement
-   [ ] **Milestone:** 2 working Go programs on GitHub
-   [ ] **YouTube Video #1:** "I built my first Go program - here's what surprised me"

---

### **PHASE 2: REST API Project (Week 3-4)**

**Chosen Project:** `[ ] GitHub Activity Dashboard | [ ] Log Aggregator | [ ] Webhook Relay`

#### Week 3: Core API Development

-   [ ] **Day 15-16:** Setup project structure (`/cmd`, `/internal`, `/pkg`)
-   [ ] **Day 15-16:** Basic API running with `net/http`
-   [ ] **Day 17-19:** Implement 3-4 REST endpoints
-   [ ] **Day 17-19:** Add request validation and JSON handling
-   [ ] **Day 17-19:** Create Postman collection
-   [ ] **Day 20-21:** Integrate PostgreSQL/SQLite
-   [ ] **Day 20-21:** Write database migrations
-   [ ] **Milestone:** Functional REST API with database

#### Week 4: Production Readiness

-   [ ] **Day 22-23:** Write unit tests for handlers
-   [ ] **Day 22-23:** Write integration tests with test database
-   [ ] **Day 22-23:** Achieve 70%+ test coverage
-   [ ] **Day 24-25:** Add logging middleware (`slog`)
-   [ ] **Day 24-25:** Add authentication middleware (JWT/API keys)
-   [ ] **Day 24-25:** Add recovery middleware
-   [ ] **Day 26-28:** Create multi-stage Dockerfile
-   [ ] **Day 26-28:** Setup docker-compose (app + database)
-   [ ] **Day 26-28:** Write complete README
-   [ ] **Milestone:** Production-ready API with full documentation
-   [ ] **YouTube Video #2:** "Building a production REST API in Go (from Django)"

---

### **PHASE 3: Observability & Deployment (Week 5-6)**

#### Week 5: Make It Observable

-   [ ] **Day 29-31:** Add Prometheus metrics endpoint
-   [ ] **Day 29-31:** Instrument requests, duration, error rate
-   [ ] **Day 29-31:** Run Prometheus locally and verify metrics
-   [ ] **Day 32-33:** Implement request ID tracing
-   [ ] **Day 32-33:** Add structured JSON logging
-   [ ] **Day 32-33:** Configure log levels
-   [ ] **Day 34-35:** Create `/health` endpoint (liveness)
-   [ ] **Day 34-35:** Create `/ready` endpoint (readiness)
-   [ ] **Day 34-35:** Implement graceful shutdown
-   [ ] **Milestone:** Fully observable API

#### Week 6: Kubernetes Deployment

-   [ ] **Day 36-38:** Write Deployment YAML
-   [ ] **Day 36-38:** Write Service YAML
-   [ ] **Day 36-38:** Create ConfigMap and Secret
-   [ ] **Day 36-38:** Deploy to local K8s (minikube/Docker Desktop)
-   [ ] **Day 39-40:** Convert manifests to Helm chart
-   [ ] **Day 39-40:** Parameterize values (replicas, image, resources)
-   [ ] **Day 39-40:** Test with different values files
-   [ ] **Day 41-42:** Create GitHub Actions workflow
-   [ ] **Day 41-42:** Automate Docker build and tests
-   [ ] **Day 41-42:** Push images to registry
-   [ ] **Milestone:** API deployed to K8s with CI/CD
-   [ ] **YouTube Video #3:** "Deploying Go to Kubernetes - complete setup"

---

### **PHASE 4: Specialization Project (Week 7-8)**

**Chosen Direction:** `[ ] Cloud-Native Tool | [ ] High-Performance Service | [ ] Developer CLI`

**Project Name:** `_______________`

#### Week 7-8: Build & Ship

-   [ ] **Day 43-46:** Build core functionality
-   [ ] **Day 43-46:** Get end-to-end working
-   [ ] **Day 47-50:** Handle edge cases
-   [ ] **Day 47-50:** Add comprehensive error handling
-   [ ] **Day 47-50:** Write tests
-   [ ] **Day 51-52:** Write complete README
-   [ ] **Day 51-52:** Add usage examples
-   [ ] **Day 51-52:** Create CONTRIBUTING guide
-   [ ] **Day 53-56:** Tag v1.0.0 release
-   [ ] **Day 53-56:** Write release notes
-   [ ] **Day 53-56:** Share on r/golang and Twitter
-   [ ] **Milestone:** Second major project released
-   [ ] **YouTube Video #4:** "8-week Go journey - [Project] showcase"

---

## 📚 Learning Resources

### Essential (Using Now)

-   [Tour of Go](https://go.dev/tour/) - Interactive tutorial
-   [Go by Example](https://gobyexample.com/) - Quick reference
-   [Effective Go](https://go.dev/doc/effective_go) - Official style guide

### Reference (As Needed)

-   [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) - TDD approach
-   [Practical Go](https://dave.cheney.net/practical-go) - Production patterns
-   [Go Blog](https://go.dev/blog/) - Deep dives on topics
-   Go standard library source code - Best learning tool

### Community

-   [r/golang](https://reddit.com/r/golang) - Daily browse
-   Gophers Slack - Questions when stuck
-   [Go Wiki](https://github.com/golang/go/wiki) - Comprehensive guides

---

## ⏰ Daily Routine

### Weekdays (2 hours)

```
08:00-08:30  →  Read/watch learning material
08:30-09:30  →  Hands-on coding on current milestone
09:30-10:00  →  Document (comments, README, video notes)
```

### Weekends (3-4 hours)

```
Saturday   →  Deep work on current project
Sunday     →  Code review, refactor, tests, next week planning
```

---

## 📊 Progress Tracking

### Week 1-2: Foundations

-   [ ] Tour of Go completed
-   [ ] CLI tool #1 shipped (`repo-name: _______________`)
-   [ ] Concurrent program working (`repo-name: _______________`)
-   [ ] Video #1 published (`link: _______________`)

### Week 3-4: REST API

-   [ ] API endpoints working
-   [ ] Database integrated
-   [ ] Tests written (70%+ coverage)
-   [ ] Dockerized & documented
-   [ ] Video #2 published (`link: _______________`)

### Week 5-6: K8s & Observability

-   [ ] Prometheus metrics added
-   [ ] Structured logging implemented
-   [ ] Deployed to local K8s
-   [ ] Helm chart created
-   [ ] CI/CD pipeline working
-   [ ] Video #3 published (`link: _______________`)

### Week 7-8: Specialization

-   [ ] Second project shipped (`repo-name: _______________`)
-   [ ] Open-source ready (README, tests, docs)
-   [ ] Video #4 published (`link: _______________`)

---

## 🎥 YouTube Channel

**Channel Name:** `_______________`  
**Channel Link:** `_______________`

### Published Videos

1. [ ] Week 2: "I built my first Go program - here's what surprised me"
2. [ ] Week 4: "Building a production REST API in Go (from Django)"
3. [ ] Week 6: "Deploying Go to Kubernetes - complete setup"
4. [ ] Week 8: "8-week Go journey - [Project] showcase"

---

## 📈 Success Metrics (8-Week Goal)

-   [ ] 3+ working Go projects on GitHub (all public)
-   [ ] Comfortable reading Go stdlib source code
-   [ ] Can build & deploy REST API from scratch
-   [ ] Solid understanding of concurrency patterns
-   [ ] Basic K8s deployment skills in Go context
-   [ ] 4 YouTube videos published
-   [ ] Active GitHub profile (green commit graph)
-   [ ] At least 1 open-source contribution to Go project

---

## 💭 Weekly Reflections

### Week 1

**What I learned:**  
`_______________`

**Challenges faced:**  
`_______________`

**Next week focus:**  
`_______________`

### Week 2

**What I learned:**  
`_______________`

**Challenges faced:**  
`_______________`

**Next week focus:**  
`_______________`

_(Continue for all 8 weeks)_

---

## 🔗 Project Links

### Week 1-2 Projects

-   CLI Tool #1: `[repo link]`
-   Concurrent Program: `[repo link]`

### Week 3-4 Projects

-   REST API: `[repo link]`

### Week 5-6 Projects

-   (Same API with K8s deployment)

### Week 7-8 Projects

-   Specialization Project: `[repo link]`

---

## 📝 Notes & Tips

### Key Learnings

-   `Add learnings as you go`

### Common Pitfalls

-   `Document mistakes to avoid`

### Useful Code Snippets

-   `Save reusable patterns`

---

## 🏆 Final Thoughts

_To be filled after completing 8 weeks_

**What worked well:**  
`_______________`

**What I'd do differently:**  
`_______________`

**Next steps in my Go journey:**  
`_______________`

---

## 📞 Connect

-   GitHub: `[your-github]`
-   YouTube: `[your-channel]`
-   Twitter: `[your-twitter]`
-   LinkedIn: `[your-linkedin]`

---

**Last Updated:** `[date]`  
**Current Week:** `[week number]`  
**Status:** 🟢 On Track | 🟡 Behind Schedule | 🔴 Need Help
