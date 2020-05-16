import pytest
import requests

import loki
from loki.scanners.integrators import GitHubAdvisory
from loki.environment import GITHUB_PAT


class TestGitHubAdvisoryAPI():

    @pytest.mark.parametrize("query, exp_error", [
        (
            """
            query retrieveVulnerability { 
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
            }
            """,
            None
        ),
    ])
    def test_post(self, query, exp_error):
        api = GitHubAdvisory(GITHUB_PAT)
        (data, error) = api.run(query)
        assert error == exp_error
        if exp_error is not None:
            assert isinstance(data, dict)
            vulnerabilities = data.get("securityVulnerabilities").get("nodes")
            assert isinstance(vulnerabilities, list)
            assert len(vulnerabilities) > 0
