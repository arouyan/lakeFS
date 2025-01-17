---
layout: default
title: More Quickstart Options
description: This section outlines additional quickstart options to deploying lakeFS.
parent: Quickstart
nav_order: 50
has_children: false
---

# More Quickstart Options
{: .no_toc}

{% include learn_only.html %} 

{% include toc.html %}

## Docker on Windows

To run a local lakeFS instance using [Docker Compose](https://docs.docker.com/compose/){:target="_blank"}:

1. Ensure that you have Docker installed on your computer and that compose version is 1.25.04 or higher. For more information, please see this [issue](https://github.com/treeverse/lakeFS/issues/894).

1. Run the following command in your terminal:

   ```bash
   Invoke-WebRequest https://compose.lakefs.io | Select-Object -ExpandProperty Content | docker-compose -f - up
   ```

1. Check your installation by opening [http://127.0.0.1:8000/setup](http://127.0.0.1:8000/setup){:target="_blank"} in your web browser.

1. You're now ready to [create your first repository](repository.md) in lakeFS.

## On Kubernetes with Helm

1. Install lakeFS on a Kubernetes cluster using Helm:
   ```bash
   # Add the lakeFS Helm repository
   helm repo add lakefs https://charts.lakefs.io
   # Deploy lakeFS with helm release "my-lakefs"
   helm install my-lakefs lakefs/lakefs
   ```

1. The printed output will help you forward a port to lakeFS, so you can access it from your browser at [http://127.0.0.1:8000/setup](http://127.0.0.1:8000/setup){:target="_blank"}.

1. Move on to [create your first repository](repository.md) in lakeFS.

## Using the Binary 

Alternatively, you may opt to run the lakefs binary directly on your computer.

1. Download the lakeFS binary for your operating system:

   [Download lakefs](../index.md#downloads){: .btn .btn-green target="_blank"}

1. Install and configure [PostgreSQL](https://www.postgresql.org/download/){:target="_blank"}

1. Create a configuration file:
    
   ```yaml
   ---
   database:
     connection_string: "postgres://localhost:5432/postgres?sslmode=disable"
    
   blockstore: 
     type: "local"
     local:
       path: "~/lakefs_data"
    
   auth:
     encrypt:
       secret_key: "a random string that should be kept secret"
   ```

1. Create a local directory to store objects:

   ```sh
   mkdir ~/lakefs_data
   ```

1. Run the server:
    
   ```bash
   ./lakefs --config /path/to/config.yaml run
   ```

1. Check your installation by opening [http://127.0.0.1:8000/setup](http://127.0.0.1:8000/setup){:target="_blank"} in your web browser.

1. You are now ready to [create your first repository](repository.md) in lakeFS.
