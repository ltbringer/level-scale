# 📦 E-commerce System Benchmark: Scaling by Design

This project simulates how a modern **e-commerce backend** might scale from a **solo developer prototype** to a **globally distributed, high-consistency platform**. It's structured around **realistic growth levels**, each introducing new system constraints, architectural shifts, and operational challenges.

You'll work through increasingly complex scenarios covering:

- 🔧 Schema and query design
- 📊 Write-heavy database loads
- 🔁 Replication and partitioning
- 🛠️ Observability and incident response
- 🌐 Global data distribution and fault-tolerance

---

## 🎯 Why E-commerce?

E-commerce systems must prioritize **strong consistency**: orders, inventory, and payment records 
**must not be lost or duplicated**. As traffic grows, so do the demands for availability, observability, 
and fault isolation. This project reflects real operational trade-offs developers and SREs face daily 
when building production-grade systems.

Whether you're tuning WAL settings, scaling read replicas, defining SLAs, or simulating geo-redundant clusters, 
this project grows with you -- just like your business would.

---

## 🚀 Level-Based System Scaling Plan (10 RPS → 1M RPS)

| Level | Target RPS | Estimated DAU | Company Stage       | Engineering Level                  | Key Focus Areas                                            |
|-------|------------|---------------|---------------------|------------------------------------|------------------------------------------------------------|
| 1     | 10 RPS     | ~1,000 DAU    | Solo Dev / MVP      | Mid-level engineer                 | Schema design, single-node PostgreSQL, basic metrics       |
| 2     | 100 RPS    | ~10,000 DAU   | Seed / Pre-Series A | Senior engineer                    | Indexing, connection pooling, basic joins, dashboards      |
| 3     | 1K RPS     | ~100,000 DAU  | Series A–B          | Staff engineer                     | Partitioning, WAL tuning, write-heavy simulation           |
| 4     | 10K RPS    | ~1M DAU       | Post-Series C       | Staff or Principal engineer        | Replication, cache, long-join optimization, VPCs, K8s      |
| 5     | 100K RPS   | ~10M DAU      | Pre-global scale    | Staff + SRE collaboration          | Queues, read/write split, Redis, SLOs, distributed metrics |
| 6     | 1M RPS     | ~100M+ DAU    | Global distributed  | Principal engineer + platform team | Sharding, async writes, geo-routing, chaos testing         |

---

## 🧠 Learning Objectives Per Level

### 🟢 Level 1: MVP (10 RPS)
- Functional relational schema
- Basic insert + join operations
- Prometheus + Grafana

### 🔵 Level 2: Small Production (100 RPS)
- Identify slow queries
- Add indexes
- Tune connection pool (e.g., PgBouncer)
- Measure disk vs memory pressure
- SLAs and Incident Management

### 🟡 Level 3: Write Scale (1K RPS)
- Add read replicas
- Add partitions
- WAL tuning
- Richer dashboards

### 🟠 Level 4: Growth Phase (10K RPS)
- Monitor replica lag, cache hit rate
- Compare MongoDB vs Postgres for feed model
- Optimize long joins
- Use query analysis tools
- Custom networking: VPCs
- Kubernetes

### 🔴 Level 5: Pre-Global Scale (100K RPS)
- Introduce query and write split (read/write path separation)
- Redis for caching and rate-limiting
- Optimize background jobs and queues (e.g., like/comment workers)
- Scale monitoring pipelines (Prometheus federation, Grafana Loki)
- External SLA enforcement (e.g., alert budget, SLO tracking)
- Chaos testing

### 🟣 Level 6: Global Distributed (1M RPS)
- Shard database (Citus/Postgres or Mongo)
- Simulate multi-region writes and reads
- Use async models (e.g., event sourcing, log-based ingestion)
- Load test geo-aware routing and failover
- Introduce chaos engineering basics (network partition, node loss)

---

## 🔧 Tooling Overview

| Category         | Tool(s)                          |
|------------------|----------------------------------|
| Metrics          | Prometheus, Grafana              |
| DB Engines       | PostgreSQL, MongoDB (opt-in)     |
| Caching/Queue    | Redis, Celery/RQ (optional)      |
| Load Simulation  | Python/Go scripts, ThreadPool    |
| Containerization | Docker, Kubernetes, Helm         |
| Networking       | Custom Docker bridges, mock VPCs |
| Alerts/SLOs      | Alertmanager, Loki, Blackbox     |

---
