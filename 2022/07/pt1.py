from __future__ import annotations

import os
from pathlib import PurePath


from node import Node, dfs
from rich import print

ROOT = Node("root")
CURR_DIR = PurePath("/")


def make_or_get_node(path: PurePath, item: str):
    p = os.path.normpath(path)
    keys = [k for k in p.split("/") if k != ""]

    # Navigate to the right dir node
    curr = ROOT
    for k in keys:
        child = curr.children.get(k)
        if not child:
            child = Node(k, parent=curr)
            curr.children[k] = child
        curr = child

    # Add item (file or dir)
    item_desc, item_name = item.split()
    node = Node(item_name, parent=curr)
    curr.children[item_name] = node
    if item_desc != "dir":
        node.update_size(int(item_desc))


def parse_cmd(line: str):
    cmd, arg = None, None
    res = line.split()
    if len(res) == 3:
        _, cmd, arg = res
    else:
        _, cmd = res
    return cmd, arg


# with open("2022/07/input_test.txt") as f:
# txt = f.read()

with open("input_test.txt") as f:
    txt = f.read()

lines = txt.split("\n")
for line in lines:
    if line.startswith("$ "):
        cmd, arg = parse_cmd(line)
        if cmd == "cd":
            if arg == "/":
                CURR_DIR = PurePath("/")
            else:
                CURR_DIR = PurePath(CURR_DIR, arg)
        continue
    make_or_get_node(CURR_DIR, line)

# print(ROOT)

dfs(ROOT)
