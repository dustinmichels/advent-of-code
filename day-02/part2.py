from part1 import convert_opp, load_guide, score_outcome, score_shape

convert_me = {
    "X": "L",  # X = lose
    "Y": "T",  # Y = tie
    "Z": "W",  # Z = win
}


def what_do_i_play(round_):
    # they play rock...
    if round_ == "r L":
        return "s"
    if round_ == "r W":
        return "p"
    if round_ == "r T":
        return "r"

    # they play paper...
    if round_ == "p W":
        return "s"
    if round_ == "p T":
        return "p"
    if round_ == "p L":
        return "r"

    # they play scissors...
    if round_ == "s T":
        return "s"
    if round_ == "s L":
        return "p"
    if round_ == "s W":
        return "r"


def score_row(row):

    # trim whitespace and split on space -> ["A", "X"]
    opp, me = row.strip().split(" ")

    # standardize symbols and concat again -> "r L"
    round_ = " ".join((convert_opp[opp], convert_me[me]))

    # score round
    my_shape = what_do_i_play(round_)
    my_outcome = round_.split(" ")[1]

    return score_shape[my_shape] + score_outcome[my_outcome]


def test_cases():
    assert score_row("A Y") == 4
    assert score_row("B X") == 1
    assert score_row("C Z") == 7
    print("tests pass...")


def main():
    guide = load_guide()
    scores = [score_row(x) for x in guide]
    print("total score:", sum(scores))


if __name__ == "__main__":
    test_cases()
    main()
