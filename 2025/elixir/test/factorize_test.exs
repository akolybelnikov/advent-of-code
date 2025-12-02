defmodule FactorizeTest do
  use ExUnit.Case

  describe "factorize/1" do
    test "factorizes 60 into prime factors" do
      assert Factors.factorize(60) == [{2, 2}, {3, 1}, {5, 1}]
    end

    test "factorizes 1 (edge case)" do
      assert Factors.factorize(1) == []
    end

    test "factorizes prime number 7" do
      assert Factors.factorize(7) == [{7, 1}]
    end

    test "factorizes prime number 13" do
      assert Factors.factorize(13) == [{13, 1}]
    end

    test "factorizes 100 (perfect square)" do
      assert Factors.factorize(100) == [{2, 2}, {5, 2}]
    end

    test "factorizes 128 (power of 2)" do
      assert Factors.factorize(128) == [{2, 7}]
    end

    test "factorizes 2310 (product of first 5 primes)" do
      assert Factors.factorize(2310) == [{2, 1}, {3, 1}, {5, 1}, {7, 1}, {11, 1}]
    end

    test "factorizes 1000" do
      assert Factors.factorize(1000) == [{2, 3}, {5, 3}]
    end

    test "factorizes large prime 97" do
      assert Factors.factorize(97) == [{97, 1}]
    end
  end

  describe "get_factors/1 with caching" do
    test "retrieves cached factors on second call" do
      # First call computes and caches
      factors1 = Factors.get_factors(60)
      assert factors1 == [{2, 2}, {3, 1}, {5, 1}]

      # Second call retrieves from cache
      factors2 = Factors.get_factors(60)
      assert factors2 == [{2, 2}, {3, 1}, {5, 1}]
    end

    test "caches different numbers independently" do
      assert Factors.get_factors(12) == [{2, 2}, {3, 1}]
      assert Factors.get_factors(18) == [{2, 1}, {3, 2}]
      assert Factors.get_factors(12) == [{2, 2}, {3, 1}]
    end
  end
end
