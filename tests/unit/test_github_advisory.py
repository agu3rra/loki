import pytest
import requests

import loki
from loki.scanners.integrators import GitHubAdvisory
from loki import GITHUB_PAT


class TestGitHubAdvisoryAPI():

    @pytest.mark.parametrize("query", [
        ({"query":
        """
        securityVulnerabilities (package:"flask", ecosystem: PIP, last:100) {
            nodes {
            package {
                name
                ecosystem
            }
            vulnerableVersionRange
            advisory {
                summary
                severity
                ghsaId
                description
                permalink
                publishedAt
                updatedAt
            }
            }
        }
        """}),
    ])
    def test_post(self, query):
        api = GitHubAdvisory(GITHUB_PAT)
        response = api.post(query)
        data = response.get("data", None)
        assert data is not None
        vulnerabilities = data.get("securityVulnerabilities", None)
        assert vulnerabilities is not None
        nodes = vulnerabilities.get("nodes", None)
        assert nodes is not None
        assert len(nodes) > 0
