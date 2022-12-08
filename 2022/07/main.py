import os
from pathlib import PurePath
from typing import List

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


def process_input(input: str):
    curr_dir = PurePath("/")
    root_node = Node("root")

    lines = input.split("\n")
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

    return root_node


def part1(node_list: List[Node]):
    node_list = [x for x in node_list if x.type == "dir" and x.size <= 100000]
    return sum(x.size for x in node_list)


def part2(node_list: List[Node]):
    total_space = 70000000
    needed_space = 30000000

    unused_space = total_space - node_list[0].size
    diff = needed_space - unused_space

    node_list = sorted(
        [x for x in node_list if x.type == "dir" and x.size >= diff],
        key=lambda x: x.size,
    )

    return node_list[0].size


def main(filename: str):

    with open(filename) as f:
        txt = f.read()

    root_node = process_input(txt)
    node_list = make_list_dfs(root_node)

    res1 = part1(node_list)
    res2 = part2(node_list)

    return {"part1": res1, "part2": res2}


if __name__ == "__main__":

    test_res = main("input_test.txt")
    assert (test_res["part1"]) == 95437
    assert (test_res["part2"]) == 24933642

    print("[tests pass]")

    res = main("input.txt")
    print("Part 1:", res["part1"])
    print("Part 2:", res["part2"])
