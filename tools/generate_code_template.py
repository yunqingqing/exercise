# generate code directories and files

import os
import argparse

cur_dir_path = os.path.dirname(os.path.realpath(__file__))
base_code_path = os.path.join(cur_dir_path.rsplit(os.sep, 1)[0], 'code')


def touch(path):
    with open(path, 'a'):
        os.utime(path, None)


def mkdir(path):
    if not os.path.isdir(path):
        os.mkdir(path)


def get_args_parser():
    parser = argparse.ArgumentParser(description='auto generate code')
    parser.add_argument('name', help='directories and file name flag.')
    parser.add_argument('--lang', dest='lang',
                        help='language type.')
    return parser.parse_args()


def gen_go(base_path, name):
    dir_path = os.path.join(base_path, name)
    mkdir(dir_path)
    main_file = os.path.join(dir_path, "{}.go".format(name))
    test_file = os.path.join(dir_path, "{}_test.go".format(name))
    touch(main_file)
    touch(test_file)


lang_hanler_map = {
    "go": gen_go
}  # language_type => generator


if __name__ == '__main__':
    args = get_args_parser()
    lang = args.lang
    name = args.name
    lang_folder_path = os.path.join(base_code_path, lang)
    if lang:
        handler = lang_hanler_map.get(lang)
        if callable(handler):
            handler(lang_folder_path, name)
