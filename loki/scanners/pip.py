class Pip():
    """
    Scanner class for Python packages in the PyPi.org ecosystem
    """
    
    def __init__(self, dependencies):
        """
        Arguments:
            (str) dependencies: string representing the dependencies file.
                                This class can parse regular requirements.txt
                                and Pipfiles.
        """
        self.dependencies = dependencies

    def get_advisories(self):
        """
        Checks GitHub's advisory DB for open vulnerabilities in the dependencies
        listed in the provided init file.

        Returns:
            (dict) packages and related advisories
        """
        return None