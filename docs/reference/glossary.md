---
layout: default
title: Glossary
description: Take a look here for clear definitions of basic comcepts used in the lakeFS documentation.
parent: Reference
nav_order: 6
has_children: false
---

# Glossary

## Auditing

Data auditing is data assessment to ensure its accuracy, security, and efficacy for specific usage. It also involves assessing data quality through its lifecycle and understanding the impact of poor quality data on the organization's performance and revenue. Ensuring data reproducibility, auditability, and governance is one of the key concerns of data engineers today.

## Branch

A branch is a mutable pointer to a commit and its staging area (i.e., mutable storage where developers can create, update, and delete objects). When a user creates a commit from a branch, all the files from the staging area will be merged into the contents of the current branch, generating a new set of objects. The pointer is updated to reference the new set of objects. The new branch tip is set to the latest commit and the previous branch tip serves as the commit's parent. Learn more about the lakeFS branching model [here](https://docs.lakefs.io/branching/model.html).

## Commit

A commit is a point-in-time snapshot of the branch. It's a collection of object metadata and data, including paths and the object contents and metadata. Commits have their own commit metadata. Every repository has one initial commit with no parent commits. If a commit has more than one parent, it is a merge commit. lakeFS supports only merge commits with two parents.

## Cross-collection consistency

In lakeFS, a repository (and thus a branch) can span multiple collections. Branches provide lakeFS users with consistency guarantees across different logical collections.

## Data lifecycle management

In data-intensive applications, data should be managed through its entire lifecycle similar to how teams use to manage code. By doing so, we could leverage the best practices and tools from application lifecycle management (like CI/CD operations) and apply them to data. lakeFS offers data lifecycle management via [isolated data development environments](https://docs.lakefs.io/use_cases/iso_env.html) instead of shared buckets.

## Data pipeline reproducibility

Reproducibility in data pipelines is the ability to recreate an issue that occurred in the pipeline. Reproducibility allows for the controlled manufacture of an error to debug and troubleshoot it at a later point in time. Reproducing a data pipeline issue is a challenge that most data engineers face on a daily basis. Learn more about how lakeFS supports data pipeline [reproducibilit](https://docs.lakefs.io/use_cases/reproducibility.html)y.

## Data quality testing

This term describes ways to test data for its accuracy, completeness, consistency, timeliness, validity, and integrity.

## Data versioning

To version data means creating a unique reference for data collection. This reference can take the form of a query, an ID, or also commonly, a DateTime identifier. Data versioning may also include saving an entire copy of the data under a new name or file path every time you want to create a version of it. More advanced versioning solutions like lakeFS optimize storage usage between versions and expose special operations to manage them.

## Environment variables

To support container-based environments, lakeFS can be configured using environment variables. See the [reference](https://docs.lakefs.io/reference/configuration.html#using-environment-variables) for a complete list of environment variables.

## Git-like operations

lakeFS allows teams to treat their data lake as a Git repository. Git-like operations over data are infrastructure, not a feature provided by an application, with data versioning as an essential part of that infrastructure.

## Graveler

Graveler is responsible for handling versioning in lakeFS by translating lakeFS addresses to the actual stored objects. See the [data model section](https://docs.lakefs.io/understand/versioning-internals.html) to learn about the data model used to store lakeFS metadata.

## Hooks

lakeFS hooks allow you to automate and ensure that a given set of checks and validations happens before important lifecycle events. They are similar conceptually to [Git Hooks](https://git-scm.com/docs/githooks), but in contrast, they run remotely on a server, so they are guaranteed to happen when the right event is triggered. Currently, lakeFS allows executing hooks when two types of events occur: pre-commit events that run before a commit is acknowledged and pre-merge events that trigger right before a merge operation.

## Isolated data snapshot

Creating a branch in lakeFS provides an isolated environment containing a snapshot of your repository. While working on your branch in isolation, all other data users will be looking at the repository's main branch. So they won't see your changes, and you also won't see the changes applied to the main branch. All of this happens without any data duplication but metadata management.

## Main branch

Every Git repository has the main branch (unless you take explicit steps to remove it). The main branch plays a key role in the software development process. In most projects, it represents the source of truth - all the code that works has been tested and is ready to be pushed to production.

## Metadata management

​​Where there's data, there's also metadata. lakeFS uses metadata to define schema, data types, [data versions](https://lakefs.io/data-versioning/), relations to other datasets, etc. This helps to improve discoverability and manageability. 

## Merge

In lakeFS, once you commit both the code and data, you can review them together and then merge the new data into the main branch. A merge generates a commit on the target branch with all your changes. Committing is a fast and atomic metadata operation in lakeFS as it doesn't involve copying data.

## Repository

A repository is a collection of objects with common history tracking. lakeFS manages versions of the repository identified by their commits.

## Rollback

If a developer introduces a new code version to production and discovers that it has a critical bug, they can simply roll back to the previous version. In lakeFS, a rollback is an atomic action that prevents the data consumers from receiving low-quality data until the issue is resolved. Learn more about how lakeFS supports the [rollback](https://docs.lakefs.io/use_cases/rollback.html) operation.