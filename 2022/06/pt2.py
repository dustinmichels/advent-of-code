from base import find_seq


def test(stream, expected):
    _, i = find_seq(stream, 14)
    assert i == expected


def test_cases():
    test("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19)
    test("bvwbjplbgvbhsrlpgdmjqwftvncz", 23)
    test("nppdvjthqldpwncqszvftbrmjlhg", 23)
    test("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29)
    test("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26)
    print("Tests passed.")


test_cases()


with open("input.txt") as f:
    txt = f.read()

seq, i = find_seq(txt, 14)
print(seq, i)
