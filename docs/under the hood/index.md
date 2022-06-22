---
layout: default
title: Under the Hood
description: Learn more about how lakeFS works under the hood
nav_order: 30
has_children: true
---

# Under the hood

This section explores all the key concepts and mechanisms underlying lakeFS. Take a look at these sections to learn how lakeFs works under the hood.

## Object model

lakeFS brings together concepts from object stores such as S3 with concepts from Git. 
This reference defines the common concepts of lakeFS.

[Learn more](https://docs.lakefs.io/understand/object-model.html)

## Protected branches

lakeFS users can define branch protection rules to prevent direct changes and commits to specific branches. Only merges are allowed into protected branches. 
Together with pre-merge hooks, users can run validations on their data before it reaches important branches and is exposed to consumers.

[Learn more](https://docs.lakefs.io/understand/object-model.html)

## Exporting data

This features comes in handy when external data consumers don’t have access to your lakeFS installation, some data pipelines aren't fully migrated with lakeFS,
or you'd like to experiment with lakeFS as a side-by-side installation first. You can use the export operation to copy all data from a given lakeFS commit to a designated object store location.

[Learn more](https://docs.lakefs.io/reference/export.html)

## Garbage collection

By default, lakeFS keeps all your objects forever, allowing you to travel back in time to the previous versions of your data. 
But sometimes you may want to hard-delete your objects from the underlying storage for cost reduction or compliance. This is possible in lakeFS. 

[Learn more](https://docs.lakefs.io/reference/garbage-collection.html)

