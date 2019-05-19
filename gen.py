# -*- coding:utf-8 -*-
"""
This module is used to generate README file content list
"""

import os
import re
_relative_path = "./Algorithm"

_README_TEMPLE = """
# LeetCode Solution

[Leetcode](https://leetcode-cn.com/problemset/all/) algorithm or database problemset.


Problem | Solution | Level
:---|:---:|:--:
%s
"""


_LINE_TEMPLE = "[%s %s](%s)|[%s](./Algorithm/%s)|%s"

_HyperLink_Pattern = r"((http|ftp|https)://)(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?"
_Language_Pattern = r"lang=.*"
_Level_pattern = r"(Easy|Medium|Hard)"


def generate():
    rows = []
    for file_name in os.listdir(_relative_path):
        segments = file_name.split(".")
        index = segments[0]
        title = segments[1]
        with open(os.path.join(_relative_path, file_name), "r") as f:
            content = f.read()
            result = re.search(_HyperLink_Pattern, content)
            hyperlink = content[result.start():result.end()]
            result = re.search(_Language_Pattern, content)
            language = content[result.start():result.end()][5:]
            result = re.search(_Level_pattern, content)
            level = content[result.start():result.end()]
            rows.append(_LINE_TEMPLE %
                        (index, title, hyperlink,  language, file_name, level))
    sorted_rows = sorted(rows, key=lambda x: int(x.split(" ")[0][1:]))
    readme_content = _README_TEMPLE % ("\n".join(sorted_rows))
    with open("./README.md", 'w') as f:
        f.write(readme_content)


if __name__ == "__main__":
    generate()
