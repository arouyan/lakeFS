---
layout: default
title: What is lakeFS
description: A lakeFS documentation website that provides information on how to use lakeFS to deliver resilience and manageability to data lakes.
nav_order: 0
redirect_from: ./downloads.html
---

# What is lakeFS?
{: .no_toc }  

lakeFS is an open-source project that provides a Git-like version control interface for data lakes, with seamless integration with most data tools and frameworks. lakeFS enables you to easily implement staging environments through branching, parallel pipelines for experimentation, data reproducibility, rollback, and data quality validation before it's marked production-ready with pre-merge hooks.

lakeFS supports AWS S3, Azure Blob Storage, and Google Cloud Storage (GCS) as its underlying storage service. It's [API-compatible with S3](reference/s3.md) and works seamlessly with all modern data frameworks such as Spark, Hive, AWS Athena, Presto, etc.

<img src="{{ site.baseurl }}/assets/img/wrapper.png" alt="lakeFS" width="650px"/>


{: .pb-5 }


## New! lakeFS Playground

Experience lakeFS first hand with your own isolated environment. You can easily integrate it with your existing tools and see lakeFS in action - in an environment similar to your own.

<p style="margin-top: 40px;">
    <a class="btn btn-green" href="https://demo.lakefs.io/" target="_blank">
        Try lakeFS now without installing
    </a>
</p>

## Why use lakeFS?

With lakeFS, data teams can stop worrying about issues that occupy much of their time:
* Roll back changes to data: recover quickly from bugs and mistakes on your production data.
* Test and validate data before propagating it to consumers.
* Reproduce and debug data issues by traveling back in time, across data versions and data collections.

Since lakeFS is compatible with the S3 object storage API, all popular applications will work without modification simply by adding the branch name to the object path:

<img src="{{ site.baseurl }}/assets/img/s3_branch.png" alt="lakeFS s3 addressing" width="60%" height="60%" />

## What you can do with lakeFS

lakeFS enhances processing workflows at each step of the data lifecycle:

### In development
* **Experiment** - try new tools, upgrade versions, and evaluate code changes in isolation. By creating a branch of the data you get an isolated snapshot to run experiments over, while others are not exposed. Compare between branches with different experiments or to the main branch of the repository to understand the impact of a change.  
* **Debug** - check out specific commits in a repository's commit history to materialize consistent, historical versions of your data. See the exact state of your data at the point-in-time of an error to understand its root cause.
* **Collaborate** - avoid managing data access at the two extremes of either 1) treating your data lake like a shared folder, or 2) creating multiple copies of the data to safely collaborate. Instead, leverage isolated branches managed by metadata (not copies of files) to work in parallel.

[Learn more](https://docs.lakefs.io/usecases/data-devenv.html)

### During deployment
* **Version Control** - retain commits for a configurable duration, so readers are able to query data from the latest commit or any other point in time. Writers introduce new data atomically, preventing inconsistent data views.
* **Test** - define pre-merge and pre-commit hooks to run tests that enforce schema and validate properties of the data to catch issues before they reach production.

[Learn more](https://docs.lakefs.io/usecases/ci.html)

### In production
* **Roll Back** - recover from errors by instantly reverting data to a former, consistent snapshot of the data lake. Choose any commit in a repository's commit history to revert in one atomic action.
* **Troubleshoot** - investigate production errors by starting with a snapshot of the inputs to the failed process. Spend less time re-creating the state of datasets at the time of failure, and more time finding the solution.
* **Cross-collection Consistency** - provide consumers multiple synchronized collections of data in one atomic, revertable action. Using branches, writers provide consistency guarantees across different logical collections - merging to the main branch only after all relevant datasets have been created or updated successfully.
   
## Additional things you should know about lakeFS: 
* It'is **format agnostic**
* Your data **stays in place**
* It helps you prevent data duplication by using **copy-on-write**
* It's **highly performant** for huge data lakes
* It includes **configurable garbage collection** capabilities
* lakeFS is **highly available and production ready**

<img src="{{ site.baseurl }}/assets/img/lakeFS_integration.png" alt="lakeFS integration into data lake" width="60%" height="60%" />

[Learn more](https://docs.lakefs.io/usecases/production.html)

## Downloads

### Binary Releases

Binary packages are available for Linux/macOS/Windows on [GitHub Releases](https://github.com/treeverse/lakeFS/releases){: target="_blank" }

### Docker Images

Official Docker images are available at [https://hub.docker.com/r/treeverse/lakefs](https://hub.docker.com/r/treeverse/lakefs){: target="_blank" }

## Features and roadmap

Take a deeper look into the [features of lakeFS](https://docs.lakefs.io/understand/) and check out our [roadmap](https://docs.lakefs.io/understand/roadmap.html) to see what's in store for lakeFS in the future.

## Community and support

Stay up to date and get lakeFS support via:

- [Slack](https://lakefs.io/slack) (to get help from our team and other users)
- [Twitter](https://twitter.com/lakeFS) (follow for updates and news)
- [YouTube](https://lakefs.io/youtube) (learn from video tutorials)
- [Contact us](https://lakefs.io/contact-us/) (reach out to us)

## Integrations 

lakeFS provides many [integrations](https://docs.lakefs.io/integrations/) that come in handy, be sure to check them out.

## How to contribute

Woudl you like to contribute to lakeFS? Please read [this guide](https://docs.lakefs.io/contributing.html) - we look forward to seeing your contributions!

## FAQ

Have a question about lakeFS? Check our [FAQ](https://docs.lakefs.io/faq.html) to learn the essentials of this project.

## License

lakeFS is completely free and open-source and licensed under the [Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).

## Get started

Get started and [set up lakeFS on your preferred cloud environment](https://docs.lakefs.io/deploy/).
