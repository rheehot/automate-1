# Automate Compliance (and dependencies) Knowledge Transfer notes

## Overview

Compliance in Automate helps customer manage security and compliance risk of:
 * Servers (Windows, Linux, etc)
 * Cloud Accounts (AWS, Azure, GCP, Digital Ocean, etc)
 * Network devices
 * Containers
 * Anything else that has an API or CLI if a custom plugin is created. Not necessary security or compliance related. Used frequently for common IT automated testing because of it's versatility.

Automate ships with 399 compliance profiles, majority of them automatically converted from industry leading compliance benchmarks like **CIS** and **STIG**.
 * `Automate UI > Compliance > Profiles > Available tab` also called in the code as the *market* profiles
 * Once imported with the `Get` button, they can be used to trigger scans

Chef has certified a lot of these profiles with **CIS**: https://www.cisecurity.org/partner/chef/

The profiles are executed a [Ruby](http://www.ruby-lang.org) based tool called **InSpec**. The profile execution generates a report, which is then used by Automate for the Compliance reporting data.
InSpec is owned by Chef/Progress, with the source code being available in this public repository:
 - https://github.com/inspec/inspec

InSpec uses a ruby gem called [train](https://github.com/inspec/train) to communicate with the targets.

Starting with InSpec v4 (released in April 2019), packages of InSpec and Chef Infra require a Chef license acceptance. For commercial use, a paid subscription.

Independent of Chef, **[CINC group](https://cinc.sh)** releases open source builds of *Chef Infra* and *Chef InSpec* at.

---

**Show slides of Infra vs InSpec and basic operation of InSpec**

---

## How InSpec reports make it to Automate:

1. **[audit cookbook](https://github.com/chef-cookbooks/audit)**, executed by Chef Infra. The most common execution modes:

 - **Chef Infra** (runs **audit cookbook** [runs **InSpec**]) -----> **Chef Server** -----> **Chef Automate**
 - **Chef Infra** (runs **audit cookbook** [runs **InSpec**]) ---*data collection token*---> **Chef Automate**


2. **Habitat Effortless package** (with bundled InSpec) ---*data collection token*---> **Chef Automate**

3. **Automate** (with bundled InSpec) executes remote scans of VMs, cloud accounts(AWS, Azure, GCP, SSL certificates)

4. **InSpec** (with the `automate` reporter) executed manually or from Ansible/Puppet/Cron job. In this mode, InSpec posts the report to an Automate URL with a **data collection token**

Example InSpec exec and json report:
```
$ inspec exec ~/git/mycompliance-profile/mylinux --reporter json | jq .
{
  "platform": {
    "name": "debian",
    "release": "bullseye/sid"
  },
  "profiles": [
    {
      "name": "mylinux",
      "version": "0.1.4",
      "sha256": "1a319ce266574c9a6f121245be8e0cf355ee6c498d556e0173207e8a71fbe908",
      "title": "My Demo Linux Profile",
...
```

## Audit cookbook

Main driver of InSpec report data that is sent to Automate.
Executed by Chef Infra client with attributes: profiles, inputs, report target.

Provided over the years a stop gap, a place to quick fix things or make sure InSpec and Automate stay "friends".

Example: InSpec generates a report of **100 MB** based on the profile being executed and the target node. Automate has a hard limit of **4 MB** inherited from GRPC. Audit cookbook takes the **100 MB** report, reduces profile metadata (if Automate already has it), truncates results, and send to Automate an acceptable report of **0.1 MB** :tada: :tada: :tada:

## Ingestion

Stored in ElasticSearch in 3 index types:

 * summary (e.g. `comp-7-r-2020.11.07`)
 * report (e.g. `comp-7-s-2020.11.07`)
 * profiles (e.g. `comp-3-profiles`)

Used to normalize data, for deduplication reasons and to avoid unnecessary ingestion load.
Calculate report stats for performant and rich data aggregations.

## Node manager

## Scan jobs

---

## Reporting with extensive filtering and deep filtering
Surface level filters:
- environment
- organization
- chef_server
- chef_tags
- policy_group
- policy_name
- status
- node_name
- platform
- platform_with_version
- role
- recipe
- inspec_version
- ipaddress
- node_id
- start_time
- end_time
- job_id

Nested filters:
- control
- control_tag
- profile_id

Example of some filters used across most reporting apis
```json
{
  "filters": [
    {
      "type": "start_time",
      "values": [
        "2019-01-26T00:00:00Z"
      ]
    },
    {
      "type": "end_time",
      "values": [
        "2019-02-05T23:59:59Z"
      ]
    },
    {
      "type": "profile_id",
      "values": [
        "f42d2f48c9acd48f52324d52ec575ca9028e405eb303f69cb34d79eb0e588b5c"
      ]
    },
    {
      "type": "control",
      "values": [
        "sshd-02"
      ]
    },
    {
      "type": "platform",
      "values": [
        "centos"
      ]
    },
    {
      "type": "environment",
      "values": [
        "DevSec Prod Zeta"
      ]
    }
  ]
}
```

Deep filtering allows us to see how our infra is doing with respect to the overall node, a specific profile, or a specific control


## The four states of a control / profile / node:
 * passed
 * skipped
 * failed
 * waived
```$go
func computeStatus(failed int32, passed int32, skipped int32, waived int32) string {
	if failed > 0 {
		return "failed"
	} else if passed == 0 && skipped == 0 && waived > 0 {
		return "waived"
	} else if passed > 0 || skipped == 0 {
		return "passed"
	} else if passed == 0 && skipped > 0 {
		return "skipped"
	}
	return "unknown"
}
```

## Suggestions
Reports contain certain fields that we use to filter on
the suggestions api allows us to list close matches as we type for more efficient selection
The reporting data is stored in ES in summary and also in detailed form..
Currently the summary is only used for surface level suggestions


## ES Migrations
* index versioning is separate for report indices and profiles indices which minimizes migrations
* the migration flow chart shows the process by which we migrate time series for maximum up time
* the interface driven approach to migrations (esMigratable interface)
