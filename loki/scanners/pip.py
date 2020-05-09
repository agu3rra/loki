from .utils import read_file
from .integrators import GitHubAdvisory


class Pip():
    """
    Scanner class for Python packages in the PyPi.org ecosystem
    """
    
    def __init__(self, dependencies, github_pat):
        """
        Arguments:
            (str) dependencies: string representing the dependencies file.
                                This class can parse regular requirements.txt
                                and Pipfiles.
        """
        dependencies = read_file(dependencies)
        if dependencies is None:
            return None
        self.dependencies = dependencies
        self.github_pat = github_pat

    def get_advisories(self):
        """
        Checks GitHub's advisory DB for open vulnerabilities in the dependencies
        listed in the provided init file.

        Returns:
            (dict) packages and related advisories
        """
        api = GitHubAdvisory(self.github_pat)
        return None