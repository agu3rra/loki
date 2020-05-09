import os


def read_file(filename):
    # Check dependencies file
    try:
        with open(args.dependencies, 'r') as df:
            return df.readlines()
    except Exception as e:
        print("Error while accessing dependencies file")
        print(str(e))
        return None