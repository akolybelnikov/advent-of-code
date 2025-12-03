defmodule Day03Test do
  use ExUnit.Case
  doctest Day03

  @sample_input File.read!("test/inputs/day03_sample.txt")
  @input File.read!("inputs/day03.txt")

  test "part 1 sample" do
    expected = 357
    actual = Day03.part1(@sample_input)
    assert actual == expected
  end

  test "part 1" do
    expected = 17034
    actual = Day03.part1(@input)
    assert actual == expected
  end

  test "part 2 sample" do
    # Sum of picked-only 12-digit prefixes per line
    # 987654321111 + 811111111119 + 434234234278 + 888911112111
    expected = 3_121_910_778_619
    actual = Day03.part2(@sample_input)
    assert actual == expected
  end

  test "part 2" do
    expected = 168_798_209_663_590
    actual = Day03.part2(@input)
    assert actual == expected
  end

  describe "find_max_joltage_2/1" do
    test "line 1: 987654321111111" do
      result = Day03.find_max_joltage_2("987654321111111")
      assert List.to_integer(result) == 98
    end

    test "line 2: 811111111111119" do
      result = Day03.find_max_joltage_2("811111111111119")
      assert List.to_integer(result) == 89
    end

    test "line 3: 234234234234278" do
      result = Day03.find_max_joltage_2("234234234234278")
      assert List.to_integer(result) == 78
    end

    test "line 4: 818181911112111" do
      result = Day03.find_max_joltage_2("818181911112111")
      assert List.to_integer(result) == 92
    end
  end

  describe "find_max_joltage_12/1 (picked-only prefix)" do
    test "line 1: 987654321111111 -> 987654321111" do
      result = Day03.find_max_joltage_12("987654321111111")
      assert List.to_integer(result) == 987_654_321_111
    end

    test "line 2: 811111111111119 -> 811111111119" do
      result = Day03.find_max_joltage_12("811111111111119")
      assert List.to_integer(result) == 811_111_111_119
    end

    test "line 3: 234234234234278 -> 434234234234" do
      result = Day03.find_max_joltage_12("234234234234278")
      assert List.to_integer(result) == 434_234_234_278
    end

    test "line 4: 818181911112111 -> 888911112111" do
      result = Day03.find_max_joltage_12("818181911112111")
      assert List.to_integer(result) == 888_911_112_111
    end
  end
end
