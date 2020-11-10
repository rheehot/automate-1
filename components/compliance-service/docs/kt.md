# Automate Compliance (and dependencies) Knowledge Transfer notes

## Overview

Compliance in Automate helps customer manage security and compliance risk of:
 * Cloud Accounts (AWS, Azure, GCP, Digital Ocean, etc)
 * Servers (Windows, Linux, etc)
 * Network devices
 * Containers
 * Anything else that has an API or CLI if a custom plugin is created. Not necessary security or compliance related. Used frequently for common IT automated testing because of it's versatility.

Automate ships with 399 compliance profiles, majority of them automatically converted from industry leading compliance benchmarks like CIS and STIG.
 * `Automate UI > Compliance > Profiles > Available tab` also called in the code as the *market* profiles
 * Once imported with the "Get" button, they can be used to trigger scans

Chef has certified a lot of these profiles with CIS: https://www.cisecurity.org/partner/chef/

The profiles are executed a Ruby based tool called Chef InSpec. The profile execution generates a report, which is then used by Automate for the Compliance reporting data.
InSpec is owned by Chef/Progress, with the source code being available in this public repository:
 - https://github.com/inspec/inspec

Starting with InSpec v4 (released in April 2019), packages of InSpec and Chef Infra require a Chef license acceptance. For commercial use, a paid subscription.

Independent of Chef, **[CINC group](https://cinc.sh)** releases open source builds of *Chef Infra* and *Chef InSpec* at:

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

## Node manager

## Scan jobs

## Reporting with extensive filtering and deep filtering

## Suggestions

## ES Migrations
