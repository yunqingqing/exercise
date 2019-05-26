# generate code directories and files

import os
import argparse

cur_dir_path = os.path.dirname(os.path.realpath(__file__))
base_code_path = os.path.join(cur_dir_path.rsplit(os.sep, 1)[0], 'code')


def touch(path, content=None):
    with open(path, 'a') as f:
        os.utime(path, None)
        if content:
            f.write(content)


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
    touch(main_file, content=code_template['go'].format(name))
    touch(test_file, content=code_template['go_test'] % name)


def gen_py(base_path, name):
    dir_path = os.path.join(base_path, name)
    mkdir(dir_path)
    main_file = os.path.join(dir_path, "{}.py".format(name))
    touch(main_file, content=code_template['py'].format(name))

# language_type => generator
lang_hanler_map = {
    "go": gen_go,
    'python': gen_py,
}

# golang test file code template.
go_test_template = '''package %s

import "testing"

func TestFunc(t *testing.T) {

}'''

py_main_template = '''


def func():
    pass


if __name__ == '__main__':
    func()
'''

code_template = {
    "go": "package {}",
    "go_test": go_test_template,
    'py': py_main_template,
}


if __name__ == '__main__':
    args = get_args_parser()
    lang = args.lang
    name = args.name
    lang_folder_path = os.path.join(base_code_path, lang)
    if lang:
        handler = lang_hanler_map.get(lang)
        if callable(handler):
            handler(lang_folder_path, name)
