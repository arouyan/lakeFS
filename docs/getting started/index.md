---
layout: default
title: Getting started
description: What is lakeFS and what can you do with it? Find out in this section.
nav_order: 4
has_children: true
has_toc: true
---

## What is lakeFS?

lakeFS transofrms an object storage into a data lake repository that exposes a Git-like interface to the underlying data so that users can use the same development workflows for code and data. The open-source lakeFS project supports AWS S3, Azure Blob Storage, and Google Cloud Storage (GCS) 
as its underlying storage service. It's API-compatible with S3 and integrates seamlessly with popular data frameworks 
such as Spark, Hive, dbt, Trino, and many others.

## Why does lakeFS exist?

In today's data world, there are many moving parts. There are developers writing machine learning code, business analysts pulling reports, 
and more consumers of data in every company. 

lakeFS allows you to keep all your data in one place. Since the data stays in object storage, there’s no need to lift and shift it anywhere. 
Another benefit is that you don’t need to be competitive and run applications over object storage and serve data to the organization. 
lakeFS works like a wrapper around parts of the data lake that you’d like to version control. 

lakeFS comes in handy when you want to develop and test in isolation over object storage, manage a resilient production environment through versioning, 
and achieve great collaboration between team members.

## What can you do with lakeFS?

You can use lakeFS across the entire lifecycle - from development, through deployment, to production. Check out our detailed use cases that show lake FS in action.

[Explore lakeFS use cases](../using_lakeFS/index.md) 

## Commitment to open source

lakeFS is an open source project under the Apache 2.0 license. It was created by Treeverse, a commercial company founded by engineers passionate 
about providing solutions to the evolving world of data engineering. 

We chose to open-source our core capabilities because we believe in bottom-up adoption of technologies, 
collaborative communities have the power to bring the best solutions to the community and small organizations 
should be able to use cutting-edge technologies for free so they can innovate in their domain and 
compete with well-established companies that can afford paying for technology. 

As a commercial organization, we intend to use an open-core model.

[Learn more](../commitment.md)

## FAQ

Take a look at our FAQ for more information about lakeFS.

[Jump to the FAQ](../faq.html)

