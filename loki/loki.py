import argparse
import os
import enum


# Program exit codes
class ExitCode(enum.Enum):
    OK=0
    MISSING_ARGUMENTS=1

# Language support
languages = {
    #language:ecosystem
    "python":"pip",
    "javascript":"npm",
}

if __name__ == "__main__":
    # Command line arguments
    parser = argparse.ArgumentParser()
    help_text = ""
    for lang in languages.keys():
        help_text += lang + ";"
    parser.add_argument(
        'language',
        help="Language to use. Supported: {}".format(help_text)
    )
    parser.add_argument('dependencies', help="Path to dependencies file.")
    args = parser.parse_args()

    if args.language is None or args.dependencies is None:
        print("Missing arguments")
        help()
        os._exit(ExitCode.MISSING_ARGUMENTS)
