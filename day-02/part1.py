convert_opp = {
    "A": "r",  # A = rock
    "B": "p",  # B = paper
    "C": "s",  # C = scissors
}

convert_me = {
    "X": "r",  # X = rock
    "Y": "p",  # Y = paper
    "Z": "s",  # Z = scissors
}

score_shape = {
    "r": 1,
    "p": 2,
    "s": 3,
}

score_outcome = {
    "L": 0,
    "T": 3,
    "W": 6,
}


def did_i_win(round_):
    # they play rock...
    if round_ == "r s":
        return "L"
    if round_ == "r p":
        return "W"
    if round_ == "r r":
        return "T"

    # they play paper...
    if round_ == "p s":
        return "W"
    if round_ == "p p":
        return "T"
    if round_ == "p r":
        return "L"

    # they play scissors...
    if round_ == "s s":
        return "T"
    if round_ == "s p":
        return "L"
    if round_ == "s r":
        return "W"


def score_row(row):

    # trim whitespace and split on space -> ["A", "X"]
    opp, me = row.strip().split(" ")

    # standardize symbols and recombine -> "r r"
    round_ = " ".join((convert_opp[opp], convert_me[me]))

    # score round
    my_outcome = did_i_win(round_)
    my_shape = round_.split(" ")[1]
    return score_shape[my_shape] + score_outcome[my_outcome]


def load_guide():
    with open("input.txt") as f:
        lines = f.readlines()
    return lines


def test_cases():
    assert score_row("A Y") == 8
    assert score_row("B X") == 1
    assert score_row("C Z") == 6
    print("tests pass...")


def main():
    guide = load_guide()
    scores = [score_row(x) for x in guide]
    print("total score:", sum(scores))


if __name__ == "__main__":
    test_cases()
    main()
