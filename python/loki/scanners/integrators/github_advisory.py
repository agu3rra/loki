import gql
from gql import gql, Client
from gql.transport.requests import RequestsHTTPTransport


class GitHubAdvisory():
    """
    Interacts with GitHub's advisory DB API
    """
    def __init__(self, pat, base_url="https://api.github.com/graphql"):
        """
        Arguments:
            (str) pat: Personal Access Token.
            (str) base_url (optional)
        """
        transport = RequestsHTTPTransport(
            url=base_url,
            use_json=True,
            headers={
                "Content-type": "application/json",
                "Authorization": "bearer {}".format(pat)
            },
            retries=3,
        )
        self.client = Client(
            transport=transport,
            fetch_schema_from_transport=True,
        )

    def run(self, query):
        """
        Queries Github's GraphQL API

        Arguments:
            (dict) query: dictionary with GraphQL query request.

        Returns:
            (dict, str), results, error
        """
        try:
            results = self.client.execute(gql(query))
            errors = results.get("errors", None)
            if errors is not None:
                return (None, errors)
            return (results, None)
        except Exception as e:
            return (None, str(e))