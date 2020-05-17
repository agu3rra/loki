import argparse
import os
import enum

from . import scanners
from .environment import GITHUB_PAT

# Program exit codes


class ExitCode(enum.Enum):
    OK = 0
    MISSING_ARGUMENTS = 1
    LANGUAGE_NOT_SUPPORTED = 2
    ERROR_OPENING_DEPENDENCIES = 3


languages = {
    # language:ecosystem
    "python": "pip",
}

if __name__ == "__main__":
    # Command line arguments
    parser = argparse.ArgumentParser()
    help_text = ""
    for lang in languages.keys():
        help_text += lang + ";"
    parser.add_argument(
        "language",
        help="Language to use. Supported: {}".format(help_text)
    )
    parser.add_argument("dependencies", help="Path to dependencies file.")
    args = parser.parse_args()

    # Check language support
    args.language = args.language.lower()
    if args.language not in languages.keys():
        print("Language not supported.")
        os._exit(ExitCode.LANGUAGE_NOT_SUPPORTED)

    if args.language == "python":
        scanner = scanners.Pip(args.dependencies, GITHUB_PAT)
        if scanner is None:
            os._exit(ExitCode.ERROR_OPENING_DEPENDENCIES)
        _ = scanner.get_advisories()