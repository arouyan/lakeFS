---
layout: default
title: Versioning Internals
parent: Understanding lakeFS
description: How does versioning work in lakeFS? Read this to find out.
nav_order: 20
has_children: false
redirect_from: ["../architecture/data-model.html","../understand/data-model.html"]
--- 
# Versioning Internals
{: .no_toc }

{% include toc.html %}

## Overview

Since commits in lakeFS are immutable, they are easy to store on an immutable object store.

Older commits are rarely accessed, while newer commits are accessed very frequently - so a tiered storage approach can work very well here. The object store is the source of truth, while the local disk and even RAM can be used to cache the more frequently accessed commits.

Since commits are immutable, once they're cached, you only need to evict them when you're running out of space. There’s no complex invalidation that needs to happen here.

In terms of the storage format, commits are be stored as [SSTables](https://en.wikipedia.org/wiki/Log-structured_merge-tree){: target="_blank" }, compatible with [RocksDB](https://rocksdb.org/){: target="_blank" }.

SSTables were chosen as a storage format for three major reasons:

1. Extremely high read throughput on modern hardware: using commits representing a 200m object repository (modeled after the S3 inventory of one of our design partners), we were able to achieve close to 500k random GetObject calls/second. This provides a very high throughput/cost ratio, probably as high as can be achieved on public cloud services.
1. This storage format is well-known, which means that it’s relatively easy to generate and consume. Storing it in the object store makes it accessible to data engineering tools for analysis and distributed computation, effectively reducing the silo effect of storing it in an operational database.
1. The SSTable format supports [delta encoding for keys](https://github.com/facebook/rocksdb/wiki/PlainTable-Format#prefix-encoding){: target="_blank" } which makes them very space-efficient for data lakes, where many keys share the same common prefixes.

Each lakeFS commit is represented as a set of contiguous, non-overlapping SSTables that make up the entire keyspace of a repository at that commit.

## SSTable File Format ("Graveler File")

The lakeFS metadata is encoded into a format called **"Graveler"** - a standardized way to encode content-addressable key-value pairs. This is what a Graveler file looks like:

![Graveler File]({{ site.baseurl }}/assets/img/graveler1.png)

Each Key/Value pair (**"ValueRecord"**) is constructed of a `key`, an `identity` and a `value`.

A simple identity could be, for example, a sha256 hash of the value’s bytes. But it could be any sequence of bytes that uniquely identifies the value. As far as Graveler is concerned, two `ValueRecord`s are considered identical if their key and identity fields are equal.

A Graveler file itself is content-addressable, i.e., similarly to Git, the name of the file is its identity.
The file identity is calculated based on the identity of the ValueRecords the file contains:

<b>valueRecordID = h(h(valueRecord.key) || h(valueRecord.Identity))</b><br/>
<b>fileID = h(valueRecordID<sub>1</sub> + … + valueRecordID<sub>N</sub>)</b>

## Constructing a consistent view of the keyspace (i.e., a commit)

We have two additional requirements for the storage format:

1. Be space- and time-efficient when creating a commit - assuming that a commit changes a single object out of a billion, we don’t want to write a full snapshot of the entire repository. Ideally, we can reuse some data files that haven’t changed to make the commit operations (in both space and time) proportional to the size of the difference, as opposed to the total size of the repository.
1. Allow an efficient diff between commits, which runs in time proportional to the size of their difference and not their absolute sizes.

To support these requirements, we decided to build a 2-layer [Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree){: target="_blank" } composed of a set of leaf nodes (**"Range"**) addressed by their content address, and a **"Meta Range"**, which is a special range containing all ranges, thus representing an entire consistent view of the keyspace:

![Metarange to ranges relationship]({{ site.baseurl }}/assets/img/graveler2.png)

Assuming that commit B is derived from commit A, and only changed files in range `e-f`, it can reuse all ranges except for SSTable #N (the one that contains the modified range of keys). They will be recreated with a new hash, representing the state as exists after applying commit B’s changes. This will, in turn, also create a new Metarange since its hash has now changed as well (as it's derived from the hash of all contained ranges).

Assuming that most commits usually change related objects (i.e., that are likely to share some common prefix), the reuse ratio could be very high. We tested this assumption using S3 inventory from two design partners - we partitioned the keyspace to an arbitrary number of simulated blocks and measured their change over time. We saw a daily change rate of about 5-20%.

Given the size of the repositories, it's safe to assume that a single day would translate into multiple commits. At a modest 20 commits per day, a commit is expected to reuse >= 99% of the previous commit blocks, so acceptable in terms of write amplification generated on commit.

On the object store, ranges are stored in the following hierarchy:

```
<lakefs root>
    _lakefs/
        <range hash1>
        <range hash2>
        <range hashN>
        ...
        <metarange hash1>
        <metarange hash2>
        <metarange hashN>
        ...
    <data object hash1>
    <data object hash2>
    <data object hashN>
    ...
```

*Note: This relatively flat structure could be modified in the future. Looking at the diagram above, it imposes no real limitations on the depth of the tree. A tree could easily be made recursive by having Meta Ranges point to other Meta Ranges - and still provide all the same characteristics. For simplicity, we decided to start with a fixed 2-level hierarchy.*

## Representing references and uncommitted data

Unlike committed data which is immutable, uncommitted (or "staged") data experiences frequent random writes and is very mutable by nature. This is also true for "refs" - in particular, branches that are simply pointers to an underlying commit and modified frequently (in every commit or merge operation).

Both these types of data are not only mutable but also require strong consistency guarantees - while also being fault-tolerant. If we can’t access the current pointer of the main branch, a big portion of the system is essentially down. 

Luckily, this is also much smaller data compared to the committed dataset.

References and uncommitted data are currently stored on PostgreSQL for its strong consistency and transactional guarantees.

[In the future](roadmap.md#lakefs-on-the-rocks-milestone-3---remove-postgresql), we plan to eliminate the need for an RDBMS by embedding [Raft](https://raft.github.io/){: target="_blank" } to replicate these writes across a cluster of machines, with the data itself being stored in RocksDB. To make operations easier, the replicated RocksDB database will be periodically snapshotted to the underlying object store.

For extremely large installations ( >= millions of read/write operations per second), it will be possible to utilize [multi-Raft](https://pingcap.com/blog/2017-08-15-multi-raft/){: target="_blank" } to shard references across a wider fleet of machines.

