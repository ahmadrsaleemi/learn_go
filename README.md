# Learn Go

Hands-on Go practice built by a software engineer with **9 years of experience in .NET and PHP**.

## Why this repo exists

Coming from mature ecosystems, I already know how to design APIs, services, and production systems. What I needed first was fluency in Go itself — syntax, idioms, and the language’s concurrency model — before building REST APIs, gRPC services, or microservices.

This repository is that foundation: deliberate, file-by-file practice from basics through intermediate and advanced topics, so the next step (real APIs and services) rests on solid muscle memory rather than guesswork.

## Approach

- Transfer existing engineering judgment into Go, instead of relearning software development from scratch.
- Prefer working examples over long notes — each `.txt` file is runnable Go practice I wrote while learning.
- Cover what matters for production readiness first; revisit niche lecture topics (for example **signals** and **reflect**) after shipping API projects.

## Repository layout

### [`basics/`](basics/)

Language fundamentals: packages, imports, types, control flow, slices, maps, functions, and early struct usage.  
This is where Go’s syntax and day-to-day patterns become familiar when coming from .NET or PHP.

### [`intermediate/`](intermediate/)

Deeper language and standard-library skills: pointers, interfaces, generics, strings, time, JSON, files, hashing, CLI, and related utilities.  
These topics bridge “I can write Go” and “I can structure real application code.”

### [`advance/`](advance/)

Concurrency and coordination: goroutines, channels, `select`, context, wait groups, worker pools, mutexes, atomics, timers/tickers, and rate limiting.  
This is the core of thinking in Go for scalable services — before wiring that knowledge into APIs and microservices.

## About the practice files

The `.txt` files are intentional learning artifacts: small, focused programs that capture each concept as I practiced it.  
They are not production modules — they are the working notes that made Go syntax and concurrency concrete. Together, they form a personal path toward writing idiomatic, production-ready Go.

## What’s next

- Build REST and gRPC APIs using this foundation.
- Apply concurrency, context, and rate-limiting patterns in service design.
- Return to deferred topics (signals, reflect, and similar) once API work is underway.

## Resources

- [LinkedIn](https://www.linkedin.com/in/ahmad-raza-77805313a/)
