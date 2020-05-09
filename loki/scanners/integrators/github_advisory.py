class GitHubAdvisory():
    """
    Interacts with GitHub's advisory DB API
    """
    def __init__(self, pat):
        """
        Arguments:
            (str) pat: Personal Access Token.
        """
        self.pat = pat
        self.auth_header = {
            "Authorization": "bearer {}".format(pat)
        }