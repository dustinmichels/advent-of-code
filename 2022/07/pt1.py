import os
from pathlib import PurePath

from node import Node, make_list_dfs
from rich import print


def add_node(root_node: Node, path: PurePath, item: str):
    p = os.path.normpath(path)
    keys = [k for k in p.split("/") if k != ""]

    # Navigate to the right dir node,
    # creating nodes along the way if necessary
    curr = root_node
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
        node.type = "file"
        node.update_size(int(item_desc))


def main(filename: str):

    curr_dir = PurePath("/")
    root_node = Node("root")

    with open(filename) as f:
        txt = f.read()

    lines = txt.split("\n")
    for line in lines:
        # if it's a "cd" command, update curr_dir
        if line.startswith("$ cd"):
            _, _, *args = line.split()
            if args[0] == "/":
                curr_dir = PurePath("/")
            else:
                curr_dir = PurePath(curr_dir, args[0])

        # if it's not a command, add node to tree
        if not line.startswith("$"):
            add_node(root_node, curr_dir, line)

    res = make_list_dfs(root_node)
    res = [x for x in res if x.type == "dir" and x.size <= 100000]
    return sum(x.size for x in res)


if __name__ == "__main__":
    test_sum = main("input_test.txt")
    assert (test_sum) == 95437

    real_sum = main("input.txt")
    print("Sum:", real_sum)
