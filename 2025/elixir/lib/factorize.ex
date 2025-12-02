defmodule Factors do
  # Ensure factor cache ETS table exists
  defp ensure_cache do
    if :ets.whereis(:factor_cache) == :undefined do
      try do
        :ets.new(:factor_cache, [:set, :public, :named_table])
      rescue
        # Ignore :already_exists error if another process created the table
        ArgumentError -> :ok
      end
    end
  end

  @doc """
  Retrieves the prime factors of n from cache or computes them
  """
  def get_factors(n) do
    ensure_cache()

    case :ets.lookup(:factor_cache, n) do
      [{^n, factors}] ->
        factors

      [] ->
        factors = factorize(n)
        :ets.insert(:factor_cache, {n, factors})
        factors
    end
  end

  @doc """
  Factorizes n into its prime factors
  """
  def factorize(n), do: factorize(n, 2, [])
  defp factorize(1, _, acc), do: Enum.reverse(acc)

  defp factorize(n, p, acc) when p * p > n,
    do: Enum.reverse([{n, 1} | acc])

  defp factorize(n, p, acc) do
    if rem(n, p) == 0 do
      {count, rest} = count_power(n, p, 0)
      factorize(rest, p + 1, [{p, count} | acc])
    else
      factorize(n, p + 1, acc)
    end
  end

  # Counts how many times n is divisible by p
  defp count_power(n, p, count) do
    if rem(n, p) == 0 do
      count_power(div(n, p), p, count + 1)
    else
      {count, n}
    end
  end
end
