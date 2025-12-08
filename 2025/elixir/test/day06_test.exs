defmodule Day06Test do
  use ExUnit.Case
  doctest Day06

  @sample_input File.read!("test/inputs/day06_sample.txt")
  @input File.read!("inputs/day06.txt")

  describe "part 1" do
    test "sample input" do
      expected = 4_277_556
      actual = Day06.part1(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 5_524_274_308_182
      actual = Day06.part1(@input)
      assert actual == expected
    end
  end

  describe "part 2" do
    test "sample input" do
      expected = 3263827
      actual = Day06.part2(@sample_input)
      assert actual == expected
    end

    test "real input" do
      expected = 8843673199391
      actual = Day06.part2(@input)
      assert actual == expected
    end
  end

  describe "parse_columns" do
    test "sample input" do
      expected = [
              {"*", [1, 24, 356]},
              {"+", [369, 248, 8]},
              {"*", [32, 581, 175]},
              {"+", [623, 431, 4]}
            ]

      actual = Day06.parse_columns(@sample_input)
      assert actual == expected
    end
  end
end
