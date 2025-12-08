defmodule Day07 do
  import Bitwise

  @moduledoc """
  Advent of Code - Day 07
  """
  @doc """
  Solves Part 1
  """
  def part1(input) do
    input
    |> String.replace("\r", "")
    |> String.split("\n", trim: true)
    |> Enum.map(&String.to_charlist/1)
    |> then(fn [head | tail] ->
      entry = Enum.find_index(head, fn x -> x == ?S end)

      tail
      |> Enum.with_index()
      |> Enum.reduce_while({0, [entry]}, fn {level, idx}, {acc, indices} ->
        if idx == length(tail) - 1 do
          {:halt, {acc, indices}}
        else
          {splits, beams} = pass_level(level, indices)
          {:cont, {acc + splits, beams}}
        end
      end)
    end)
    |> elem(0)
  end

  def pass_level(level, indices) do
    indices
    |> Enum.reduce({0, []}, fn idx, {splits, beams} ->
      result = neighbours(level, idx)

      if length(result) == 0 do
        {splits, [idx | beams]}
      else
        {splits + 1, result ++ beams}
      end
    end)
    |> then(fn {splits, beams} -> {splits, Enum.uniq(beams)} end)
  end

  defp neighbours(level, idx) do
    if Enum.at(level, idx) == ?^ do
      neighbors = []
      neighbors = if idx > 0, do: [idx - 1 | neighbors], else: neighbors
      neighbors = if idx < length(level) - 1, do: [idx + 1 | neighbors], else: neighbors
      neighbors
    else
      []
    end
  end

  @doc """
  Solves Part 2
  """
  def part2(input) do
    input
    |> String.replace("\r", "")
    |> String.split("\n", trim: true)
    |> Enum.map(&String.to_charlist/1)
    |> then(fn [head | tail] ->
      entry = head |> Enum.find_index(fn x -> x == ?S end)

      rows =
        tail
        |> Enum.map(fn row ->
          # Convert each row to an integer with bits representing positions
          Enum.reduce(row, 0, fn char, acc ->
            bit = if char == ?^, do: 1, else: 0
            acc <<< 1 ||| bit
          end)
        end)

      num_cols = length(head)
      count_states(rows, num_cols, entry)
    end)
  end

  def count_states(rows, num_cols, start_col, element_bit \\ 1) do
    rows
    |> Enum.drop(-1)
    |> Enum.reduce(%{start_col => 1}, fn row_bits, dp ->
      dp
      |> Enum.reject(fn {_col, ways} -> ways == 0 end)
      |> Enum.reduce(%{}, fn {col, ways}, next_dp ->
        bit = (row_bits >>> col) &&& 1

        next_cols =
          if bit == element_bit do
            [col - 1, col + 1] |> Enum.filter(&(&1 >= 0 and &1 < num_cols))
          else
            [col]
          end

        Enum.reduce(next_cols, next_dp, fn next_col, acc ->
          Map.update(acc, next_col, ways, &(&1 + ways))
        end)
      end)
    end)
    |> Map.values()
    |> Enum.sum()
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
