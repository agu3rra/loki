# Loki Open Source Scanner (work in progress...)
A Python program that displays security advisories for a given dependencies file leveraging [GitHub's Security Advisory API](https://developer.github.com/v4/object/securityadvisory/).  
For ad-hod advisory look-ups: https://github.com/advisories

# About the name
In Norse mythology Loki is the son of giants and a trickster/shapeshifter. Likewise, open-source is ever-evolving. What is considered safe today will not remain so indefinetly. This is why automatically receiving advisories on new vulnerabilities that affect your software projets is key to ensure your customers don't fall victim to exploitation.

Through his cunning, Loki helped the Aesir (Northern Gods) in difficult situations. So will **LOKI** help you, the Developers!

![loki](docs/loki.jpg)

# Usage
```terminal
> loki --dependencies "requirements.txt" --language "python"

#####################################################
# Welcome to the LOKI Open Source Scanner           #
#####################################################

Scan information:
Language ecosystem selected: Python
Dependencies file: requirements.txt
To supress and justify (optionally) findings, simply provide a gos3ignore.yml file.

-----------------------------------
Package 1: flask
Pinned version: None. Selecting latest: 0.12.2.
Total vulnerabilities: 2

Identified vulnerabilities
<< Vulnerability 1 >>
Severity: MODERATE
Summary: Moderate severity vulnerability that affects flask
Age: 201 days
--
Description:
The Pallets Project flask version Before 0.12.3 contains a CWE-20: Improper Input Validation vulnerability in flask that can result in Large amount of memory usage possibly leading to denial of service. This attack appear to be exploitable via Attacker provides JSON data in incorrect encoding. This vulnerability appears to have been fixed in 0.12.3.
--
Vulnerable version range: < 0.12.3
GitHub Security Advisory ID: GHSA-562c-5r94-xh97
Published date: 2018-08-23T19:10:40Z
Last updated: 2019-07-03T21:02:03Z
Reference: https://github.com/advisories/GHSA-562c-5r94-xh97

<< Vulnerability 2 >>
Severity: LOW
Summary: Low severity vulnerability that affects flask
Age: 869 days
--
Description:
The Pallets Project Flask before 1.0 is affected by: unexpected memory usage. The impact is: denial of service. The attack vector is: crafted encoded JSON data. The fixed version is: 1. NOTE: this may overlap CVE-2018-1000656.
--
Vulnerable version range: < 1.0.0
GitHub Security Advisory ID: GHSA-5wv5-4vpf-pj6m
Published date: 2019-07-19T16:12:46Z
Last updated: 2019-08-27T15:31:30Z
Reference: https://github.com/advisories/GHSA-5wv5-4vpf-pj6m
-----------------------------------

```
