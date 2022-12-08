from __future__ import annotations

from typing import Dict, Literal


class Node:
    name: str
    type: Literal["dir", "file"]
    parent: Node | None
    children: Dict[str, Node]
    size: int

    def __init__(self, name: str, parent: Node | None = None) -> None:
        self.name = name
        self.type = "dir"  # default, can be changed to file
        self.parent = parent
        self.children = dict()
        self.size = 0

    def __repr__(self) -> str:
        return f"<{self.name} [{self.size}]> {list(self.children.keys())}"

    def update_size(self, size: int):
        self.size = size
        node = self
        while node.parent:
            node = node.parent
            node.size = sum([node.children[x].size for x in node.children])
            # node = node.parent


def make_list_dfs(root: Node):
    res = [root]
    for _, v in root.children.items():
        res.extend(make_list_dfs(v))
    return res
