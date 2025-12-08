defmodule Day05 do
  @moduledoc """
  Advent of Code - Day 05
  """

  @doc """
  Solves Part 1
  """
  def part1(input) do
    [l, r] = String.split(input, "\n\n", trim: true)

    ranges =
      l
      |> String.split("\n", trim: true)
      |> Enum.map(fn line ->
        [start, finish] =
          line
          |> String.split("-", trim: true)
          |> Enum.map(&String.to_integer/1)

        start..finish
      end)

    ids =
      r
      |> String.split("\n", trim: true)
      |> Enum.map(&String.to_integer/1)

    Enum.count(ids, fn x -> Enum.any?(ranges, fn range -> x in range end) end)
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    [l, _r] = String.split(input, "\n\n", trim: true)

    l
    |> String.split("\n", trim: true)
    |> Stream.map(fn line ->
      line
      |> String.split("-", trim: true)
      |> Enum.map(&String.to_integer/1)
    end)
    |> Enum.sort(&(hd(&1) <= hd(&2)))
    |> Stream.map(fn [start, finish] -> start..finish end)
    |> Enum.reduce([], fn range, acc ->
      case acc do
        [] ->
          [range]

        [last_range | rest] ->
          if Range.disjoint?(last_range, range) do
            [range | acc]
          else
            [min(last_range.first, range.first)..max(last_range.last, range.last) | rest]
          end
      end
    end)
    |> Enum.reduce(0, fn r, acc -> Enum.count(r) + acc end)
  end

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
    File.read!("inputs/day05.txt")
  end
end
