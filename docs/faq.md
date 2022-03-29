---
layout: default
title: FAQ
description: Frequently Asked Questions (FAQ). Have a question about lakeFS? Find our what others where asking
nav_order: 60
has_children: false
---

# FAQ

### 1. Is lakeFS open source?
lakeFS is open-source and available to use free of charge. lakeFS is licensed under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) License. We maintain a public [product roadmap](https://docs.lakefs.io/understand/roadmap.html) and [Slack channel](https://lakefs.io/slack) for open discussions among the lakeFS community.

### 2. How does lakeFS data versioning work?
lakeFS uses a copy-on-write mechanism to avoid data duplication. For example, creating a new branch is a metadata-only operation: no objects are actually copied in this operation. Only when an object changes does lakeFS create another [version of the data](https://lakefs.io/data-versioning/) in the storage. For more information, see [Versioning internals](./understand/versioning-internals.md).

### 3. How do I get support for my lakeFS installation?
We are extremely responsive on our Slack channel, and we make sure to prioritize the most urgent issues for our community. For SLA-based support, please contact us at [support@treeverse.io](mailto:support@treeverse.io).

### 4. Do you collect data from your active installations?
We collect anonymous usage statistics to understand the usage patterns better and identify product gaps to be fixed. This is completely optional and you can turn it off by setting `stats.enabled` to `false`. See the [configuration reference](reference/configuration.md#reference) for more details.

The data we gather is limited to the following:
1. The UUID generated when you set up lakeFS for the first time. It contains no personal or otherwise identifiable information,
1. The lakeFS version currently running,
1. The OS and architecture lakeFS is running on,
1. Metadata regarding the database used (version, installed extensions and parameters such as DB Timezone and work memory),
1. Periodic aggregated action counters (e.g., how many "get_object" operations occurred).

### 5. How is lakeFS different from Delta Lake / Hudi / Iceberg?
Delta Lake, Hudi, and Iceberg all define dedicated, structured data formats that allow deletes and upserts. lakeFS is format-agnostic and enables consistent cross-collection versioning of your data using Git-like operations. Read our [blog](https://lakefs.io/hudi-iceberg-and-delta-lake-data-lake-table-formats-compared/) for a more detailed comparison. 

### 6. What inspired the lakeFS logo?
The [Axolotl](https://en.wikipedia.org/wiki/Axolotl){:target="_blank"} is a species of salamander, also known as the Mexican Lake Monster or the Peter Pan of the animal kingdom. It's a magical creature living in a lake, just like us. :-)

![Axolotl](https://upload.wikimedia.org/wikipedia/commons/f/f6/AxolotlBE.jpg)

<small>
    [copyright](https://en.wikipedia.org/wiki/Axolotl#/media/File:AxolotlBE.jpg)
</small>
