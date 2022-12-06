from base import find_seq


def test(stream, expected):
    _, i = find_seq(stream, 4)
    assert i == expected


def test_cases():
    test("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7)
    test("bvwbjplbgvbhsrlpgdmjqwftvncz", 5)
    test("nppdvjthqldpwncqszvftbrmjlhg", 6)
    test("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10)
    test("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11)
    print("Tests passed.")


test_cases()


with open("input.txt") as f:
    txt = f.read()

seq, i = find_seq(txt, 4)
print(seq, i)
