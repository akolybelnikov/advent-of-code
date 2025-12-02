defmodule Divisors do
  @doc """
  Given a factorization like [{2, 2}, {3, 1}, {5, 1}],
  returns the sorted list of all divisors in desc order.
  """
  def from_factors([]), do: [1]

  def from_factors(factors) do
    factors
    |> Enum.reduce([1], fn {p, e}, divisors ->
      powers = prime_powers(p, e)
      for d <- divisors, pow <- powers, do: d * pow
    end)
    |> Enum.sort(:desc)
  end

  # helper: [p^0, p^1, ..., p^e]
  defp prime_powers(p, e) do
    Stream.iterate(1, &(&1 * p))
    |> Enum.take(e + 1)
  end
end
