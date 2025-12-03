defmodule Day03 do
  @moduledoc """
  Advent of Code - Day 03
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.to_charlist/1)
    |> Stream.map(&pick_max_k_prefix(&1, 2))
    |> Stream.map(&List.to_integer/1)
    |> Enum.sum()
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    input
    |> String.trim()
    |> String.split("\n")
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.to_charlist/1)
    |> Stream.map(&pick_max_k_prefix(&1, 12))
    |> Stream.map(&List.to_integer/1)
    |> Enum.sum()
  end

  def find_max_joltage_2(battery) do
    String.trim(battery) |> String.to_charlist() |> pick_max_k_prefix(2)
  end

  def find_max_joltage_12(battery) do
    String.trim(battery) |> String.to_charlist() |> pick_max_k_prefix(12)
  end

  # Generalized: picks the lexicographically largest subsequence of length k using a greedy stack.
  defp pick_max_k_prefix(charlist, k) do
    removals = length(charlist) - k

    {stack_rev, _rem} =
      Enum.reduce(charlist, {[], removals}, fn ch, {stack, rem} ->
        {stack_after, rem_after} = pop_smaller(stack, rem, ch)
        {[ch | stack_after], rem_after}
      end)

    stack = Enum.reverse(stack_rev)
    Enum.take(stack, k)
  end

  defp pop_smaller([s | rest], rem, ch) when s < ch and rem > 0 do
    pop_smaller(rest, rem - 1, ch)
  end

  defp pop_smaller(stack, rem, _ch), do: {stack, rem}

  @doc """
  Reads input file and runs both parts
  """
  def solve do
    input = read_input()

    IO.puts("--- Part One ---")
    IO.puts("Result: #{part1(input)}")

    IO.puts("--- Part Two ---")
    IO.puts("Result: #{part2(input)}")
  end

  defp read_input do
    File.read!("inputs/day03.txt")
  end
end
