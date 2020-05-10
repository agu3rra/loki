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

    def post(self, query):
        """
        Sends POST request with API query

        Arguments:
            (dict) query: dictionary with GraphQL query request.
        """
        return None