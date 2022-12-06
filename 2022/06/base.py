# Woops, doesn't work because of recursion limit...
# def find_seq_recursive(s, i=4):
#     if len(set(s[0:4])) == 4:
#         return s[0:4], i
#     return find_seq_recursive(s[1:], i + 1)


def find_seq(buffer: str, n: int):
    """
    Look for a series of _n_ distinct characters
    in the given data buffer.
    """
    for i in range(n, len(buffer)):
        seq = buffer[i - n : i]
        if len(set(seq)) == n:
            return seq, i
