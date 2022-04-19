---
layout: default
title: Set up lakeFS
description: This section show how to set up lakeFS, from preparing your storage to importing data from S3.
nav_order: 15
has_children: true
---

# Set up lakeFS

Once you install lakeFS, it's just a few simple steps to start working. The first step is to [create a bucket/container](./storage/index.md) in your underlying storage. If you already have one, you are ready to [create your first lakeFS repository](./create-repo.md). Take a look at the guides below to learn more.


## Prepare your storage

A production installation of lakeFS will usually use your cloud provider’s object storage as the underlying storage layer. Check out this guide to learn how to prepare your storage.

[Learn more](https://docs.lakefs.io/setup/storage/)

## Create a repository

Creating your first repository in lakeFS is easy - this guide shows you how to do it.

[Learn more](https://docs.lakefs.io/setup/storage/s3.html)

## Import data into lakeFS

This is a key step in your setup journey, make sure to read this guide and learn how to import data into lakeFS.

[Learn more](https://docs.lakefs.io/setup/storage/gcs.html)

## Hooks

Like other version control systems, lakeFS allows the configuration of  `Actions`  to trigger when predefined events occur. Here's a guide to setting up hooks.

[Learn more](https://docs.lakefs.io/setup/hooks.html)

## Importing data from S3 (MVCC)

To import existing data to lakeFS, you may choose to copy it using  [S3 CLI] or tools like [Apache DistCp](https://docs.lakefs.io/integrations/distcp.html#from-s3-to-lakefs). Learn how to import data from S3 in this guide.

[Learn more](https://docs.lakefs.io/setup/import-mvcc.html)

## S3 Virtual-host addressing (advanced)

Some systems require S3 endpoints to support [virtual-host style addressing](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).
lakeFS supports this, but requires some configuration - here's a handy guide.

[Learn more](https://docs.lakefs.io/setup/virtual-host-addressing.html)


