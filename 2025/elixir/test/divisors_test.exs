defmodule DivisorsTest do
  use ExUnit.Case, async: true

  test "from_factors([]) returns [1]" do
    assert Divisors.from_factors([]) == [1]
  end

  test "divisors of 60 are correct and sorted desc" do
    factors = [{2, 2}, {3, 1}, {5, 1}]
    assert Divisors.from_factors(factors) ==
             [60, 30, 20, 15, 12, 10, 6, 5, 4, 3, 2, 1]
  end

  test "single prime power (2^7 = 128)" do
    assert Divisors.from_factors([{2, 7}]) ==
             [128, 64, 32, 16, 8, 4, 2, 1]
  end

  test "unsorted factors still work (60 again)" do
    factors = [{5, 1}, {2, 2}, {3, 1}]
    assert Divisors.from_factors(factors) ==
             [60, 30, 20, 15, 12, 10, 6, 5, 4, 3, 2, 1]
  end

  test "divisor count matches product of (e_i + 1), all divide n, no duplicates" do
    factors = [{2, 3}, {3, 2}] # n = 2^3 * 3^2 = 72
    n = Enum.reduce(factors, 1, fn {p, e}, acc -> acc * Integer.pow(p, e) end)
    divs = Divisors.from_factors(factors)

    # Count == (3+1)*(2+1) = 12
    assert length(divs) == 12

    # All divide n
    assert Enum.all?(divs, fn d -> rem(n, d) == 0 end)

    # No duplicates
    assert length(divs) == divs |> Enum.uniq() |> length()

    # Includes extrema
    assert hd(divs) == n
    assert List.last(divs) == 1
  end
end
