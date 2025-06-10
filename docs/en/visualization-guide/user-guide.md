Language: English | [日本語](TBA)

# User guide

> [!NOTE]
> This document explains how to use KHI after running it. If you haven't run KHI, please check [README](/README.md).

## Query logs with KHI

Once you run KHI and open the page served by it, you will see the welcome page.
You can query your logs with clicking the `New inspection` button, or you can open KHI file.
After completing querying your logs with the later instruction, you would see the inspection list on right side.

![User guide: start screen](/docs/en/images/guide-start-screen.png)

To query your cluster logs, you would need to fill few information with 3 steps.

1. Select the cluster type
1. Select log types to query
1. Fill parameters needed for composing log filter

![User guide: query logs](/docs/en/images/guide-query.png)

After clicking the `Run` button on the query dialog, you will see progress bars on the start screen.
Once the query done, you can open the visualization result with hitting the `Open` button.

## Understanding the visualization

After opening the log, you will see a colorful visualization.
This is the main view of KHI, and you can see the timeline diagram on the left side, which reveals the macroscopic behavior of resources in your cluster.
The right side is dedicated to details of each resource. You can see detailed information by selecting a resource in the timeline.

![User guide timeline screens](/docs/en/images/guide-timeline-screen.png)

### History view

The history view shows the modification history of the selected timeline.

![User guide history-view](/docs/en/images/guide-history-view.png)

### Filtering features

You can filter timelines by Kind, Namespace, and Resource name using the input fields at the top left of the timeline diagram.
This allows you to quickly narrow down the displayed timelines to only those relevant to your current investigation.

You can also filter logs using a regular expression in the input field at the top right. This will also filter out timelines that do not have any logs matching the regex.

![User guide filtering](/docs/en/images/guide-filtering.png)

### Topology view (alpha)

The Topology view provides a visual representation of the relationships between different Kubernetes resources.
This can be helpful for understanding the resource topology at a specific point in time within a log.

![User guide topology-view](/docs/en/images/guide-topology-view.png)
