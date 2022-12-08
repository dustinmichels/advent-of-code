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
        self.type = "dir"  # default
        self.parent = parent
        self.children = dict()
        self.size = 0

    def __repr__(self) -> str:
        return f"<{self.name}[{self.size}]>{list(self.children.keys())}"

    def update_size(self, size: int):
        self.size = size
        node = self
        while node.parent:
            node.parent.size += node.size
            node = node.parent


def dfs(root: Node):

    fmt = lambda node: [{"name": node.name, "size": node.size, "type": node.type}]

    if not root.children:
        return fmt(root)

    res = []
    for _, v in root.children.items():
        res.extend(dfs(v))
        res.extend(fmt(v))

    return res
