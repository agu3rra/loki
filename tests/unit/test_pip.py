import pytest
from .pip import Pip


class TestPipScan():

    @pytest.mark.parametrize("dependencies, expected", [
        ("samples/pip_requirements.txt", dict),
        ("", None),
        ("dummy.txt", None),
        ("samples/sample_pip_Pipfile.lock", dict),
        ("sample_pip_malformed", None),
    ])
    def test_scan(self, dependencies, expected):
        scanner = Pip(dependencies)
        advisories = scanner.get_advisories()
        if advisories is None:
            assert isinstance(type(advisories), expected)
        else:
            assert isinstance(advisories, expected)