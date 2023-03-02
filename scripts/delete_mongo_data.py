import os
import shutil

def delete(path):
    """path could either be relative or absolute. """
    # check if file or directory exists
    for file_name in os.listdir(path):
        if file_name == '.gitkeep':
            continue
        file = path + os.sep + file_name
        if os.path.isfile(file) or os.path.islink(file):
            # remove file
            os.remove(file)
        elif os.path.isdir(file):
            # remove directory and all its content
            shutil.rmtree(file)
        else:
            raise ValueError("Path {} is not a file or dir.".format(file))

if __name__ == '__main__':
    delete('..' + os.sep + 'data')