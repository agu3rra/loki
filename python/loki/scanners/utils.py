import os


def read_file(filename):
    """
    Opens as file and returns it as a single string

    Arguments:
        filename (str): full path to the file
    """
    try:
        with open(filename, 'r') as df:
            return df.readlines()
    except Exception as e:
        print("Error while accessing dependencies file")
        print(str(e))
        return None
